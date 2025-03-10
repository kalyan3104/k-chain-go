package factory

import (
	"fmt"
	"os"
	"time"

	"github.com/kalyan3104/k-chain-core-go/core"
	"github.com/kalyan3104/k-chain-core-go/core/check"
	"github.com/kalyan3104/k-chain-core-go/marshal"
	"github.com/kalyan3104/k-chain-go/config"
	"github.com/kalyan3104/k-chain-go/dataRetriever"
	"github.com/kalyan3104/k-chain-go/dataRetriever/dataPool"
	"github.com/kalyan3104/k-chain-go/dataRetriever/dataPool/headersCache"
	"github.com/kalyan3104/k-chain-go/dataRetriever/shardedData"
	"github.com/kalyan3104/k-chain-go/dataRetriever/txpool"
	"github.com/kalyan3104/k-chain-go/process"
	"github.com/kalyan3104/k-chain-go/sharding"
	"github.com/kalyan3104/k-chain-go/storage"
	"github.com/kalyan3104/k-chain-go/storage/cache"
	"github.com/kalyan3104/k-chain-go/storage/disabled"
	"github.com/kalyan3104/k-chain-go/storage/factory"
	"github.com/kalyan3104/k-chain-go/storage/storageunit"
	trieFactory "github.com/kalyan3104/k-chain-go/trie/factory"
	logger "github.com/kalyan3104/k-chain-logger-go"
)

const (
	peerAuthenticationCacheRefresh = time.Minute * 5
	peerAuthExpiryMultiplier       = time.Duration(2) // 2 times the computed duration
)

var log = logger.GetOrCreate("dataRetriever/factory")

// ArgsDataPool holds the arguments needed for NewDataPoolFromConfig function
type ArgsDataPool struct {
	Config           *config.Config
	EconomicsData    process.EconomicsDataHandler
	ShardCoordinator sharding.Coordinator
	Marshalizer      marshal.Marshalizer
	PathManager      storage.PathManagerHandler
}

// NewDataPoolFromConfig will return a new instance of a PoolsHolder
func NewDataPoolFromConfig(args ArgsDataPool) (dataRetriever.PoolsHolder, error) {
	log.Debug("creatingDataPool from config")

	if args.Config == nil {
		return nil, dataRetriever.ErrNilConfig
	}
	if check.IfNil(args.EconomicsData) {
		return nil, dataRetriever.ErrNilEconomicsData
	}
	if check.IfNil(args.ShardCoordinator) {
		return nil, dataRetriever.ErrNilShardCoordinator
	}
	if check.IfNil(args.PathManager) {
		return nil, dataRetriever.ErrNilPathManager
	}

	mainConfig := args.Config

	txPool, err := txpool.NewShardedTxPool(txpool.ArgShardedTxPool{
		Config:         factory.GetCacherFromConfig(mainConfig.TxDataPool),
		NumberOfShards: args.ShardCoordinator.NumberOfShards(),
		SelfShardID:    args.ShardCoordinator.SelfId(),
		TxGasHandler:   args.EconomicsData,
	})
	if err != nil {
		return nil, fmt.Errorf("%w while creating the cache for the transactions", err)
	}

	uTxPool, err := shardedData.NewShardedData(dataRetriever.UnsignedTxPoolName, factory.GetCacherFromConfig(mainConfig.UnsignedTransactionDataPool))
	if err != nil {
		return nil, fmt.Errorf("%w while creating the cache for the unsigned transactions", err)
	}

	rewardTxPool, err := shardedData.NewShardedData(dataRetriever.RewardTxPoolName, factory.GetCacherFromConfig(mainConfig.RewardTransactionDataPool))
	if err != nil {
		return nil, fmt.Errorf("%w while creating the cache for the rewards", err)
	}

	hdrPool, err := headersCache.NewHeadersPool(mainConfig.HeadersPoolConfig)
	if err != nil {
		return nil, fmt.Errorf("%w while creating the cache for the headers", err)
	}

	cacherCfg := factory.GetCacherFromConfig(mainConfig.TxBlockBodyDataPool)
	txBlockBody, err := storageunit.NewCache(cacherCfg)
	if err != nil {
		return nil, fmt.Errorf("%w while creating the cache for the miniblocks", err)
	}

	cacherCfg = factory.GetCacherFromConfig(mainConfig.PeerBlockBodyDataPool)
	peerChangeBlockBody, err := storageunit.NewCache(cacherCfg)
	if err != nil {
		return nil, fmt.Errorf("%w while creating the cache for the peer mini block body", err)
	}

	cacher, err := cache.NewCapacityLRU(
		int(mainConfig.TrieSyncStorage.Capacity),
		int64(mainConfig.TrieSyncStorage.SizeInBytes),
	)
	if err != nil {
		return nil, fmt.Errorf("%w while creating the cache for the trie nodes", err)
	}

	trieSyncDB, err := createTrieSyncDB(args)
	if err != nil {
		return nil, err
	}

	tnf := trieFactory.NewTrieNodeFactory()
	adaptedTrieNodesStorage, err := storageunit.NewStorageCacherAdapter(cacher, trieSyncDB, tnf, args.Marshalizer)
	if err != nil {
		return nil, fmt.Errorf("%w while creating the adapter for the trie nodes", err)
	}

	cacherCfg = factory.GetCacherFromConfig(mainConfig.TrieNodesChunksDataPool)
	trieNodesChunks, err := storageunit.NewCache(cacherCfg)
	if err != nil {
		return nil, fmt.Errorf("%w while creating the cache for the trie chunks", err)
	}

	cacherCfg = factory.GetCacherFromConfig(mainConfig.SmartContractDataPool)
	smartContracts, err := storageunit.NewCache(cacherCfg)
	if err != nil {
		return nil, fmt.Errorf("%w while creating the cache for the smartcontract results", err)
	}

	peerAuthPool, err := cache.NewTimeCacher(cache.ArgTimeCacher{
		DefaultSpan: time.Duration(mainConfig.HeartbeatV2.PeerAuthenticationTimeBetweenSendsInSec) * time.Second * peerAuthExpiryMultiplier,
		CacheExpiry: peerAuthenticationCacheRefresh,
	})
	if err != nil {
		return nil, fmt.Errorf("%w while creating the cache for the peer authentication messages", err)
	}

	cacherCfg = factory.GetCacherFromConfig(mainConfig.HeartbeatV2.HeartbeatPool)
	heartbeatPool, err := storageunit.NewCache(cacherCfg)
	if err != nil {
		return nil, fmt.Errorf("%w while creating the cache for the heartbeat messages", err)
	}

	validatorsInfo, err := shardedData.NewShardedData(dataRetriever.ValidatorsInfoPoolName, factory.GetCacherFromConfig(mainConfig.ValidatorInfoPool))
	if err != nil {
		return nil, fmt.Errorf("%w while creating the cache for the validator info results", err)
	}

	currBlockTransactions := dataPool.NewCurrentBlockTransactionsPool()
	currEpochValidatorInfo := dataPool.NewCurrentEpochValidatorInfoPool()
	dataPoolArgs := dataPool.DataPoolArgs{
		Transactions:              txPool,
		UnsignedTransactions:      uTxPool,
		RewardTransactions:        rewardTxPool,
		Headers:                   hdrPool,
		MiniBlocks:                txBlockBody,
		PeerChangesBlocks:         peerChangeBlockBody,
		TrieNodes:                 adaptedTrieNodesStorage,
		TrieNodesChunks:           trieNodesChunks,
		CurrentBlockTransactions:  currBlockTransactions,
		CurrentEpochValidatorInfo: currEpochValidatorInfo,
		SmartContracts:            smartContracts,
		PeerAuthentications:       peerAuthPool,
		Heartbeats:                heartbeatPool,
		ValidatorsInfo:            validatorsInfo,
	}
	return dataPool.NewDataPool(dataPoolArgs)
}

func createTrieSyncDB(args ArgsDataPool) (storage.Persister, error) {
	mainConfig := args.Config

	if !mainConfig.TrieSyncStorage.EnableDB {
		log.Debug("no DB for the intercepted trie nodes")
		return disabled.NewPersister(), nil
	}

	shardId := core.GetShardIDString(args.ShardCoordinator.SelfId())
	path := args.PathManager.PathForStatic(shardId, mainConfig.TrieSyncStorage.DB.FilePath)

	dbConfigHandler := factory.NewDBConfigHandler(mainConfig.TrieSyncStorage.DB)
	persisterFactory, err := factory.NewPersisterFactory(dbConfigHandler)
	if err != nil {
		return nil, err
	}

	if mainConfig.TrieSyncStorage.DB.UseTmpAsFilePath {
		filePath, errTempDir := os.MkdirTemp("", "trieSyncStorage")
		if errTempDir != nil {
			return nil, errTempDir
		}

		path = filePath
	}

	db, err := storageunit.NewDB(persisterFactory, path)
	if err != nil {
		return nil, fmt.Errorf("%w while creating the db for the trie nodes", err)
	}

	return db, nil
}
