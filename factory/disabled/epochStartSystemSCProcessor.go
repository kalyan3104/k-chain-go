package disabled

import (
	"github.com/kalyan3104/k-chain-core-go/data"
	"github.com/kalyan3104/k-chain-core-go/data/block"
	"github.com/kalyan3104/k-chain-go/epochStart"
	"github.com/kalyan3104/k-chain-go/state"
)

type epochStartSystemSCProcessor struct {
}

// NewDisabledEpochStartSystemSC creates a new disabled EpochStartSystemSCProcessor instance
func NewDisabledEpochStartSystemSC() *epochStartSystemSCProcessor {
	return &epochStartSystemSCProcessor{}
}

// ToggleUnStakeUnBond returns nil
func (e *epochStartSystemSCProcessor) ToggleUnStakeUnBond(_ bool) error {
	return nil
}

// ProcessSystemSmartContract returns nil
func (e *epochStartSystemSCProcessor) ProcessSystemSmartContract(
	_ state.ShardValidatorsInfoMapHandler,
	_ data.HeaderHandler,
) error {
	return nil
}

// ProcessDelegationRewards returns nil
func (e *epochStartSystemSCProcessor) ProcessDelegationRewards(
	_ block.MiniBlockSlice,
	_ epochStart.TransactionCacher,
) error {
	return nil
}

// IsInterfaceNil returns true if there is no value under the interface
func (e *epochStartSystemSCProcessor) IsInterfaceNil() bool {
	return e == nil
}
