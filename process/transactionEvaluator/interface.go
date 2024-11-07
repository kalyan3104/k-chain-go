package transactionEvaluator

import (
	"github.com/kalyan3104/k-chain-core-go/data/transaction"
	vmcommon "github.com/kalyan3104/k-chain-vm-common-go"
	datafield "github.com/kalyan3104/k-chain-vm-common-go/parsers/dataField"
)

// TransactionProcessor defines the operations needed to be done by a transaction processor
type TransactionProcessor interface {
	ProcessTransaction(transaction *transaction.Transaction) (vmcommon.ReturnCode, error)
	VerifyTransaction(transaction *transaction.Transaction) error
	IsInterfaceNil() bool
}

// DataFieldParser defines what a data field parser should be able to do
type DataFieldParser interface {
	Parse(dataField []byte, sender, receiver []byte, numOfShards uint32) *datafield.ResponseParseData
}
