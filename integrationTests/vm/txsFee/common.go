package txsFee

import (
	"testing"

	"github.com/kalyan3104/k-chain-go/common"
	"github.com/kalyan3104/k-chain-go/integrationTests/vm"
	"github.com/kalyan3104/k-chain-go/state"
	"github.com/stretchr/testify/require"
)

func getAccount(tb testing.TB, testContext *vm.VMTestContext, scAddress []byte) state.UserAccountHandler {
	scAcc, err := testContext.Accounts.LoadAccount(scAddress)
	require.Nil(tb, err)
	acc, ok := scAcc.(state.UserAccountHandler)
	require.True(tb, ok)

	return acc
}

func getAccountDataTrie(tb testing.TB, testContext *vm.VMTestContext, address []byte) common.Trie {
	acc := getAccount(tb, testContext, address)
	dataTrieInstance, ok := acc.DataTrie().(common.Trie)
	require.True(tb, ok)

	return dataTrieInstance
}
