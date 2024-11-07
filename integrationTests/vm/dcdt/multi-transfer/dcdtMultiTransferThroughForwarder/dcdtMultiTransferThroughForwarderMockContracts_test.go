package dcdtMultiTransferThroughForwarder

import (
	"testing"

	"github.com/kalyan3104/k-chain-go/integrationTests"
	"github.com/kalyan3104/k-chain-go/integrationTests/vm/dcdt"
	wasmvm "github.com/kalyan3104/k-chain-go/integrationTests/vm/wasm/wasmvm"
	vmcommon "github.com/kalyan3104/k-chain-vm-common-go"
	test "github.com/kalyan3104/k-chain-vm-go/testcommon"
)

func TestDCDTMultiTransferThroughForwarder_LegacyAsync_MockContracts(t *testing.T) {
	DCDTMultiTransferThroughForwarder_MockContracts(t, true)
}

func TestDCDTMultiTransferThroughForwarder_NewAsync_MockContracts(t *testing.T) {
	DCDTMultiTransferThroughForwarder_MockContracts(t, false)
}

func DCDTMultiTransferThroughForwarder_MockContracts(t *testing.T, legacyAsync bool) {
	if testing.Short() {
		t.Skip("this is not a short test")
	}

	net, ownerShard1, ownerShard2, senderNode, forwarder, vaultShard1, vaultShard2 :=
		DCDTMultiTransferThroughForwarder_MockContracts_SetupNetwork(t)
	defer net.Close()

	DCDTMultiTransferThroughForwarder_MockContracts_Deploy(t,
		legacyAsync,
		net,
		ownerShard1,
		ownerShard2,
		forwarder,
		vaultShard1,
		vaultShard2)

	DCDTMultiTransferThroughForwarder_RunStepsAndAsserts(t,
		net,
		senderNode,
		ownerShard1,
		ownerShard2,
		forwarder,
		vaultShard1,
		vaultShard2,
	)
}

func DCDTMultiTransferThroughForwarder_MockContracts_SetupNetwork(t *testing.T) (*integrationTests.TestNetwork, *integrationTests.TestWalletAccount, *integrationTests.TestWalletAccount, *integrationTests.TestProcessorNode, []byte, []byte, []byte) {
	net := integrationTests.NewTestNetworkSized(t, 2, 1, 1)
	net.Start().Step()

	net.CreateUninitializedWallets(2)
	owner := net.CreateWalletOnShard(0, 0)
	owner2 := net.CreateWalletOnShard(1, 1)

	initialVal := uint64(1000000000)
	net.MintWalletsUint64(initialVal)

	node0shard0 := net.NodesSharded[0][0]
	node0shard1 := net.NodesSharded[1][0]

	forwarder, forwarderAccount := wasmvm.GetAddressForNewAccountOnWalletAndNode(t, net, owner, node0shard0)
	wasmvm.SetCodeMetadata(t, []byte{0, vmcommon.MetadataPayable}, node0shard0, forwarderAccount)

	vaultShard1, vaultShard1Account := wasmvm.GetAddressForNewAccountOnWalletAndNode(t, net, owner, node0shard0)
	wasmvm.SetCodeMetadata(t, []byte{0, vmcommon.MetadataPayable}, node0shard0, vaultShard1Account)

	vaultShard2, _ := wasmvm.GetAddressForNewAccountOnWalletAndNode(t, net, owner2, node0shard1)

	return net, owner, owner2, node0shard0, forwarder, vaultShard1, vaultShard2
}

func DCDTMultiTransferThroughForwarder_MockContracts_Deploy(t *testing.T, legacyAsync bool, net *integrationTests.TestNetwork, ownerShard1 *integrationTests.TestWalletAccount, ownerShard2 *integrationTests.TestWalletAccount, forwarder []byte, vaultShard1 []byte, vaultShard2 []byte) {
	testConfig := &test.TestConfig{
		IsLegacyAsync: legacyAsync,
		// used for new async
		SuccessCallback:    "callBack",
		ErrorCallback:      "callBack",
		GasProvidedToChild: 500_000,
		GasToLock:          300_000,
	}

	wasmvm.InitializeMockContractsWithVMContainer(
		t, net,
		net.NodesSharded[0][0].VMContainer,
		test.CreateMockContractOnShard(forwarder, 0).
			WithOwnerAddress(ownerShard1.Address).
			WithConfig(testConfig).
			WithMethods(
				dcdt.MultiTransferViaAsyncMock,
				dcdt.SyncMultiTransferMock,
				dcdt.MultiTransferExecuteMock,
				dcdt.EmptyCallbackMock),
		test.CreateMockContractOnShard(vaultShard1, 0).
			WithOwnerAddress(ownerShard1.Address).
			WithConfig(testConfig).
			WithMethods(dcdt.AcceptFundsEchoMock),
		test.CreateMockContractOnShard(vaultShard2, 1).
			WithOwnerAddress(ownerShard2.Address).
			WithConfig(testConfig).
			WithMethods(dcdt.AcceptMultiFundsEchoMock),
	)
}

func TestDCDTMultiTransferWithWrongArgumentsSFT_MockContracts(t *testing.T) {
	if testing.Short() {
		t.Skip("this is not a short test")
	}

	net, ownerShard1, ownerShard2, senderNode, forwarder, _, vaultShard2 :=
		DCDTMultiTransferThroughForwarder_MockContracts_SetupNetwork(t)
	defer net.Close()

	DCDTMultiTransferWithWrongArguments_MockContracts_Deploy(t, net, ownerShard1, forwarder, vaultShard2, ownerShard2)

	DCDTMultiTransferWithWrongArgumentsSFT_RunStepsAndAsserts(t, net, senderNode, ownerShard1, forwarder, vaultShard2)
}

func TestDCDTMultiTransferWithWrongArgumentsFungible_MockContracts(t *testing.T) {
	if testing.Short() {
		t.Skip("this is not a short test")
	}

	net, ownerShard1, ownerShard2, senderNode, forwarder, _, vaultShard2 :=
		DCDTMultiTransferThroughForwarder_MockContracts_SetupNetwork(t)
	defer net.Close()

	DCDTMultiTransferWithWrongArguments_MockContracts_Deploy(t, net, ownerShard1, forwarder, vaultShard2, ownerShard2)

	DCDTMultiTransferWithWrongArgumentsFungible_RunStepsAndAsserts(t, net, senderNode, ownerShard1, forwarder, vaultShard2)
}

func DCDTMultiTransferWithWrongArguments_MockContracts_Deploy(t *testing.T, net *integrationTests.TestNetwork, ownerShard1 *integrationTests.TestWalletAccount, forwarder []byte, vaultShard2 []byte, ownerShard2 *integrationTests.TestWalletAccount) {
	testConfig := &test.TestConfig{
		IsLegacyAsync: true,
	}

	wasmvm.InitializeMockContractsWithVMContainer(
		t, net,
		net.NodesSharded[0][0].VMContainer,
		test.CreateMockContractOnShard(forwarder, 0).
			WithOwnerAddress(ownerShard1.Address).
			WithConfig(testConfig).
			WithMethods(dcdt.DoAsyncCallMock),
		test.CreateMockContractOnShard(vaultShard2, 1).
			WithOwnerAddress(ownerShard2.Address).
			WithConfig(testConfig).
			WithMethods(dcdt.AcceptFundsEchoMock),
	)
}
