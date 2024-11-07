package mock

import (
	"github.com/kalyan3104/k-chain-core-go/data"
	"github.com/kalyan3104/k-chain-go/state"
)

// ValidatorInfoSyncerStub -
type ValidatorInfoSyncerStub struct {
}

// SyncMiniBlocks -
func (vip *ValidatorInfoSyncerStub) SyncMiniBlocks(_ data.HeaderHandler) ([][]byte, data.BodyHandler, error) {
	return nil, nil, nil
}

// SyncValidatorsInfo -
func (vip *ValidatorInfoSyncerStub) SyncValidatorsInfo(_ data.BodyHandler) ([][]byte, map[string]*state.ShardValidatorInfo, error) {
	return nil, nil, nil
}

// IsInterfaceNil -
func (vip *ValidatorInfoSyncerStub) IsInterfaceNil() bool {
	return vip == nil
}
