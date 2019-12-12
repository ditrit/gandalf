package aggregator

import (
	"fmt"

	zmq "github.com/zeromq/goczmq"
)

type AggregatorEventRoutine struct {
	aggregatorEventSendToCluster              zmq.Sock
	aggregatorEventSendToClusterConnection    string
	aggregatorEventReceiveFromConnector           zmq.Sock
	aggregatorEventReceiveFromConnectorConnection string
	aggregatorEventSendToConnector              zmq.Sock
	aggregatorEventSendToConnectorConnection    string
	aggregatorEventReceiveFromCluster           zmq.Sock
	aggregatorEventReceiveFromClusterConnection string
	identity                             string
}

func (r AggregatorEventRoutine) New(identity, aggregatorEventSendToClusterConnection, aggregatorEventReceiveFromConnectorConnection, aggregatorEventSendToConnectorConnection, aggregatorEventReceiveFromClusterConnection string) err error {
	r.identity = identity

	r.aggregatorEventSendToClusterConnection = aggregatorEventSendToClusterConnection
	r.aggregatorEventSendToCluster = zmq.NewDealer(aggregatorEventSendToClusterConnection)
	r.aggregatorEventSendToCluster.Identity(r.identity)
	fmt.Printf("aggregatorEventSendToCluster connect : " + aggregatorEventSendToClusterConnection)

	r.aggregatorEventReceiveFromConnectorConnection = aggregatorEventReceiveFromConnectorConnection
	r.aggregatorEventReceiveFromConnector = zmq.NewSub(aggregatorEventReceiveFromConnectorConnection)
	r.aggregatorEventReceiveFromConnector.Identity(r.identity)
	fmt.Printf("aggregatorEventReceiveFromConnector connect : " + aggregatorEventReceiveFromConnectorConnection)

	r.aggregatorEventSendToConnectorConnection = aggregatorEventSendToConnectorConnection
	r.aggregatorEventSendToConnector = zmq.NewSub(aggregatorEventSendToConnectorConnection)
	r.aggregatorEventSendToConnector.Identity(r.identity)
	fmt.Printf("aggregatorEventSendToConnector connect : " + aggregatorEventSendToConnectorConnection)

	r.aggregatorEventReceiveFromClusterConnection = aggregatorEventReceiveFromClusterConnection
	r.aggregatorEventReceiveFromCluster = zmq.NewSub(aggregatorEventReceiveFromClusterConnection)
	r.aggregatorEventReceiveFromCluster.Identity(r.identity)
	fmt.Printf("aggregatorEventReceiveFromCluster connect : " + aggregatorEventReceiveFromClusterConnection)
}

func (r AggregatorEventRoutine) close() err error {
	r.aggregatorEventSendToCluster.close()
	r.aggregatorEventReceiveFromConnector.close()
	r.aggregatorEventSendToConnector.close()
	r.aggregatorEventReceiveFromCluster.close()
	r.Context.close()
}

func (r AggregatorEventRoutine) run() err error {
	pi := zmq.PollItems{
		zmq.PollItem{Socket: aggregatorEventSendToCluster, Events: zmq.POLLIN},
		zmq.PollItem{Socket: aggregatorEventReceiveFromConnector, Events: zmq.POLLIN},
		zmq.PollItem{Socket: aggregatorEventSendToConnector, Events: zmq.POLLIN},
		zmq.PollItem{Socket: aggregatorEventReceiveFromCluster, Events: zmq.POLLIN}}

	var event = [][]byte{}

	for {
		r.sendReadyCommand()

		_, _ = zmq.Poll(pi, -1)

		switch {
		case pi[0].REvents&zmq.POLLIN != 0:

			event, err := pi[0].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			err = r.processEventSendC2CL(event)
			if err != nil {
				panic(err)
			}

		case pi[1].REvents&zmq.POLLIN != 0:

			event, err := pi[1].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			err = r.processEventReceiveC2CL(event)
			if err != nil {
				panic(err)
			}

		case pi[2].REvents&zmq.POLLIN != 0:

			event, err := pi[1].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			err = r.processEventSendC2CL(event)
			if err != nil {
				panic(err)
			}

		case pi[3].REvents&zmq.POLLIN != 0:

			event, err := pi[1].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			err = r.processEventReceiveC2CL(event)
			if err != nil {
				panic(err)
			}
		}
	}
	fmt.Println("done")
}

func (r AggregatorEventRoutine) processEventSendC2CL(event [][]byte) err error {
	eventMessage = EventMessage.decodeEvent(event[1])
	go eventMessage.sendEventWith(r.aggregatorEventSendToCluster)

}

func (r AggregatorEventRoutine) processEventReceiveC2CL(event [][]byte) err error {
	eventMessage = EventMessage.decodeEvent(event[1])	
	go eventMessage.sendEventWith(r.aggregatorEventReceiveFromConnector)

}

func (r AggregatorEventRoutine) processEventSendC2CL(event [][]byte) err error {
	eventMessage = EventMessage.decodeEvent(event[1])
	go eventMessage.sendEventWith(r.aggregatorEventSendToConnector)
}

func (r AggregatorEventRoutine) processEventReceiveC2CL(event [][]byte) err error {
	eventMessage = EventMessage.decodeEvent(event[1])
	go eventMessage.sendEventWith(r.aggregatorEventReceiveFromCluster)
}

