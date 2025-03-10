package disabled

import (
	"github.com/kalyan3104/k-chain-go/storage"
	"github.com/kalyan3104/k-chain-go/storage/database"
	"github.com/kalyan3104/k-chain-go/storage/storageunit"
)

const defaultCapacity = 10000
const defaultNumShards = 1
const zeroSize = 0

// CreateMemUnit creates an in-memory storer unit using maps
func CreateMemUnit() storage.Storer {
	cache, err := storageunit.NewCache(storageunit.CacheConfig{Type: storageunit.LRUCache, Capacity: defaultCapacity, Shards: defaultNumShards, SizeInBytes: zeroSize})
	if err != nil {
		return nil
	}

	unit, err := storageunit.NewStorageUnit(cache, database.NewMemDB())
	if err != nil {
		return nil
	}

	return unit
}
