package resolvers

import (
	"github.com/kalyan3104/k-chain-go/storage"
)

type baseSimpleResolver struct {
	storer storage.Storer
}

func (bsr *baseSimpleResolver) getFromStorage(key []byte, _ uint32) ([]byte, error) {
	return bsr.storer.SearchFirst(key)
}

func (bsr *baseSimpleResolver) searchFirst(key []byte) ([]byte, error) {
	return bsr.storer.SearchFirst(key)
}
