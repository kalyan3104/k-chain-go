package mock

import (
	"github.com/kalyan3104/k-chain-go/common"
	"github.com/kalyan3104/k-chain-go/config"
)

// TrieFactoryStub -
type TrieFactoryStub struct {
	CreateCalled func(config config.StorageConfig, s string, b bool) (common.StorageManager, common.Trie, error)
}

// Create -
func (t *TrieFactoryStub) Create(config config.StorageConfig, s string, b bool) (common.StorageManager, common.Trie, error) {
	if t.CreateCalled != nil {
		return t.CreateCalled(config, s, b)
	}
	return nil, nil, nil
}

// IsInterfaceNil -
func (t *TrieFactoryStub) IsInterfaceNil() bool {
	return t == nil
}
