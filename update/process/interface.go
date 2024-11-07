package process

import (
	"github.com/kalyan3104/k-chain-core-go/data"
	"github.com/kalyan3104/k-chain-go/common"
)

type receiptsRepository interface {
	SaveReceipts(holder common.ReceiptsHolder, header data.HeaderHandler, headerHash []byte) error
	IsInterfaceNil() bool
}
