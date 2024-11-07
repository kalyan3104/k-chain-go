package rewards

import (
	"bytes"
	"testing"

	"github.com/kalyan3104/k-chain-core-go/core"
	"github.com/kalyan3104/k-chain-go/integrationTests/resolvers"
	"github.com/kalyan3104/k-chain-go/process"
	"github.com/kalyan3104/k-chain-go/process/factory"
)

func TestRequestResolveLargeSCRByHashRequestingShardResolvingOtherShard(t *testing.T) {
	if testing.Short() {
		t.Skip("this is not a short test")
	}

	rm := resolvers.NewReceiverMonitor(t)
	shardIdResolver := core.MetachainShardId
	shardIdRequester := uint32(0)
	nResolver, nRequester := resolvers.CreateResolverRequester(shardIdResolver, shardIdRequester)
	defer func() {
		nRequester.Close()
		nResolver.Close()
	}()
	scr, hash := resolvers.CreateLargeSmartContractResults()

	cacheId := process.ShardCacherIdentifier(shardIdRequester, core.MetachainShardId)
	nResolver.DataPool.UnsignedTransactions().AddData(hash, scr, scr.Size(), cacheId)

	// setup header received event
	nRequester.DataPool.UnsignedTransactions().RegisterOnAdded(
		func(key []byte, value interface{}) {
			if bytes.Equal(key, hash) {
				resolvers.Log.Info("received reward tx", "hash", key)
				rm.Done()
			}
		},
	)

	// request by hash should work
	requester, err := nRequester.RequestersFinder.CrossShardRequester(factory.UnsignedTransactionTopic, core.MetachainShardId)
	resolvers.Log.LogIfError(err)
	err = requester.RequestDataFromHash(hash, 0)
	resolvers.Log.LogIfError(err)

	rm.WaitWithTimeout()
}
