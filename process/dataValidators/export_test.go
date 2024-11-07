package dataValidators

import (
	"github.com/kalyan3104/k-chain-go/process"
	vmcommon "github.com/kalyan3104/k-chain-vm-common-go"
)

// CheckAccount -
func (txv *txValidator) CheckAccount(
	interceptedTx process.InterceptedTransactionHandler,
	accountHandler vmcommon.AccountHandler,
) error {
	return txv.checkAccount(interceptedTx, accountHandler)
}

// GetTxData -
func GetTxData(interceptedTx process.InterceptedTransactionHandler) ([]byte, error) {
	return getTxData(interceptedTx)
}
