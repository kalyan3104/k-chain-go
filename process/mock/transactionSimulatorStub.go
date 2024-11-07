package mock

import (
	"github.com/kalyan3104/k-chain-core-go/data"
	"github.com/kalyan3104/k-chain-core-go/data/transaction"
	txSimData "github.com/kalyan3104/k-chain-go/process/transactionEvaluator/data"
)

// TransactionSimulatorStub -
type TransactionSimulatorStub struct {
	ProcessTxCalled func(tx *transaction.Transaction, currentHeader data.HeaderHandler) (*txSimData.SimulationResultsWithVMOutput, error)
}

// ProcessTx -
func (tss *TransactionSimulatorStub) ProcessTx(tx *transaction.Transaction, currentHeader data.HeaderHandler) (*txSimData.SimulationResultsWithVMOutput, error) {
	if tss.ProcessTxCalled != nil {
		return tss.ProcessTxCalled(tx, currentHeader)
	}

	return nil, nil
}

// IsInterfaceNil -
func (tss *TransactionSimulatorStub) IsInterfaceNil() bool {
	return tss == nil
}
