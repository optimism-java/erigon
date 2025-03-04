package downloader

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"

	"github.com/anacrolix/torrent"

	"github.com/c2h5oh/datasize"
	"github.com/ledgerwatch/erigon-lib/chain/snapcfg"
	"golang.org/x/sync/errgroup"

	"github.com/anacrolix/torrent/bencode"
	"github.com/anacrolix/torrent/metainfo"
	"github.com/ledgerwatch/erigon-lib/downloader/snaptype"
	"github.com/ledgerwatch/log/v3"
	"github.com/pelletier/go-toml/v2"
)

// WebSeeds - allow use HTTP-based infrastrucutre to support Bittorrent network
// it allows download .torrent files and data files from trusted url's (for example: S3 signed url)
type WebSeeds struct {
	lock sync.Mutex

	byFileName          snaptype.WebSeedUrls // HTTP urls of data files
	torrentUrls         snaptype.TorrentUrls // HTTP urls of .torrent files
	downloadTorrentFile bool
	torrentsWhitelist   snapcfg.Preverified

	seeds []*url.URL

	logger    log.Logger
	verbosity log.Lvl

	torrentFiles *TorrentFiles
}

func NewWebSeeds(seeds []*url.URL, verbosity log.Lvl, logger log.Logger) *WebSeeds {
	return &WebSeeds{
		seeds:     seeds,
		logger:    logger,
		verbosity: verbosity,
	}
}

func (d *WebSeeds) getWebDownloadInfo(ctx context.Context, t *torrent.Torrent) (infos []webDownloadInfo, seedHashMismatches []*seedHash, err error) {
	torrentHash := t.InfoHash().Bytes()

	for _, webseed := range d.seeds {
		downloadUrl := webseed.JoinPath(t.Name())

		if headRequest, err := http.NewRequestWithContext(ctx, http.MethodHead, downloadUrl.String(), nil); err == nil {
			headResponse, err := http.DefaultClient.Do(headRequest)
			if err != nil {
				continue
			}
			headResponse.Body.Close()

			if headResponse.StatusCode != http.StatusOK {
				d.logger.Debug("[snapshots.webseed] getWebDownloadInfo: HEAD request failed",
					"webseed", webseed.String(), "name", t.Name(), "status", headResponse.Status)
				continue
			}
			if meta, err := getWebpeerTorrentInfo(ctx, downloadUrl); err == nil {
				if bytes.Equal(torrentHash, meta.HashInfoBytes().Bytes()) {
					md5tag := headResponse.Header.Get("Etag")
					if md5tag != "" {
						md5tag = strings.Trim(md5tag, "\"")
					}

					infos = append(infos, webDownloadInfo{
						url:     downloadUrl,
						length:  headResponse.ContentLength,
						md5:     md5tag,
						torrent: t,
					})
				} else {
					hash := meta.HashInfoBytes()
					seedHashMismatches = append(seedHashMismatches, &seedHash{url: webseed, hash: &hash})
				}
			}
		}
		seedHashMismatches = append(seedHashMismatches, &seedHash{url: webseed})
	}

	return infos, seedHashMismatches, nil
}

func (d *WebSeeds) SetTorrent(t *TorrentFiles, whiteList snapcfg.Preverified, downloadTorrentFile bool) {
	d.downloadTorrentFile = downloadTorrentFile
	d.torrentsWhitelist = whiteList
	d.torrentFiles = t
}

func (d *WebSeeds) checkHasTorrents(manifestResponse snaptype.WebSeedsFromProvider, report *webSeedCheckReport) {
	// check that for each file in the manifest, there is a corresponding .torrent file
	torrentNames := make(map[string]struct{})
	for name := range manifestResponse {
		if strings.HasSuffix(name, ".torrent") {
			torrentNames[name] = struct{}{}
		}
	}
	hasTorrents := len(torrentNames) > 0
	report.missingTorrents = make([]string, 0)
	for name := range manifestResponse {
		// todo extract list of extensions which are
		//  seeded as torrents (kv, ef, v, seg)
		//  seeded as is (.txt, efi)
		//  temporarily not seedable (.idx)
		if !strings.HasSuffix(name, ".torrent") && !strings.HasSuffix(name, ".txt") {
			tname := name + ".torrent"
			if _, ok := torrentNames[tname]; !ok {
				report.missingTorrents = append(report.missingTorrents, name)
				continue
			}
			delete(torrentNames, tname)
		}
	}

	if len(torrentNames) > 0 {
		report.danglingTorrents = make([]string, 0, len(torrentNames))
		for file := range torrentNames {
			report.danglingTorrents = append(report.danglingTorrents, file)
		}
	}
	report.torrentsOK = len(report.missingTorrents) == 0 && len(report.danglingTorrents) == 0 && hasTorrents
}

func (d *WebSeeds) fetchFileEtags(ctx context.Context, manifestResponse snaptype.WebSeedsFromProvider) (tags map[string]string, invalidTags, etagFetchFailed []string, err error) {
	etagFetchFailed = make([]string, 0)
	tags = make(map[string]string)
	invalidTagsMap := make(map[string]string)

	for name, wurl := range manifestResponse {
		u, err := url.Parse(wurl)
		if err != nil {
			return nil, nil, nil, fmt.Errorf("webseed.fetchFileEtags: %w", err)
		}
		md5Tag, err := d.retrieveFileEtag(ctx, u)
		if err != nil {
			if errors.Is(err, ErrInvalidEtag) {
				invalidTagsMap[name] = md5Tag
				continue
			}
			if errors.Is(err, ErrEtagNotFound) {
				etagFetchFailed = append(etagFetchFailed, name)
				continue
			}
			d.logger.Debug("[snapshots.webseed] get file ETag", "err", err, "url", u.String())
			return nil, nil, nil, fmt.Errorf("webseed.fetchFileEtags: %w", err)
		}
		tags[name] = md5Tag
	}

	invalidTags = make([]string, 0)
	if len(invalidTagsMap) > 0 {
		for name, tag := range invalidTagsMap {
			invalidTags = append(invalidTags, fmt.Sprintf("%-50s %s", name, tag))
		}
	}
	return tags, invalidTags, etagFetchFailed, nil
}

func (d *WebSeeds) VerifyManifestedBuckets(ctx context.Context, failFast bool) error {
	var supErr error
	for _, webSeedProviderURL := range d.seeds {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		d.logger.Debug("[snapshots.webseed] verify manifest", "url", webSeedProviderURL.String())

		if err := d.VerifyManifestedBucket(ctx, webSeedProviderURL); err != nil {
			d.logger.Warn("[snapshots.webseed] verify manifest", "err", err)
			if failFast {
				return err
			} else {
				supErr = err
			}
		}
	}
	return supErr
}

type webSeedCheckReport struct {
	seed             *url.URL
	manifestExist    bool
	torrentsOK       bool
	missingTorrents  []string
	danglingTorrents []string
	totalEtags       int
	invalidEtags     []string
	etagFetchFailed  []string
}

func (w *webSeedCheckReport) sort() {
	sort.Strings(w.missingTorrents)
	sort.Strings(w.invalidEtags)
	sort.Strings(w.etagFetchFailed)
	sort.Strings(w.danglingTorrents)
}

func (w *webSeedCheckReport) String() string {
	if !w.manifestExist {
		return fmt.Sprintf("## REPORT on %s: manifest not found\n", w.seed)
	}
	w.sort()
	var b strings.Builder
	b.WriteString(fmt.Sprintf("## REPORT on %s\n", w.seed))
	b.WriteString(fmt.Sprintf(" - manifest exist: %t\n", w.manifestExist))
	b.WriteString(fmt.Sprintf(" - missing torrents (files without torrents): %d\n", len(w.missingTorrents)))
	b.WriteString(fmt.Sprintf(" - dangling (data file not found) torrents: %d\n", len(w.danglingTorrents)))
	b.WriteString(fmt.Sprintf(" - invalid ETags format: %d/%d\n", len(w.invalidEtags), w.totalEtags))
	b.WriteString(fmt.Sprintf(" - ETag fetch failed: %d/%d\n", len(w.etagFetchFailed), w.totalEtags))

	titles := []string{
		"Missing torrents",
		"Dangling torrents",
		"Invalid ETags format",
		"ETag fetch failed",
	}

	fnamess := [][]string{
		w.missingTorrents,
		w.danglingTorrents,
		w.invalidEtags,
		w.etagFetchFailed,
	}

	var printedAnything bool
	for ti, names := range fnamess {
		if len(names) == 0 {
			continue
		}
		if ti == 0 {
			b.WriteByte(10)
		}
		printedAnything = true
		b.WriteString(fmt.Sprintf("# %s\n", titles[ti]))
		for _, name := range names {
			b.WriteString(fmt.Sprintf("%s\n", name))
		}
		if ti != len(fnamess)-1 {
			b.WriteByte(10)
		}
	}
	if !printedAnything {
		b.WriteString(fmt.Sprintf("== OK  %s\n", w.seed.String()))
	} else {
		b.WriteString(fmt.Sprintf("== BAD %sn", w.seed.String()))
	}
	return b.String()
}

func (d *WebSeeds) VerifyManifestedBucket(ctx context.Context, webSeedProviderURL *url.URL) error {
	report := &webSeedCheckReport{seed: webSeedProviderURL}
	manifestResponse, err := d.retrieveManifest(ctx, webSeedProviderURL)
	report.manifestExist = len(manifestResponse) != 0
	defer func() { fmt.Printf("%s\n", report.String()) }()
	if err != nil {
		return err
	}

	d.checkHasTorrents(manifestResponse, report)
	remoteTags, invalidTags, noTags, err := d.fetchFileEtags(ctx, manifestResponse)
	if err != nil {
		return err
	}

	report.invalidEtags = invalidTags
	report.etagFetchFailed = noTags
	report.totalEtags = len(remoteTags) + len(noTags)
	return nil
}

func (d *WebSeeds) Discover(ctx context.Context, files []string, rootDir string) {
	listsOfFiles := d.constructListsOfFiles(ctx, d.seeds, files)
	torrentMap := d.makeTorrentUrls(listsOfFiles)
	webSeedMap := d.downloadTorrentFilesFromProviders(ctx, rootDir, torrentMap)
	d.makeWebSeedUrls(listsOfFiles, webSeedMap)
}

func (d *WebSeeds) constructListsOfFiles(ctx context.Context, httpProviders []*url.URL, diskProviders []string) []snaptype.WebSeedsFromProvider {
	log.Debug("[snapshots.webseed] providers", "http", len(httpProviders), "disk", len(diskProviders))
	listsOfFiles := make([]snaptype.WebSeedsFromProvider, 0, len(httpProviders)+len(diskProviders))

	for _, webSeedProviderURL := range httpProviders {
		select {
		case <-ctx.Done():
			return listsOfFiles
		default:
		}
		manifestResponse, err := d.retrieveManifest(ctx, webSeedProviderURL)
		if err != nil { // don't fail on error
			d.logger.Debug("[snapshots.webseed] get from HTTP provider", "err", err, "url", webSeedProviderURL.EscapedPath())
			continue
		}
		// check if we need to prohibit new downloads for some files
		for name := range manifestResponse {
			prohibited, err := d.torrentFiles.newDownloadsAreProhibited(name)
			if prohibited || err != nil {
				delete(manifestResponse, name)
			}
		}

		listsOfFiles = append(listsOfFiles, manifestResponse)
	}

	// add to list files from disk
	for _, webSeedFile := range diskProviders {
		response, err := d.readWebSeedsFile(webSeedFile)
		if err != nil { // don't fail on error
			d.logger.Debug("[snapshots.webseed] get from File provider", "err", err)
			continue
		}
		// check if we need to prohibit new downloads for some files
		for name := range response {
			prohibited, err := d.torrentFiles.newDownloadsAreProhibited(name)
			if prohibited || err != nil {
				delete(response, name)
			}
		}
		listsOfFiles = append(listsOfFiles, response)
	}
	return listsOfFiles
}

func (d *WebSeeds) makeTorrentUrls(listsOfFiles []snaptype.WebSeedsFromProvider) map[url.URL]string {
	torrentMap := map[url.URL]string{}
	torrentUrls := snaptype.TorrentUrls{}
	for _, urls := range listsOfFiles {
		for name, wUrl := range urls {
			if !strings.HasSuffix(name, ".torrent") {
				continue
			}
			if !nameWhitelisted(name, d.torrentsWhitelist) {
				continue
			}
			uri, err := url.ParseRequestURI(wUrl)
			if err != nil {
				d.logger.Debug("[snapshots] url is invalid", "url", wUrl, "err", err)
				continue
			}
			torrentUrls[name] = append(torrentUrls[name], uri)
			torrentMap[*uri] = strings.TrimSuffix(name, ".torrent")
		}
	}

	d.lock.Lock()
	defer d.lock.Unlock()
	d.torrentUrls = torrentUrls
	return torrentMap
}

func (d *WebSeeds) makeWebSeedUrls(listsOfFiles []snaptype.WebSeedsFromProvider, webSeedMap map[string]struct{}) {
	webSeedUrls := snaptype.WebSeedUrls{}
	for _, urls := range listsOfFiles {
		for name, wUrl := range urls {
			if strings.HasSuffix(name, ".torrent") {
				continue
			}
			if _, ok := webSeedMap[name]; ok {
				webSeedUrls[name] = append(webSeedUrls[name], wUrl)
			}
		}
	}

	d.lock.Lock()
	defer d.lock.Unlock()
	d.byFileName = webSeedUrls
}

func (d *WebSeeds) TorrentUrls() snaptype.TorrentUrls {
	d.lock.Lock()
	defer d.lock.Unlock()
	return d.torrentUrls
}

func (d *WebSeeds) Len() int {
	d.lock.Lock()
	defer d.lock.Unlock()
	return len(d.byFileName)
}

func (d *WebSeeds) ByFileName(name string) (metainfo.UrlList, bool) {
	d.lock.Lock()
	defer d.lock.Unlock()
	v, ok := d.byFileName[name]
	return v, ok
}

var ErrInvalidEtag = fmt.Errorf("invalid etag")
var ErrEtagNotFound = fmt.Errorf("not found")

func (d *WebSeeds) retrieveFileEtag(ctx context.Context, file *url.URL) (string, error) {
	request, err := http.NewRequest(http.MethodHead, file.String(), nil)
	if err != nil {
		return "", err
	}

	request = request.WithContext(ctx)
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return "", fmt.Errorf("webseed.http: %w, url=%s", err, file.String())
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return "", ErrEtagNotFound
		}
		return "", fmt.Errorf("webseed.http: status code %d, url=%s", resp.StatusCode, file.String())
	}

	etag := resp.Header.Get("Etag") // file md5
	if etag == "" {
		return "", fmt.Errorf("webseed.http: file has no etag, url=%s", file.String())
	}
	etag = strings.Trim(etag, "\"")
	if strings.Contains(etag, "-") {
		return etag, ErrInvalidEtag
	}
	return etag, nil
}

func (d *WebSeeds) retrieveManifest(ctx context.Context, webSeedProviderUrl *url.URL) (snaptype.WebSeedsFromProvider, error) {
	baseUrl := webSeedProviderUrl.String()
	ref, err := url.Parse("manifest.txt")
	if err != nil {
		return nil, err
	}
	u := webSeedProviderUrl.ResolveReference(ref)
	request, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	request = request.WithContext(ctx)
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("webseed.http: make request: %w, url=%s", err, u.String())
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		d.logger.Debug("[snapshots.webseed] /manifest.txt retrieval failed, no downloads from this webseed",
			"webseed", webSeedProviderUrl.String(), "status", resp.Status)
		return nil, fmt.Errorf("webseed.http: status=%d, url=%s", resp.StatusCode, u.String())
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("webseed.http: read: %w, url=%s, ", err, u.String())
	}

	response := snaptype.WebSeedsFromProvider{}
	fileNames := strings.Split(string(b), "\n")
	for fi, f := range fileNames {
		if strings.TrimSpace(f) == "" {
			if fi != len(fileNames)-1 {
				d.logger.Debug("[snapshots.webseed] empty line in manifest.txt", "webseed", webSeedProviderUrl.String(), "lineNum", fi)
			}
			continue
		}

		response[f], err = url.JoinPath(baseUrl, f)
		if err != nil {
			return nil, err
		}
	}
	d.logger.Debug("[snapshots.webseed] get from HTTP provider", "urls", len(response), "url", webSeedProviderUrl.EscapedPath())
	return response, nil
}

func (d *WebSeeds) readWebSeedsFile(webSeedProviderPath string) (snaptype.WebSeedsFromProvider, error) {
	_, fileName := filepath.Split(webSeedProviderPath)
	data, err := os.ReadFile(webSeedProviderPath)
	if err != nil {
		return nil, fmt.Errorf("webseed.readWebSeedsFile: file=%s, %w", fileName, err)
	}
	response := snaptype.WebSeedsFromProvider{}
	if err := toml.Unmarshal(data, &response); err != nil {
		return nil, fmt.Errorf("webseed.readWebSeedsFile: file=%s, %w", fileName, err)
	}
	d.logger.Debug("[snapshots.webseed] get from File provider", "urls", len(response), "file", fileName)
	return response, nil
}

// downloadTorrentFilesFromProviders - if they are not exist on file-system
func (d *WebSeeds) downloadTorrentFilesFromProviders(ctx context.Context, rootDir string, torrentMap map[url.URL]string) map[string]struct{} {
	// TODO: need more tests, need handle more forward-compatibility and backward-compatibility case
	//  - now, if add new type of .torrent files to S3 bucket - existing nodes will start downloading it. maybe need whitelist of file types
	//  - maybe need download new files if --snap.stop=true
	webSeedMap := map[string]struct{}{}
	var webSeeMapLock sync.RWMutex
	if !d.downloadTorrentFile {
		return webSeedMap
	}
	if len(d.TorrentUrls()) == 0 {
		return webSeedMap
	}
	var addedNew int
	e, ctx := errgroup.WithContext(ctx)
	e.SetLimit(1024)
	urlsByName := d.TorrentUrls()

	for fileName, tUrls := range urlsByName {
		name := fileName
		addedNew++
		if !strings.HasSuffix(name, ".seg.torrent") {
			_, fName := filepath.Split(name)
			d.logger.Log(d.verbosity, "[snapshots] webseed has .torrent, but we skip it because this file-type not supported yet", "name", fName)
			continue
		}

		tUrls := tUrls
		e.Go(func() error {
			for _, url := range tUrls {
				//validation happens inside
				_, err := d.callTorrentHttpProvider(ctx, url, name)
				if err != nil {
					d.logger.Log(d.verbosity, "[snapshots] got from webseed", "name", name, "err", err, "url", url)
					continue
				}
				//don't save .torrent here - do it inside downloader.Add
				webSeeMapLock.Lock()
				webSeedMap[torrentMap[*url]] = struct{}{}
				webSeeMapLock.Unlock()
				return nil
			}
			return nil
		})
	}
	if err := e.Wait(); err != nil {
		d.logger.Debug("[snapshots] webseed discover", "err", err)
	}
	return webSeedMap
}

func (d *WebSeeds) DownloadAndSaveTorrentFile(ctx context.Context, name string) (bool, error) {
	urls, ok := d.ByFileName(name)
	if !ok {
		return false, nil
	}
	for _, urlStr := range urls {
		parsedUrl, err := url.Parse(urlStr)
		if err != nil {
			continue
		}
		res, err := d.callTorrentHttpProvider(ctx, parsedUrl, name)
		if err != nil {
			return false, err
		}
		if d.torrentFiles.Exists(name) {
			continue
		}
		if err := d.torrentFiles.Create(name, res); err != nil {
			d.logger.Log(d.verbosity, "[snapshots] .torrent from webseed rejected", "name", name, "err", err)
			continue
		}
		return true, nil
	}

	return false, nil
}

func (d *WebSeeds) callTorrentHttpProvider(ctx context.Context, url *url.URL, fileName string) ([]byte, error) {
	request, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}
	request = request.WithContext(ctx)
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("webseed.downloadTorrentFile: host=%s, url=%s, %w", url.Hostname(), url.EscapedPath(), err)
	}
	defer resp.Body.Close()
	//protect against too small and too big data
	if resp.ContentLength == 0 || resp.ContentLength > int64(128*datasize.MB) {
		return nil, nil
	}
	res, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("webseed.downloadTorrentFile: host=%s, url=%s, %w", url.Hostname(), url.EscapedPath(), err)
	}
	if err = validateTorrentBytes(fileName, res, d.torrentsWhitelist); err != nil {
		return nil, fmt.Errorf("webseed.downloadTorrentFile: host=%s, url=%s, %w", url.Hostname(), url.EscapedPath(), err)
	}
	return res, nil
}

func validateTorrentBytes(fileName string, b []byte, whitelist snapcfg.Preverified) error {
	var mi metainfo.MetaInfo
	if err := bencode.NewDecoder(bytes.NewBuffer(b)).Decode(&mi); err != nil {
		return err
	}
	torrentHash := mi.HashInfoBytes()
	// files with different names can have same hash. means need check AND name AND hash.
	if !nameAndHashWhitelisted(fileName, torrentHash.String(), whitelist) {
		return fmt.Errorf(".torrent file is not whitelisted")
	}
	return nil
}

func nameWhitelisted(fileName string, whitelist snapcfg.Preverified) bool {
	return whitelist.Contains(strings.TrimSuffix(fileName, ".torrent"))
}

func nameAndHashWhitelisted(fileName, fileHash string, whitelist snapcfg.Preverified) bool {
	fileName = strings.TrimSuffix(fileName, ".torrent")

	for i := 0; i < len(whitelist); i++ {
		if whitelist[i].Name == fileName && whitelist[i].Hash == fileHash {
			return true
		}
	}
	return false
}
