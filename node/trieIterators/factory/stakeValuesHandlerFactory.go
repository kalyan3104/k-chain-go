package factory

import (
	"github.com/kalyan3104/k-chain-core-go/core"
	"github.com/kalyan3104/k-chain-go/node/external"
	"github.com/kalyan3104/k-chain-go/node/trieIterators"
	"github.com/kalyan3104/k-chain-go/node/trieIterators/disabled"
)

// CreateTotalStakedValueHandler will create a new instance of TotalStakedValueHandler
func CreateTotalStakedValueHandler(args trieIterators.ArgTrieIteratorProcessor) (external.TotalStakedValueHandler, error) {
	if args.ShardID != core.MetachainShardId {
		return disabled.NewDisabledStakeValuesProcessor(), nil
	}

	return trieIterators.NewTotalStakedValueProcessor(args)
}
