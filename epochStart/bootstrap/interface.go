package bootstrap

import (
	"context"

	"github.com/kalyan3104/k-chain-core-go/core"
	"github.com/kalyan3104/k-chain-core-go/data"
	"github.com/kalyan3104/k-chain-core-go/data/block"
	"github.com/kalyan3104/k-chain-go/dataRetriever"
	"github.com/kalyan3104/k-chain-go/process"
	"github.com/kalyan3104/k-chain-go/sharding/nodesCoordinator"
)

// StartOfEpochNodesConfigHandler defines the methods to process nodesConfig from epoch start metablocks
type StartOfEpochNodesConfigHandler interface {
	NodesConfigFromMetaBlock(currMetaBlock data.HeaderHandler, prevMetaBlock data.HeaderHandler) (nodesCoordinator.NodesCoordinatorRegistryHandler, uint32, []*block.MiniBlock, error)
	IsInterfaceNil() bool
}

// EpochStartMetaBlockInterceptorProcessor defines the methods to sync an epoch start metablock
type EpochStartMetaBlockInterceptorProcessor interface {
	process.InterceptorProcessor
	GetEpochStartMetaBlock(ctx context.Context) (data.MetaHeaderHandler, error)
}

// StartInEpochNodesCoordinator defines the methods to process and save nodesCoordinator information to storage
type StartInEpochNodesCoordinator interface {
	EpochStartPrepare(metaHdr data.HeaderHandler, body data.BodyHandler)
	NodesCoordinatorToRegistry(epoch uint32) nodesCoordinator.NodesCoordinatorRegistryHandler
	ShardIdForEpoch(epoch uint32) (uint32, error)
	IsInterfaceNil() bool
}

// Messenger defines which methods a p2p messenger should implement
type Messenger interface {
	dataRetriever.MessageHandler
	dataRetriever.TopicHandler
	UnregisterMessageProcessor(topic string, identifier string) error
	UnregisterAllMessageProcessors() error
	UnJoinAllTopics() error
	ConnectedPeers() []core.PeerID
	Verify(payload []byte, pid core.PeerID, signature []byte) error
	Broadcast(topic string, buff []byte)
	BroadcastUsingPrivateKey(topic string, buff []byte, pid core.PeerID, skBytes []byte)
	Sign(payload []byte) ([]byte, error)
	SignUsingPrivateKey(skBytes []byte, payload []byte) ([]byte, error)
}

// RequestHandler defines which methods a request handler should implement
type RequestHandler interface {
	RequestStartOfEpochMetaBlock(epoch uint32)
	SetNumPeersToQuery(topic string, intra int, cross int) error
	GetNumPeersToQuery(topic string) (int, int, error)
	IsInterfaceNil() bool
}

// NodeTypeProviderHandler defines the actions needed for a component that can handle the node type
type NodeTypeProviderHandler interface {
	SetType(nodeType core.NodeType)
	GetType() core.NodeType
	IsInterfaceNil() bool
}
