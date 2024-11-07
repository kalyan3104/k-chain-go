package mock

import (
	"github.com/kalyan3104/k-chain-go/sharding/nodesCoordinator"
)

// NodesCoordinatorStub -
type NodesCoordinatorStub struct {
	GetValidatorWithPublicKeyCalled func(publicKey []byte) (validator nodesCoordinator.Validator, shardId uint32, err error)
}

// GetValidatorWithPublicKey -
func (nc *NodesCoordinatorStub) GetValidatorWithPublicKey(publicKey []byte) (validator nodesCoordinator.Validator, shardId uint32, err error) {
	if nc.GetValidatorWithPublicKeyCalled != nil {
		return nc.GetValidatorWithPublicKeyCalled(publicKey)
	}
	return nil, 0, nil
}

// IsInterfaceNil -
func (nc *NodesCoordinatorStub) IsInterfaceNil() bool {
	return false
}
