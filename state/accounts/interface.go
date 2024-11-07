package accounts

import (
	"github.com/kalyan3104/k-chain-go/state"
)

type dataTrieInteractor interface {
	state.DataTrieTracker
}
