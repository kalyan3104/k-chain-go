package helpers

import (
	"github.com/kalyan3104/k-chain-core-go/data"
	"github.com/kalyan3104/k-chain-go/common"
)

// ComputeRandomnessForTxSorting returns the randomness for transactions sorting
func ComputeRandomnessForTxSorting(header data.HeaderHandler, enableEpochsHandler common.EnableEpochsHandler) []byte {
	if enableEpochsHandler.IsFlagEnabled(common.CurrentRandomnessOnSortingFlag) {
		return header.GetRandSeed()
	}

	return header.GetPrevRandSeed()
}
