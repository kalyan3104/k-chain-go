package factory

import (
	"fmt"
	"testing"

	"github.com/kalyan3104/k-chain-go/config"
	"github.com/kalyan3104/k-chain-go/testscommon/marshallerMock"
	"github.com/kalyan3104/k-chain-communication-go/websocket/data"
	"github.com/stretchr/testify/require"
)

func TestCreateHostDriver(t *testing.T) {
	t.Parallel()

	args := ArgsHostDriverFactory{
		HostConfig: config.HostDriversConfig{
			URL:                "localhost",
			RetryDurationInSec: 1,
			MarshallerType:     "json",
			Mode:               data.ModeClient,
		},
		Marshaller: &marshallerMock.MarshalizerStub{},
	}

	driver, err := CreateHostDriver(args)
	require.Nil(t, err)
	require.NotNil(t, driver)
	require.Equal(t, "*host.hostDriver", fmt.Sprintf("%T", driver))
}
