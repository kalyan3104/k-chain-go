package data

import (
	"github.com/kalyan3104/k-chain-core-go/data/transaction"
	vmcommon "github.com/kalyan3104/k-chain-vm-common-go"
)

// SimulationResultsWithVMOutput is the data transfer object which will hold results for simulation a transaction's execution
type SimulationResultsWithVMOutput struct {
	transaction.SimulationResults
	VMOutput *vmcommon.VMOutput `json:"-"`
}
