package dcdtMultiTransferToVaultSameShard

import (
	"testing"

	multitransfer "github.com/kalyan3104/k-chain-go/integrationTests/vm/dcdt/multi-transfer"
)

func TestDCDTMultiTransferToVaultSameShard(t *testing.T) {
	multitransfer.DcdtMultiTransferToVault(t, false, "../../testdata/vaultV2.wasm")
}
