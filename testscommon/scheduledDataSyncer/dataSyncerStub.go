package scheduledDataSyncer

import (
	"github.com/kalyan3104/k-chain-core-go/data"
	"github.com/kalyan3104/k-chain-core-go/data/block"
)

// ScheduledSyncerStub -
type ScheduledSyncerStub struct {
	UpdateSyncDataIfNeededCalled func(notarizedShardHeader data.ShardHeaderHandler) (data.ShardHeaderHandler, map[string]data.HeaderHandler, map[string]*block.MiniBlock, error)
	GetRootHashToSyncCalled      func(notarizedShardHeader data.ShardHeaderHandler) []byte
}

// UpdateSyncDataIfNeeded -
func (sdss *ScheduledSyncerStub) UpdateSyncDataIfNeeded(notarizedShardHeader data.ShardHeaderHandler) (data.ShardHeaderHandler, map[string]data.HeaderHandler, map[string]*block.MiniBlock, error) {
	if sdss.UpdateSyncDataIfNeededCalled != nil {
		return sdss.UpdateSyncDataIfNeededCalled(notarizedShardHeader)
	}
	return nil, nil, nil, nil
}

// GetRootHashToSync -
func (sdss *ScheduledSyncerStub) GetRootHashToSync(notarizedShardHeader data.ShardHeaderHandler) []byte {
	if sdss.GetRootHashToSyncCalled != nil {
		return sdss.GetRootHashToSyncCalled(notarizedShardHeader)
	}
	return []byte("rootHash to sync")
}

// IsInterfaceNil -
func (sdss *ScheduledSyncerStub) IsInterfaceNil() bool {
	return sdss == nil
}
