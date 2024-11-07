package process

import (
	"math/big"

	"github.com/kalyan3104/k-chain-core-go/core"
	"github.com/kalyan3104/k-chain-core-go/data"
	"github.com/kalyan3104/k-chain-core-go/data/typeConverters"
	"github.com/kalyan3104/k-chain-core-go/hashing"
	"github.com/kalyan3104/k-chain-core-go/marshal"
	"github.com/kalyan3104/k-chain-go/common"
	"github.com/kalyan3104/k-chain-go/config"
	"github.com/kalyan3104/k-chain-go/dataRetriever"
	"github.com/kalyan3104/k-chain-go/dblookupext"
	"github.com/kalyan3104/k-chain-go/genesis"
	"github.com/kalyan3104/k-chain-go/process"
	"github.com/kalyan3104/k-chain-go/sharding"
	"github.com/kalyan3104/k-chain-go/state"
	"github.com/kalyan3104/k-chain-go/update"
	crypto "github.com/kalyan3104/k-chain-crypto-go"
)

type coreComponentsHandler interface {
	InternalMarshalizer() marshal.Marshalizer
	TxMarshalizer() marshal.Marshalizer
	Hasher() hashing.Hasher
	AddressPubKeyConverter() core.PubkeyConverter
	Uint64ByteSliceConverter() typeConverters.Uint64ByteSliceConverter
	TxVersionChecker() process.TxVersionCheckerHandler
	ChainID() string
	EnableEpochsHandler() common.EnableEpochsHandler
	IsInterfaceNil() bool
}

type dataComponentsHandler interface {
	StorageService() dataRetriever.StorageService
	Blockchain() data.ChainHandler
	Datapool() dataRetriever.PoolsHolder
	SetBlockchain(chain data.ChainHandler) error
	Clone() interface{}
	IsInterfaceNil() bool
}

// ArgsGenesisBlockCreator holds the arguments which are needed to create a genesis block
type ArgsGenesisBlockCreator struct {
	GenesisTime             uint64
	GenesisNonce            uint64
	GenesisRound            uint64
	StartEpochNum           uint32
	GenesisEpoch            uint32
	Data                    dataComponentsHandler
	Core                    coreComponentsHandler
	Accounts                state.AccountsAdapter
	ValidatorAccounts       state.AccountsAdapter
	InitialNodesSetup       genesis.InitialNodesHandler
	Economics               process.EconomicsDataHandler
	ShardCoordinator        sharding.Coordinator
	AccountsParser          genesis.AccountsParser
	SmartContractParser     genesis.InitialSmartContractParser
	GasSchedule             core.GasScheduleNotifier
	TxLogsProcessor         process.TransactionLogProcessor
	VirtualMachineConfig    config.VirtualMachineConfig
	HardForkConfig          config.HardforkConfig
	TrieStorageManagers     map[string]common.StorageManager
	SystemSCConfig          config.SystemSmartContractsConfig
	RoundConfig             config.RoundConfig
	EpochConfig             config.EpochConfig
	HeaderVersionConfigs    config.VersionsConfig
	WorkingDir              string
	BlockSignKeyGen         crypto.KeyGenerator
	HistoryRepository       dblookupext.HistoryRepository
	TxExecutionOrderHandler common.TxExecutionOrderHandler

	GenesisNodePrice *big.Int
	GenesisString    string

	// created components
	importHandler          update.ImportHandler
	versionedHeaderFactory genesis.VersionedHeaderFactory
}
