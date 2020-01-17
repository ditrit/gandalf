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

	event := [][]byte{}
	err := errors.New("")
	for {
		fmt.Println("Running AggregatorEventRoutine")
		sockets, _ := poller.Poll(-1)
		for _, socket := range sockets {
			switch currentSocket := socket.Socket; currentSocket {
			case r.AggregatorEventSendToCluster:
				fmt.Println("Send Cluster")
				event, err = currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}
				r.processEventSendToCluster(event)

			case r.AggregatorEventReceiveFromConnector:
				fmt.Println("Receive connector")
				event, err = currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}
				r.processEventReceiveFromConnector(event)

			case r.AggregatorEventSendToConnector:
				fmt.Println("send connector")
				event, err = currentSocket.RecvMessageBytes(0)
				fmt.Println(event)
				if err != nil {
					panic(err)
				}
				r.processEventSendToConnector(event)

			case r.AggregatorEventReceiveFromCluster:
				fmt.Println("receive cluster")
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
	fmt.Println("processEventSendToCluster")
	fmt.Println(event)

	if len(event) == 1 {
		topic := event[0]
		fmt.Println("SUB")
		fmt.Println("AggregatorEventReceiveFromConnector")
		fmt.Println(topic)
		fmt.Println(string(topic))
		//r.AggregatorEventReceiveFromConnector.SetSubscribe(string(topic))
		//go message.SendSubscribeTopic(r.AggregatorEventReceiveFromConnector, topic)
	} else {
		eventMessage, _ := message.DecodeEventMessage(event[1])
		go eventMessage.SendMessageWith(r.AggregatorEventReceiveFromConnector)
	}
}

func (r AggregatorEventRoutine) processEventReceiveFromCluster(event [][]byte) {
	fmt.Println(event)
	eventMessage, _ := message.DecodeEventMessage(event[1])
	fmt.Println(eventMessage)
	go eventMessage.SendMessageWith(r.AggregatorEventSendToConnector)
}

func (r AggregatorEventRoutine) processEventSendToConnector(event [][]byte) {
	fmt.Println("processEventSendToConnector")
	fmt.Println(event)

	if len(event) == 1 {
		topic := event[0]
		fmt.Println("SUB")
		fmt.Println("AggregatorEventReceiveFromCluster")
		fmt.Println(topic)
		fmt.Println(string(topic))
		//r.AggregatorEventReceiveFromCluster.SetSubscribe(string(topic))
		//go message.SendSubscribeTopic(r.AggregatorEventReceiveFromCluster, topic)
	} else {
		eventMessage, _ := message.DecodeEventMessage(event[0])
		fmt.Println("SEN")
		fmt.Println(event)
		fmt.Println(eventMessage)
		go eventMessage.SendMessageWith(r.AggregatorEventReceiveFromCluster)
	}
}

func (r AggregatorEventRoutine) processEventReceiveFromConnector(event [][]byte) {
	fmt.Println(event)
	fmt.Println(event[0])
	fmt.Println(event[1])
	eventMessage, _ := message.DecodeEventMessage(event[1])
	go eventMessage.SendMessageWith(r.AggregatorEventSendToCluster)
}
