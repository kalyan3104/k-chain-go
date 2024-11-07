package epochProviders

import (
	"github.com/kalyan3104/k-chain-go/config"
	"github.com/kalyan3104/k-chain-go/dataRetriever"
	"github.com/kalyan3104/k-chain-go/dataRetriever/resolvers/epochproviders"
	"github.com/kalyan3104/k-chain-go/dataRetriever/resolvers/epochproviders/disabled"
)

// CreateCurrentEpochProvider will create an instance of dataRetriever.CurrentNetworkEpochProviderHandler
func CreateCurrentEpochProvider(
	generalConfigs config.Config,
	roundTimeInMilliseconds uint64,
	startTime int64,
	isFullArchive bool,
) (dataRetriever.CurrentNetworkEpochProviderHandler, error) {
	if !isFullArchive {
		return disabled.NewEpochProvider(), nil
	}

	arg := epochproviders.ArgArithmeticEpochProvider{
		RoundsPerEpoch:          uint32(generalConfigs.EpochStartConfig.RoundsPerEpoch),
		RoundTimeInMilliseconds: roundTimeInMilliseconds,
		StartTime:               startTime,
	}

	return epochproviders.NewArithmeticEpochProvider(arg)
}
