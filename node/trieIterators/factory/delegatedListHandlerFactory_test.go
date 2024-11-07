package factory

import (
	"fmt"
	"sync"
	"testing"

	"github.com/kalyan3104/k-chain-core-go/core"
	"github.com/kalyan3104/k-chain-go/node/mock"
	"github.com/kalyan3104/k-chain-go/node/trieIterators"
	"github.com/kalyan3104/k-chain-go/testscommon"
	stateMock "github.com/kalyan3104/k-chain-go/testscommon/state"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateDelegatedListHandlerHandler_Disabled(t *testing.T) {
	t.Parallel()

	args := trieIterators.ArgTrieIteratorProcessor{
		ShardID: 0,
	}

	delegatedListHandler, err := CreateDelegatedListHandler(args)
	require.Nil(t, err)
	assert.Equal(t, "*disabled.delegatedListProcessor", fmt.Sprintf("%T", delegatedListHandler))
}

func TestCreateDelegatedListHandlerHandler_DelegatedListProcessor(t *testing.T) {
	t.Parallel()

	args := trieIterators.ArgTrieIteratorProcessor{
		ShardID: core.MetachainShardId,
		Accounts: &trieIterators.AccountsWrapper{
			Mutex:           &sync.Mutex{},
			AccountsAdapter: &stateMock.AccountsStub{},
		},
		PublicKeyConverter: &testscommon.PubkeyConverterMock{},
		QueryService:       &mock.SCQueryServiceStub{},
	}

	delegatedListHandler, err := CreateDelegatedListHandler(args)
	require.Nil(t, err)
	assert.Equal(t, "*trieIterators.delegatedListProcessor", fmt.Sprintf("%T", delegatedListHandler))
}
