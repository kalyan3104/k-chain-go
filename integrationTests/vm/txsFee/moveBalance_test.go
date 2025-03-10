package txsFee

import (
	"errors"
	"math"
	"math/big"
	"testing"

	"github.com/kalyan3104/k-chain-core-go/data/receipt"
	"github.com/kalyan3104/k-chain-go/config"
	"github.com/kalyan3104/k-chain-go/integrationTests/vm"
	"github.com/kalyan3104/k-chain-go/integrationTests/vm/txsFee/utils"
	"github.com/kalyan3104/k-chain-go/process"
	vmcommon "github.com/kalyan3104/k-chain-vm-common-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const gasPrice = uint64(10)

// minGasPrice = 1, gasPerDataByte = 1, minGasLimit = 1
func TestMoveBalanceSelfShouldWorkAndConsumeTxFee(t *testing.T) {
	if testing.Short() {
		t.Skip("this is not a short test")
	}

	testContext, err := vm.CreatePreparedTxProcessorWithVMs(config.EnableEpochs{})
	require.Nil(t, err)
	defer testContext.Close()

	sndAddr := []byte("12345678901234567890123456789012")
	senderNonce := uint64(0)
	senderBalance := big.NewInt(10000)
	gasLimit := uint64(100)

	_, _ = vm.CreateAccount(testContext.Accounts, sndAddr, 0, senderBalance)
	tx := vm.CreateTransaction(senderNonce, big.NewInt(100), sndAddr, sndAddr, gasPrice, gasLimit, []byte("aaaa"))

	returnCode, err := testContext.TxProcessor.ProcessTransaction(tx)
	require.Nil(t, err)
	require.Equal(t, vmcommon.Ok, returnCode)

	_, err = testContext.Accounts.Commit()
	require.Nil(t, err)

	// 10_000 - gasPrice(10)*gasLimit(1) + gasPerDataByte(1)*gasPrice(10) = 9950
	expectedBalance := big.NewInt(9950)
	vm.TestAccount(t, testContext.Accounts, sndAddr, senderNonce+1, expectedBalance)

	// check receipts
	require.Equal(t, 1, len(testContext.GetIntermediateTransactions(t)))
	rcpt := testContext.GetIntermediateTransactions(t)[0].(*receipt.Receipt)
	assert.Equal(t, "950", rcpt.Value.String())

	// check accumulated fees
	accumulatedFees := testContext.TxFeeHandler.GetAccumulatedFees()
	require.Equal(t, big.NewInt(50), accumulatedFees)
}

// minGasPrice = 1, gasPerDataByte = 1, minGasLimit = 1
func TestMoveBalanceAllFlagsEnabledLessBalanceThanGasLimitMulGasPrice(t *testing.T) {
	if testing.Short() {
		t.Skip("this is not a short test")
	}

	testContext, err := vm.CreatePreparedTxProcessorWithVMs(config.EnableEpochs{})
	require.Nil(t, err)
	defer testContext.Close()

	sndAddr := []byte("12345678901234567890123456789012")
	senderNonce := uint64(0)
	senderBalance := big.NewInt(10000)
	gasLimit := uint64(10000)

	_, _ = vm.CreateAccount(testContext.Accounts, sndAddr, 0, senderBalance)
	tx := vm.CreateTransaction(senderNonce, big.NewInt(100), sndAddr, sndAddr, gasPrice, gasLimit, []byte("aaaa"))

	_, err = testContext.TxProcessor.ProcessTransaction(tx)
	require.True(t, errors.Is(err, process.ErrInsufficientFee))
}

func TestMoveBalanceShouldWork(t *testing.T) {
	if testing.Short() {
		t.Skip("this is not a short test")
	}

	testContext, err := vm.CreatePreparedTxProcessorWithVMs(config.EnableEpochs{})
	require.Nil(t, err)
	defer testContext.Close()

	sndAddr := []byte("12345678901234567890123456789012")
	rcvAddr := []byte("12345678901234567890123456789022")
	senderNonce := uint64(0)
	senderBalance := big.NewInt(10000)
	gasLimit := uint64(100)

	_, _ = vm.CreateAccount(testContext.Accounts, sndAddr, 0, senderBalance)
	tx := vm.CreateTransaction(senderNonce, big.NewInt(100), sndAddr, rcvAddr, gasPrice, gasLimit, []byte("aaaa"))

	returnCode, err := testContext.TxProcessor.ProcessTransaction(tx)
	require.Nil(t, err)
	require.Equal(t, vmcommon.Ok, returnCode)

	_, err = testContext.Accounts.Commit()
	require.Nil(t, err)

	// 10_000 - gasPrice(10)*gasLimit(1) + gasPerDataByte(1)*gasPrice(10) - transferredValue(100) = 9850
	// verify balance of sender
	expectedBalance := big.NewInt(9850)
	vm.TestAccount(
		t,
		testContext.Accounts,
		sndAddr,
		senderNonce+1,
		expectedBalance)

	// verify receiver
	expectedBalanceReceiver := big.NewInt(100)
	vm.TestAccount(t, testContext.Accounts, rcvAddr, 0, expectedBalanceReceiver)

	// check accumulated fees
	accumulatedFees := testContext.TxFeeHandler.GetAccumulatedFees()
	require.Equal(t, big.NewInt(50), accumulatedFees)
}

func TestMoveBalanceInvalidHasGasButNoValueShouldConsumeGas(t *testing.T) {
	if testing.Short() {
		t.Skip("this is not a short test")
	}

	testContext, err := vm.CreatePreparedTxProcessorWithVMs(config.EnableEpochs{})
	require.Nil(t, err)
	defer testContext.Close()

	sndAddr := []byte("12345678901234567890123456789012")
	rcvAddr := []byte("12345678901234567890123456789022")
	senderBalance := big.NewInt(100)
	gasPriceLocal := uint64(1)
	gasLimit := uint64(20)
	_, _ = vm.CreateAccount(testContext.Accounts, sndAddr, 0, senderBalance)
	tx := vm.CreateTransaction(0, big.NewInt(100), sndAddr, rcvAddr, gasPriceLocal, gasLimit, []byte("aaaa"))

	returnCode, err := testContext.TxProcessor.ProcessTransaction(tx)
	require.Equal(t, process.ErrFailedTransaction, err)
	require.Equal(t, vmcommon.Ok, returnCode)

	_, err = testContext.Accounts.Commit()
	require.Nil(t, err)

	// check sender balance
	expectedBalance := big.NewInt(80)
	vm.TestAccount(t, testContext.Accounts, sndAddr, 1, expectedBalance)

	// check accumulated fees
	accumulatedFees := testContext.TxFeeHandler.GetAccumulatedFees()
	require.Equal(t, big.NewInt(20), accumulatedFees)
}

func TestMoveBalanceHigherNonceShouldNotConsumeGas(t *testing.T) {
	if testing.Short() {
		t.Skip("this is not a short test")
	}

	testContext, err := vm.CreatePreparedTxProcessorWithVMs(config.EnableEpochs{})
	require.Nil(t, err)
	defer testContext.Close()

	sndAddr := []byte("12345678901234567890123456789012")
	rcvAddr := []byte("12345678901234567890123456789022")

	senderBalance := big.NewInt(100)
	gasPriceLocal := uint64(1)
	gasLimit := uint64(20)

	_, _ = vm.CreateAccount(testContext.Accounts, sndAddr, 0, senderBalance)
	tx := vm.CreateTransaction(1, big.NewInt(500), sndAddr, rcvAddr, gasPriceLocal, gasLimit, []byte("aaaa"))

	_, err = testContext.TxProcessor.ProcessTransaction(tx)
	require.Equal(t, process.ErrHigherNonceInTransaction, err)

	_, err = testContext.Accounts.Commit()
	require.Nil(t, err)

	// check sender balance
	expectedBalance := big.NewInt(0).Set(senderBalance)
	vm.TestAccount(t, testContext.Accounts, sndAddr, 0, expectedBalance)

	// check accumulated fees
	accumulatedFees := testContext.TxFeeHandler.GetAccumulatedFees()
	require.Equal(t, big.NewInt(0), accumulatedFees)
}

func TestMoveBalanceMoreGasThanGasLimitPerMiniBlockForSafeCrossShard(t *testing.T) {
	if testing.Short() {
		t.Skip("this is not a short test")
	}

	testContext, err := vm.CreatePreparedTxProcessorWithVMs(config.EnableEpochs{})
	require.Nil(t, err)
	defer testContext.Close()

	sndAddr := []byte("12345678901234567890123456789012")
	rcvAddr := []byte("12345678901234567890123456789022")

	senderBalance := big.NewInt(0).SetUint64(math.MaxUint64)
	gasPriceLocal := uint64(1)
	gasLimit := uint64(math.MaxUint64)

	_, _ = vm.CreateAccount(testContext.Accounts, sndAddr, 0, senderBalance)
	tx := vm.CreateTransaction(0, big.NewInt(500), sndAddr, rcvAddr, gasPriceLocal, gasLimit, []byte("aaaa"))

	returnCode, err := testContext.TxProcessor.ProcessTransaction(tx)
	require.Equal(t, process.ErrMoreGasThanGasLimitPerBlock, err)
	require.Equal(t, vmcommon.UserError, returnCode)

	_, err = testContext.Accounts.Commit()
	require.Nil(t, err)

	// check sender balance
	expectedBalance := big.NewInt(0).Set(senderBalance)
	vm.TestAccount(t, testContext.Accounts, sndAddr, 0, expectedBalance)

	// check accumulated fees
	accumulatedFees := testContext.TxFeeHandler.GetAccumulatedFees()
	require.Equal(t, big.NewInt(0), accumulatedFees)
}

func TestMoveBalanceInvalidUserNames(t *testing.T) {
	if testing.Short() {
		t.Skip("this is not a short test")
	}

	testContext, err := vm.CreatePreparedTxProcessorWithVMs(config.EnableEpochs{})
	require.Nil(t, err)
	defer testContext.Close()

	sndAddr := []byte("12345678901234567890123456789012")
	rcvAddr := []byte("12345678901234567890123456789022")
	senderNonce := uint64(0)
	senderBalance := big.NewInt(10000)
	gasLimit := uint64(100)

	_, _ = vm.CreateAccount(testContext.Accounts, sndAddr, 0, senderBalance)
	tx := vm.CreateTransaction(senderNonce, big.NewInt(100), sndAddr, rcvAddr, gasPrice, gasLimit, []byte("aaaa"))
	tx.SndUserName = []byte("invalidUserName")
	tx.RcvUserName = []byte("invalidRcvUserName")

	_, err = testContext.TxProcessor.ProcessTransaction(tx)
	require.Equal(t, process.ErrFailedTransaction, err)

	_, err = testContext.Accounts.Commit()
	require.Nil(t, err)

	expectedBalance := big.NewInt(9000)
	utils.TestAccount(
		t,
		testContext.Accounts,
		sndAddr,
		1,
		expectedBalance)

	// check accumulated fees
	accumulatedFees := testContext.TxFeeHandler.GetAccumulatedFees()
	require.Equal(t, big.NewInt(1000), accumulatedFees)
}
