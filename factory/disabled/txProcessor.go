package disabled

import (
	"github.com/kalyan3104/k-chain-core-go/data/transaction"
	vmcommon "github.com/kalyan3104/k-chain-vm-common-go"
)

// TxProcessor implements the TransactionProcessor interface but does nothing as it is disabled
type TxProcessor struct {
}

// ProcessTransaction does nothing as it is disabled
func (txProc *TxProcessor) ProcessTransaction(_ *transaction.Transaction) (vmcommon.ReturnCode, error) {
	return 0, nil
}

// VerifyTransaction does nothing as it is disabled
func (txProc *TxProcessor) VerifyTransaction(_ *transaction.Transaction) error {
	return nil
}

// IsInterfaceNil returns true if there is no value under the interface
func (txProc *TxProcessor) IsInterfaceNil() bool {
	return txProc == nil
}
