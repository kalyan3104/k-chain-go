package factory

import (
	"github.com/kalyan3104/k-chain-core-go/hashing"
	"github.com/kalyan3104/k-chain-core-go/marshal"
	"github.com/kalyan3104/k-chain-go/common"
	"github.com/kalyan3104/k-chain-go/storage"
)

type coreComponentsHandler interface {
	InternalMarshalizer() marshal.Marshalizer
	Hasher() hashing.Hasher
	PathHandler() storage.PathManagerHandler
	ProcessStatusHandler() common.ProcessStatusHandler
	EnableEpochsHandler() common.EnableEpochsHandler
}
