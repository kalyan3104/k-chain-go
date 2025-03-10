package mock

import (
	"context"

	"github.com/kalyan3104/k-chain-core-go/data/api"
)

// StakeValuesProcessorStub -
type StakeValuesProcessorStub struct {
	GetTotalStakedValueCalled func(ctx context.Context) (*api.StakeValues, error)
}

// GetTotalStakedValue -
func (svps *StakeValuesProcessorStub) GetTotalStakedValue(ctx context.Context) (*api.StakeValues, error) {
	if svps.GetTotalStakedValueCalled != nil {
		return svps.GetTotalStakedValueCalled(ctx)
	}

	return nil, nil
}

// IsInterfaceNil -
func (svps *StakeValuesProcessorStub) IsInterfaceNil() bool {
	return svps == nil
}
