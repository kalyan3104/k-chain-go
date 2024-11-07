package processor

import (
	"github.com/kalyan3104/k-chain-go/state"
)

type interceptedDataSizeHandler interface {
	SizeInBytes() int
}

type interceptedHeartbeatMessageHandler interface {
	interceptedDataSizeHandler
	Message() interface{}
}

type interceptedPeerAuthenticationMessageHandler interface {
	interceptedDataSizeHandler
	Message() interface{}
	Payload() []byte
	Pubkey() []byte
}

type interceptedValidatorInfo interface {
	Hash() []byte
	ValidatorInfo() *state.ShardValidatorInfo
}
