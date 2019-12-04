package aggregator

import (
	"fmt"

	zmq "github.com/zeromq/goczmq"
)

type AggregatorCommandRoutine struct {
	aggregatorCommandSendC2CL              zmq.Sock
	aggregatorCommandSendC2CLConnection    string
	aggregatorCommandReceiveC2CL           zmq.Sock
	aggregatorCommandReceiveC2CLConnection string
	aggregatorCommandSendC2CL              zmq.Sock
	aggregatorCommandSendC2CLConnection    string
	aggregatorCommandReceiveC2CL           zmq.Sock
	aggregatorCommandReceiveC2CLConnection string
	Identity                               string
}

func (r AggregatorCommandRoutine) new(identity, aggregatorCommandSendC2CLConnection, aggregatorCommandReceiveC2CLConnection, aggregatorCommandSendC2CLConnection, aggregatorCommandReceiveC2CLConnection string) {
	r.Identity = identity

	r.aggregatorCommandSendC2CLConnection = aggregatorCommandSendC2CLConnection
	r.aggregatorCommandSendC2CL = zmq.NewDealer(aggregatorCommandSendC2CLConnection)
	r.aggregatorCommandSendC2CL.Identity(w.identity)
	fmt.Printf("woraggregatorCommandSendC2CLkerCommandReceiveC2W connect : " + aggregatorCommandSendC2CLConnection)

	r.workerEventReceiveC2WConnection = aggregatorCommandReceiveC2CLConnection
	r.aggregatorCommandReceiveC2CL = zmq.NewSub(aggregatorCommandReceiveC2CLConnection)
	r.aggregatorCommandReceiveC2CL.Identity(w.identity)
	fmt.Printf("aggregatorCommandReceiveC2CL connect : " + aggregatorCommandReceiveC2CLConnection)

	r.aggregatorCommandSendC2CLConnection = aggregatorCommandSendC2CLConnection
	r.aggregatorCommandSendC2CL = zmq.NewSub(aggregatorCommandSendC2CLConnection)
	r.aggregatorCommandSendC2CL.Identity(w.identity)
	fmt.Printf("aggregatorCommandSendC2CL connect : " + aggregatorCommandSendC2CLConnection)

	r.aggregatorCommandReceiveC2CLConnection = aggregatorCommandReceiveC2CLConnection
	r.aggregatorCommandReceiveC2CL = zmq.NewSub(aggregatorCommandReceiveC2CLConnection)
	r.aggregatorCommandReceiveC2CL.Identity(w.identity)
	fmt.Printf("aggregatorCommandReceiveC2CL connect : " + aggregatorCommandReceiveC2CLConnection)
}

func (r AggregatorCommandRoutine) close() {
	r.aggregatorCommandSendC2CL.close()
	r.aggregatorCommandReceiveC2CL.close()
	r.aggregatorCommandSendC2CL.close()
	r.aggregatorCommandReceiveC2CL.close()
	r.Context.close()
}

func (r AggregatorCommandRoutine) run() {
	pi := zmq.PollItems{
		zmq.PollItem{Socket: aggregatorCommandSendC2CL, Events: zmq.POLLIN},
		zmq.PollItem{Socket: aggregatorCommandReceiveC2CL, Events: zmq.POLLIN},
		zmq.PollItem{Socket: aggregatorCommandSendC2CL, Events: zmq.POLLIN},
		zmq.PollItem{Socket: aggregatorCommandReceiveC2CL, Events: zmq.POLLIN}}

	var command = [][]byte{}

	for {
		r.sendReadyCommand()

		_, _ = zmq.Poll(pi, -1)

		switch {
		case pi[0].REvents&zmq.POLLIN != 0:

			command, err := pi[0].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			//PROCESS SEND COMMAND TO CLUSTER
			err = routerSock.SendMessage(msg)
			if err != nil {
				panic(err)
			}

		case pi[1].REvents&zmq.POLLIN != 0:

			command, err := pi[1].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			//PROCESS RECEIVE COMMAND TO CLUSTER
			err = routerSock.SendMessage(msg)
			if err != nil {
				panic(err)
			}

		case pi[2].REvents&zmq.POLLIN != 0:

			command, err := pi[1].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			//PROCESS SEND COMMAND TO CONNECTOR
			err = routerSock.SendMessage(msg)
			if err != nil {
				panic(err)
			}

		case pi[3].REvents&zmq.POLLIN != 0:

			command, err := pi[1].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			//PROCESS RECEIVE COMMAND TO CONNECTOR
			err = routerSock.SendMessage(msg)
			if err != nil {
				panic(err)
			}
		}
	}

	fmt.Println("done")

}
