package factory

import (
	"github.com/kalyan3104/k-chain-go/state"
	"github.com/kalyan3104/k-chain-go/state/accounts"
	vmcommon "github.com/kalyan3104/k-chain-vm-common-go"
)

// PeerAccountCreator has a method to create a new peer account
type PeerAccountCreator struct {
}

// NewPeerAccountCreator creates a peer account creator
func NewPeerAccountCreator() state.AccountFactory {
	return &PeerAccountCreator{}
}

// CreateAccount calls the new Account creator and returns the result
func (pac *PeerAccountCreator) CreateAccount(address []byte) (vmcommon.AccountHandler, error) {
	return accounts.NewPeerAccount(address)
}

// IsInterfaceNil returns true if there is no value under the interface
func (pac *PeerAccountCreator) IsInterfaceNil() bool {
	return pac == nil
}
