package mock

import (
	"github.com/kalyan3104/k-chain-core-go/data"
	"github.com/kalyan3104/k-chain-core-go/data/block"
	"github.com/kalyan3104/k-chain-go/dataRetriever"
)

// MiniBlockProvider defines what a miniblock data provider should do
type MiniBlockProvider interface {
	GetMiniBlocks(hashes [][]byte) ([]*block.MiniblockAndHash, [][]byte)
	GetMiniBlocksFromPool(hashes [][]byte) ([]*block.MiniblockAndHash, [][]byte)
	GetMiniBlocksFromStorer(hashes [][]byte) ([]*block.MiniblockAndHash, [][]byte)
	IsInterfaceNil() bool
}

// DataComponentsMock -
type DataComponentsMock struct {
	Storage           dataRetriever.StorageService
	Blkc              data.ChainHandler
	DataPool          dataRetriever.PoolsHolder
	MiniBlockProvider MiniBlockProvider
}

// StorageService -
func (dcm *DataComponentsMock) StorageService() dataRetriever.StorageService {
	return dcm.Storage
}

// Blockchain -
func (dcm *DataComponentsMock) Blockchain() data.ChainHandler {
	return dcm.Blkc
}

// Clone -
func (dcm *DataComponentsMock) Clone() interface{} {
	return &DataComponentsMock{
		Storage:  dcm.StorageService(),
		Blkc:     dcm.Blockchain(),
		DataPool: dcm.Datapool(),
	}
}

// Datapool -
func (dcm *DataComponentsMock) Datapool() dataRetriever.PoolsHolder {
	return dcm.DataPool
}

// MiniBlocksProvider -
func (dcm *DataComponentsMock) MiniBlocksProvider() MiniBlockProvider {
	return dcm.MiniBlockProvider
}

// SetBlockchain -
func (dcm *DataComponentsMock) SetBlockchain(chain data.ChainHandler) error {
	dcm.Blkc = chain
	return nil
}

// IsInterfaceNil -
func (dcm *DataComponentsMock) IsInterfaceNil() bool {
	return dcm == nil
}
