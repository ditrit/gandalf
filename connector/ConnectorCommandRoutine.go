package connector

import (
	"fmt"

	zmq "github.com/zeromq/goczmq"
)

type ConnectorCommandRoutine struct {
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

func (r ConnectorCommandRoutine) new(identity, connectorCommandSendCL2CConnection, connectorCommandReceiveCL2CConnection, connectorCommandSendC2CLConnection, connectorCommandReceiveC2CLConnection string) {
	r.identity = identity
	r.connectorCommandSendCL2CConnection = connectorCommandSendCL2CConnection
	r.connectorCommandSendCL2C = zmq.NewDealer(r.connectorCommandSendCL2CConnection)
	r.connectorCommandSendCL2C.Identity(r.identity)
	fmt.Printf("connectorCommandSendCL2C connect : " + connectorCommandSendCL2CConnection)

	r.connectorCommandReceiveCL2CConnection = connectorCommandReceiveCL2CConnection
	r.connectorCommandReceiveCL2C = zmq.NewRouter(connectorCommandReceiveCL2CConnection)
	r.connectorCommandReceiveCL2C.Identity(r.identity)
	fmt.Printf("connectorCommandReceiveCL2C connect : " + connectorCommandReceiveCL2CConnection)

	r.connectorCommandSendC2CLConnection = connectorCommandSendC2CLConnection
	r.connectorCommandSendC2CL = zmq.NewDealer(connectorCommandSendC2CLConnection)
	r.connectorCommandSendC2CL.Identity(r.identity)
	fmt.Printf("connectorCommandSendC2CL connect : " + connectorCommandSendC2CLConnection)

	r.connectorCommandReceiveC2CLConnection = connectorCommandReceiveC2CLConnection
	r.connectorCommandReceiveC2CL = zmq.NewRouter(connectorCommandReceiveC2CLConnection)
	r.connectorCommandReceiveC2CL.Identity(r.identity)
	fmt.Printf("connectorCommandReceiveC2CL connect : " + connectorCommandReceiveC2CLConnection)
}

func (r ConnectorCommandRoutine) close() {
}

func (r ConnectorCommandRoutine) reconnectToProxy() {

}

func (r ConnectorCommandRoutine) run() {
	pi := zmq.PollItems{
		zmq.PollItem{Socket: connectorCommandSendCL2C, Events: zmq.POLLIN},
		zmq.PollItem{Socket: connectorCommandReceiveCL2C, Events: zmq.POLLIN},
		zmq.PollItem{Socket: connectorCommandSendC2CL, Events: zmq.POLLIN},
		zmq.PollItem{Socket: connectorCommandReceiveC2CL, Events: zmq.POLLIN},

		var command = [][]byte{}

	for {

		_, _ = zmq.Poll(pi, -1)

		switch {
		case pi[0].REvents&zmq.POLLIN != 0:

			command, err := pi[0].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}

			err = r.processCommandSendCL2C(command)
			if err != nil {
				panic(err)
			}

		case pi[1].REvents&zmq.POLLIN != 0:

			command, err := pi[1].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			err = r.processCommandReceiveCL2C(command)
			if err != nil {
				panic(err)
			}

		case pi[2].REvents&zmq.POLLIN != 0:

			command, err := pi[2].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			err = r.processCommandSendC2CL(command)
			if err != nil {
				panic(err)
			}

		case pi[3].REvents&zmq.POLLIN != 0:

			command, err := pi[3].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			err = r.processCommandReceiveC2CL(command)
			if err != nil {
				panic(err)
			}
		}
	}
}

func (r ConnectorCommandRoutine) processCommandSendCL2C(command [][]byte) {
	command = r.updateHeaderCommandSendCL2C(command)
}

func (r ConnectorCommandRoutine) updateHeaderCommandSendCL2C(command [][]byte) {

}

func (r ConnectorCommandRoutine) processCommandReceiveCL2C(command [][]byte) {
	command = r.updateHeaderCommandReceiveCL2C(command)
}

func (r ConnectorCommandRoutine) updateHeaderCommandReceiveCL2C(command [][]byte) {

}

func (r ConnectorCommandRoutine) processCommandSendC2CL(command [][]byte) {
	command = r.updateHeaderCommandSendC2CL(command)
}

func (r ConnectorCommandRoutine) updateHeaderCommandSendC2CL(command [][]byte) {

}

func (r ConnectorCommandRoutine) processCommandReceiveC2CL(command [][]byte) {
	command = r.updateHeaderCommandSendC2CL(command)
}

func (r ConnectorCommandRoutine) updateHeaderCommandReceiveC2CL(command [][]byte) {

}
