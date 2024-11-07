package process

import (
	"github.com/kalyan3104/k-chain-core-go/core"
	"github.com/kalyan3104/k-chain-core-go/hashing"
	"github.com/kalyan3104/k-chain-core-go/marshal"
	"github.com/kalyan3104/k-chain-go/common"
	"github.com/kalyan3104/k-chain-go/state"
	disabledState "github.com/kalyan3104/k-chain-go/state/disabled"
	"github.com/kalyan3104/k-chain-go/state/storagePruningManager/disabled"
	"github.com/kalyan3104/k-chain-go/trie"
)

const maxTrieLevelInMemory = uint(5)

func createAccountAdapter(
	marshaller marshal.Marshalizer,
	hasher hashing.Hasher,
	accountFactory state.AccountFactory,
	trieStorage common.StorageManager,
	addressConverter core.PubkeyConverter,
	enableEpochsHandler common.EnableEpochsHandler,
) (state.AccountsAdapter, error) {
	tr, err := trie.NewTrie(trieStorage, marshaller, hasher, enableEpochsHandler, maxTrieLevelInMemory)
	if err != nil {
		return nil, err
	}

	args := state.ArgsAccountsDB{
		Trie:                  tr,
		Hasher:                hasher,
		Marshaller:            marshaller,
		AccountFactory:        accountFactory,
		StoragePruningManager: disabled.NewDisabledStoragePruningManager(),
		AddressConverter:      addressConverter,
		SnapshotsManager:      disabledState.NewDisabledSnapshotsManager(),
	}

	adb, err := state.NewAccountsDB(args)
	if err != nil {
		return nil, err
	}

	return adb, nil
}
