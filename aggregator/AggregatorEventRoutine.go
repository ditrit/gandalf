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

func (r AggregatorEventRoutine) New(identity, aggregatorEventSendToClusterConnection, aggregatorEventReceiveFromConnectorConnection, aggregatorEventSendToConnectorConnection, aggregatorEventReceiveFromClusterConnection string) err error {
	r.identity = identity

	r.context, _ := zmq4.NewContext()
	r.aggregatorEventSendToClusterConnection = aggregatorEventSendToClusterConnection
	r.aggregatorEventSendToCluster = r.context.NewXPub(aggregatorEventSendToClusterConnection)
	r.aggregatorEventSendToCluster.Identity(r.identity)
	fmt.Printf("aggregatorEventSendToCluster connect : " + aggregatorEventSendToClusterConnection)

	r.aggregatorEventReceiveFromClusterConnection = aggregatorEventReceiveFromClusterConnection
	r.aggregatorEventReceiveFromCluster = r.context.NewXSub(aggregatorEventReceiveFromClusterConnection)
	r.aggregatorEventReceiveFromCluster.Identity(r.identity)
	fmt.Printf("aggregatorEventReceiveFromCluster connect : " + aggregatorEventReceiveFromClusterConnection)

	r.aggregatorEventSendToConnectorConnection = aggregatorEventSendToConnectorConnection
	r.aggregatorEventSendToConnector = r.context.NewXPub(aggregatorEventSendToConnectorConnection)
	r.aggregatorEventSendToConnector.Identity(r.identity)
	fmt.Printf("aggregatorEventSendToConnector connect : " + aggregatorEventSendToConnectorConnection)

	r.aggregatorEventReceiveFromConnectorConnection = aggregatorEventReceiveFromConnectorConnection
	r.aggregatorEventReceiveFromConnector = r.context.NewSub(aggregatorEventReceiveFromConnectorConnection)
	r.aggregatorEventReceiveFromConnector.Identity(r.identity)
	fmt.Printf("aggregatorEventReceiveFromConnector connect : " + aggregatorEventReceiveFromConnectorConnection)
}

func (r AggregatorEventRoutine) close() err error {
	r.aggregatorEventSendToCluster.close()
	r.aggregatorEventReceiveFromConnector.close()
	r.aggregatorEventSendToConnector.close()
	r.aggregatorEventReceiveFromCluster.close()
	r.Context.close()
}

func (r AggregatorEventRoutine) run() err error {
	pi := zmq4.PollItems{
		zmq4.PollItem{Socket: aggregatorEventSendToCluster, Events: zmq4.POLLIN},
		zmq4.PollItem{Socket: aggregatorEventReceiveFromConnector, Events: zmq4.POLLIN},
		zmq4.PollItem{Socket: aggregatorEventSendToConnector, Events: zmq4.POLLIN},
		zmq4.PollItem{Socket: aggregatorEventReceiveFromCluster, Events: zmq4.POLLIN}}

	var event = [][]byte{}

	for {
		r.sendReadyCommand()

		_, _ = zmq4.Poll(pi, -1)

		switch {
		case pi[0].REvents&zmq4.POLLIN != 0:

			event, err := pi[0].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			err = r.processEventSendToCluster(event)
			if err != nil {
				panic(err)
			}

		case pi[1].REvents&zmq4.POLLIN != 0:

			event, err := pi[1].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			err = r.processEventReceiveFromConnector(event)
			if err != nil {
				panic(err)
			}

		case pi[2].REvents&zmq4.POLLIN != 0:

			event, err := pi[1].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			err = r.processEventSendToConnector(event)
			if err != nil {
				panic(err)
			}

		case pi[3].REvents&zmq4.POLLIN != 0:

			event, err := pi[1].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			err = r.processEventReceiveFromCluster(event)
			if err != nil {
				panic(err)
			}
		}
	}
	fmt.Println("done")
}

func (r AggregatorEventRoutine) processEventSendToCluster(event [][]byte) err error {
	eventMessage = EventMessage.decodeEvent(event[1])
	go eventMessage.sendEventWith(r.aggregatorEventReceiveFromConnector)

}

func (r AggregatorEventRoutine) processEventReceiveFromCluster(event [][]byte) err error {
	eventMessage = EventMessage.decodeEvent(event[1])	
	go eventMessage.sendEventWith(r.aggregatorEventSendToConnector)

}

func (r AggregatorEventRoutine) processEventSendToConnector(event [][]byte) err error {
	eventMessage = EventMessage.decodeEvent(event[1])
	go eventMessage.sendEventWith(r.aggregatorEventReceiveFromCluster)
}

func (r AggregatorEventRoutine) processEventReceiveFromConnector(event [][]byte) err error {
	eventMessage = EventMessage.decodeEvent(event[1])
	go eventMessage.sendEventWith(r.aggregatorEventSendToCluster)
}

