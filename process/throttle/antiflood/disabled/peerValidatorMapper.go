package disabled

import "github.com/kalyan3104/k-chain-core-go/core"

// PeerValidatorMapper is a disabled validator mapper
type PeerValidatorMapper struct {
}

// GetPeerInfo returns validator type for all peers
func (p *PeerValidatorMapper) GetPeerInfo(_ core.PeerID) core.P2PPeerInfo {
	return core.P2PPeerInfo{PeerType: core.ValidatorPeer}
}

// IsInterfaceNil returns true if underlying object is nil
func (p *PeerValidatorMapper) IsInterfaceNil() bool {
	return p == nil
}
