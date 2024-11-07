package stakingcommon

import "github.com/kalyan3104/k-chain-go/state"

// AuctionListSelectorStub -
type AuctionListSelectorStub struct {
	SelectNodesFromAuctionListCalled func(validatorsInfoMap state.ShardValidatorsInfoMapHandler, randomness []byte) error
}

// SelectNodesFromAuctionList -
func (als *AuctionListSelectorStub) SelectNodesFromAuctionList(
	validatorsInfoMap state.ShardValidatorsInfoMapHandler,
	randomness []byte,
) error {
	if als.SelectNodesFromAuctionListCalled != nil {
		return als.SelectNodesFromAuctionListCalled(validatorsInfoMap, randomness)
	}

	return nil
}

// IsInterfaceNil -
func (als *AuctionListSelectorStub) IsInterfaceNil() bool {
	return als == nil
}
