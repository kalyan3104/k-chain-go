package mock

import (
	"github.com/kalyan3104/k-chain-core-go/data"
	"github.com/kalyan3104/k-chain-core-go/data/block"
)

// ValidatorInfoSyncerStub -
type ValidatorInfoSyncerStub struct {
}

// SyncMiniBlocks -
func (vip *ValidatorInfoSyncerStub) SyncMiniBlocks(*block.MetaBlock) ([][]byte, data.BodyHandler, error) {
	return nil, nil, nil
}

// IsInterfaceNil -
func (vip *ValidatorInfoSyncerStub) IsInterfaceNil() bool {
	return vip == nil
}
