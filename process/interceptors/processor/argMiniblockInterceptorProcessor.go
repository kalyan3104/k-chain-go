package processor

import (
	"github.com/kalyan3104/k-chain-core-go/hashing"
	"github.com/kalyan3104/k-chain-core-go/marshal"
	"github.com/kalyan3104/k-chain-go/process"
	"github.com/kalyan3104/k-chain-go/sharding"
	"github.com/kalyan3104/k-chain-go/storage"
)

// ArgMiniblockInterceptorProcessor is the argument for the interceptor processor used for miniblocks
type ArgMiniblockInterceptorProcessor struct {
	MiniblockCache   storage.Cacher
	Marshalizer      marshal.Marshalizer
	Hasher           hashing.Hasher
	ShardCoordinator sharding.Coordinator
	WhiteListHandler process.WhiteListHandler
}
