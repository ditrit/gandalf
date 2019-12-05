package connector

import (
	"fmt"

	zmq "github.com/zeromq/goczmq"
)

type ConnectorEventRoutine struct {
	connectorEventSendA2W              zmq.Sock
	connectorEventSendA2WConnection    string
	connectorEventReceiveA2W           zmq.Sock
	connectorEventReceiveA2WConnection string
	connectorEventSendW2A              zmq.Sock
	connectorEventSendW2AConnection    string
	connectorEventReceiveW2A           zmq.Sock
	connectorEventReceiveW2AConnection string
	identity                            string
}

func (r ConnectorEventRoutine) New(identity, connectorEventSendA2WConnection, connectorEventReceiveA2WConnection, connectorEventSendW2AConnection, connectorEventReceiveW2AConnection string) err error {
	r.identity = identity
	r.connectorEventSendA2WConnection = connectorEventSendA2WConnection
	r.connectorEventSendA2W = zmq.NewDealer(connectorEventSendA2WConnection)
	r.connectorEventSendA2W.Identity(r.Identity)
	fmt.Printf("connectorEventSendA2W connect : " + connectorEventSendA2WConnection)

	r.connectorEventReceiveA2WConnection = connectorEventReceiveA2WConnection
	r.connectorEventReceiveA2W = zmq.NewRouter(connectorEventReceiveA2WConnection)
	r.connectorEventReceiveA2W.Identity(r.Identity)
	fmt.Printf("connectorEventReceiveA2W connect : " + connectorEventReceiveA2WConnection)

	r.connectorEventSendW2AConnection = connectorEventSendW2AConnection
	r.connectorEventSendW2A = zmq.NewDealer(connectorEventSendW2AConnection)
	r.connectorEventSendW2A.Identity(r.Identity)
	fmt.Printf("connectorEventSendW2A connect : " + connectorEventSendW2AConnection)

	r.connectorEventReceiveW2AConnection = connectorEventReceiveW2AConnection
	r.connectorEventReceiveW2A = zmq.NewRouter(connectorEventReceiveW2AConnection)
	r.connectorEventReceiveW2A.Identity(r.Identity)
	fmt.Printf("connectorEventReceiveW2A connect : " + connectorEventReceiveW2AConnection)
}

func (r ConnectorEventRoutine) close() err error {
}

func (r ConnectorEventRoutine) reconnectToProxy() err error {

}

func (r ConnectorEventRoutine) run() err error {
	pi := zmq.PollItems{
		zmq.PollItem{Socket: connectorEventSendA2W, Events: zmq.POLLIN},
		zmq.PollItem{Socket: connectorEventReceiveA2W, Events: zmq.POLLIN},
		zmq.PollItem{Socket: connectorEventSendW2A, Events: zmq.POLLIN},
		zmq.PollItem{Socket: connectorEventReceiveW2A, Events: zmq.POLLIN}}

	var event = [][]byte{}

	for {

		_, _ = zmq.Poll(pi, -1)

		switch {
		case pi[0].REvents&zmq.POLLIN != 0:

			event, err := pi[0].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			err = r.processEventSendA2W(event)
			if err != nil {
				panic(err)
			}

		case pi[1].REvents&zmq.POLLIN != 0:

			event, err := pi[1].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			err = r.processEventReceiveA2W(event)
			if err != nil {
				panic(err)
			}

		case pi[2].REvents&zmq.POLLIN != 0:

			event, err := pi[2].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			err = r.processEventSendW2A(event)
			if err != nil {
				panic(err)
			}

		case pi[3].REvents&zmq.POLLIN != 0:

			event, err := pi[3].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			err = r.processEventReceiveW2A(event)
			if err != nil {
				panic(err)
			}
		}
	}
}

func (r ConnectorCommandRoutine) processEventSendA2W(event [][]byte) err error {
	event = r.updateHeaderEventSendA2W(event)
	r.connectorEventSendW2A.SendMessage(event)
}

func (r ConnectorCommandRoutine) updateHeaderEventSendA2W(event [][]byte) err error {
    //TODO NOTHING
}

func (r ConnectorCommandRoutine) processEventReceiveA2W(event [][]byte) err error {
	event = r.updateHeaderEventReceiveA2W(event)
	r.connectorEventReceiveW2A.SendMessage(event)
}

func (r ConnectorCommandRoutine) updateHeaderEventReceiveA2W(event [][]byte) err error {
    //TODO NOTHING
}

func (r ConnectorCommandRoutine) processEventSendW2A(event [][]byte) err error {
	event = r.updateHeaderEventSendW2A(event)
	r.connectorEventSendA2W.SendMessage(event)
}

func (r ConnectorCommandRoutine) updateHeaderEventSendW2A(event [][]byte) err error {
        //TODO NOTHING
}

func (r ConnectorCommandRoutine) processEventReceiveW2A(event [][]byte) err error {
	event = r.updateHeaderEventReceiveW2A(event)
	r.connectorEventReceiveA2W.SendMessage(event)
}

func (r ConnectorCommandRoutine) updateHeaderEventReceiveW2A(event [][]byte) err error {
    //TODO NOTHING
}
