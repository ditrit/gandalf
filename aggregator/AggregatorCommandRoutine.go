package aggregator

import (
	"fmt"
    "message"
	zmq "github.com/zeromq/goczmq"
)

type AggregatorCommandRoutine struct {
	aggregatorCommandSendC2CL              zmq.Sock
	aggregatorCommandSendC2CLConnections   *string
	aggregatorCommandReceiveC2CL           zmq.Sock
	aggregatorCommandReceiveC2CLConnection string
	aggregatorCommandSendCL2C              zmq.Sock
	aggregatorCommandSendCL2CConnections   *string
	aggregatorCommandReceiveCL2C           zmq.Sock
	aggregatorCommandReceiveCL2CConnection string
	identity                               string
}

func (r AggregatorCommandRoutine) New(identity, aggregatorCommandSendC2CLConnections, aggregatorCommandReceiveC2CLConnection, aggregatorCommandSendCL2CConnections, aggregatorCommandReceiveCL2CConnection string) err error {
	r.identity = identity

	r.aggregatorCommandSendC2CLConnections = aggregatorCommandSendC2CLConnections
	r.aggregatorCommandSendC2CL = zmq.NewDealer(aggregatorCommandSendC2CLConnections)
	r.aggregatorCommandSendC2CL.Identity(r.identity)
	fmt.Printf("aggregatorCommandSendC2CL connect : " + aggregatorCommandSendC2CLConnections)

	r.workerEventReceiveC2WConnection = aggregatorCommandReceiveC2CLConnection
	r.aggregatorCommandReceiveC2CL = zmq.NewSub(aggregatorCommandReceiveC2CLConnection)
	r.aggregatorCommandReceiveC2CL.Identity(r.identity)
	fmt.Printf("aggregatorCommandReceiveC2CL connect : " + aggregatorCommandReceiveC2CLConnection)

	r.aggregatorCommandSendCL2CConnections = aggregatorCommandSendCL2CConnections
	r.aggregatorCommandSendC2CL = zmq.NewSub(aggregatorCommandSendCL2CConnections)
	r.aggregatorCommandSendC2CL.Identity(r.identity)
	fmt.Printf("aggregatorCommandSendC2CL connect : " + aggregatorCommandSendCL2CConnections)

	r.aggregatorCommandReceiveC2CLConnection = aggregatorCommandReceiveC2CLConnection
	r.aggregatorCommandReceiveC2CL = zmq.NewSub(aggregatorCommandReceiveC2CLConnection)
	r.aggregatorCommandReceiveC2CL.Identity(r.identity)
	fmt.Printf("aggregatorCommandReceiveC2CL connect : " + aggregatorCommandReceiveC2CLConnection)
}

func (r AggregatorCommandRoutine) close() err error {
	r.aggregatorCommandSendC2CL.close()
	r.aggregatorCommandReceiveC2CL.close()
	r.aggregatorCommandSendC2CL.close()
	r.aggregatorCommandReceiveC2CL.close()
	r.Context.close()
}

func (r AggregatorCommandRoutine) run() err error {
	pi := zmq.PollItems{
		zmq.PollItem{Socket: aggregatorCommandSendC2CL, Events: zmq.POLLIN},
		zmq.PollItem{Socket: aggregatorCommandReceiveC2CL, Events: zmq.POLLIN},
		zmq.PollItem{Socket: aggregatorCommandSendCL2C, Events: zmq.POLLIN},
		zmq.PollItem{Socket: aggregatorCommandReceiveCL2C, Events: zmq.POLLIN}}

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
			err = r.processCommandSendC2CL(command)
			if err != nil {
				panic(err)
			}

		case pi[1].REvents&zmq.POLLIN != 0:

			command, err := pi[1].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			err = r.processCommandReceiveC2CL(command)
			if err != nil {
				panic(err)
			}

		case pi[2].REvents&zmq.POLLIN != 0:

			command, err := pi[1].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			err = r.processCommandSendCL2C(command)
			if err != nil {
				panic(err)
			}

		case pi[3].REvents&zmq.POLLIN != 0:

			command, err := pi[1].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			err = r.processCommandReceiveC2CL(command)
			if err != nil {
				panic(err)
			}
		}
	}

	fmt.Println("done")
}

func (r AggregatorCommandRoutine) processCommandSendC2CL(command [][]byte) err error {
	command = r.updateHeaderCommandSendC2CL(command)
	 r.connectorCommandReceiveC2CL.SendMessage(command)
	 //RESULT TO CLUSTER
}

func (r AggregatorCommandRoutine) updateHeaderCommandSendC2CL(command [][]byte) (command [][]byte, err error) {
    sourceConnector := command[0]
    commandMessage := message.CommandMessage.decode(command[1])
    commandMessage.sourceConnector = sourceConnector
    commandMessage.sourceAggreagator = r.identity
}

func (r AggregatorCommandRoutine) processCommandReceiveC2CL(command [][]byte) err error {
	command = r.updateHeaderCommandReceiveC2CL(command)
	 r.connectorCommandSendC2CL.SendMessage(command)
	 //COMMAND FROM CLUSTER
}

func (r AggregatorCommandRoutine) updateHeaderCommandReceiveC2CL(command [][]byte) (command [][]byte, err error) {
    commandMessage := message.CommandMessage.decode(command[1])
    command[0] = commandMessage.targetConnector
}

func (r AggregatorCommandRoutine) processCommandSendCL2C(command [][]byte) err error {
	command = r.updateHeaderCommandSendCL2C(command)
	 r.connectorCommandReceiveC2CL.SendMessage(command)
        //COMMAND
}

func (r AggregatorCommandRoutine) updateHeaderCommandSendCL2C(command [][]byte) (command [][]byte, err error {
    commandMessage := message.CommandMessage.decode(command[1])
    command[0] = commandMessage.sourceConnector
}

func (r AggregatorCommandRoutine) processCommandReceiveCL2C(command [][]byte) err error {
	command = r.updateHeaderCommandReceiveC2CL(command)
	 r.connectorCommandReceiveC2CL.SendMessage(command)
	 //RECEIVE FROM CONNECTOR
}

func (r AggregatorCommandRoutine) updateHeaderCommandReceiveC2CL(command [][]byte) (command [][]byte, err error {
    commandMessage := message.CommandMessage.decode(command[1])
    command[0] = commandMessage.targetConnector
}
