package caplin1

import (
	"context"
	"fmt"
	"os"
	"path"
	"time"

	"google.golang.org/grpc/credentials"

	proto_downloader "github.com/ledgerwatch/erigon-lib/gointerfaces/downloader"
	"github.com/optimism-java/erigon/cl/antiquary"
	"github.com/optimism-java/erigon/cl/beacon"
	"github.com/optimism-java/erigon/cl/beacon/beacon_router_configuration"
	"github.com/optimism-java/erigon/cl/beacon/beaconevents"
	"github.com/optimism-java/erigon/cl/beacon/handler"
	"github.com/optimism-java/erigon/cl/beacon/synced_data"
	"github.com/optimism-java/erigon/cl/clparams/initial_state"
	"github.com/optimism-java/erigon/cl/cltypes"
	"github.com/optimism-java/erigon/cl/cltypes/solid"
	"github.com/optimism-java/erigon/cl/fork"
	"github.com/optimism-java/erigon/cl/persistence"
	"github.com/optimism-java/erigon/cl/rpc"
	"github.com/optimism-java/erigon/cl/sentinel"
	"github.com/optimism-java/erigon/cl/sentinel/service"
	"github.com/optimism-java/erigon/cl/validator/attestation_producer"
	"github.com/optimism-java/erigon/cl/validator/validator_params"
	"github.com/optimism-java/erigon/eth/ethconfig"
	"github.com/optimism-java/erigon/params"
	"github.com/optimism-java/erigon/turbo/snapshotsync/freezeblocks"

	"github.com/spf13/afero"

	"github.com/optimism-java/erigon/cl/persistence/beacon_indicies"
	"github.com/optimism-java/erigon/cl/persistence/blob_storage"
	"github.com/optimism-java/erigon/cl/persistence/db_config"
	"github.com/optimism-java/erigon/cl/persistence/format/snapshot_format"
	state_accessors "github.com/optimism-java/erigon/cl/persistence/state"
	"github.com/optimism-java/erigon/cl/persistence/state/historical_states_reader"
	"github.com/optimism-java/erigon/cl/phase1/core/state"
	"github.com/optimism-java/erigon/cl/phase1/execution_client"
	"github.com/optimism-java/erigon/cl/phase1/forkchoice"
	"github.com/optimism-java/erigon/cl/phase1/forkchoice/fork_graph"
	"github.com/optimism-java/erigon/cl/phase1/network"
	"github.com/optimism-java/erigon/cl/phase1/stages"
	"github.com/optimism-java/erigon/cl/pool"

	"github.com/Giulio2002/bls"
	"github.com/ledgerwatch/log/v3"

	"github.com/ledgerwatch/erigon-lib/common/datadir"
	"github.com/ledgerwatch/erigon-lib/kv"
	"github.com/ledgerwatch/erigon-lib/kv/mdbx"
	"github.com/optimism-java/erigon/cl/clparams"
)

func OpenCaplinDatabase(ctx context.Context,
	databaseConfig db_config.DatabaseConfiguration,
	beaconConfig *clparams.BeaconChainConfig,
	genesisConfig *clparams.GenesisConfig,
	dbPath string,
	blobDir string,
	engine execution_client.ExecutionEngine,
	wipeout bool,
	blobPruneDistance uint64,
) (kv.RwDB, blob_storage.BlobStorage, error) {
	dataDirIndexer := path.Join(dbPath, "beacon_indicies")
	blobDbPath := path.Join(blobDir, "chaindata")

	if wipeout {
		os.RemoveAll(dataDirIndexer)
		os.RemoveAll(blobDbPath)
	}

	os.MkdirAll(dbPath, 0700)
	os.MkdirAll(dataDirIndexer, 0700)

	db := mdbx.MustOpen(dataDirIndexer)
	blobDB := mdbx.MustOpen(blobDbPath)

	tx, err := db.BeginRw(ctx)
	if err != nil {
		return nil, nil, err
	}
	defer tx.Rollback()

	if err := db_config.WriteConfigurationIfNotExist(ctx, tx, databaseConfig); err != nil {
		return nil, nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, nil, err
	}
	{ // start ticking forkChoice
		go func() {
			<-ctx.Done()
			db.Close()     // close sql database here
			blobDB.Close() // close blob database here
		}()
	}
	return db, blob_storage.NewBlobStore(blobDB, afero.NewBasePathFs(afero.NewOsFs(), blobDir), blobPruneDistance, beaconConfig, genesisConfig), nil
}

func RunCaplinPhase1(ctx context.Context, engine execution_client.ExecutionEngine, config *ethconfig.Config, networkConfig *clparams.NetworkConfig,
	beaconConfig *clparams.BeaconChainConfig, genesisConfig *clparams.GenesisConfig, state *state.CachingBeaconState, dirs datadir.Dirs, cfg beacon_router_configuration.RouterConfiguration, eth1Getter snapshot_format.ExecutionBlockReaderByNumber,
	snDownloader proto_downloader.DownloaderClient, backfilling, blobBackfilling bool, states bool, indexDB kv.RwDB, blobStorage blob_storage.BlobStorage, creds credentials.TransportCredentials) error {
	ctx, cn := context.WithCancel(ctx)
	defer cn()

	logger := log.New("app", "caplin")

	csn := freezeblocks.NewCaplinSnapshots(ethconfig.BlocksFreezing{}, beaconConfig, dirs, logger)
	rcsn := freezeblocks.NewBeaconSnapshotReader(csn, eth1Getter, beaconConfig)

	pool := pool.NewOperationsPool(beaconConfig)
	attestationProducer := attestation_producer.New(beaconConfig)

	caplinFcuPath := path.Join(dirs.Tmp, "caplin-forkchoice")
	os.RemoveAll(caplinFcuPath)
	err := os.MkdirAll(caplinFcuPath, 0o755)
	if err != nil {
		return err
	}
	fcuFs := afero.NewBasePathFs(afero.NewOsFs(), caplinFcuPath)
	syncedDataManager := synced_data.NewSyncedDataManager(true, beaconConfig)

	emitters := beaconevents.NewEmitters()
	forkChoice, err := forkchoice.NewForkChoiceStore(state, engine, pool, fork_graph.NewForkGraphDisk(state, fcuFs, cfg), emitters, syncedDataManager, blobStorage)
	if err != nil {
		logger.Error("Could not create forkchoice", "err", err)
		return err
	}
	bls.SetEnabledCaching(true)
	state.ForEachValidator(func(v solid.Validator, idx, total int) bool {
		pk := v.PublicKey()
		if err := bls.LoadPublicKeyIntoCache(pk[:], false); err != nil {
			panic(err)
		}
		return true
	})

	forkDigest, err := fork.ComputeForkDigest(beaconConfig, genesisConfig)
	if err != nil {
		return err
	}

	sentinel, err := service.StartSentinelService(&sentinel.SentinelConfig{
		IpAddr:        config.LightClientDiscoveryAddr,
		Port:          int(config.LightClientDiscoveryPort),
		TCPPort:       uint(config.LightClientDiscoveryTCPPort),
		GenesisConfig: genesisConfig,
		NetworkConfig: networkConfig,
		BeaconConfig:  beaconConfig,
		TmpDir:        dirs.Tmp,
		EnableBlocks:  true,
	}, rcsn, blobStorage, indexDB, &service.ServerConfig{
		Network:   "tcp",
		Addr:      fmt.Sprintf("%s:%d", config.SentinelAddr, config.SentinelPort),
		Creds:     creds,
		Validator: config.BeaconRouter.Validator,
		InitialStatus: &cltypes.Status{
			ForkDigest:     forkDigest,
			FinalizedRoot:  state.FinalizedCheckpoint().BlockRoot(),
			FinalizedEpoch: state.FinalizedCheckpoint().Epoch(),
			HeadSlot:       state.FinalizedCheckpoint().Epoch() * beaconConfig.SlotsPerEpoch,
			HeadRoot:       state.FinalizedCheckpoint().BlockRoot(),
		},
	}, forkChoice, logger)
	if err != nil {
		return err
	}

	beaconRpc := rpc.NewBeaconRpcP2P(ctx, sentinel, beaconConfig, genesisConfig)
	gossipSource := persistence.NewGossipSource(ctx)

	gossipManager := network.NewGossipReceiver(sentinel, forkChoice, beaconConfig, genesisConfig, emitters, gossipSource)
	{ // start ticking forkChoice
		go func() {
			tickInterval := time.NewTicker(2 * time.Millisecond)
			for {
				select {
				case <-tickInterval.C:
					forkChoice.OnTick(uint64(time.Now().Unix()))
				case <-ctx.Done():
					return
				}
			}
		}()
	}

	{ // start the gossip manager
		go gossipManager.Start(ctx)
		logger.Info("Started Ethereum 2.0 Gossip Service")
	}

	{ // start logging peers
		go func() {
			logIntervalPeers := time.NewTicker(1 * time.Minute)
			for {
				select {
				case <-logIntervalPeers.C:
					if peerCount, err := beaconRpc.Peers(); err == nil {
						logger.Info("P2P", "peers", peerCount)
					}
				case <-ctx.Done():
					return
				}
			}
		}()
	}

	tx, err := indexDB.BeginRw(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	dbConfig, err := db_config.ReadConfiguration(ctx, tx)
	if err != nil {
		return err
	}

	if err := state_accessors.InitializeStaticTables(tx, state); err != nil {
		return err
	}
	if err := beacon_indicies.WriteHighestFinalized(tx, 0); err != nil {
		return err
	}

	vTables := state_accessors.NewStaticValidatorTable()
	// Read the current table
	if states {
		if err := state_accessors.ReadValidatorsTable(tx, vTables); err != nil {
			return err
		}
	}
	// get the initial state
	genesisState, err := initial_state.GetGenesisState(clparams.NetworkType(beaconConfig.DepositNetworkID))
	if err != nil {
		return err
	}
	antiq := antiquary.NewAntiquary(ctx, blobStorage, genesisState, vTables, beaconConfig, dirs, snDownloader, indexDB, csn, rcsn, logger, states, backfilling, blobBackfilling)
	// Create the antiquary
	go func() {
		if err := antiq.Loop(); err != nil {
			logger.Error("Antiquary failed", "err", err)
		}
	}()

	if err := tx.Commit(); err != nil {
		return err
	}

	statesReader := historical_states_reader.NewHistoricalStatesReader(beaconConfig, rcsn, vTables, genesisState)
	validatorParameters := validator_params.NewValidatorParams()
	if cfg.Active {
		apiHandler := handler.NewApiHandler(logger, genesisConfig, beaconConfig, indexDB, forkChoice, pool, rcsn, syncedDataManager, statesReader, sentinel, params.GitTag, &cfg, emitters, blobStorage, csn, validatorParameters, attestationProducer, engine)
		go beacon.ListenAndServe(&beacon.LayeredBeaconHandler{
			ArchiveApi: apiHandler,
		}, cfg)
		log.Info("Beacon API started", "addr", cfg.Address)
	}

	forkChoice.StartJobsRTT(ctx)

	stageCfg := stages.ClStagesCfg(beaconRpc, antiq, genesisConfig, beaconConfig, state, engine, gossipManager, forkChoice, indexDB, csn, rcsn, dirs.Tmp, dbConfig, backfilling, blobBackfilling, syncedDataManager, emitters, gossipSource, blobStorage, attestationProducer)
	sync := stages.ConsensusClStages(ctx, stageCfg)

	logger.Info("[Caplin] starting clstages loop")
	err = sync.StartWithStage(ctx, "DownloadHistoricalBlocks", logger, stageCfg)
	logger.Info("[Caplin] exiting clstages loop")
	if err != nil {
		return err
	}
	return err
}
