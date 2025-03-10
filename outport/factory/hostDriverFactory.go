package factory

import (
	"github.com/kalyan3104/k-chain-go/config"
	"github.com/kalyan3104/k-chain-go/outport"
	"github.com/kalyan3104/k-chain-go/outport/host"
	"github.com/kalyan3104/k-chain-communication-go/websocket/data"
	"github.com/kalyan3104/k-chain-communication-go/websocket/factory"
	"github.com/kalyan3104/k-chain-core-go/marshal"
	logger "github.com/kalyan3104/k-chain-logger-go"
)

type ArgsHostDriverFactory struct {
	HostConfig config.HostDriversConfig
	Marshaller marshal.Marshalizer
}

var log = logger.GetOrCreate("outport/factory/hostdriver")

// CreateHostDriver will create a new instance of outport.Driver
func CreateHostDriver(args ArgsHostDriverFactory) (outport.Driver, error) {
	wsHost, err := factory.CreateWebSocketHost(factory.ArgsWebSocketHost{
		WebSocketConfig: data.WebSocketConfig{
			URL:                        args.HostConfig.URL,
			WithAcknowledge:            args.HostConfig.WithAcknowledge,
			Mode:                       args.HostConfig.Mode,
			RetryDurationInSec:         args.HostConfig.RetryDurationInSec,
			BlockingAckOnError:         args.HostConfig.BlockingAckOnError,
			DropMessagesIfNoConnection: args.HostConfig.DropMessagesIfNoConnection,
			AcknowledgeTimeoutInSec:    args.HostConfig.AcknowledgeTimeoutInSec,
			Version:                    args.HostConfig.Version,
		},
		Marshaller: args.Marshaller,
		Log:        log,
	})
	if err != nil {
		return nil, err
	}

	return host.NewHostDriver(host.ArgsHostDriver{
		Marshaller: args.Marshaller,
		SenderHost: wsHost,
		Log:        log,
	})
}
