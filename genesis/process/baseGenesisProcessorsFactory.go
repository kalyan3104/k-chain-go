package process

import (
	"github.com/kalyan3104/k-chain-go/node/external"
	"github.com/kalyan3104/k-chain-go/process"
	"github.com/kalyan3104/k-chain-go/vm"
)

type genesisProcessors struct {
	txCoordinator       process.TransactionCoordinator
	systemSCs           vm.SystemSCContainer
	txProcessor         process.TransactionProcessor
	scProcessor         process.SmartContractProcessor
	scrProcessor        process.SmartContractResultProcessor
	rwdProcessor        process.RewardTransactionProcessor
	blockchainHook      process.BlockChainHookHandler
	queryService        external.SCQueryService
	vmContainersFactory process.VirtualMachinesContainerFactory
	vmContainer         process.VirtualMachinesContainer
}
