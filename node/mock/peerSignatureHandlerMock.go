package mock

import (
	"github.com/kalyan3104/k-chain-core-go/core"
	"github.com/kalyan3104/k-chain-crypto-go"
)

// PeerSignatureHandler -
type PeerSignatureHandler struct{}

// VerifyPeerSignature -
func (p *PeerSignatureHandler) VerifyPeerSignature(_ []byte, _ core.PeerID, _ []byte) error {
	return nil
}

// GetPeerSignature -
func (p *PeerSignatureHandler) GetPeerSignature(_ crypto.PrivateKey, _ []byte) ([]byte, error) {
	return nil, nil
}

// IsInterfaceNil -
func (p *PeerSignatureHandler) IsInterfaceNil() bool {
	return p == nil
}
