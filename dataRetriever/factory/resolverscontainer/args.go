package resolverscontainer

import (
	"github.com/kalyan3104/k-chain-core-go/data/typeConverters"
	"github.com/kalyan3104/k-chain-core-go/marshal"
	"github.com/kalyan3104/k-chain-go/common"
	"github.com/kalyan3104/k-chain-go/dataRetriever"
	"github.com/kalyan3104/k-chain-go/p2p"
	"github.com/kalyan3104/k-chain-go/sharding"
)

// FactoryArgs will hold the arguments for ResolversContainerFactory for both shard and meta
type FactoryArgs struct {
	NumConcurrentResolvingJobs          int32
	NumConcurrentResolvingTrieNodesJobs int32
	ShardCoordinator                    sharding.Coordinator
	MainMessenger                       p2p.Messenger
	FullArchiveMessenger                p2p.Messenger
	Store                               dataRetriever.StorageService
	Marshalizer                         marshal.Marshalizer
	DataPools                           dataRetriever.PoolsHolder
	Uint64ByteSliceConverter            typeConverters.Uint64ByteSliceConverter
	DataPacker                          dataRetriever.DataPacker
	TriesContainer                      common.TriesHolder
	InputAntifloodHandler               dataRetriever.P2PAntifloodHandler
	OutputAntifloodHandler              dataRetriever.P2PAntifloodHandler
	MainPreferredPeersHolder            p2p.PreferredPeersHolderHandler
	FullArchivePreferredPeersHolder     p2p.PreferredPeersHolderHandler
	SizeCheckDelta                      uint32
	IsFullHistoryNode                   bool
	PayloadValidator                    dataRetriever.PeerAuthenticationPayloadValidator
}
