package txpool

import (
	"github.com/kalyan3104/k-chain-go/storage"
	"github.com/kalyan3104/k-chain-go/storage/txcache"
)

type txCache interface {
	storage.Cacher

	AddTx(tx *txcache.WrappedTransaction) (ok bool, added bool)
	GetByTxHash(txHash []byte) (*txcache.WrappedTransaction, bool)
	RemoveTxByHash(txHash []byte) bool
	ImmunizeTxsAgainstEviction(keys [][]byte)
	ForEachTransaction(function txcache.ForEachTransaction)
	NumBytes() int
	Diagnose(deep bool)
	GetTransactionsPoolForSender(sender string) []*txcache.WrappedTransaction
}
