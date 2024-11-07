package blockInfoProviders

import (
	"github.com/kalyan3104/k-chain-core-go/core/check"
	chainData "github.com/kalyan3104/k-chain-core-go/data"
	"github.com/kalyan3104/k-chain-go/common"
	"github.com/kalyan3104/k-chain-go/common/holders"
)

type finalBlockInfo struct {
	chainHandler chainData.ChainHandler
}

// NewFinalBlockInfo creates a new instance of type finalBlockInfo
func NewFinalBlockInfo(chainHandler chainData.ChainHandler) (*finalBlockInfo, error) {
	if check.IfNil(chainHandler) {
		return nil, ErrNilChainHandler
	}

	return &finalBlockInfo{
		chainHandler: chainHandler,
	}, nil
}

// GetBlockInfo returns the current block info
func (provider *finalBlockInfo) GetBlockInfo() common.BlockInfo {
	nonce, hash, rootHash := provider.chainHandler.GetFinalBlockInfo()

	return holders.NewBlockInfo(hash, nonce, rootHash)
}

// IsInterfaceNil returns true if there is no value under the interface
func (provider *finalBlockInfo) IsInterfaceNil() bool {
	return provider == nil
}
