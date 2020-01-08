package aggregator

import (
	"errors"
	"fmt"
	"gandalf-go/message"

	"github.com/pebbe/zmq4"
)

type AggregatorEventRoutine struct {
	context                                       *zmq4.Context
	aggregatorEventSendToCluster                  *zmq4.Socket
	aggregatorEventSendToClusterConnection        string
	aggregatorEventReceiveFromConnector           *zmq4.Socket
	aggregatorEventReceiveFromConnectorConnection string
	aggregatorEventSendToConnector                *zmq4.Socket
	aggregatorEventSendToConnectorConnection      string
	aggregatorEventReceiveFromCluster             *zmq4.Socket
	aggregatorEventReceiveFromClusterConnection   string
	identity                                      string
}

func NewAggregatorEventRoutine(identity, aggregatorEventSendToClusterConnection, aggregatorEventReceiveFromConnectorConnection, aggregatorEventSendToConnectorConnection, aggregatorEventReceiveFromClusterConnection string) (aggregatorEventRoutine *AggregatorEventRoutine) {
	aggregatorEventRoutine = new(AggregatorEventRoutine)

	aggregatorEventRoutine.identity = identity

	aggregatorEventRoutine.context, _ = zmq4.NewContext()
	aggregatorEventRoutine.aggregatorEventSendToClusterConnection = aggregatorEventSendToClusterConnection
	aggregatorEventRoutine.aggregatorEventSendToCluster, _ = aggregatorEventRoutine.context.NewSocket(zmq4.XPUB)
	aggregatorEventRoutine.aggregatorEventSendToCluster.SetIdentity(aggregatorEventRoutine.identity)
	aggregatorEventRoutine.aggregatorEventSendToCluster.Connect(aggregatorEventRoutine.aggregatorEventSendToClusterConnection)
	fmt.Println("aggregatorEventSendToCluster connect : " + aggregatorEventSendToClusterConnection)

	aggregatorEventRoutine.aggregatorEventReceiveFromClusterConnection = aggregatorEventReceiveFromClusterConnection
	aggregatorEventRoutine.aggregatorEventReceiveFromCluster, _ = aggregatorEventRoutine.context.NewSocket(zmq4.XSUB)
	aggregatorEventRoutine.aggregatorEventReceiveFromCluster.SetIdentity(aggregatorEventRoutine.identity)
	aggregatorEventRoutine.aggregatorEventSendToCluster.Connect(aggregatorEventRoutine.aggregatorEventReceiveFromClusterConnection)
	fmt.Println("aggregatorEventReceiveFromCluster connect : " + aggregatorEventReceiveFromClusterConnection)
	aggregatorEventRoutine.aggregatorEventReceiveFromCluster.SendBytes([]byte{0x01}, 0) //SUBSCRIBE ALL

	aggregatorEventRoutine.aggregatorEventSendToConnectorConnection = aggregatorEventSendToConnectorConnection
	aggregatorEventRoutine.aggregatorEventSendToConnector, _ = aggregatorEventRoutine.context.NewSocket(zmq4.XPUB)
	aggregatorEventRoutine.aggregatorEventSendToConnector.SetIdentity(aggregatorEventRoutine.identity)
	aggregatorEventRoutine.aggregatorEventSendToCluster.Bind(aggregatorEventRoutine.aggregatorEventSendToConnectorConnection)
	fmt.Println("aggregatorEventSendToConnector Bind : " + aggregatorEventSendToConnectorConnection)

	aggregatorEventRoutine.aggregatorEventReceiveFromConnectorConnection = aggregatorEventReceiveFromConnectorConnection
	aggregatorEventRoutine.aggregatorEventReceiveFromConnector, _ = aggregatorEventRoutine.context.NewSocket(zmq4.XSUB)
	aggregatorEventRoutine.aggregatorEventReceiveFromConnector.SetIdentity(aggregatorEventRoutine.identity)
	aggregatorEventRoutine.aggregatorEventSendToCluster.Bind(aggregatorEventRoutine.aggregatorEventReceiveFromConnectorConnection)
	fmt.Println("aggregatorEventReceiveFromConnector Bind : " + aggregatorEventReceiveFromConnectorConnection)
	aggregatorEventRoutine.aggregatorEventReceiveFromConnector.SendBytes([]byte{0x01}, 0) //SUBSCRIBE ALL

	return
}

func (r AggregatorEventRoutine) close() {
	r.aggregatorEventSendToCluster.Close()
	r.aggregatorEventReceiveFromConnector.Close()
	r.aggregatorEventSendToConnector.Close()
	r.aggregatorEventReceiveFromCluster.Close()
	r.context.Term()
}

func (r AggregatorEventRoutine) run() {
	poller := zmq4.NewPoller()
	poller.Add(r.aggregatorEventSendToCluster, zmq4.POLLIN)
	poller.Add(r.aggregatorEventReceiveFromConnector, zmq4.POLLIN)
	poller.Add(r.aggregatorEventSendToConnector, zmq4.POLLIN)
	poller.Add(r.aggregatorEventReceiveFromCluster, zmq4.POLLIN)

	topic := []byte{}
	event := [][]byte{}
	err := errors.New("")
	for {
		fmt.Println("Running AggregatorEventRoutine")
		sockets, _ := poller.Poll(-1)
		for _, socket := range sockets {

			switch currentSocket := socket.Socket; currentSocket {
			case r.aggregatorEventSendToCluster:
				fmt.Println("TOTO0")
				topic, err = currentSocket.RecvBytes(0)
				fmt.Println(string(topic))
				if err != nil {
					panic(err)
				}
				event, err = currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}
				r.processEventSendToCluster(topic, event)

			case r.aggregatorEventReceiveFromConnector:
				fmt.Println("TOTO1")
				topic, err = currentSocket.RecvBytes(0)
				fmt.Println(string(topic))
				if err != nil {
					panic(err)
				}
				event, err = currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}
				r.processEventReceiveFromConnector(topic, event)

			case r.aggregatorEventSendToConnector:
				fmt.Println("TOTO2")
				topic, err = currentSocket.RecvBytes(0)
				fmt.Println(string(topic))
				if err != nil {
					panic(err)
				}
				event, err = currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}
				r.processEventSendToConnector(topic, event)

			case r.aggregatorEventReceiveFromCluster:
				fmt.Println("TOTO3")
				topic, err = currentSocket.RecvBytes(0)
				fmt.Println(string(topic))
				if err != nil {
					panic(err)
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
	fmt.Println("OH NO")
	eventMessage, _ := message.DecodeEventMessage(event[0])
	go eventMessage.SendEventWith(r.aggregatorEventReceiveFromConnector)
}

func (r AggregatorEventRoutine) processEventReceiveFromCluster(topic []byte, event [][]byte) {
	eventMessage, _ := message.DecodeEventMessage(event[0])
	go eventMessage.SendEventWith(r.aggregatorEventSendToConnector)
}

func (r AggregatorEventRoutine) processEventSendToConnector(topic []byte, event [][]byte) {
	eventMessage, _ := message.DecodeEventMessage(event[0])
	go eventMessage.SendEventWith(r.aggregatorEventReceiveFromCluster)
}

func (r AggregatorEventRoutine) processEventReceiveFromConnector(topic []byte, event [][]byte) {
	eventMessage, _ := message.DecodeEventMessage(event[0])
	go eventMessage.SendEventWith(r.aggregatorEventSendToCluster)
}
