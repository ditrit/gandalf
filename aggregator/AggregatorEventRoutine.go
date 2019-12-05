package aggregator

import (
	"fmt"

	zmq "github.com/zeromq/goczmq"
)

type AggregatorEventRoutine struct {
	aggregatorEventSendC2CL              zmq.Sock
	aggregatorEventSendC2CLConnection    string
	aggregatorEventReceiveC2CL           zmq.Sock
	aggregatorEventReceiveC2CLConnection string
	aggregatorEventSendCL2C              zmq.Sock
	aggregatorEventSendCL2CConnection    string
	aggregatorEventReceiveCL2C           zmq.Sock
	aggregatorEventReceiveCL2CConnection string
	identity                             string
}

func (r AggregatorEventRoutine) New(identity, aggregatorEventSendC2CLConnection, aggregatorEventReceiveC2CLConnection, aggregatorEventSendCL2CConnection, aggregatorEventReceiveCL2CConnection string) err error {
	r.identity = identity

	r.aggregatorEventSendC2CLConnection = aggregatorEventSendC2CLConnection
	r.aggregatorEventSendC2CL = zmq.NewDealer(aggregatorEventSendC2CLConnection)
	r.aggregatorEventSendC2CL.Identity(r.identity)
	fmt.Printf("aggregatorEventSendC2CL connect : " + aggregatorEventSendC2CLConnection)

	r.aggregatorEventReceiveC2CLConnection = aggregatorEventReceiveC2CLConnection
	r.aggregatorEventReceiveC2CL = zmq.NewSub(aggregatorEventReceiveC2CLConnection)
	r.aggregatorEventReceiveC2CL.Identity(r.identity)
	fmt.Printf("aggregatorEventReceiveC2CL connect : " + aggregatorEventReceiveC2CLConnection)

	r.aggregatorEventSendCL2CConnection = aggregatorEventSendCL2CConnection
	r.aggregatorEventSendCL2C = zmq.NewSub(aggregatorEventSendCL2CConnection)
	r.aggregatorEventSendCL2C.Identity(r.identity)
	fmt.Printf("aggregatorEventSendCL2C connect : " + aggregatorEventSendCL2CConnection)

	r.aggregatorEventReceiveCL2CConnection = aggregatorEventReceiveCL2CConnection
	r.aggregatorEventReceiveCL2C = zmq.NewSub(aggregatorEventReceiveCL2CConnection)
	r.aggregatorEventReceiveCL2C.Identity(r.identity)
	fmt.Printf("aggregatorEventReceiveCL2C connect : " + aggregatorEventReceiveCL2CConnection)
}

func (r AggregatorEventRoutine) close() err error {
	r.aggregatorEventSendC2CL.close()
	r.aggregatorEventReceiveC2CL.close()
	r.aggregatorEventSendCL2C.close()
	r.aggregatorEventReceiveCL2C.close()
	r.Context.close()
}

func (r AggregatorEventRoutine) run() err error {
	pi := zmq.PollItems{
		zmq.PollItem{Socket: aggregatorEventSendC2CL, Events: zmq.POLLIN},
		zmq.PollItem{Socket: aggregatorEventReceiveC2CL, Events: zmq.POLLIN},
		zmq.PollItem{Socket: aggregatorEventSendCL2C, Events: zmq.POLLIN},
		zmq.PollItem{Socket: aggregatorEventReceiveCL2C, Events: zmq.POLLIN}}

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
	event = r.updateHeaderEventSendC2CL(command)
}

func (r AggregatorEventRoutine) updateHeaderEventSendC2CL(event [][]byte) err error {

}

func (r AggregatorEventRoutine) processEventReceiveC2CL(event [][]byte) err error {
	event = r.updateHeaderEventReceiveC2CL(event)
}

func (r AggregatorEventRoutine) updateHeaderEventReceiveC2CL(event [][]byte) err error {

}

func (r AggregatorEventRoutine) processEventSendC2CL(event [][]byte) err error {
	event = r.updateHeaderEventSendC2CL(event)
}

func (r AggregatorEventRoutine) updateHeaderEventSendC2CL(event [][]byte) err error {

}

func (r AggregatorEventRoutine) processEventReceiveC2CL(event [][]byte) err error {
	event = r.updateHeaderEventReceiveC2CL(event)
}

func (r AggregatorEventRoutine) updateHeaderEventReceiveC2CL(event [][]byte) err error {

}
