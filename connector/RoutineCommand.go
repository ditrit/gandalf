package connector

import (
	"fmt"

	zmq "github.com/zeromq/goczmq"
)

type RoutineCommand struct {
	connectorCommandSendCL2C              zmq.Sock
	connectorCommandSendCL2CConnection    string
	connectorCommandReceiveCL2C           zmq.Sock
	connectorCommandReceiveCL2CConnection string
	connectorCommandSendC2CL              zmq.Sock
	connectorCommandSendC2CLConnection    string
	connectorCommandReceiveC2CL           zmq.Sock
	connectorCommandReceiveC2CLConnection string
	identity                              string
}

func (rt RoutineCommand) new(identity, connectorCommandSendCL2CConnection, connectorCommandReceiveCL2CConnection, connectorCommandSendC2CLConnection, connectorCommandReceiveC2CLConnection string) {
	rt.identity = identity
	rt.connectorCommandReceiveC2CLConnection = connectorCommandReceiveC2CLConnection
	rt.connectorCommandSendCL2C = zmq.NewDealer(connectorCommandSendCL2CConnection)
	rt.connectorCommandSendCL2C.Identity(rt.Identity)
	fmt.Printf("connectorCommandSendCL2C connect : " + connectorCommandSendCL2CConnection)

	rt.connectorCommandReceiveCL2CConnection = connectorCommandReceiveCL2CConnection
	rt.connectorCommandReceiveCL2C = zmq.NewRouter(connectorCommandReceiveCL2CConnection)
	rt.connectorCommandReceiveCL2C.Identity(rt.Identity)
	fmt.Printf("connectorCommandReceiveCL2C connect : " + connectorCommandReceiveCL2CConnection)

	rt.connectorCommandSendC2CLConnection = connectorCommandSendC2CLConnection
	rt.connectorCommandSendC2CL = zmq.NewDealer(connectorCommandSendC2CLConnection)
	rt.connectorCommandSendC2CL.Identity(rt.Identity)
	fmt.Printf("connectorCommandSendC2CL connect : " + connectorCommandSendC2CLConnection)

	rt.connectorCommandReceiveC2CLConnection = connectorCommandReceiveC2CLConnection
	rt.connectorCommandReceiveC2CL = zmq.NewRouter(connectorCommandReceiveC2CLConnection)
	rt.connectorCommandReceiveC2CL.Identity(rt.Identity)
	fmt.Printf("connectorCommandReceiveC2CL connect : " + connectorCommandReceiveC2CLConnection)
}

func (rt RoutineCommand) close() {
}

func (rt RoutineCommand) reconnectToProxy() {

}

func (rt RoutineCommand) run() {
	pi := zmq.PollItems{
		zmq.PollItem{Socket: connectorCommandSendCL2C, Events: zmq.POLLIN},
		zmq.PollItem{Socket: connectorCommandReceiveCL2C, Events: zmq.POLLIN},
		zmq.PollItem{Socket: connectorCommandSendC2CL, Events: zmq.POLLIN},
		zmq.PollItem{Socket: connectorCommandReceiveC2CL, Events: zmq.POLLIN}}

	var command = [][]byte{}

	for {

		_, _ = zmq.Poll(pi, -1)

		switch {
		case pi[0].REvents&zmq.POLLIN != 0:

			command, err := pi[0].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}

			//PROCESS FORMATAGE TO WORKER
			err = command.send(pi[0].Socket)
			if err != nil {
				panic(err)
			}

		case pi[1].REvents&zmq.POLLIN != 0:

			command, err := pi[1].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			err = command.send(pi[0].Socket)
			if err != nil {
				panic(err)
			}

		case pi[2].REvents&zmq.POLLIN != 0:

			command, err := pi[2].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			//PROCESS FORMATAGE TO CLUSTER
			err = command.send(pi[2].Socket)
			if err != nil {
				panic(err)
			}

		case pi[3].REvents&zmq.POLLIN != 0:

			command, err := pi[3].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			err = command.send(pi[2].Socket)
			if err != nil {
				panic(err)
			}
		}

	}
}
