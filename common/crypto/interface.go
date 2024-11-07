package crypto

import crypto "github.com/kalyan3104/k-chain-crypto-go"

// MultiSignerContainer defines the container for different versioned multiSigner instances
type MultiSignerContainer interface {
	GetMultiSigner(epoch uint32) (crypto.MultiSigner, error)
	IsInterfaceNil() bool
}
