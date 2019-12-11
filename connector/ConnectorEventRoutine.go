package connector

import (
	"fmt"

	zmq "github.com/zeromq/goczmq"
)

type ConnectorEventRoutine struct {
	connectorMapUUIDEventMessage	   map[string][]EventMessage					
	connectorMapWorkerEvents 		   map[string][]string	
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
	go cleanEventsByTimeout()

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

func (r ConnectorEventRoutine) processEventSendA2W(event [][]byte) err error {
	eventMessage := EventMessage.decodeEvent(event[1])
	r.addEvents(eventMessage)
	for {
		isSend = eventMessage.sendEventWith(r.connectorEventReceiveA2W)
		if isSend {
			break
		}
		time.Sleep(2 * time.Second)
	}
}

func (r ConnectorEventRoutine) processEventReceiveA2W(event [][]byte) err error {
	eventMessage := EventMessage.decodeEvent(event[1])
	for {
		isSend = eventMessage.sendEventWith(r.connectorEventSendA2W)
		if isSend {
			break
		}
		time.Sleep(2 * time.Second)
	}
}

func (r ConnectorEventRoutine) processEventSendW2A(event [][]byte) err error {
	eventMessage := EventMessage.decodeEvent(event[1])
	for {
		isSend = eventMessage.sendEventWith(r.connectorEventReceiveW2A)
		if isSend {
			break
		}
		time.Sleep(2 * time.Second)
	}
}

func (r ConnectorEventRoutine) processEventReceiveW2A(event [][]byte) err error {
	
	if event[0] == Constant.COMMAND_VALIDATION_FUNCTIONS {
		commandFunctions := decodeCommandCommandsEvents(command[2])
		result := r.validationEvents(workerSource, commandFunctions.events)
        if result {
			r.connectorMapWorkerEvents[workerSource] = events
			commandFunctionReply := CommandFunctionReply.New(result)
			for {
				isSend = commandFunctionReply.sendCommandFunctionReplyWith(r.connectorCommandSendA2W)
				if isSend {
					break
				}
				time.Sleep(2 * time.Second)
			}
        }
	}
	else {
		eventMessage := EventMessage.decodeEvent(event[1])
		for {
			isSend = eventMessage.sendEventWith(r.connectorEventSendW2A)
			if isSend {
				break
			}
			time.Sleep(2 * time.Second)
		}
	}
}

func (r ConnectorEventRoutine) validationEvents(workerSource string, events []string) (result bool, err error) {
	//TODO
	result := true
	return
}

func (r ConnectorEventRoutine) addEvents(eventMessage EventMessage) {
	if val, ok := r.connectorMapUUIDEventMessage[eventMessage.uuid]; ok {
		if !ok {
			r.connectorMapUUIDEventMessage[eventMessage.uuid] = eventMessage
		}
	}
}

func (r ConnectorEventRoutine) cleanEventsByTimeout() {
	maxTimeout = 0
	for {
		for uuid, eventMessage := range r.connectorMapUUIDEventMessage { 
			if commandMessage.timestamp - commandMessage.timeout == 0 {
				delete(r.commandUUIDCommandMessage, uuid) 	
			}
			else {
				if commandMessage.timeout >= maxTimeout {
					maxTimeout = commandMessage.timeout
				}
			}
		}
		time.Sleep(maxTimeout * time.Millisecond)
	}
}