package factory

import (
	"github.com/kalyan3104/k-chain-core-go/core"
	"github.com/kalyan3104/k-chain-core-go/core/check"
	"github.com/kalyan3104/k-chain-go/errors"
)

// DecodeAddresses will decode the provided string addresses
func DecodeAddresses(pkConverter core.PubkeyConverter, stringAddresses []string) ([][]byte, error) {
	if check.IfNil(pkConverter) {
		return nil, errors.ErrNilPubKeyConverter
	}
	decodedAddresses := make([][]byte, len(stringAddresses))
	for i, stringAddress := range stringAddresses {
		decodedAddress, errDecode := pkConverter.Decode(stringAddress)
		if errDecode != nil {
			return nil, errDecode
		}
		decodedAddresses[i] = decodedAddress
	}
	return decodedAddresses, nil
}
