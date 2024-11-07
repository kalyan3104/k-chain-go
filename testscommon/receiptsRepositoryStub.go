package testscommon

import (
	"github.com/kalyan3104/k-chain-core-go/data"
	"github.com/kalyan3104/k-chain-go/common"
	"github.com/kalyan3104/k-chain-go/common/holders"
)

// ReceiptsRepositoryStub -
type ReceiptsRepositoryStub struct {
	SaveReceiptsCalled func(holder common.ReceiptsHolder, header data.HeaderHandler, headerHash []byte) error
	LoadReceiptsCalled func(header data.HeaderHandler, headerHash []byte) (common.ReceiptsHolder, error)
}

// SaveReceipts -
func (stub *ReceiptsRepositoryStub) SaveReceipts(holder common.ReceiptsHolder, header data.HeaderHandler, headerHash []byte) error {
	if stub.SaveReceiptsCalled != nil {
		return stub.SaveReceiptsCalled(holder, header, headerHash)
	}

	return nil
}

// LoadReceipts -
func (stub *ReceiptsRepositoryStub) LoadReceipts(header data.HeaderHandler, headerHash []byte) (common.ReceiptsHolder, error) {
	if stub.LoadReceiptsCalled != nil {
		return stub.LoadReceiptsCalled(header, headerHash)
	}

	return holders.NewReceiptsHolder(nil), nil
}

// IsInterfaceNil -
func (stub *ReceiptsRepositoryStub) IsInterfaceNil() bool {
	return stub == nil
}
