package logs

import (
	"github.com/kalyan3104/k-chain-core-go/core"
	"github.com/kalyan3104/k-chain-core-go/core/check"
	"github.com/kalyan3104/k-chain-core-go/marshal"
	"github.com/kalyan3104/k-chain-go/dataRetriever"
)

// ArgsNewLogsFacade holds the arguments for constructing a logsFacade
type ArgsNewLogsFacade struct {
	StorageService  dataRetriever.StorageService
	Marshaller      marshal.Marshalizer
	PubKeyConverter core.PubkeyConverter
}

func (args *ArgsNewLogsFacade) check() error {
	if check.IfNil(args.StorageService) {
		return core.ErrNilStore
	}
	if check.IfNil(args.Marshaller) {
		return core.ErrNilMarshalizer
	}
	if check.IfNil(args.PubKeyConverter) {
		return core.ErrNilPubkeyConverter
	}

	return nil
}
