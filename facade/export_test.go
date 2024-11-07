package facade

import (
	"github.com/kalyan3104/k-chain-go/ntp"
)

// GetSyncer returns the current syncer
func (nf *nodeFacade) GetSyncer() ntp.SyncTimer {
	return nf.syncer
}
