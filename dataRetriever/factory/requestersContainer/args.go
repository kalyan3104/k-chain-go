package requesterscontainer

import (
	"github.com/kalyan3104/k-chain-core-go/data/typeConverters"
	"github.com/kalyan3104/k-chain-core-go/marshal"
	"github.com/kalyan3104/k-chain-go/config"
	"github.com/kalyan3104/k-chain-go/dataRetriever"
	"github.com/kalyan3104/k-chain-go/p2p"
	"github.com/kalyan3104/k-chain-go/sharding"
)

// FactoryArgs will hold the arguments for RequestersContainerFactory for both shard and meta
type FactoryArgs struct {
	RequesterConfig                 config.RequesterConfig
	ShardCoordinator                sharding.Coordinator
	MainMessenger                   p2p.Messenger
	FullArchiveMessenger            p2p.Messenger
	Marshaller                      marshal.Marshalizer
	Uint64ByteSliceConverter        typeConverters.Uint64ByteSliceConverter
	OutputAntifloodHandler          dataRetriever.P2PAntifloodHandler
	CurrentNetworkEpochProvider     dataRetriever.CurrentNetworkEpochProviderHandler
	MainPreferredPeersHolder        p2p.PreferredPeersHolderHandler
	FullArchivePreferredPeersHolder p2p.PreferredPeersHolderHandler
	PeersRatingHandler              dataRetriever.PeersRatingHandler
	SizeCheckDelta                  uint32
}
