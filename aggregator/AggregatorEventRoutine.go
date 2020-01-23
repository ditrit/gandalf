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
	aggregatorEventSendToClusterConnections       []string
	aggregatorEventReceiveFromConnector           *zmq4.Socket
	aggregatorEventReceiveFromConnectorConnection string
	aggregatorEventSendToConnector                *zmq4.Socket
	aggregatorEventSendToConnectorConnection      string
	aggregatorEventReceiveFromCluster             *zmq4.Socket
	aggregatorEventReceiveFromClusterConnections  []string
	identity                                      string
	tenant                                        string
}

func NewAggregatorEventRoutine(identity, tenant, aggregatorEventReceiveFromConnectorConnection, aggregatorEventSendToConnectorConnection string, aggregatorEventSendToClusterConnections, aggregatorEventReceiveFromClusterConnections []string) (aggregatorEventRoutine *AggregatorEventRoutine) {
	aggregatorEventRoutine = new(AggregatorEventRoutine)

	aggregatorEventRoutine.identity = identity
	aggregatorEventRoutine.tenant = tenant

	aggregatorEventRoutine.context, _ = zmq4.NewContext()
	aggregatorEventRoutine.aggregatorEventSendToClusterConnections = aggregatorEventSendToClusterConnections
	aggregatorEventRoutine.aggregatorEventSendToCluster, _ = aggregatorEventRoutine.context.NewSocket(zmq4.XPUB)
	aggregatorEventRoutine.aggregatorEventSendToCluster.SetIdentity(aggregatorEventRoutine.identity)
	//aggregatorEventRoutine.aggregatorEventSendToCluster.Connect(aggregatorEventRoutine.aggregatorEventSendToClusterConnections)
	//fmt.Println("aggregatorEventSendToCluster connect : " + aggregatorEventSendToClusterConnections)
	for _, connection := range aggregatorEventRoutine.aggregatorEventSendToClusterConnections {
		aggregatorEventRoutine.aggregatorEventSendToCluster.Connect(connection)
		fmt.Println("aggregatorEventSendToCluster connect : " + connection)
	}

	aggregatorEventRoutine.aggregatorEventReceiveFromClusterConnections = aggregatorEventReceiveFromClusterConnections
	aggregatorEventRoutine.aggregatorEventReceiveFromCluster, _ = aggregatorEventRoutine.context.NewSocket(zmq4.XSUB)
	aggregatorEventRoutine.aggregatorEventReceiveFromCluster.SetIdentity(aggregatorEventRoutine.identity)
	//aggregatorEventRoutine.aggregatorEventSendToCluster.Connect(aggregatorEventRoutine.aggregatorEventReceiveFromClusterConnections)
	//fmt.Println("aggregatorEventReceiveFromCluster connect : " + aggregatorEventReceiveFromClusterConnections)
	for _, connection := range aggregatorEventRoutine.aggregatorEventReceiveFromClusterConnections {
		aggregatorEventRoutine.aggregatorEventReceiveFromCluster.Connect(connection)
		fmt.Println("aggregatorEventReceiveFromCluster connect : " + connection)
	}
	aggregatorEventRoutine.aggregatorEventReceiveFromCluster.SendBytes([]byte{0x01}, 0) //SUBSCRIBE ALL

	aggregatorEventRoutine.aggregatorEventSendToConnectorConnection = aggregatorEventSendToConnectorConnection
	aggregatorEventRoutine.aggregatorEventSendToConnector, _ = aggregatorEventRoutine.context.NewSocket(zmq4.XPUB)
	aggregatorEventRoutine.aggregatorEventSendToConnector.SetIdentity(aggregatorEventRoutine.identity)
	aggregatorEventRoutine.aggregatorEventSendToConnector.Bind(aggregatorEventRoutine.aggregatorEventSendToConnectorConnection)
	fmt.Println("aggregatorEventSendToConnector Bind : " + aggregatorEventSendToConnectorConnection)

	aggregatorEventRoutine.aggregatorEventReceiveFromConnectorConnection = aggregatorEventReceiveFromConnectorConnection
	aggregatorEventRoutine.aggregatorEventReceiveFromConnector, _ = aggregatorEventRoutine.context.NewSocket(zmq4.XSUB)
	aggregatorEventRoutine.aggregatorEventReceiveFromConnector.SetIdentity(aggregatorEventRoutine.identity)
	aggregatorEventRoutine.aggregatorEventReceiveFromConnector.Bind(aggregatorEventRoutine.aggregatorEventReceiveFromConnectorConnection)
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

	event := [][]byte{}
	err := errors.New("")
	for {
		fmt.Println("Running AggregatorEventRoutine")
		sockets, _ := poller.Poll(-1)
		for _, socket := range sockets {
			switch currentSocket := socket.Socket; currentSocket {
			case r.aggregatorEventSendToCluster:
				fmt.Println("Send Cluster")
				event, err = currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}
				r.processEventSendToCluster(event)

			case r.aggregatorEventReceiveFromConnector:
				fmt.Println("Receive Connector")
				event, err = currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}
				r.processEventReceiveFromConnector(event)

			case r.aggregatorEventSendToConnector:
				fmt.Println("Send Connector")
				event, err = currentSocket.RecvMessageBytes(0)
				fmt.Println(event)
				if err != nil {
					panic(err)
				}
				r.processEventSendToConnector(event)

			case r.aggregatorEventReceiveFromCluster:
				fmt.Println("Receive Cluster")
				event, err = currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}
				r.processEventReceiveFromCluster(event)
			}
		}
	}
	fmt.Println("done")
}

func (r AggregatorEventRoutine) processEventSendToCluster(event [][]byte) {
	if len(event) > 1 {
		eventMessage, _ := message.DecodeEventMessage(event[1])
		go eventMessage.SendMessageWith(r.aggregatorEventReceiveFromConnector)
	}
}

func (r AggregatorEventRoutine) processEventReceiveFromCluster(event [][]byte) {
	fmt.Println(event)
	eventMessage, _ := message.DecodeEventMessage(event[1])
	fmt.Println(eventMessage)
	go eventMessage.SendMessageWith(r.aggregatorEventSendToConnector)
}

func (r AggregatorEventRoutine) processEventSendToConnector(event [][]byte) {
	if len(event) > 1 {
		eventMessage, _ := message.DecodeEventMessage(event[0])
		go eventMessage.SendMessageWith(r.aggregatorEventReceiveFromCluster)
	}
}

func (r AggregatorEventRoutine) processEventReceiveFromConnector(event [][]byte) {
	eventMessage, _ := message.DecodeEventMessage(event[1])
	eventMessage.Tenant = r.tenant
	go eventMessage.SendMessageWith(r.aggregatorEventSendToCluster)
}
