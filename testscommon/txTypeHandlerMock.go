package testscommon

import (
	"github.com/kalyan3104/k-chain-core-go/data"
	"github.com/kalyan3104/k-chain-go/process"
)

// TxTypeHandlerMock -
type TxTypeHandlerMock struct {
	ComputeTransactionTypeCalled func(tx data.TransactionHandler) (process.TransactionType, process.TransactionType)
}

// ComputeTransactionType -
func (th *TxTypeHandlerMock) ComputeTransactionType(tx data.TransactionHandler) (process.TransactionType, process.TransactionType) {
	if th.ComputeTransactionTypeCalled == nil {
		return process.MoveBalance, process.MoveBalance
	}

	return th.ComputeTransactionTypeCalled(tx)
}

// IsInterfaceNil -
func (th *TxTypeHandlerMock) IsInterfaceNil() bool {
	return th == nil
}
