package disabled

import (
	"github.com/kalyan3104/k-chain-core-go/data"
	"github.com/kalyan3104/k-chain-go/epochStart"
)

// EpochStartNotifier -
type EpochStartNotifier struct {
}

// NewEpochStartNotifier returns the disabled epoch start notifier
func NewEpochStartNotifier() *EpochStartNotifier {
	return &EpochStartNotifier{}
}

// RegisterHandler -
func (desn *EpochStartNotifier) RegisterHandler(_ epochStart.ActionHandler) {
}

// UnregisterHandler -
func (desn *EpochStartNotifier) UnregisterHandler(_ epochStart.ActionHandler) {
}

// NotifyAllPrepare -
func (desn *EpochStartNotifier) NotifyAllPrepare(_ data.HeaderHandler, _ data.BodyHandler, _ epochStart.ValidatorInfoCacher) {
}

// NotifyAll -
func (desn *EpochStartNotifier) NotifyAll(_ data.HeaderHandler) {
}

// IsInterfaceNil -
func (desn *EpochStartNotifier) IsInterfaceNil() bool {
	return desn == nil
}
