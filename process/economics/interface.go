package economics

import (
	vmcommon "github.com/kalyan3104/k-chain-vm-common-go"
)

// EpochNotifier raises epoch change events
type EpochNotifier interface {
	RegisterNotifyHandler(handler vmcommon.EpochSubscriberHandler)
	IsInterfaceNil() bool
}
