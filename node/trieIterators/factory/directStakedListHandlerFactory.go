package factory

import (
	"github.com/kalyan3104/k-chain-core-go/core"
	"github.com/kalyan3104/k-chain-go/node/external"
	"github.com/kalyan3104/k-chain-go/node/trieIterators"
	"github.com/kalyan3104/k-chain-go/node/trieIterators/disabled"
)

// CreateDirectStakedListHandler will create a new instance of DirectStakedListHandler
func CreateDirectStakedListHandler(args trieIterators.ArgTrieIteratorProcessor) (external.DirectStakedListHandler, error) {
	//TODO add unit tests
	if args.ShardID != core.MetachainShardId {
		return disabled.NewDisabledDirectStakedListProcessor(), nil
	}

	return trieIterators.NewDirectStakedListProcessor(args)
}
