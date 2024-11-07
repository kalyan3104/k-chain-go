package factory

import (
	"github.com/kalyan3104/k-chain-core-go/core"
	"github.com/kalyan3104/k-chain-go/node/external"
	"github.com/kalyan3104/k-chain-go/node/trieIterators"
	"github.com/kalyan3104/k-chain-go/node/trieIterators/disabled"
)

// CreateDelegatedListHandler will create a new instance of DirectStakedListHandler
func CreateDelegatedListHandler(args trieIterators.ArgTrieIteratorProcessor) (external.DelegatedListHandler, error) {
	//TODO add unit tests
	if args.ShardID != core.MetachainShardId {
		return disabled.NewDisabledDelegatedListProcessor(), nil
	}

	return trieIterators.NewDelegatedListProcessor(args)
}
