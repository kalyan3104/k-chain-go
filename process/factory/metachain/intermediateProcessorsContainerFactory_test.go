package metachain_test

import (
	"testing"

	"github.com/kalyan3104/k-chain-go/common"
	"github.com/kalyan3104/k-chain-go/process"
	"github.com/kalyan3104/k-chain-go/process/factory/metachain"
	"github.com/kalyan3104/k-chain-go/process/mock"
	"github.com/kalyan3104/k-chain-go/testscommon"
	txExecOrderStub "github.com/kalyan3104/k-chain-go/testscommon/common"
	dataRetrieverMock "github.com/kalyan3104/k-chain-go/testscommon/dataRetriever"
	"github.com/kalyan3104/k-chain-go/testscommon/economicsmocks"
	"github.com/kalyan3104/k-chain-go/testscommon/enableEpochsHandlerMock"
	"github.com/kalyan3104/k-chain-go/testscommon/hashingMocks"
	storageStubs "github.com/kalyan3104/k-chain-go/testscommon/storage"
	"github.com/stretchr/testify/assert"
)

func createMockPubkeyConverter() *testscommon.PubkeyConverterMock {
	return testscommon.NewPubkeyConverterMock(32)
}

func createMockArgsNewIntermediateProcessorsFactory() metachain.ArgsNewIntermediateProcessorsContainerFactory {
	args := metachain.ArgsNewIntermediateProcessorsContainerFactory{
		Hasher:                  &hashingMocks.HasherMock{},
		Marshalizer:             &mock.MarshalizerMock{},
		ShardCoordinator:        mock.NewMultiShardsCoordinatorMock(5),
		PubkeyConverter:         createMockPubkeyConverter(),
		Store:                   &storageStubs.ChainStorerStub{},
		PoolsHolder:             dataRetrieverMock.NewPoolsHolderMock(),
		EconomicsFee:            &economicsmocks.EconomicsHandlerStub{},
		EnableEpochsHandler:     enableEpochsHandlerMock.NewEnableEpochsHandlerStub(common.KeepExecOrderOnCreatedSCRsFlag),
		TxExecutionOrderHandler: &txExecOrderStub.TxExecutionOrderHandlerStub{},
	}
	return args
}

func TestNewIntermediateProcessorsContainerFactory_NilShardCoord(t *testing.T) {
	t.Parallel()

	args := createMockArgsNewIntermediateProcessorsFactory()
	args.ShardCoordinator = nil
	ipcf, err := metachain.NewIntermediateProcessorsContainerFactory(args)

	assert.Nil(t, ipcf)
	assert.Equal(t, process.ErrNilShardCoordinator, err)
}

func TestNewIntermediateProcessorsContainerFactory_NilMarshalizer(t *testing.T) {
	t.Parallel()

	args := createMockArgsNewIntermediateProcessorsFactory()
	args.Marshalizer = nil
	ipcf, err := metachain.NewIntermediateProcessorsContainerFactory(args)

	assert.Nil(t, ipcf)
	assert.Equal(t, process.ErrNilMarshalizer, err)
}

func TestNewIntermediateProcessorsContainerFactory_NilHasher(t *testing.T) {
	t.Parallel()

	args := createMockArgsNewIntermediateProcessorsFactory()
	args.Hasher = nil
	ipcf, err := metachain.NewIntermediateProcessorsContainerFactory(args)

	assert.Nil(t, ipcf)
	assert.Equal(t, process.ErrNilHasher, err)
}

func TestNewIntermediateProcessorsContainerFactory_NilAdrConv(t *testing.T) {
	t.Parallel()

	args := createMockArgsNewIntermediateProcessorsFactory()
	args.PubkeyConverter = nil
	ipcf, err := metachain.NewIntermediateProcessorsContainerFactory(args)

	assert.Nil(t, ipcf)
	assert.Equal(t, process.ErrNilPubkeyConverter, err)
}

func TestNewIntermediateProcessorsContainerFactory_NilStorer(t *testing.T) {
	t.Parallel()

	args := createMockArgsNewIntermediateProcessorsFactory()
	args.Store = nil
	ipcf, err := metachain.NewIntermediateProcessorsContainerFactory(args)

	assert.Nil(t, ipcf)
	assert.Equal(t, process.ErrNilStorage, err)
}

func TestNewIntermediateProcessorsContainerFactory_NilPoolsHolder(t *testing.T) {
	t.Parallel()

	args := createMockArgsNewIntermediateProcessorsFactory()
	args.PoolsHolder = nil
	ipcf, err := metachain.NewIntermediateProcessorsContainerFactory(args)

	assert.Nil(t, ipcf)
	assert.Equal(t, process.ErrNilPoolsHolder, err)
}

func TestNewIntermediateProcessorsContainerFactory_NilEconomicsFeeHandler(t *testing.T) {
	t.Parallel()

	args := createMockArgsNewIntermediateProcessorsFactory()
	args.EconomicsFee = nil
	ipcf, err := metachain.NewIntermediateProcessorsContainerFactory(args)

	assert.Nil(t, ipcf)
	assert.Equal(t, process.ErrNilEconomicsFeeHandler, err)
}

func TestNewIntermediateProcessorsContainerFactory_NilEnableEpochHandler(t *testing.T) {
	t.Parallel()

	args := createMockArgsNewIntermediateProcessorsFactory()
	args.EnableEpochsHandler = nil
	ipcf, err := metachain.NewIntermediateProcessorsContainerFactory(args)

	assert.Nil(t, ipcf)
	assert.Equal(t, process.ErrNilEnableEpochsHandler, err)
}

func TestNewIntermediateProcessorsContainerFactory(t *testing.T) {
	t.Parallel()

	args := createMockArgsNewIntermediateProcessorsFactory()
	ipcf, err := metachain.NewIntermediateProcessorsContainerFactory(args)

	assert.Nil(t, err)
	assert.NotNil(t, ipcf)
	assert.False(t, ipcf.IsInterfaceNil())
}

func TestIntermediateProcessorsContainerFactory_Create(t *testing.T) {
	t.Parallel()

	args := createMockArgsNewIntermediateProcessorsFactory()
	ipcf, err := metachain.NewIntermediateProcessorsContainerFactory(args)

	assert.Nil(t, err)
	assert.NotNil(t, ipcf)

	container, err := ipcf.Create()
	assert.Nil(t, err)
	assert.Equal(t, 2, container.Len())
}
