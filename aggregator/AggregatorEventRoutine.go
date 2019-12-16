package aggregator

import (
	"fmt"
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

func (r AggregatorEventRoutine) New(identity, aggregatorEventSendToClusterConnection, aggregatorEventReceiveFromConnectorConnection, aggregatorEventSendToConnectorConnection, aggregatorEventReceiveFromClusterConnection string) {
	r.identity = identity

	r.context, _ = zmq4.NewContext()
	r.aggregatorEventSendToClusterConnection = aggregatorEventSendToClusterConnection
	r.aggregatorEventSendToCluster, _ = r.context.NewSocket(zmq4.PUB)
	r.aggregatorEventSendToCluster.SetIdentity(r.identity)
	r.aggregatorEventSendToCluster.Connect(r.aggregatorEventSendToClusterConnection)
	fmt.Printf("aggregatorEventSendToCluster connect : " + aggregatorEventSendToClusterConnection)

	r.aggregatorEventReceiveFromClusterConnection = aggregatorEventReceiveFromClusterConnection
	r.aggregatorEventReceiveFromCluster, _ = r.context.NewSocket(zmq4.SUB)
	r.aggregatorEventReceiveFromCluster.SetIdentity(r.identity)
	r.aggregatorEventSendToCluster.Connect(r.aggregatorEventReceiveFromClusterConnection)
	fmt.Printf("aggregatorEventReceiveFromCluster connect : " + aggregatorEventReceiveFromClusterConnection)

	r.aggregatorEventSendToConnectorConnection = aggregatorEventSendToConnectorConnection
	r.aggregatorEventSendToConnector, _ = r.context.NewSocket(zmq4.PUB)
	r.aggregatorEventSendToConnector.SetIdentity(r.identity)
	r.aggregatorEventSendToCluster.Bind(r.aggregatorEventSendToConnectorConnection)
	fmt.Printf("aggregatorEventSendToConnector connect : " + aggregatorEventSendToConnectorConnection)

	r.aggregatorEventReceiveFromConnectorConnection = aggregatorEventReceiveFromConnectorConnection
	r.aggregatorEventReceiveFromConnector, _ = r.context.NewSocket(zmq4.SUB)
	r.aggregatorEventReceiveFromConnector.SetIdentity(r.identity)
	r.aggregatorEventSendToCluster.Bind(r.aggregatorEventReceiveFromConnectorConnection)
	fmt.Printf("aggregatorEventReceiveFromConnector connect : " + aggregatorEventReceiveFromConnectorConnection)
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

	for {

		sockets, _ := poller.Poll(-1)
		for _, socket := range sockets {

			switch currentSocket := socket.Socket; currentSocket {
			case r.aggregatorEventSendToCluster:

				event, err := currentSocket.RecvBytes()
				if err != nil {
					panic(err)
				}
				err = r.processEventSendToCluster(event)
				if err != nil {
					panic(err)
				}

			case r.aggregatorEventReceiveFromConnector:

				event, err := currentSocket.RecvBytes()
				if err != nil {
					panic(err)
				}
				err = r.processEventReceiveFromConnector(event)
				if err != nil {
					panic(err)
				}

			case r.aggregatorEventSendToConnector:

				event, err := currentSocket.RecvBytes()
				if err != nil {
					panic(err)
				}
				err = r.processEventSendToConnector(event)
				if err != nil {
					panic(err)
				}

			case r.aggregatorEventReceiveFromCluster:

				event, err := currentSocket.RecvBytes()
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
	eventMessage := message.EventMessage.decodeEvent(event[1])
	go eventMessage.sendEventWith(r.aggregatorEventReceiveFromConnector)
	return
}

func (r AggregatorEventRoutine) processEventReceiveFromCluster(event [][]byte) (err error) {
	eventMessage := message.EventMessage.decodeEvent(event[1])	
	go eventMessage.sendEventWith(r.aggregatorEventSendToConnector)
	return
}

func (r AggregatorEventRoutine) processEventSendToConnector(event [][]byte) (err error) {
	eventMessage := message.EventMessage.decodeEvent(event[1])
	go eventMessage.sendEventWith(r.aggregatorEventReceiveFromCluster)
	return
}

func (r AggregatorEventRoutine) processEventReceiveFromConnector(event [][]byte) (err error) {
	eventMessage := message.EventMessage.decodeEvent(event[1])
	go eventMessage.sendEventWith(r.aggregatorEventSendToCluster)
	return
}

