package state

import (
	"github.com/kalyan3104/k-chain-core-go/data/api"
	"github.com/kalyan3104/k-chain-go/common"
	"github.com/kalyan3104/k-chain-go/state"
	vmcommon "github.com/kalyan3104/k-chain-vm-common-go"
)

// AccountsRepositoryStub -
type AccountsRepositoryStub struct {
	GetAccountWithBlockInfoCalled        func(address []byte, options api.AccountQueryOptions) (vmcommon.AccountHandler, common.BlockInfo, error)
	GetCodeWithBlockInfoCalled           func(codeHash []byte, options api.AccountQueryOptions) ([]byte, common.BlockInfo, error)
	GetCurrentStateAccountsWrapperCalled func() state.AccountsAdapterAPI
	CloseCalled                          func() error
}

// GetAccountWithBlockInfo -
func (stub *AccountsRepositoryStub) GetAccountWithBlockInfo(address []byte, options api.AccountQueryOptions) (vmcommon.AccountHandler, common.BlockInfo, error) {
	if stub.GetAccountWithBlockInfoCalled != nil {
		return stub.GetAccountWithBlockInfoCalled(address, options)
	}

	return nil, nil, nil
}

// GetCodeWithBlockInfo -
func (stub *AccountsRepositoryStub) GetCodeWithBlockInfo(codeHash []byte, options api.AccountQueryOptions) ([]byte, common.BlockInfo, error) {
	if stub.GetCodeWithBlockInfoCalled != nil {
		return stub.GetCodeWithBlockInfoCalled(codeHash, options)
	}

	return nil, nil, nil
}

// GetCurrentStateAccountsWrapper -
func (stub *AccountsRepositoryStub) GetCurrentStateAccountsWrapper() state.AccountsAdapterAPI {
	if stub.GetCurrentStateAccountsWrapperCalled != nil {
		return stub.GetCurrentStateAccountsWrapperCalled()
	}

	return nil
}

// Close -
func (stub *AccountsRepositoryStub) Close() error {
	if stub.CloseCalled != nil {
		return stub.CloseCalled()
	}

	return nil
}

// IsInterfaceNil -
func (stub *AccountsRepositoryStub) IsInterfaceNil() bool {
	return stub == nil
}
