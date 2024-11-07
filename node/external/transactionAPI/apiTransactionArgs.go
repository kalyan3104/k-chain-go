package transactionAPI

import (
	"time"

	"github.com/kalyan3104/k-chain-core-go/core"
	"github.com/kalyan3104/k-chain-core-go/data/typeConverters"
	"github.com/kalyan3104/k-chain-core-go/marshal"
	"github.com/kalyan3104/k-chain-go/dataRetriever"
	"github.com/kalyan3104/k-chain-go/dblookupext"
	"github.com/kalyan3104/k-chain-go/process"
	"github.com/kalyan3104/k-chain-go/sharding"
)

// ArgAPITransactionProcessor is structure that store components that are needed to create an api transaction processor
type ArgAPITransactionProcessor struct {
	RoundDuration            uint64
	GenesisTime              time.Time
	Marshalizer              marshal.Marshalizer
	AddressPubKeyConverter   core.PubkeyConverter
	ShardCoordinator         sharding.Coordinator
	HistoryRepository        dblookupext.HistoryRepository
	StorageService           dataRetriever.StorageService
	DataPool                 dataRetriever.PoolsHolder
	Uint64ByteSliceConverter typeConverters.Uint64ByteSliceConverter
	FeeComputer              feeComputer
	TxTypeHandler            process.TxTypeHandler
	LogsFacade               LogsFacade
	DataFieldParser          DataFieldParser
}
