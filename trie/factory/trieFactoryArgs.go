package factory

import (
	"github.com/kalyan3104/k-chain-core-go/hashing"
	"github.com/kalyan3104/k-chain-core-go/marshal"
	"github.com/kalyan3104/k-chain-go/common"
	"github.com/kalyan3104/k-chain-go/config"
	"github.com/kalyan3104/k-chain-go/storage"
)

// TrieFactoryArgs holds the arguments for creating a trie factory
type TrieFactoryArgs struct {
	Marshalizer              marshal.Marshalizer
	Hasher                   hashing.Hasher
	PathManager              storage.PathManagerHandler
	TrieStorageManagerConfig config.TrieStorageManagerConfig
	StateStatsHandler        common.StateStatisticsHandler
}
