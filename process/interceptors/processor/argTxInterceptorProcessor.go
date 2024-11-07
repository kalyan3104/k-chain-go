package processor

import (
	"github.com/kalyan3104/k-chain-go/dataRetriever"
	"github.com/kalyan3104/k-chain-go/process"
)

// ArgTxInterceptorProcessor is the argument for the interceptor processor used for transactions
// (balance txs, smart contract results, reward and so on)
type ArgTxInterceptorProcessor struct {
	ShardedDataCache dataRetriever.ShardedDataCacherNotifier
	TxValidator      process.TxValidator
}
