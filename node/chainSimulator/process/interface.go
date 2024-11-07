package process

import (
	chainData "github.com/kalyan3104/k-chain-core-go/data"
	"github.com/kalyan3104/k-chain-go/api/shared"
	"github.com/kalyan3104/k-chain-go/consensus"
	"github.com/kalyan3104/k-chain-go/factory"
	"github.com/kalyan3104/k-chain-go/node/chainSimulator/dtos"
	"github.com/kalyan3104/k-chain-go/sharding"
)

// NodeHandler defines what a node handler should be able to do
type NodeHandler interface {
	GetProcessComponents() factory.ProcessComponentsHolder
	GetChainHandler() chainData.ChainHandler
	GetBroadcastMessenger() consensus.BroadcastMessenger
	GetShardCoordinator() sharding.Coordinator
	GetCryptoComponents() factory.CryptoComponentsHolder
	GetCoreComponents() factory.CoreComponentsHolder
	GetDataComponents() factory.DataComponentsHolder
	GetStateComponents() factory.StateComponentsHolder
	GetFacadeHandler() shared.FacadeHandler
	GetStatusCoreComponents() factory.StatusCoreComponentsHolder
	SetKeyValueForAddress(addressBytes []byte, state map[string]string) error
	SetStateForAddress(address []byte, state *dtos.AddressState) error
	RemoveAccount(address []byte) error
	Close() error
	IsInterfaceNil() bool
}
