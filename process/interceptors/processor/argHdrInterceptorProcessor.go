package processor

import (
	"github.com/kalyan3104/k-chain-go/dataRetriever"
	"github.com/kalyan3104/k-chain-go/process"
)

// ArgHdrInterceptorProcessor is the argument for the interceptor processor used for headers (shard, meta and so on)
type ArgHdrInterceptorProcessor struct {
	Headers        dataRetriever.HeadersPool
	BlockBlackList process.TimeCacher
}
