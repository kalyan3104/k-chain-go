package enablers

import (
	"github.com/kalyan3104/k-chain-core-go/core/atomic"
)

type roundFlag struct {
	*atomic.Flag
	round   uint64
	options []string
}
