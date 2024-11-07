package p2pmocks

import (
	"github.com/kalyan3104/k-chain-communication-go/p2p"
	"github.com/kalyan3104/k-chain-core-go/core"
)

// MessageProcessorStub -
type MessageProcessorStub struct {
	ProcessReceivedMessageCalled func(message p2p.MessageP2P, fromConnectedPeer core.PeerID, source p2p.MessageHandler) error
}

// ProcessReceivedMessage -
func (stub *MessageProcessorStub) ProcessReceivedMessage(message p2p.MessageP2P, fromConnectedPeer core.PeerID, source p2p.MessageHandler) error {
	if stub.ProcessReceivedMessageCalled != nil {
		return stub.ProcessReceivedMessageCalled(message, fromConnectedPeer, source)
	}

	return nil
}

// IsInterfaceNil -
func (stub *MessageProcessorStub) IsInterfaceNil() bool {
	return stub == nil
}
