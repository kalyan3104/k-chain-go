package counters

import (
	"testing"

	"github.com/kalyan3104/k-chain-core-go/core/check"
	"github.com/kalyan3104/k-chain-go/process"
	"github.com/kalyan3104/k-chain-go/testscommon"
	vmcommon "github.com/kalyan3104/k-chain-vm-common-go"
	"github.com/stretchr/testify/assert"
)

func generateMapsOfValues(value uint64) map[string]uint64 {
	return map[string]uint64{
		maxTrieReads:    value,
		maxTransfers:    value,
		maxBuiltinCalls: value,
	}
}

func TestNewUsageCounter(t *testing.T) {
	t.Parallel()

	t.Run("nil dcdt transfer parser should error", func(t *testing.T) {
		t.Parallel()

		counter, err := NewUsageCounter(nil)
		assert.True(t, check.IfNil(counter))
		assert.Equal(t, process.ErrNilDCDTTransferParser, err)
	})
	t.Run("nil dcdt transfer parser should error", func(t *testing.T) {
		t.Parallel()

		counter, err := NewUsageCounter(&testscommon.DCDTTransferParserStub{})
		assert.False(t, check.IfNil(counter))
		assert.Nil(t, err)
	})
}

func TestUsageCounter_ProcessCrtNumberOfTrieReadsCounter(t *testing.T) {
	t.Parallel()

	counter, _ := NewUsageCounter(&testscommon.DCDTTransferParserStub{})
	counter.SetMaximumValues(generateMapsOfValues(3))

	assert.Nil(t, counter.ProcessCrtNumberOfTrieReadsCounter())                                 // counter is now 1
	assert.Nil(t, counter.ProcessCrtNumberOfTrieReadsCounter())                                 // counter is now 2
	assert.Nil(t, counter.ProcessCrtNumberOfTrieReadsCounter())                                 // counter is now 3
	assert.ErrorIs(t, counter.ProcessCrtNumberOfTrieReadsCounter(), process.ErrMaxCallsReached) // counter is now 4, error signalled
	err := counter.ProcessCrtNumberOfTrieReadsCounter()                                         // counter is now 5, error signalled
	assert.ErrorIs(t, err, process.ErrMaxCallsReached)
	t.Run("backwards compatibility on error message string", func(t *testing.T) {
		assert.Equal(t, "max calls reached: too many reads from trie", err.Error())
	})

	countersMap := counter.GetCounterValues()
	expectedMap := map[string]uint64{
		crtBuiltinCalls: 0,
		crtTransfers:    0,
		crtTrieReads:    5,
	}
	assert.Equal(t, expectedMap, countersMap)

	counter.ResetCounters()

	assert.Nil(t, counter.ProcessCrtNumberOfTrieReadsCounter()) // counter is now 1
	countersMap = counter.GetCounterValues()
	expectedMap = map[string]uint64{
		crtBuiltinCalls: 0,
		crtTransfers:    0,
		crtTrieReads:    1,
	}
	assert.Equal(t, expectedMap, countersMap)
}

func TestUsageCounter_ProcessMaxBuiltInCounters(t *testing.T) {
	t.Parallel()

	t.Run("builtin functions exceeded", func(t *testing.T) {
		t.Parallel()

		counter, _ := NewUsageCounter(&testscommon.DCDTTransferParserStub{})
		counter.SetMaximumValues(generateMapsOfValues(3))

		vmInput := &vmcommon.ContractCallInput{}

		assert.Nil(t, counter.ProcessMaxBuiltInCounters(vmInput))                                 // counter is now 1
		assert.Nil(t, counter.ProcessMaxBuiltInCounters(vmInput))                                 // counter is now 2
		assert.Nil(t, counter.ProcessMaxBuiltInCounters(vmInput))                                 // counter is now 3
		assert.ErrorIs(t, counter.ProcessMaxBuiltInCounters(vmInput), process.ErrMaxCallsReached) // counter is now 4, error signalled
		err := counter.ProcessMaxBuiltInCounters(vmInput)                                         // counter is now 5, error signalled
		assert.ErrorIs(t, err, process.ErrMaxCallsReached)
		t.Run("backwards compatibility on error message string", func(t *testing.T) {
			assert.Equal(t, "max calls reached: too many built-in functions calls", err.Error())
		})

		countersMap := counter.GetCounterValues()
		expectedMap := map[string]uint64{
			crtBuiltinCalls: 5,
			crtTransfers:    0,
			crtTrieReads:    0,
		}
		assert.Equal(t, expectedMap, countersMap)

		counter.ResetCounters()

		assert.Nil(t, counter.ProcessMaxBuiltInCounters(vmInput)) // counter is now 1
		countersMap = counter.GetCounterValues()
		expectedMap = map[string]uint64{
			crtBuiltinCalls: 1,
			crtTransfers:    0,
			crtTrieReads:    0,
		}
		assert.Equal(t, expectedMap, countersMap)
	})
	t.Run("number of transfers exceeded", func(t *testing.T) {
		t.Parallel()

		counter, _ := NewUsageCounter(&testscommon.DCDTTransferParserStub{
			ParseDCDTTransfersCalled: func(sndAddr []byte, rcvAddr []byte, function string, args [][]byte) (*vmcommon.ParsedDCDTTransfers, error) {
				return &vmcommon.ParsedDCDTTransfers{
					DCDTTransfers: make([]*vmcommon.DCDTTransfer, 2),
				}, nil
			},
		})
		values := generateMapsOfValues(6)
		values[maxBuiltinCalls] = 1000
		counter.SetMaximumValues(values)

		vmInput := &vmcommon.ContractCallInput{}

		assert.Nil(t, counter.ProcessMaxBuiltInCounters(vmInput))                                 // counter is now 2
		assert.Nil(t, counter.ProcessMaxBuiltInCounters(vmInput))                                 // counter is now 4
		assert.Nil(t, counter.ProcessMaxBuiltInCounters(vmInput))                                 // counter is now 6
		assert.ErrorIs(t, counter.ProcessMaxBuiltInCounters(vmInput), process.ErrMaxCallsReached) // counter is now 8, error signalled
		err := counter.ProcessMaxBuiltInCounters(vmInput)                                         // counter is now 10, error signalled
		assert.ErrorIs(t, err, process.ErrMaxCallsReached)
		t.Run("backwards compatibility on error message string", func(t *testing.T) {
			assert.Equal(t, "max calls reached: too many DCDT transfers", err.Error())
		})

		countersMap := counter.GetCounterValues()
		expectedMap := map[string]uint64{
			crtBuiltinCalls: 5,
			crtTransfers:    10,
			crtTrieReads:    0,
		}
		assert.Equal(t, expectedMap, countersMap)

		counter.ResetCounters()

		assert.Nil(t, counter.ProcessMaxBuiltInCounters(vmInput)) // counter is now 2
		countersMap = counter.GetCounterValues()
		expectedMap = map[string]uint64{
			crtBuiltinCalls: 1,
			crtTransfers:    2,
			crtTrieReads:    0,
		}
		assert.Equal(t, expectedMap, countersMap)
	})
}
