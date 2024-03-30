package stagedsync

import "github.com/optimism-java/erigon/turbo/stages/bodydownload"

type DownloaderGlue interface {
	SpawnHeaderDownloadStage([]func() error, *StageState, Unwinder) error
	SpawnBodyDownloadStage(string, string, *StageState, Unwinder, *bodydownload.PrefetchedBlocks) (bool, error)
}
