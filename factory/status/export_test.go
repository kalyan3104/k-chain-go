package status

import (
	"github.com/kalyan3104/k-chain-core-go/core"
	"github.com/kalyan3104/k-chain-go/epochStart"
	outportDriverFactory "github.com/kalyan3104/k-chain-go/outport/factory"
	"github.com/kalyan3104/k-chain-go/p2p"
)

// EpochStartEventHandler -
func (pc *statusComponents) EpochStartEventHandler() epochStart.ActionHandler {
	return pc.epochStartEventHandler()
}

// ComputeNumConnectedPeers -
func ComputeNumConnectedPeers(
	appStatusHandler core.AppStatusHandler,
	netMessenger p2p.Messenger,
	suffix string,
) {
	computeNumConnectedPeers(appStatusHandler, netMessenger, suffix)
}

// ComputeConnectedPeers -
func ComputeConnectedPeers(
	appStatusHandler core.AppStatusHandler,
	netMessenger p2p.Messenger,
	suffix string,
) {
	computeConnectedPeers(appStatusHandler, netMessenger, suffix)
}

// MakeHostDriversArgs -
func (scf *statusComponentsFactory) MakeHostDriversArgs() ([]outportDriverFactory.ArgsHostDriverFactory, error) {
	return scf.makeHostDriversArgs()
}
