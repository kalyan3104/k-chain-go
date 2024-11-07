package mock

import (
	"github.com/kalyan3104/k-chain-core-go/data/transaction"
	"github.com/kalyan3104/k-chain-go/common"
	"github.com/kalyan3104/k-chain-go/process"
	vmcommon "github.com/kalyan3104/k-chain-vm-common-go"
)

// QueryServiceStub -
type QueryServiceStub struct {
	ComputeScCallGasLimitCalled func(tx *transaction.Transaction) (uint64, error)
	ExecuteQueryCalled          func(query *process.SCQuery) (*vmcommon.VMOutput, common.BlockInfo, error)
	CloseCalled                 func() error
}

// ComputeScCallGasLimit -
func (qss *QueryServiceStub) ComputeScCallGasLimit(tx *transaction.Transaction) (uint64, error) {
	if qss.ComputeScCallGasLimitCalled != nil {
		return qss.ComputeScCallGasLimitCalled(tx)
	}

	return 0, nil
}

// ExecuteQuery -
func (qss *QueryServiceStub) ExecuteQuery(query *process.SCQuery) (*vmcommon.VMOutput, common.BlockInfo, error) {
	if qss.ExecuteQueryCalled != nil {
		return qss.ExecuteQueryCalled(query)
	}

	return &vmcommon.VMOutput{}, nil, nil
}

// Close -
func (qss *QueryServiceStub) Close() error {
	if qss.CloseCalled != nil {
		return qss.CloseCalled()
	}

	return nil
}

// IsInterfaceNil -
func (qss *QueryServiceStub) IsInterfaceNil() bool {
	return qss == nil
}
