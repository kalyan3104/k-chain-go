package mock

import (
	"github.com/kalyan3104/k-chain-core-go/data/transaction"
	"github.com/kalyan3104/k-chain-go/common"
	"github.com/kalyan3104/k-chain-go/process"
	vmcommon "github.com/kalyan3104/k-chain-vm-common-go"
)

// SCQueryServiceStub -
type SCQueryServiceStub struct {
	ExecuteQueryCalled           func(*process.SCQuery) (*vmcommon.VMOutput, common.BlockInfo, error)
	ComputeScCallGasLimitHandler func(tx *transaction.Transaction) (uint64, error)
	CloseCalled                  func() error
}

// ExecuteQuery -
func (serviceStub *SCQueryServiceStub) ExecuteQuery(query *process.SCQuery) (*vmcommon.VMOutput, common.BlockInfo, error) {
	return serviceStub.ExecuteQueryCalled(query)
}

// ComputeScCallGasLimit -
func (serviceStub *SCQueryServiceStub) ComputeScCallGasLimit(tx *transaction.Transaction) (uint64, error) {
	return serviceStub.ComputeScCallGasLimitHandler(tx)
}

// Close -
func (serviceStub *SCQueryServiceStub) Close() error {
	if serviceStub.CloseCalled != nil {
		return serviceStub.CloseCalled()
	}

	return nil
}

// IsInterfaceNil returns true if there is no value under the interface
func (serviceStub *SCQueryServiceStub) IsInterfaceNil() bool {
	return serviceStub == nil
}
