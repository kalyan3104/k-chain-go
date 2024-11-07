package mock

import (
	"github.com/kalyan3104/k-chain-core-go/data/transaction"
	"github.com/kalyan3104/k-chain-go/common"
	"github.com/kalyan3104/k-chain-go/process"
	vmcommon "github.com/kalyan3104/k-chain-vm-common-go"
)

// ScQueryStub -
type ScQueryStub struct {
	ExecuteQueryCalled           func(query *process.SCQuery) (*vmcommon.VMOutput, common.BlockInfo, error)
	ComputeScCallGasLimitHandler func(tx *transaction.Transaction) (uint64, error)
	CloseCalled                  func() error
}

// ExecuteQuery -
func (s *ScQueryStub) ExecuteQuery(query *process.SCQuery) (*vmcommon.VMOutput, common.BlockInfo, error) {
	if s.ExecuteQueryCalled != nil {
		return s.ExecuteQueryCalled(query)
	}
	return &vmcommon.VMOutput{}, nil, nil
}

// ComputeScCallGasLimit -
func (s *ScQueryStub) ComputeScCallGasLimit(tx *transaction.Transaction) (uint64, error) {
	if s.ComputeScCallGasLimitHandler != nil {
		return s.ComputeScCallGasLimitHandler(tx)
	}
	return 100, nil
}

// Close -
func (s *ScQueryStub) Close() error {
	if s.CloseCalled != nil {
		return s.CloseCalled()
	}

	return nil
}

// IsInterfaceNil -
func (s *ScQueryStub) IsInterfaceNil() bool {
	return s == nil
}
