package mock

import (
	"github.com/kalyan3104/k-chain-go/dataRetriever"
	"github.com/kalyan3104/k-chain-go/storage"
)

// StorageListProviderStub -
type StorageListProviderStub struct {
	GetAllStorersCalled func() map[dataRetriever.UnitType]storage.Storer
}

// GetAllStorers -
func (slps *StorageListProviderStub) GetAllStorers() map[dataRetriever.UnitType]storage.Storer {
	if slps.GetAllStorersCalled != nil {
		return slps.GetAllStorersCalled()
	}

	return nil
}

// IsInterfaceNil -
func (slps *StorageListProviderStub) IsInterfaceNil() bool {
	return slps == nil
}
