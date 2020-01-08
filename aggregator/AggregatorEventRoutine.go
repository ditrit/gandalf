package aggregator

import (
	"errors"
	"fmt"
	"gandalf-go/message"

	"github.com/pebbe/zmq4"
)

type AggregatorEventRoutine struct {
	Context                                       *zmq4.Context
	AggregatorEventSendToCluster                  *zmq4.Socket
	AggregatorEventSendToClusterConnections       []string
	AggregatorEventReceiveFromConnector           *zmq4.Socket
	AggregatorEventReceiveFromConnectorConnection string
	AggregatorEventSendToConnector                *zmq4.Socket
	AggregatorEventSendToConnectorConnection      string
	AggregatorEventReceiveFromCluster             *zmq4.Socket
	AggregatorEventReceiveFromClusterConnections  []string
	Identity                                      string
}

func NewAggregatorEventRoutine(identity, aggregatorEventReceiveFromConnectorConnection, aggregatorEventSendToConnectorConnection string, aggregatorEventSendToClusterConnections, aggregatorEventReceiveFromClusterConnections []string) (aggregatorEventRoutine *AggregatorEventRoutine) {
	aggregatorEventRoutine = new(AggregatorEventRoutine)

	aggregatorEventRoutine.Identity = identity

	aggregatorEventRoutine.Context, _ = zmq4.NewContext()
	aggregatorEventRoutine.AggregatorEventSendToClusterConnections = aggregatorEventSendToClusterConnections
	aggregatorEventRoutine.AggregatorEventSendToCluster, _ = aggregatorEventRoutine.Context.NewSocket(zmq4.XPUB)
	aggregatorEventRoutine.AggregatorEventSendToCluster.SetIdentity(aggregatorEventRoutine.Identity)
	//aggregatorEventRoutine.aggregatorEventSendToCluster.Connect(aggregatorEventRoutine.aggregatorEventSendToClusterConnections)
	//fmt.Println("aggregatorEventSendToCluster connect : " + aggregatorEventSendToClusterConnections)
	for _, connection := range aggregatorEventRoutine.AggregatorEventSendToClusterConnections {
		aggregatorEventRoutine.AggregatorEventSendToCluster.Connect(connection)
		fmt.Println("aggregatorEventSendToCluster connect : " + connection)
	}

	aggregatorEventRoutine.AggregatorEventReceiveFromClusterConnections = aggregatorEventReceiveFromClusterConnections
	aggregatorEventRoutine.AggregatorEventReceiveFromCluster, _ = aggregatorEventRoutine.Context.NewSocket(zmq4.XSUB)
	aggregatorEventRoutine.AggregatorEventReceiveFromCluster.SetIdentity(aggregatorEventRoutine.Identity)
	//aggregatorEventRoutine.aggregatorEventSendToCluster.Connect(aggregatorEventRoutine.aggregatorEventReceiveFromClusterConnections)
	//fmt.Println("aggregatorEventReceiveFromCluster connect : " + aggregatorEventReceiveFromClusterConnections)
	for _, connection := range aggregatorEventRoutine.AggregatorEventReceiveFromClusterConnections {
		aggregatorEventRoutine.AggregatorEventReceiveFromCluster.Connect(connection)
		fmt.Println("aggregatorEventReceiveFromCluster connect : " + connection)
	}
	aggregatorEventRoutine.AggregatorEventReceiveFromCluster.SendBytes([]byte{0x01}, 0) //SUBSCRIBE ALL

	aggregatorEventRoutine.AggregatorEventSendToConnectorConnection = aggregatorEventSendToConnectorConnection
	aggregatorEventRoutine.AggregatorEventSendToConnector, _ = aggregatorEventRoutine.Context.NewSocket(zmq4.XPUB)
	aggregatorEventRoutine.AggregatorEventSendToConnector.SetIdentity(aggregatorEventRoutine.Identity)
	aggregatorEventRoutine.AggregatorEventSendToConnector.Bind(aggregatorEventRoutine.AggregatorEventSendToConnectorConnection)
	fmt.Println("aggregatorEventSendToConnector Bind : " + aggregatorEventSendToConnectorConnection)

	aggregatorEventRoutine.AggregatorEventReceiveFromConnectorConnection = aggregatorEventReceiveFromConnectorConnection
	aggregatorEventRoutine.AggregatorEventReceiveFromConnector, _ = aggregatorEventRoutine.Context.NewSocket(zmq4.XSUB)
	aggregatorEventRoutine.AggregatorEventReceiveFromConnector.SetIdentity(aggregatorEventRoutine.Identity)
	aggregatorEventRoutine.AggregatorEventReceiveFromConnector.Bind(aggregatorEventRoutine.AggregatorEventReceiveFromConnectorConnection)
	fmt.Println("aggregatorEventReceiveFromConnector Bind : " + aggregatorEventReceiveFromConnectorConnection)
	aggregatorEventRoutine.AggregatorEventReceiveFromConnector.SendBytes([]byte{0x01}, 0) //SUBSCRIBE ALL

	return
}

func (r AggregatorEventRoutine) close() {
	r.AggregatorEventSendToCluster.Close()
	r.AggregatorEventReceiveFromConnector.Close()
	r.AggregatorEventSendToConnector.Close()
	r.AggregatorEventReceiveFromCluster.Close()
	r.Context.Term()
}

func (r AggregatorEventRoutine) run() {
	poller := zmq4.NewPoller()
	poller.Add(r.AggregatorEventSendToCluster, zmq4.POLLIN)
	poller.Add(r.AggregatorEventReceiveFromConnector, zmq4.POLLIN)
	poller.Add(r.AggregatorEventSendToConnector, zmq4.POLLIN)
	poller.Add(r.AggregatorEventReceiveFromCluster, zmq4.POLLIN)

	topic := []byte{}
	event := [][]byte{}
	err := errors.New("")
	for {
		fmt.Println("Running AggregatorEventRoutine")
		sockets, _ := poller.Poll(-1)
		for _, socket := range sockets {
			switch currentSocket := socket.Socket; currentSocket {
			case r.AggregatorEventSendToCluster:
				topic, err = currentSocket.RecvBytes(0)
				if err != nil {
					panic(err)
				}
				if len(topic) <= 1 {
					break
				}
				event, err = currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}
				r.processEventSendToCluster(topic, event)

			case r.AggregatorEventReceiveFromConnector:
				topic, err = currentSocket.RecvBytes(0)
				if err != nil {
					panic(err)
				}
				if len(topic) <= 1 {
					break
				}
				event, err = currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}
				r.processEventReceiveFromConnector(topic, event)

			case r.AggregatorEventSendToConnector:
				topic, err = currentSocket.RecvBytes(0)
				if err != nil {
					panic(err)
				}
				if len(topic) <= 1 {
					break
				}
				event, err = currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}
				r.processEventSendToConnector(topic, event)

			case r.AggregatorEventReceiveFromCluster:
				topic, err = currentSocket.RecvBytes(0)
				if err != nil {
					panic(err)
				}
				if len(topic) <= 1 {
					break
				}
				event, err = currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}
				r.processEventReceiveFromCluster(topic, event)
			}
		}
	}
	fmt.Println("done")
}

func (r AggregatorEventRoutine) processEventSendToCluster(topic []byte, event [][]byte) {
	eventMessage, _ := message.DecodeEventMessage(event[0])
	go eventMessage.SendEventWith(r.AggregatorEventReceiveFromConnector)
}

func (r AggregatorEventRoutine) processEventReceiveFromCluster(topic []byte, event [][]byte) {
	eventMessage, _ := message.DecodeEventMessage(event[0])
	go eventMessage.SendEventWith(r.AggregatorEventSendToConnector)
}

func (r AggregatorEventRoutine) processEventSendToConnector(topic []byte, event [][]byte) {
	eventMessage, _ := message.DecodeEventMessage(event[0])
	go eventMessage.SendEventWith(r.AggregatorEventReceiveFromCluster)
}

func (r AggregatorEventRoutine) processEventReceiveFromConnector(topic []byte, event [][]byte) {
	eventMessage, _ := message.DecodeEventMessage(event[0])
	go eventMessage.SendEventWith(r.AggregatorEventSendToCluster)
}
