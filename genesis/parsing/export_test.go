package parsing

import (
	"math/big"

	"github.com/kalyan3104/k-chain-core-go/core"
	coreData "github.com/kalyan3104/k-chain-core-go/data"
	"github.com/kalyan3104/k-chain-core-go/data/block"
	"github.com/kalyan3104/k-chain-core-go/data/outport"
	transactionData "github.com/kalyan3104/k-chain-core-go/data/transaction"
	"github.com/kalyan3104/k-chain-go/genesis"
	"github.com/kalyan3104/k-chain-go/genesis/data"
	"github.com/kalyan3104/k-chain-go/genesis/mock"
	"github.com/kalyan3104/k-chain-go/sharding"
	"github.com/kalyan3104/k-chain-go/testscommon/hashingMocks"
	crypto "github.com/kalyan3104/k-chain-crypto-go"
)

func (ap *accountsParser) SetInitialAccounts(initialAccounts []*data.InitialAccount) {
	ap.initialAccounts = initialAccounts
}

func (ap *accountsParser) SetEntireSupply(entireSupply *big.Int) {
	ap.entireSupply = entireSupply
}

func (ap *accountsParser) Process() error {
	return ap.process()
}

func (ap *accountsParser) SetPukeyConverter(pubkeyConverter core.PubkeyConverter) {
	ap.pubkeyConverter = pubkeyConverter
}

func (ap *accountsParser) SetKeyGenerator(keyGen crypto.KeyGenerator) {
	ap.keyGenerator = keyGen
}

func (ap *accountsParser) CreateMintTransactions() []coreData.TransactionHandler {
	return ap.createMintTransactions()
}

func (ap *accountsParser) SetScrsTxsPool(
	shardCoordinator sharding.Coordinator,
	indexingData map[uint32]*genesis.IndexingData,
	txsPoolPerShard map[uint32]*outport.TransactionPool,
) {
	ap.setScrsTxsPool(shardCoordinator, indexingData, txsPoolPerShard)
}

func (ap *accountsParser) CreateMintTransaction(ia genesis.InitialAccountHandler, nonce uint64) *transactionData.Transaction {
	return ap.createMintTransaction(ia, nonce)
}

func NewTestAccountsParser(pubkeyConverter core.PubkeyConverter) *accountsParser {
	addrBytes, _ := pubkeyConverter.Decode("moa17rc0pu8s7rc0pu8s7rc0pu8s7rc0pu8s7rc0pu8s7rc0pu8s7rcqdw3ycp")
	return &accountsParser{
		pubkeyConverter:    pubkeyConverter,
		initialAccounts:    make([]*data.InitialAccount, 0),
		minterAddressBytes: addrBytes,
		keyGenerator:       &mock.KeyGeneratorStub{},
		hasher:             &hashingMocks.HasherMock{},
		marshalizer:        &mock.MarshalizerMock{},
	}
}

func NewTestSmartContractsParser(pubkeyConverter core.PubkeyConverter) *smartContractParser {
	scp := &smartContractParser{
		pubkeyConverter:       pubkeyConverter,
		keyGenerator:          &mock.KeyGeneratorStub{},
		initialSmartContracts: make([]*data.InitialSmartContract, 0),
	}
	//mock implementation, assumes the files are present
	scp.checkForFileHandler = func(filename string) error {
		return nil
	}

	return scp
}

func (scp *smartContractParser) SetInitialSmartContracts(initialSmartContracts []*data.InitialSmartContract) {
	scp.initialSmartContracts = initialSmartContracts
}

func (scp *smartContractParser) Process() error {
	return scp.process()
}

func (scp *smartContractParser) SetFileHandler(handler func(string) error) {
	scp.checkForFileHandler = handler
}

func (scp *smartContractParser) SetKeyGenerator(keyGen crypto.KeyGenerator) {
	scp.keyGenerator = keyGen
}

func CreateMiniBlocks(shardIDs []uint32, blockType block.Type) []*block.MiniBlock {
	return createMiniBlocks(shardIDs, blockType)
}

func GetShardIDs(shardCoordinator sharding.Coordinator) []uint32 {
	return getShardIDs(shardCoordinator)
}
