package spos

import (
	"github.com/kalyan3104/k-chain-go/common"
	"github.com/kalyan3104/k-chain-go/sharding"
)

// GetConsensusTopicID will construct and return the topic ID based on shard coordinator
func GetConsensusTopicID(shardCoordinator sharding.Coordinator) string {
	return common.ConsensusTopic + shardCoordinator.CommunicationIdentifier(shardCoordinator.SelfId())
}
