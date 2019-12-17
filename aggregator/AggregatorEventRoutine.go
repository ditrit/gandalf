package aggregator

import (
	"fmt"
	"errors"
	"gandalfgo/message"
	"github.com/pebbe/zmq4"
)

type AggregatorEventRoutine struct {
	context									 		*zmq4.Context
	aggregatorEventSendToCluster              		*zmq4.Socket
	aggregatorEventSendToClusterConnection    		string
	aggregatorEventReceiveFromConnector           	*zmq4.Socket
	aggregatorEventReceiveFromConnectorConnection 	string
	aggregatorEventSendToConnector              	*zmq4.Socket
	aggregatorEventSendToConnectorConnection    	string
	aggregatorEventReceiveFromCluster           	*zmq4.Socket
	aggregatorEventReceiveFromClusterConnection 	string
	identity                             			string
}

func NewAggregatorEventRoutine(identity, aggregatorEventSendToClusterConnection, aggregatorEventReceiveFromConnectorConnection, aggregatorEventSendToConnectorConnection, aggregatorEventReceiveFromClusterConnection string) (aggregatorEventRoutine *AggregatorEventRoutine) {
	aggregatorEventRoutine = new(AggregatorEventRoutine)

	aggregatorEventRoutine.identity = identity

	aggregatorEventRoutine.context, _ = zmq4.NewContext()
	aggregatorEventRoutine.aggregatorEventSendToClusterConnection = aggregatorEventSendToClusterConnection
	aggregatorEventRoutine.aggregatorEventSendToCluster, _ = aggregatorEventRoutine.context.NewSocket(zmq4.PUB)
	aggregatorEventRoutine.aggregatorEventSendToCluster.SetIdentity(aggregatorEventRoutine.identity)
	aggregatorEventRoutine.aggregatorEventSendToCluster.Connect(aggregatorEventRoutine.aggregatorEventSendToClusterConnection)
	fmt.Printf("aggregatorEventSendToCluster connect : " + aggregatorEventSendToClusterConnection)

	aggregatorEventRoutine.aggregatorEventReceiveFromClusterConnection = aggregatorEventReceiveFromClusterConnection
	aggregatorEventRoutine.aggregatorEventReceiveFromCluster, _ = aggregatorEventRoutine.context.NewSocket(zmq4.SUB)
	aggregatorEventRoutine.aggregatorEventReceiveFromCluster.SetIdentity(aggregatorEventRoutine.identity)
	aggregatorEventRoutine.aggregatorEventSendToCluster.Connect(aggregatorEventRoutine.aggregatorEventReceiveFromClusterConnection)
	fmt.Printf("aggregatorEventReceiveFromCluster connect : " + aggregatorEventReceiveFromClusterConnection)

	aggregatorEventRoutine.aggregatorEventSendToConnectorConnection = aggregatorEventSendToConnectorConnection
	aggregatorEventRoutine.aggregatorEventSendToConnector, _ = aggregatorEventRoutine.context.NewSocket(zmq4.PUB)
	aggregatorEventRoutine.aggregatorEventSendToConnector.SetIdentity(aggregatorEventRoutine.identity)
	aggregatorEventRoutine.aggregatorEventSendToCluster.Bind(aggregatorEventRoutine.aggregatorEventSendToConnectorConnection)
	fmt.Printf("aggregatorEventSendToConnector connect : " + aggregatorEventSendToConnectorConnection)

	aggregatorEventRoutine.aggregatorEventReceiveFromConnectorConnection = aggregatorEventReceiveFromConnectorConnection
	aggregatorEventRoutine.aggregatorEventReceiveFromConnector, _ = aggregatorEventRoutine.context.NewSocket(zmq4.SUB)
	aggregatorEventRoutine.aggregatorEventReceiveFromConnector.SetIdentity(aggregatorEventRoutine.identity)
	aggregatorEventRoutine.aggregatorEventSendToCluster.Bind(aggregatorEventRoutine.aggregatorEventReceiveFromConnectorConnection)
	fmt.Printf("aggregatorEventReceiveFromConnector connect : " + aggregatorEventReceiveFromConnectorConnection)

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

	event := [][]byte{}
	err := errors.New("")
	for {

		sockets, _ := poller.Poll(-1)
		for _, socket := range sockets {

			switch currentSocket := socket.Socket; currentSocket {
			case r.aggregatorEventSendToCluster:

				event, err = currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}
				err = r.processEventSendToCluster(event)
				if err != nil {
					panic(err)
				}

			case r.aggregatorEventReceiveFromConnector:

				event, err = currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}
				err = r.processEventReceiveFromConnector(event)
				if err != nil {
					panic(err)
				}

			case r.aggregatorEventSendToConnector:

				event, err = currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}
				err = r.processEventSendToConnector(event)
				if err != nil {
					panic(err)
				}

			case r.aggregatorEventReceiveFromCluster:

				event, err = currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}
				err = r.processEventReceiveFromCluster(event)
				if err != nil {
					panic(err)
				}
			}
		}
	}
	fmt.Println("done")
}

func (r AggregatorEventRoutine) processEventSendToCluster(event [][]byte) (err error) {
	eventMessage, err := message.DecodeEventMessage(event[1])
	go eventMessage.SendEventWith(r.aggregatorEventReceiveFromConnector)
	return
}

func (r AggregatorEventRoutine) processEventReceiveFromCluster(event [][]byte) (err error) {
	eventMessage, err := message.DecodeEventMessage(event[1])	
	go eventMessage.SendEventWith(r.aggregatorEventSendToConnector)
	return
}

func (r AggregatorEventRoutine) processEventSendToConnector(event [][]byte) (err error) {
	eventMessage, err := message.DecodeEventMessage(event[1])
	go eventMessage.SendEventWith(r.aggregatorEventReceiveFromCluster)
	return
}

func (r AggregatorEventRoutine) processEventReceiveFromConnector(event [][]byte) (err error) {
	eventMessage, err := message.DecodeEventMessage(event[1])
	go eventMessage.SendEventWith(r.aggregatorEventSendToCluster)
	return
}

