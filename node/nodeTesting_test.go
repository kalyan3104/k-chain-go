package node_test

import (
	"errors"
	"math/big"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/kalyan3104/k-chain-core-go/data/batch"
	"github.com/kalyan3104/k-chain-core-go/data/transaction"
	"github.com/kalyan3104/k-chain-go/common"
	"github.com/kalyan3104/k-chain-go/dataRetriever"
	"github.com/kalyan3104/k-chain-go/node"
	"github.com/kalyan3104/k-chain-go/node/mock"
	factoryMock "github.com/kalyan3104/k-chain-go/node/mock/factory"
	"github.com/kalyan3104/k-chain-go/process/factory"
	"github.com/kalyan3104/k-chain-go/storage"
	"github.com/kalyan3104/k-chain-go/testscommon"
	"github.com/kalyan3104/k-chain-go/testscommon/cryptoMocks"
	dataRetrieverMock "github.com/kalyan3104/k-chain-go/testscommon/dataRetriever"
	factoryMocks "github.com/kalyan3104/k-chain-go/testscommon/factory"
	"github.com/kalyan3104/k-chain-go/testscommon/p2pmocks"
	stateMock "github.com/kalyan3104/k-chain-go/testscommon/state"
	"github.com/kalyan3104/k-chain-go/testscommon/storageManager"
	trieMock "github.com/kalyan3104/k-chain-go/testscommon/trie"
	crypto "github.com/kalyan3104/k-chain-crypto-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var timeoutWait = time.Second

// ------- GenerateAndSendBulkTransactions

func TestGenerateAndSendBulkTransactions_ZeroTxShouldErr(t *testing.T) {
	n, _ := node.NewNode()

	err := n.GenerateAndSendBulkTransactions("", big.NewInt(0), 0, &mock.PrivateKeyStub{}, nil, []byte("chainID"), 1)
	assert.NotNil(t, err)
	assert.Equal(t, "can not generate and broadcast 0 transactions", err.Error())
}

func TestGenerateAndSendBulkTransactions_NilAccountAdapterShouldErr(t *testing.T) {
	marshalizer := &mock.MarshalizerFake{}

	keyGen := &mock.KeyGenMock{}
	sk, _ := keyGen.GeneratePair()
	singleSigner := &mock.SinglesignMock{}

	coreComponents := getDefaultCoreComponents()
	coreComponents.IntMarsh = marshalizer
	coreComponents.AddrPubKeyConv = createMockPubkeyConverter()
	processComponents := getDefaultProcessComponents()
	processComponents.ShardCoord = mock.NewOneShardCoordinatorMock()
	cryptoComponents := getDefaultCryptoComponents()
	cryptoComponents.TxSig = singleSigner
	stateComponents := getDefaultStateComponents()

	n, _ := node.NewNode(
		node.WithCoreComponents(coreComponents),
		node.WithStateComponents(stateComponents),
		node.WithProcessComponents(processComponents),
		node.WithCryptoComponents(cryptoComponents),
	)

	stateComponents.AccountsAPI = nil
	err := n.GenerateAndSendBulkTransactions(createDummyHexAddress(64), big.NewInt(0), 1, sk, nil, []byte("chainID"), 1)
	assert.Equal(t, node.ErrNilAccountsAdapter, err)
}

func TestGenerateAndSendBulkTransactions_NilSingleSignerShouldErr(t *testing.T) {
	marshalizer := &mock.MarshalizerFake{}

	keyGen := &mock.KeyGenMock{}
	sk, _ := keyGen.GeneratePair()
	accAdapter := getAccAdapter(big.NewInt(0))
	coreComponents := getDefaultCoreComponents()
	coreComponents.IntMarsh = marshalizer
	coreComponents.AddrPubKeyConv = createMockPubkeyConverter()
	processComponents := getDefaultProcessComponents()
	processComponents.ShardCoord = mock.NewOneShardCoordinatorMock()
	stateComponents := getDefaultStateComponents()
	stateComponents.Accounts = accAdapter
	cryptoComponents := getDefaultCryptoComponents()

	n, _ := node.NewNode(
		node.WithCoreComponents(coreComponents),
		node.WithProcessComponents(processComponents),
		node.WithStateComponents(stateComponents),
		node.WithCryptoComponents(cryptoComponents),
	)

	cryptoComponents.TxSig = nil
	err := n.GenerateAndSendBulkTransactions(createDummyHexAddress(64), big.NewInt(0), 1, sk, nil, []byte("chainID"), 1)
	assert.Equal(t, node.ErrNilSingleSig, err)
}

func TestGenerateAndSendBulkTransactions_NilShardCoordinatorShouldErr(t *testing.T) {
	marshalizer := &mock.MarshalizerFake{}

	keyGen := &mock.KeyGenMock{}
	sk, _ := keyGen.GeneratePair()
	accAdapter := getAccAdapter(big.NewInt(0))
	singleSigner := &mock.SinglesignMock{}
	coreComponents := getDefaultCoreComponents()
	coreComponents.IntMarsh = marshalizer
	coreComponents.AddrPubKeyConv = createMockPubkeyConverter()
	cryptoComponents := getDefaultCryptoComponents()
	cryptoComponents.TxSig = singleSigner
	stateComponents := getDefaultStateComponents()
	stateComponents.Accounts = accAdapter
	processComponents := getDefaultProcessComponents()

	n, _ := node.NewNode(
		node.WithCoreComponents(coreComponents),
		node.WithCryptoComponents(cryptoComponents),
		node.WithStateComponents(stateComponents),
		node.WithProcessComponents(processComponents),
	)

	processComponents.ShardCoord = nil
	err := n.GenerateAndSendBulkTransactions(createDummyHexAddress(64), big.NewInt(0), 1, sk, nil, []byte("chainID"), 1)
	assert.Equal(t, node.ErrNilShardCoordinator, err)
}

func TestGenerateAndSendBulkTransactions_NilPubkeyConverterShouldErr(t *testing.T) {
	marshalizer := &mock.MarshalizerFake{}
	accAdapter := getAccAdapter(big.NewInt(0))
	keyGen := &mock.KeyGenMock{}
	sk, _ := keyGen.GeneratePair()
	singleSigner := &mock.SinglesignMock{}
	coreComponents := getDefaultCoreComponents()
	coreComponents.IntMarsh = marshalizer
	cryptoComponents := getDefaultCryptoComponents()
	cryptoComponents.TxSig = singleSigner
	stateComponents := getDefaultStateComponents()
	stateComponents.Accounts = accAdapter
	processComponents := getDefaultProcessComponents()

	n, _ := node.NewNode(
		node.WithCoreComponents(coreComponents),
		node.WithCryptoComponents(cryptoComponents),
		node.WithStateComponents(stateComponents),
		node.WithProcessComponents(processComponents),
	)

	coreComponents.AddrPubKeyConv = nil
	err := n.GenerateAndSendBulkTransactions(createDummyHexAddress(64), big.NewInt(0), 1, sk, nil, []byte("chainID"), 1)
	assert.Equal(t, node.ErrNilPubkeyConverter, err)
}

func TestGenerateAndSendBulkTransactions_NilPrivateKeyShouldErr(t *testing.T) {
	accAdapter := getAccAdapter(big.NewInt(0))
	singleSigner := &mock.SinglesignMock{}
	dataPool := &dataRetrieverMock.PoolsHolderStub{
		TransactionsCalled: func() dataRetriever.ShardedDataCacherNotifier {
			return &testscommon.ShardedDataStub{
				ShardDataStoreCalled: func(cacheId string) (c storage.Cacher) {
					return nil
				},
			}
		},
	}
	coreComponents := getDefaultCoreComponents()
	coreComponents.IntMarsh = &mock.MarshalizerFake{}
	coreComponents.AddrPubKeyConv = createMockPubkeyConverter()
	processComponents := getDefaultProcessComponents()
	processComponents.ShardCoord = mock.NewOneShardCoordinatorMock()
	dataComponents := getDefaultDataComponents()
	dataComponents.DataPool = dataPool
	cryptoComponents := getDefaultCryptoComponents()
	cryptoComponents.TxSig = singleSigner
	stateComponents := getDefaultStateComponents()
	stateComponents.Accounts = accAdapter

	n, _ := node.NewNode(
		node.WithCoreComponents(coreComponents),
		node.WithProcessComponents(processComponents),
		node.WithDataComponents(dataComponents),
		node.WithCryptoComponents(cryptoComponents),
		node.WithStateComponents(stateComponents),
	)

	err := n.GenerateAndSendBulkTransactions(createDummyHexAddress(64), big.NewInt(0), 1, nil, nil, []byte("chainID"), 1)
	assert.NotNil(t, err)
	assert.True(t, strings.Contains(err.Error(), "trying to set nil private key"))
}

func TestGenerateAndSendBulkTransactions_InvalidReceiverAddressShouldErr(t *testing.T) {
	accAdapter := getAccAdapter(big.NewInt(0))

	sk := &mock.PrivateKeyStub{GeneratePublicHandler: func() crypto.PublicKey {
		return &mock.PublicKeyMock{
			ToByteArrayHandler: func() (bytes []byte, err error) {
				return []byte("key"), nil
			},
		}
	}}
	singleSigner := &mock.SinglesignMock{}
	dataPool := &dataRetrieverMock.PoolsHolderStub{
		TransactionsCalled: func() dataRetriever.ShardedDataCacherNotifier {
			return &testscommon.ShardedDataStub{
				ShardDataStoreCalled: func(cacheId string) (c storage.Cacher) {
					return nil
				},
			}
		},
	}
	expectedErr := errors.New("expected error")
	coreComponents := getDefaultCoreComponents()
	coreComponents.AddrPubKeyConv = &testscommon.PubkeyConverterStub{
		DecodeCalled: func(humanReadable string) ([]byte, error) {
			if len(humanReadable) == 0 {
				return nil, expectedErr
			}

			return []byte("1234"), nil
		},
	}
	processComponents := getDefaultProcessComponents()
	processComponents.ShardCoord = mock.NewOneShardCoordinatorMock()
	dataComponents := getDefaultDataComponents()
	dataComponents.DataPool = dataPool
	cryptoComponents := getDefaultCryptoComponents()
	cryptoComponents.TxSig = singleSigner
	stateComponents := getDefaultStateComponents()
	stateComponents.Accounts = accAdapter

	n, _ := node.NewNode(
		node.WithCoreComponents(coreComponents),
		node.WithProcessComponents(processComponents),
		node.WithDataComponents(dataComponents),
		node.WithCryptoComponents(cryptoComponents),
		node.WithStateComponents(stateComponents),
	)

	err := n.GenerateAndSendBulkTransactions("", big.NewInt(0), 1, sk, nil, []byte("chainID"), 1)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "could not create receiver address from provided param")
}

func TestGenerateAndSendBulkTransactions_MarshalizerErrorsShouldErr(t *testing.T) {
	accAdapter := getAccAdapter(big.NewInt(0))
	marshalizer := &mock.MarshalizerFake{}
	marshalizer.Fail = true
	sk := &mock.PrivateKeyStub{GeneratePublicHandler: func() crypto.PublicKey {
		return &mock.PublicKeyMock{
			ToByteArrayHandler: func() (bytes []byte, err error) {
				return []byte("key"), nil
			},
		}
	}}
	singleSigner := &mock.SinglesignMock{}
	dataPool := &dataRetrieverMock.PoolsHolderStub{
		TransactionsCalled: func() dataRetriever.ShardedDataCacherNotifier {
			return &testscommon.ShardedDataStub{
				ShardDataStoreCalled: func(cacheId string) (c storage.Cacher) {
					return nil
				},
			}
		},
	}

	coreComponents := getDefaultCoreComponents()
	coreComponents.IntMarsh = marshalizer
	coreComponents.AddrPubKeyConv = createMockPubkeyConverter()
	processComponents := getDefaultProcessComponents()
	processComponents.ShardCoord = mock.NewOneShardCoordinatorMock()
	dataComponents := getDefaultDataComponents()
	dataComponents.DataPool = dataPool
	cryptoComponents := getDefaultCryptoComponents()
	cryptoComponents.TxSig = singleSigner
	stateComponents := getDefaultStateComponents()
	stateComponents.AccountsAPI = accAdapter

	n, _ := node.NewNode(
		node.WithCoreComponents(coreComponents),
		node.WithProcessComponents(processComponents),
		node.WithDataComponents(dataComponents),
		node.WithCryptoComponents(cryptoComponents),
		node.WithStateComponents(stateComponents),
	)

	err := n.GenerateAndSendBulkTransactions(createDummyHexAddress(64), big.NewInt(1), 1, sk, nil, []byte("chainID"), 1)
	assert.NotNil(t, err)
	assert.True(t, strings.Contains(err.Error(), "MarshalizerMock generic error"))
}

func TestGenerateAndSendBulkTransactions_ShouldWork(t *testing.T) {
	marshalizer := &mock.MarshalizerFake{}

	noOfTx := 1000
	mutRecoveredTransactions := &sync.RWMutex{}
	recoveredTransactions := make(map[uint64]*transaction.Transaction)
	signer := &mock.SinglesignMock{}
	shardCoordinator := mock.NewOneShardCoordinatorMock()

	wg := sync.WaitGroup{}
	wg.Add(noOfTx)

	chDone := make(chan struct{})
	go func() {
		wg.Wait()
		chDone <- struct{}{}
	}()

	mes := &p2pmocks.MessengerStub{
		BroadcastOnChannelCalled: func(pipe string, topic string, buff []byte) {
			identifier := factory.TransactionTopic + shardCoordinator.CommunicationIdentifier(shardCoordinator.SelfId())

			if topic == identifier {
				// handler to capture sent data
				b := &batch.Batch{}
				err := marshalizer.Unmarshal(b, buff)
				if err != nil {
					assert.Fail(t, err.Error())
				}
				for _, txBuff := range b.Data {
					tx := transaction.Transaction{}
					errMarshal := marshalizer.Unmarshal(&tx, txBuff)
					require.Nil(t, errMarshal)

					mutRecoveredTransactions.Lock()
					recoveredTransactions[tx.Nonce] = &tx
					mutRecoveredTransactions.Unlock()

					wg.Done()
				}
			}
		},
	}

	dataPool := &dataRetrieverMock.PoolsHolderStub{
		TransactionsCalled: func() dataRetriever.ShardedDataCacherNotifier {
			return &testscommon.ShardedDataStub{
				ShardDataStoreCalled: func(cacheId string) (c storage.Cacher) {
					return nil
				},
			}
		},
	}
	accAdapter := getAccAdapter(big.NewInt(0))
	sk := &mock.PrivateKeyStub{GeneratePublicHandler: func() crypto.PublicKey {
		return &mock.PublicKeyMock{
			ToByteArrayHandler: func() (bytes []byte, err error) {
				return []byte("key"), nil
			},
		}
	}}
	coreComponents := getDefaultCoreComponents()
	coreComponents.IntMarsh = marshalizer
	coreComponents.TxMarsh = marshalizer
	coreComponents.AddrPubKeyConv = createMockPubkeyConverter()
	processComponents := getDefaultProcessComponents()
	processComponents.ShardCoord = shardCoordinator
	dataComponents := getDefaultDataComponents()
	dataComponents.DataPool = dataPool
	cryptoComponents := getDefaultCryptoComponents()
	cryptoComponents.TxSig = signer
	stateComponents := getDefaultStateComponents()
	stateComponents.AccountsAPI = accAdapter
	networkComponents := getDefaultNetworkComponents()
	networkComponents.Messenger = mes

	n, _ := node.NewNode(
		node.WithCoreComponents(coreComponents),
		node.WithProcessComponents(processComponents),
		node.WithDataComponents(dataComponents),
		node.WithCryptoComponents(cryptoComponents),
		node.WithStateComponents(stateComponents),
		node.WithNetworkComponents(networkComponents),
	)

	err := n.GenerateAndSendBulkTransactions(createDummyHexAddress(64), big.NewInt(1), uint64(noOfTx), sk, nil, []byte("chainID"), 1)
	assert.Nil(t, err)

	select {
	case <-chDone:
	case <-time.After(timeoutWait):
		assert.Fail(t, "timout while waiting the broadcast of the generated transactions")
		return
	}

	mutRecoveredTransactions.RLock()
	assert.Equal(t, noOfTx, len(recoveredTransactions))
	mutRecoveredTransactions.RUnlock()
}

func getDefaultCryptoComponents() *factoryMock.CryptoComponentsMock {
	return &factoryMock.CryptoComponentsMock{
		PubKey:                  &mock.PublicKeyMock{},
		P2pPubKey:               &mock.PublicKeyMock{},
		PrivKey:                 &mock.PrivateKeyStub{},
		P2pPrivKey:              &mock.PrivateKeyStub{},
		PubKeyString:            "pubKey",
		PubKeyBytes:             []byte("pubKey"),
		BlockSig:                &mock.SingleSignerMock{},
		TxSig:                   &mock.SingleSignerMock{},
		MultiSigContainer:       cryptoMocks.NewMultiSignerContainerMock(cryptoMocks.NewMultiSigner()),
		PeerSignHandler:         &mock.PeerSignatureHandler{},
		BlKeyGen:                &mock.KeyGenMock{},
		TxKeyGen:                &mock.KeyGenMock{},
		P2PKeyGen:               &mock.KeyGenMock{},
		MsgSigVerifier:          &testscommon.MessageSignVerifierMock{},
		KeysHandlerField:        &testscommon.KeysHandlerStub{},
		ManagedPeersHolderField: &testscommon.ManagedPeersHolderStub{},
	}
}

func getDefaultStateComponents() *factoryMocks.StateComponentsMock {
	return &factoryMocks.StateComponentsMock{
		PeersAcc:        &stateMock.AccountsStub{},
		Accounts:        &stateMock.AccountsStub{},
		AccountsAPI:     &stateMock.AccountsStub{},
		AccountsRepo:    &stateMock.AccountsRepositoryStub{},
		Tries:           &trieMock.TriesHolderStub{},
		StorageManagers: map[string]common.StorageManager{"0": &storageManager.StorageManagerStub{}},
	}
}

func getDefaultNetworkComponents() *factoryMock.NetworkComponentsMock {
	return &factoryMock.NetworkComponentsMock{
		Messenger:       &p2pmocks.MessengerStub{},
		InputAntiFlood:  &mock.P2PAntifloodHandlerStub{},
		OutputAntiFlood: &mock.P2PAntifloodHandlerStub{},
		PeerBlackList:   &mock.PeerBlackListHandlerStub{},
	}
}
