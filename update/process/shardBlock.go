package process

import (
	"math/big"

	"github.com/kalyan3104/k-chain-core-go/core"
	"github.com/kalyan3104/k-chain-core-go/data"
	"github.com/kalyan3104/k-chain-core-go/data/block"
	"github.com/kalyan3104/k-chain-core-go/hashing"
	"github.com/kalyan3104/k-chain-core-go/marshal"
	"github.com/kalyan3104/k-chain-go/dataRetriever"
	"github.com/kalyan3104/k-chain-go/process"
	"github.com/kalyan3104/k-chain-go/sharding"
	"github.com/kalyan3104/k-chain-go/update"
)

// ArgsNewShardBlockCreatorAfterHardFork defines the arguments structure to create a new shard block creator
type ArgsNewShardBlockCreatorAfterHardFork struct {
	Hasher             hashing.Hasher
	ImportHandler      update.ImportHandler
	Marshalizer        marshal.Marshalizer
	PendingTxProcessor update.PendingTransactionProcessor
	ShardCoordinator   sharding.Coordinator
	Storage            dataRetriever.StorageService
	TxCoordinator      process.TransactionCoordinator
	ReceiptsRepository receiptsRepository
	SelfShardID        uint32
}

type shardBlockCreator struct {
	*baseProcessor
}

// NewShardBlockCreatorAfterHardFork creates a shard block processor for the first block after hardfork
func NewShardBlockCreatorAfterHardFork(args ArgsNewShardBlockCreatorAfterHardFork) (*shardBlockCreator, error) {
	err := checkBlockCreatorAfterHardForkNilParameters(
		args.Hasher,
		args.ImportHandler,
		args.Marshalizer,
		args.PendingTxProcessor,
		args.ShardCoordinator,
		args.Storage,
		args.TxCoordinator,
		args.ReceiptsRepository,
	)
	if err != nil {
		return nil, err
	}

	base := &baseProcessor{
		hasher:             args.Hasher,
		importHandler:      args.ImportHandler,
		marshalizer:        args.Marshalizer,
		pendingTxProcessor: args.PendingTxProcessor,
		shardCoordinator:   args.ShardCoordinator,
		storage:            args.Storage,
		txCoordinator:      args.TxCoordinator,
		receiptsRepository: args.ReceiptsRepository,
		selfShardID:        args.SelfShardID,
	}

	return &shardBlockCreator{
		baseProcessor: base,
	}, nil
}

// CreateBlock will create a block after hardfork import
func (s *shardBlockCreator) CreateBlock(
	body *block.Body,
	chainID string,
	round uint64,
	nonce uint64,
	epoch uint32,
) (data.HeaderHandler, error) {
	if len(chainID) == 0 {
		return nil, update.ErrEmptyChainID
	}

	rootHash, err := s.pendingTxProcessor.RootHash()
	if err != nil {
		return nil, err
	}

	shardHeader := &block.Header{
		Nonce:           nonce,
		ShardID:         s.shardCoordinator.SelfId(),
		Round:           round,
		Epoch:           epoch,
		ChainID:         []byte(chainID),
		SoftwareVersion: []byte(""),
		RootHash:        rootHash,
		RandSeed:        rootHash,
		PrevHash:        rootHash,
		PrevRandSeed:    rootHash,
		AccumulatedFees: big.NewInt(0),
		DeveloperFees:   big.NewInt(0),
		PubKeysBitmap:   []byte{1},
	}

	shardHeader.ReceiptsHash, err = s.txCoordinator.CreateReceiptsHash()
	if err != nil {
		return nil, err
	}

	totalTxCount, miniBlockHeaders, err := s.createMiniBlockHeaders(body)
	if err != nil {
		return nil, err
	}

	shardHeader.MiniBlockHeaders = miniBlockHeaders
	shardHeader.TxCount = uint32(totalTxCount)

	metaBlockHash, err := core.CalculateHash(s.marshalizer, s.hasher, s.importHandler.GetHardForkMetaBlock())
	if err != nil {
		return nil, err
	}

	shardHeader.MetaBlockHashes = [][]byte{metaBlockHash}

	s.saveAllBlockDataToStorageForSelfShard(shardHeader, body)

	return shardHeader, nil
}

// IsInterfaceNil returns true if underlying object is nil
func (s *shardBlockCreator) IsInterfaceNil() bool {
	return s == nil
}
