package testscommon

import "github.com/kalyan3104/k-chain-go/common"

// IteratorChannelsProviderStub -
type IteratorChannelsProviderStub struct {
	GetIteratorChannelsCalled func() *common.TrieIteratorChannels
}

// GetIteratorChannels -
func (icps *IteratorChannelsProviderStub) GetIteratorChannels() *common.TrieIteratorChannels {
	if icps.GetIteratorChannelsCalled != nil {
		return icps.GetIteratorChannelsCalled()
	}

	return nil
}

// IsInterfaceNil returns true if there is no value under the interface
func (icps *IteratorChannelsProviderStub) IsInterfaceNil() bool {
	return icps == nil
}
