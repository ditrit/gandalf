package connector

import (
	"fmt"

	zmq "github.com/zeromq/goczmq"
)

type ConnectorEventRoutine struct {
	connectorEventSendCL2C              zmq.Sock
	connectorEventSendCL2CConnection    string
	connectorEventReceiveCL2C           zmq.Sock
	connectorEventReceiveCL2CConnection string
	connectorEventSendC2CL              zmq.Sock
	connectorEventSendC2CLConnection    string
	connectorEventReceiveC2CL           zmq.Sock
	connectorEventReceiveC2CLConnection string
	identity                            string
}

func (r ConnectorEventRoutine) new(identity, connectorEventSendCL2CConnection, connectorEventReceiveCL2CConnection, connectorEventSendC2CLConnection, connectorEventReceiveC2CLConnection string) {
	re.identity = identity
	re.connectorEventSendCL2CConnection = connectorEventSendCL2CConnection
	re.connectorEventSendCL2C = zmq.NewDealer(connectorEventSendCL2CConnection)
	re.connectorEventSendCL2C.Identity(re.Identity)
	fmt.Printf("connectorEventSendCL2C connect : " + connectorEventSendCL2CConnection)

	re.connectorEventReceiveCL2CConnection = connectorEventReceiveCL2CConnection
	re.connectorEventReceiveCL2C = zmq.NewRouter(connectorEventReceiveCL2CConnection)
	re.connectorEventReceiveCL2C.Identity(re.Identity)
	fmt.Printf("connectorEventReceiveCL2C connect : " + connectorEventReceiveCL2CConnection)

	re.connectorEventSendC2CLConnection = connectorEventSendC2CLConnection
	re.connectorEventSendC2CL = zmq.NewDealer(connectorEventSendC2CLConnection)
	re.connectorEventSendC2CL.Identity(re.Identity)
	fmt.Printf("connectorEventSendC2CL connect : " + connectorEventSendC2CLConnection)

	re.connectorEventReceiveC2CLConnection = connectorEventReceiveC2CLConnection
	re.connectorEventReceiveC2CL = zmq.NewRouter(connectorEventReceiveC2CLConnection)
	re.connectorEventReceiveC2CL.Identity(re.Identity)
	fmt.Printf("connectorEventReceiveC2CL connect : " + connectorEventReceiveC2CLConnection)
}

func (r ConnectorEventRoutine) close() {
}

func (r ConnectorEventRoutine) reconnectToProxy() {

}

func (r ConnectorEventRoutine) run() {
	pi := zmq.PollItems{
		zmq.PollItem{Socket: connectorEventSendCL2C, Events: zmq.POLLIN},
		zmq.PollItem{Socket: connectorEventReceiveCL2C, Events: zmq.POLLIN},
		zmq.PollItem{Socket: connectorEventSendC2CL, Events: zmq.POLLIN},
		zmq.PollItem{Socket: connectorEventReceiveC2CL, Events: zmq.POLLIN}}

	var event = [][]byte{}

	for {

		_, _ = zmq.Poll(pi, -1)

		switch {
		case pi[0].REvents&zmq.POLLIN != 0:

			event, err := pi[0].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			err = r.processEventSendCL2C(event)
			if err != nil {
				panic(err)
			}

		case pi[1].REvents&zmq.POLLIN != 0:

			event, err := pi[1].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			err = r.processEventReceiveCL2C(event)
			if err != nil {
				panic(err)
			}

		case pi[2].REvents&zmq.POLLIN != 0:

			event, err := pi[2].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			err = r.processEventSendC2CL(event)
			if err != nil {
				panic(err)
			}

		case pi[3].REvents&zmq.POLLIN != 0:

			event, err := pi[3].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			err = r.processEventReceiveC2CL(event)
			if err != nil {
				panic(err)
			}
		}
	}
}

func (r ConnectorCommandRoutine) processEventSendCL2C(event [][]byte) {
	event = r.updateHeaderCommandSendCL2C(event)
}

func (r ConnectorCommandRoutine) updateHeaderEventSendCL2C(event [][]byte) {

}

func (r ConnectorCommandRoutine) processEventReceiveCL2C(event [][]byte) {
	event = r.updateHeaderEventReceiveCL2C(event)
}

func (r ConnectorCommandRoutine) updateHeaderEventReceiveCL2C(event [][]byte) {

}

func (r ConnectorCommandRoutine) processEventSendC2CL(event [][]byte) {
	event = r.updateHeaderEventSendC2CL(event)
}

func (r ConnectorCommandRoutine) updateHeaderEventSendC2CL(event [][]byte) {

}

func (r ConnectorCommandRoutine) processEventReceiveC2CL(event [][]byte) {
	event = r.updateHeaderEventSendC2CL(event)
}

func (r ConnectorCommandRoutine) updateHeaderEventReceiveC2CL(event [][]byte) {

}
