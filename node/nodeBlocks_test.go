package node_test

import (
	"testing"

	"github.com/kalyan3104/k-chain-core-go/data/block"
	"github.com/kalyan3104/k-chain-go/dataRetriever"
	"github.com/kalyan3104/k-chain-go/node"
	"github.com/kalyan3104/k-chain-go/storage"
	"github.com/kalyan3104/k-chain-go/testscommon/dblookupext"
	storageMocks "github.com/kalyan3104/k-chain-go/testscommon/storage"
	"github.com/stretchr/testify/require"
)

func TestNode_GetBlockHeaderByHash(t *testing.T) {
	coreComponents := getDefaultCoreComponents()
	stateComponents := getDefaultStateComponents()
	dataComponents := getDefaultDataComponents()
	processComponents := getDefaultProcessComponents()

	blockHash := []byte("blockHash")
	blockHeader := &block.Header{Nonce: 42}
	blockHeaderBytes, _ := coreComponents.InternalMarshalizer().Marshal(blockHeader)
	isDbLookupExtEnabled := true

	// Setup storage
	headersStorer := &storageMocks.StorerStub{}
	dataComponents.Store = &storageMocks.ChainStorerStub{
		GetStorerCalled: func(_ dataRetriever.UnitType) (storage.Storer, error) {
			return headersStorer, nil
		},
	}

	// Setup dblookupext
	processComponents.HistoryRepositoryInternal = &dblookupext.HistoryRepositoryStub{
		IsEnabledCalled: func() bool {
			return isDbLookupExtEnabled
		},
	}

	n, _ := node.NewNode(
		node.WithCoreComponents(coreComponents),
		node.WithStateComponents(stateComponents),
		node.WithDataComponents(dataComponents),
		node.WithProcessComponents(processComponents),
	)

	t.Run("with dblookupext", func(t *testing.T) {
		isDbLookupExtEnabled = true

		headersStorer.GetCalled = func(_ []byte) ([]byte, error) {
			require.Fail(t, "should not have been called")
			return nil, nil
		}
		headersStorer.GetFromEpochCalled = func(_ []byte, _ uint32) ([]byte, error) {
			return blockHeaderBytes, nil
		}

		header, err := n.GetBlockHeaderByHash(blockHash)
		require.Nil(t, err)
		require.Equal(t, blockHeader, header)
	})

	t.Run("without dblookupext", func(t *testing.T) {
		isDbLookupExtEnabled = false

		headersStorer.GetCalled = func(_ []byte) ([]byte, error) {
			return blockHeaderBytes, nil
		}
		headersStorer.GetFromEpochCalled = func(_ []byte, _ uint32) ([]byte, error) {
			require.Fail(t, "should not have been called")
			return nil, nil
		}

		header, err := n.GetBlockHeaderByHash(blockHash)
		require.Nil(t, err)
		require.Equal(t, blockHeader, header)
	})
}
