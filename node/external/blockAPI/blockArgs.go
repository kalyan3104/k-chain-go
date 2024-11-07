package blockAPI

import (
	"github.com/kalyan3104/k-chain-core-go/core"
	"github.com/kalyan3104/k-chain-core-go/data/transaction"
	"github.com/kalyan3104/k-chain-core-go/data/typeConverters"
	"github.com/kalyan3104/k-chain-core-go/hashing"
	"github.com/kalyan3104/k-chain-core-go/marshal"
	"github.com/kalyan3104/k-chain-go/common"
	"github.com/kalyan3104/k-chain-go/dataRetriever"
	"github.com/kalyan3104/k-chain-go/dblookupext"
	outportProcess "github.com/kalyan3104/k-chain-go/outport/process"
	"github.com/kalyan3104/k-chain-go/process"
	"github.com/kalyan3104/k-chain-go/state"
)

// ArgAPIBlockProcessor is structure that store components that are needed to create an api block processor
type ArgAPIBlockProcessor struct {
	SelfShardID                  uint32
	Store                        dataRetriever.StorageService
	Marshalizer                  marshal.Marshalizer
	Uint64ByteSliceConverter     typeConverters.Uint64ByteSliceConverter
	HistoryRepo                  dblookupext.HistoryRepository
	APITransactionHandler        APITransactionHandler
	StatusComputer               transaction.StatusComputerHandler
	Hasher                       hashing.Hasher
	AddressPubkeyConverter       core.PubkeyConverter
	LogsFacade                   logsFacade
	ReceiptsRepository           receiptsRepository
	AlteredAccountsProvider      outportProcess.AlteredAccountsProviderHandler
	AccountsRepository           state.AccountsRepository
	ScheduledTxsExecutionHandler process.ScheduledTxsExecutionHandler
	EnableEpochsHandler          common.EnableEpochsHandler
}
