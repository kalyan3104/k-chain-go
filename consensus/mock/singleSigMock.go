package mock

import (
	crypto "github.com/kalyan3104/k-chain-crypto-go"
)

// SingleSignerMock -
type SingleSignerMock struct {
	SignStub   func(private crypto.PrivateKey, msg []byte) ([]byte, error)
	VerifyStub func(public crypto.PublicKey, msg []byte, sig []byte) error
}

// Sign -
func (s *SingleSignerMock) Sign(private crypto.PrivateKey, msg []byte) ([]byte, error) {
	if s.SignStub != nil {
		return s.SignStub(private, msg)
	}

	return nil, nil
}

// Verify -
func (s *SingleSignerMock) Verify(public crypto.PublicKey, msg []byte, sig []byte) error {
	if s.VerifyStub != nil {
		return s.VerifyStub(public, msg, sig)
	}

	return nil
}

// IsInterfaceNil returns true if there is no value under the interface
func (s *SingleSignerMock) IsInterfaceNil() bool {
	return s == nil
}
