package mock

import (
	"github.com/kalyan3104/k-chain-go/config"
	"github.com/kalyan3104/k-chain-go/storage"
	"github.com/kalyan3104/k-chain-go/testscommon/genericMocks"
)

// UnitOpenerStub -
type UnitOpenerStub struct {
}

// OpenDB -
func (u *UnitOpenerStub) OpenDB(_ config.DBConfig, _ uint32, _ uint32) (storage.Storer, error) {
	return genericMocks.NewStorerMock(), nil
}

// GetMostRecentStorageUnit -
func (u *UnitOpenerStub) GetMostRecentStorageUnit(_ config.DBConfig) (storage.Storer, error) {
	return genericMocks.NewStorerMock(), nil
}

// IsInterfaceNil -
func (u *UnitOpenerStub) IsInterfaceNil() bool {
	return u == nil
}
