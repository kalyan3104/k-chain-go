package interceptedBlocks

import (
	"github.com/kalyan3104/k-chain-core-go/hashing"
	"github.com/kalyan3104/k-chain-core-go/marshal"
	"github.com/kalyan3104/k-chain-go/process"
	"github.com/kalyan3104/k-chain-go/sharding"
)

// ArgInterceptedBlockHeader is the argument for the intercepted header
type ArgInterceptedBlockHeader struct {
	HdrBuff                 []byte
	Marshalizer             marshal.Marshalizer
	Hasher                  hashing.Hasher
	ShardCoordinator        sharding.Coordinator
	HeaderSigVerifier       process.InterceptedHeaderSigVerifier
	HeaderIntegrityVerifier process.HeaderIntegrityVerifier
	ValidityAttester        process.ValidityAttester
	EpochStartTrigger       process.EpochStartTriggerHandler
}
