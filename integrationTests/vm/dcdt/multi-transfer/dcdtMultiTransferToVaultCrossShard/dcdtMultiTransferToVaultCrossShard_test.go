package dcdtMultiTransferToVaultCrossShard

import (
	"testing"

	multitransfer "github.com/kalyan3104/k-chain-go/integrationTests/vm/dcdt/multi-transfer"
)

func TestDCDTMultiTransferToVaultCrossShard(t *testing.T) {
	multitransfer.DcdtMultiTransferToVault(t, true, "../../testdata/vaultV2.wasm")
}
