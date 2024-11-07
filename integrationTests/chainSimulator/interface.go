package chainSimulator

import (
	"math/big"

	"github.com/kalyan3104/k-chain-core-go/data/api"
	"github.com/kalyan3104/k-chain-core-go/data/transaction"
	"github.com/kalyan3104/k-chain-go/node/chainSimulator/dtos"
	"github.com/kalyan3104/k-chain-go/node/chainSimulator/process"
	crypto "github.com/kalyan3104/k-chain-crypto-go"
)

// ChainSimulator defines the operations for an entity that can simulate operations of a chain
type ChainSimulator interface {
	GenerateBlocks(numOfBlocks int) error
	GenerateBlocksUntilEpochIsReached(targetEpoch int32) error
	AddValidatorKeys(validatorsPrivateKeys [][]byte) error
	GetNodeHandler(shardID uint32) process.NodeHandler
	SendTxAndGenerateBlockTilTxIsExecuted(txToSend *transaction.Transaction, maxNumOfBlockToGenerateWhenExecutingTx int) (*transaction.ApiTransactionResult, error)
	SendTxsAndGenerateBlocksTilAreExecuted(txsToSend []*transaction.Transaction, maxNumOfBlocksToGenerateWhenExecutingTx int) ([]*transaction.ApiTransactionResult, error)
	SetStateMultiple(stateSlice []*dtos.AddressState) error
	GenerateAndMintWalletAddress(targetShardID uint32, value *big.Int) (dtos.WalletAddress, error)
	GetInitialWalletKeys() *dtos.InitialWalletKeys
	GetAccount(address dtos.WalletAddress) (api.AccountResponse, error)
	ForceResetValidatorStatisticsCache() error
	GetValidatorPrivateKeys() []crypto.PrivateKey
}
