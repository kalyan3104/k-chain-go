package peer

import (
	"github.com/kalyan3104/k-chain-go/dataRetriever"
	"github.com/kalyan3104/k-chain-go/epochStart"
	"github.com/kalyan3104/k-chain-go/state"
)

// DataPool indicates the main functionality needed in order to fetch the required blocks from the pool
type DataPool interface {
	Headers() dataRetriever.HeadersPool
	IsInterfaceNil() bool
}

// StakingDataProviderAPI is able to provide staking data from the system smart contracts
type StakingDataProviderAPI interface {
	ComputeUnQualifiedNodes(validatorInfos state.ShardValidatorsInfoMapHandler) ([][]byte, map[string][][]byte, error)
	FillValidatorInfo(validator state.ValidatorInfoHandler) error
	GetOwnersData() map[string]*epochStart.OwnerData
	Clean()
	IsInterfaceNil() bool
}
