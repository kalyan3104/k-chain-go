package factory

import (
	"github.com/kalyan3104/k-chain-core-go/core"
	"github.com/kalyan3104/k-chain-go/process"
	"github.com/kalyan3104/k-chain-go/process/block/bootstrapStorage"
	"github.com/kalyan3104/k-chain-go/storage"
)

// BootstrapDataProviderHandler defines which actions should be done for loading bootstrap data from the boot storer
type BootstrapDataProviderHandler interface {
	LoadForPath(persisterFactory storage.PersisterFactory, path string) (*bootstrapStorage.BootstrapData, storage.Storer, error)
	GetStorer(storer storage.Storer) (process.BootStorer, error)
	IsInterfaceNil() bool
}

// NodeTypeProviderHandler defines the actions needed for a component that can handle the node type
type NodeTypeProviderHandler interface {
	SetType(nodeType core.NodeType)
	GetType() core.NodeType
	IsInterfaceNil() bool
}
