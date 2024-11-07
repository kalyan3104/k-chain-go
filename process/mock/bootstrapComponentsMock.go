package mock

import (
	"github.com/kalyan3104/k-chain-go/cmd/node/factory"
	"github.com/kalyan3104/k-chain-go/sharding"
)

// BootstrapComponentsMock -
type BootstrapComponentsMock struct {
	Coordinator          sharding.Coordinator
	HdrIntegrityVerifier factory.HeaderIntegrityVerifierHandler
	VersionedHdrFactory  factory.VersionedHeaderFactory
}

// ShardCoordinator -
func (bcm *BootstrapComponentsMock) ShardCoordinator() sharding.Coordinator {
	return bcm.Coordinator
}

// HeaderIntegrityVerifier -
func (bcm *BootstrapComponentsMock) HeaderIntegrityVerifier() factory.HeaderIntegrityVerifierHandler {
	return bcm.HdrIntegrityVerifier
}

// VersionedHeaderFactory -
func (bcm *BootstrapComponentsMock) VersionedHeaderFactory() factory.VersionedHeaderFactory {
	return bcm.VersionedHdrFactory
}

// IsInterfaceNil -
func (bcm *BootstrapComponentsMock) IsInterfaceNil() bool {
	return bcm == nil
}
