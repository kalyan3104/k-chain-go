package mock

import (
	"context"
	"encoding/hex"
	"math/big"

	"github.com/kalyan3104/k-chain-core-go/core"
	"github.com/kalyan3104/k-chain-core-go/data/api"
	"github.com/kalyan3104/k-chain-core-go/data/dcdt"
	"github.com/kalyan3104/k-chain-core-go/data/transaction"
	"github.com/kalyan3104/k-chain-core-go/data/validator"
	"github.com/kalyan3104/k-chain-go/common"
	"github.com/kalyan3104/k-chain-go/debug"
	"github.com/kalyan3104/k-chain-go/heartbeat/data"
	"github.com/kalyan3104/k-chain-go/node/external"
)

// NodeStub -
type NodeStub struct {
	ConnectToAddressesHandler                      func([]string) error
	GetBalanceCalled                               func(address string, options api.AccountQueryOptions) (*big.Int, api.BlockInfo, error)
	GenerateTransactionHandler                     func(sender string, receiver string, amount string, code string) (*transaction.Transaction, error)
	CreateTransactionHandler                       func(txArgs *external.ArgsCreateTransaction) (*transaction.Transaction, []byte, error)
	ValidateTransactionHandler                     func(tx *transaction.Transaction) error
	ValidateTransactionForSimulationCalled         func(tx *transaction.Transaction, bypassSignature bool) error
	SendBulkTransactionsHandler                    func(txs []*transaction.Transaction) (uint64, error)
	GetAccountCalled                               func(address string, options api.AccountQueryOptions) (api.AccountResponse, api.BlockInfo, error)
	GetAccountWithKeysCalled                       func(address string, options api.AccountQueryOptions, ctx context.Context) (api.AccountResponse, api.BlockInfo, error)
	GetCodeCalled                                  func(codeHash []byte, options api.AccountQueryOptions) ([]byte, api.BlockInfo)
	GetCurrentPublicKeyHandler                     func() string
	GenerateAndSendBulkTransactionsHandler         func(destination string, value *big.Int, nrTransactions uint64) error
	GenerateAndSendBulkTransactionsOneByOneHandler func(destination string, value *big.Int, nrTransactions uint64) error
	GetHeartbeatsHandler                           func() []data.PubKeyHeartbeat
	ValidatorStatisticsApiCalled                   func() (map[string]*validator.ValidatorStatistics, error)
	DirectTriggerCalled                            func(epoch uint32, withEarlyEndOfEpoch bool) error
	IsSelfTriggerCalled                            func() bool
	GetQueryHandlerCalled                          func(name string) (debug.QueryHandler, error)
	GetValueForKeyCalled                           func(address string, key string, options api.AccountQueryOptions) (string, api.BlockInfo, error)
	GetGuardianDataCalled                          func(address string, options api.AccountQueryOptions) (api.GuardianData, api.BlockInfo, error)
	GetPeerInfoCalled                              func(pid string) ([]core.QueryP2PPeerInfo, error)
	GetConnectedPeersRatingsOnMainNetworkCalled    func() (string, error)
	GetEpochStartDataAPICalled                     func(epoch uint32) (*common.EpochStartDataAPI, error)
	GetUsernameCalled                              func(address string, options api.AccountQueryOptions) (string, api.BlockInfo, error)
	GetCodeHashCalled                              func(address string, options api.AccountQueryOptions) ([]byte, api.BlockInfo, error)
	GetDCDTDataCalled                              func(address string, key string, nonce uint64, options api.AccountQueryOptions) (*dcdt.DCDigitalToken, api.BlockInfo, error)
	GetAllDCDTTokensCalled                         func(address string, options api.AccountQueryOptions, ctx context.Context) (map[string]*dcdt.DCDigitalToken, api.BlockInfo, error)
	GetNFTTokenIDsRegisteredByAddressCalled        func(address string, options api.AccountQueryOptions, ctx context.Context) ([]string, api.BlockInfo, error)
	GetDCDTsWithRoleCalled                         func(address string, role string, options api.AccountQueryOptions, ctx context.Context) ([]string, api.BlockInfo, error)
	GetDCDTsRolesCalled                            func(address string, options api.AccountQueryOptions, ctx context.Context) (map[string][]string, api.BlockInfo, error)
	GetKeyValuePairsCalled                         func(address string, options api.AccountQueryOptions, ctx context.Context) (map[string]string, api.BlockInfo, error)
	GetAllIssuedDCDTsCalled                        func(tokenType string, ctx context.Context) ([]string, error)
	GetProofCalled                                 func(rootHash string, key string) (*common.GetProofResponse, error)
	GetProofDataTrieCalled                         func(rootHash string, address string, key string) (*common.GetProofResponse, *common.GetProofResponse, error)
	VerifyProofCalled                              func(rootHash string, address string, proof [][]byte) (bool, error)
	GetTokenSupplyCalled                           func(token string) (*api.DCDTSupply, error)
	IsDataTrieMigratedCalled                       func(address string, options api.AccountQueryOptions) (bool, error)
	AuctionListApiCalled                           func() ([]*common.AuctionListValidatorAPIResponse, error)
}

// GetProof -
func (ns *NodeStub) GetProof(rootHash string, key string) (*common.GetProofResponse, error) {
	if ns.GetProofCalled != nil {
		return ns.GetProofCalled(rootHash, key)
	}

	return nil, nil
}

// GetProofDataTrie -
func (ns *NodeStub) GetProofDataTrie(rootHash string, address string, key string) (*common.GetProofResponse, *common.GetProofResponse, error) {
	if ns.GetProofDataTrieCalled != nil {
		return ns.GetProofDataTrieCalled(rootHash, address, key)
	}

	return nil, nil, nil
}

// VerifyProof -
func (ns *NodeStub) VerifyProof(rootHash string, address string, proof [][]byte) (bool, error) {
	if ns.VerifyProofCalled != nil {
		return ns.VerifyProofCalled(rootHash, address, proof)
	}

	return false, nil
}

// GetUsername -
func (ns *NodeStub) GetUsername(address string, options api.AccountQueryOptions) (string, api.BlockInfo, error) {
	if ns.GetUsernameCalled != nil {
		return ns.GetUsernameCalled(address, options)
	}

	return "", api.BlockInfo{}, nil
}

// GetCodeHash -
func (ns *NodeStub) GetCodeHash(address string, options api.AccountQueryOptions) ([]byte, api.BlockInfo, error) {
	if ns.GetCodeHashCalled != nil {
		return ns.GetCodeHashCalled(address, options)
	}

	return nil, api.BlockInfo{}, nil
}

// GetKeyValuePairs -
func (ns *NodeStub) GetKeyValuePairs(address string, options api.AccountQueryOptions, ctx context.Context) (map[string]string, api.BlockInfo, error) {
	if ns.GetKeyValuePairsCalled != nil {
		return ns.GetKeyValuePairsCalled(address, options, ctx)
	}

	return nil, api.BlockInfo{}, nil
}

// GetValueForKey -
func (ns *NodeStub) GetValueForKey(address string, key string, options api.AccountQueryOptions) (string, api.BlockInfo, error) {
	if ns.GetValueForKeyCalled != nil {
		return ns.GetValueForKeyCalled(address, key, options)
	}

	return "", api.BlockInfo{}, nil
}

// GetGuardianData -
func (ns *NodeStub) GetGuardianData(address string, options api.AccountQueryOptions) (api.GuardianData, api.BlockInfo, error) {
	if ns.GetGuardianDataCalled != nil {
		return ns.GetGuardianDataCalled(address, options)
	}
	return api.GuardianData{}, api.BlockInfo{}, nil
}

// EncodeAddressPubkey -
func (ns *NodeStub) EncodeAddressPubkey(pk []byte) (string, error) {
	return hex.EncodeToString(pk), nil
}

// DecodeAddressPubkey -
func (ns *NodeStub) DecodeAddressPubkey(pk string) ([]byte, error) {
	return hex.DecodeString(pk)
}

// GetBalance -
func (ns *NodeStub) GetBalance(address string, options api.AccountQueryOptions) (*big.Int, api.BlockInfo, error) {
	if ns.GetBalanceCalled != nil {
		return ns.GetBalanceCalled(address, options)
	}

	return nil, api.BlockInfo{}, nil
}

// CreateTransaction -
func (ns *NodeStub) CreateTransaction(txArgs *external.ArgsCreateTransaction) (*transaction.Transaction, []byte, error) {

	return ns.CreateTransactionHandler(txArgs)
}

// ValidateTransaction -
func (ns *NodeStub) ValidateTransaction(tx *transaction.Transaction) error {
	if ns.ValidateTransactionHandler != nil {
		return ns.ValidateTransactionHandler(tx)
	}

	return nil
}

// ValidateTransactionForSimulation -
func (ns *NodeStub) ValidateTransactionForSimulation(tx *transaction.Transaction, bypassSignature bool) error {
	if ns.ValidateTransactionForSimulationCalled != nil {
		return ns.ValidateTransactionForSimulationCalled(tx, bypassSignature)
	}

	return nil
}

// SendBulkTransactions -
func (ns *NodeStub) SendBulkTransactions(txs []*transaction.Transaction) (uint64, error) {
	if ns.SendBulkTransactionsHandler != nil {
		return ns.SendBulkTransactionsHandler(txs)
	}

	return 0, nil
}

// GetAccount -
func (ns *NodeStub) GetAccount(address string, options api.AccountQueryOptions) (api.AccountResponse, api.BlockInfo, error) {
	if ns.GetAccountCalled != nil {
		return ns.GetAccountCalled(address, options)
	}

	return api.AccountResponse{}, api.BlockInfo{}, nil
}

// GetAccountWithKeys -
func (ns *NodeStub) GetAccountWithKeys(address string, options api.AccountQueryOptions, ctx context.Context) (api.AccountResponse, api.BlockInfo, error) {
	if ns.GetAccountWithKeysCalled != nil {
		return ns.GetAccountWithKeysCalled(address, options, ctx)
	}

	return api.AccountResponse{}, api.BlockInfo{}, nil
}

// GetCode -
func (ns *NodeStub) GetCode(codeHash []byte, options api.AccountQueryOptions) ([]byte, api.BlockInfo) {
	if ns.GetCodeCalled != nil {
		return ns.GetCodeCalled(codeHash, options)
	}

	return nil, api.BlockInfo{}
}

// GetHeartbeats -
func (ns *NodeStub) GetHeartbeats() []data.PubKeyHeartbeat {
	if ns.GetHeartbeatsHandler != nil {
		return ns.GetHeartbeatsHandler()
	}

	return nil
}

// ValidatorStatisticsApi -
func (ns *NodeStub) ValidatorStatisticsApi() (map[string]*validator.ValidatorStatistics, error) {
	if ns.ValidatorStatisticsApiCalled != nil {
		return ns.ValidatorStatisticsApiCalled()
	}

	return nil, nil
}

// AuctionListApi -
func (ns *NodeStub) AuctionListApi() ([]*common.AuctionListValidatorAPIResponse, error) {
	if ns.AuctionListApiCalled != nil {
		return ns.AuctionListApiCalled()
	}

	return nil, nil
}

// DirectTrigger -
func (ns *NodeStub) DirectTrigger(epoch uint32, withEarlyEndOfEpoch bool) error {
	if ns.DirectTriggerCalled != nil {
		return ns.DirectTriggerCalled(epoch, withEarlyEndOfEpoch)
	}

	return nil
}

// IsSelfTrigger -
func (ns *NodeStub) IsSelfTrigger() bool {
	if ns.IsSelfTriggerCalled != nil {
		return ns.IsSelfTriggerCalled()
	}

	return false
}

// GetQueryHandler -
func (ns *NodeStub) GetQueryHandler(name string) (debug.QueryHandler, error) {
	if ns.GetQueryHandlerCalled != nil {
		return ns.GetQueryHandlerCalled(name)
	}

	return nil, nil
}

// GetPeerInfo -
func (ns *NodeStub) GetPeerInfo(pid string) ([]core.QueryP2PPeerInfo, error) {
	if ns.GetPeerInfoCalled != nil {
		return ns.GetPeerInfoCalled(pid)
	}

	return make([]core.QueryP2PPeerInfo, 0), nil
}

// GetConnectedPeersRatingsOnMainNetwork -
func (ns *NodeStub) GetConnectedPeersRatingsOnMainNetwork() (string, error) {
	if ns.GetConnectedPeersRatingsOnMainNetworkCalled != nil {
		return ns.GetConnectedPeersRatingsOnMainNetworkCalled()
	}

	return "", nil
}

// GetEpochStartDataAPI -
func (ns *NodeStub) GetEpochStartDataAPI(epoch uint32) (*common.EpochStartDataAPI, error) {
	if ns.GetEpochStartDataAPICalled != nil {
		return ns.GetEpochStartDataAPICalled(epoch)
	}

	return &common.EpochStartDataAPI{}, nil
}

// GetDCDTData -
func (ns *NodeStub) GetDCDTData(address, tokenID string, nonce uint64, options api.AccountQueryOptions) (*dcdt.DCDigitalToken, api.BlockInfo, error) {
	if ns.GetDCDTDataCalled != nil {
		return ns.GetDCDTDataCalled(address, tokenID, nonce, options)
	}

	return &dcdt.DCDigitalToken{Value: big.NewInt(0)}, api.BlockInfo{}, nil
}

// GetDCDTsRoles -
func (ns *NodeStub) GetDCDTsRoles(address string, options api.AccountQueryOptions, ctx context.Context) (map[string][]string, api.BlockInfo, error) {
	if ns.GetDCDTsRolesCalled != nil {
		return ns.GetDCDTsRolesCalled(address, options, ctx)
	}

	return map[string][]string{}, api.BlockInfo{}, nil
}

// GetDCDTsWithRole -
func (ns *NodeStub) GetDCDTsWithRole(address string, role string, options api.AccountQueryOptions, ctx context.Context) ([]string, api.BlockInfo, error) {
	if ns.GetDCDTsWithRoleCalled != nil {
		return ns.GetDCDTsWithRoleCalled(address, role, options, ctx)
	}

	return make([]string, 0), api.BlockInfo{}, nil
}

// GetAllDCDTTokens -
func (ns *NodeStub) GetAllDCDTTokens(address string, options api.AccountQueryOptions, ctx context.Context) (map[string]*dcdt.DCDigitalToken, api.BlockInfo, error) {
	if ns.GetAllDCDTTokensCalled != nil {
		return ns.GetAllDCDTTokensCalled(address, options, ctx)
	}

	return make(map[string]*dcdt.DCDigitalToken), api.BlockInfo{}, nil
}

// GetTokenSupply -
func (ns *NodeStub) GetTokenSupply(token string) (*api.DCDTSupply, error) {
	if ns.GetTokenSupplyCalled != nil {
		return ns.GetTokenSupplyCalled(token)
	}
	return nil, nil
}

// GetAllIssuedDCDTs -
func (ns *NodeStub) GetAllIssuedDCDTs(tokenType string, ctx context.Context) ([]string, error) {
	if ns.GetAllIssuedDCDTsCalled != nil {
		return ns.GetAllIssuedDCDTsCalled(tokenType, ctx)
	}
	return make([]string, 0), nil
}

// IsDataTrieMigrated -
func (ns *NodeStub) IsDataTrieMigrated(address string, options api.AccountQueryOptions) (bool, error) {
	if ns.IsDataTrieMigratedCalled != nil {
		return ns.IsDataTrieMigratedCalled(address, options)
	}
	return false, nil
}

// GetNFTTokenIDsRegisteredByAddress -
func (ns *NodeStub) GetNFTTokenIDsRegisteredByAddress(address string, options api.AccountQueryOptions, ctx context.Context) ([]string, api.BlockInfo, error) {
	if ns.GetNFTTokenIDsRegisteredByAddressCalled != nil {
		return ns.GetNFTTokenIDsRegisteredByAddressCalled(address, options, ctx)
	}

	return make([]string, 0), api.BlockInfo{}, nil
}

// IsInterfaceNil returns true if there is no value under the interface
func (ns *NodeStub) IsInterfaceNil() bool {
	return ns == nil
}
