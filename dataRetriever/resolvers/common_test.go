package resolvers_test

import (
	"github.com/kalyan3104/k-chain-go/dataRetriever"
	"github.com/kalyan3104/k-chain-go/dataRetriever/mock"
	"github.com/kalyan3104/k-chain-go/p2p"
	"github.com/kalyan3104/k-chain-go/testscommon/p2pmocks"
)

func createRequestMsg(dataType dataRetriever.RequestDataType, val []byte) p2p.MessageP2P {
	marshalizer := &mock.MarshalizerMock{}
	buff, _ := marshalizer.Marshal(&dataRetriever.RequestData{Type: dataType, Value: val})
	return &p2pmocks.P2PMessageMock{DataField: buff}
}

func createRequestMsgWithChunkIndex(dataType dataRetriever.RequestDataType, val []byte, epoch uint32, chunkIndex uint32) p2p.MessageP2P {
	marshalizer := &mock.MarshalizerMock{}
	buff, _ := marshalizer.Marshal(&dataRetriever.RequestData{
		Type:       dataType,
		Value:      val,
		Epoch:      epoch,
		ChunkIndex: chunkIndex,
	})
	return &p2pmocks.P2PMessageMock{DataField: buff}
}
