package factory

import (
	"github.com/kalyan3104/k-chain-go/config"
	"github.com/kalyan3104/k-chain-go/storage"
	"github.com/kalyan3104/k-chain-go/storage/databaseremover"
	"github.com/kalyan3104/k-chain-go/storage/databaseremover/disabled"
)

// CreateCustomDatabaseRemover will handle the creation of a custom database remover based on the configuration
func CreateCustomDatabaseRemover(storagePruningConfig config.StoragePruningConfig) (storage.CustomDatabaseRemoverHandler, error) {
	if storagePruningConfig.AccountsTrieCleanOldEpochsData {
		return databaseremover.NewCustomDatabaseRemover(storagePruningConfig)
	}

	return disabled.NewDisabledCustomDatabaseRemover(), nil
}
