package connector

import (
	"fmt"

	zmq "github.com/zeromq/goczmq"
)

type ConnectorEventRoutine struct {
	connectorMapUUIDEventMessage	   map[string][]EventMessage					
	connectorMapWorkerEvents 		   map[string][]string	
	connectorEventSendToWorker              zmq.Sock
	connectorEventSendToWorkerConnection    string
	connectorEventReceiveFromAggregator           zmq.Sock
	connectorEventReceiveFromAggregatorConnection string
	connectorEventSendToAggregator              zmq.Sock
	connectorEventSendToAggregatorConnection    string
	connectorEventReceiveFromWorker           zmq.Sock
	connectorEventReceiveFromWorkerConnection string
	identity                            string
}

func (r ConnectorEventRoutine) New(identity, connectorEventSendToWorkerConnection, connectorEventReceiveFromAggregatorConnection, connectorEventSendToAggregatorConnection, connectorEventReceiveFromWorkerConnection string) err error {
	r.identity = identity
	r.connectorEventSendToWorkerConnection = connectorEventSendToWorkerConnection
	r.connectorEventSendToWorker = zmq.NewXPub(connectorEventSendToWorkerConnection)
	r.connectorEventSendToWorker.Identity(r.Identity)
	fmt.Printf("connectorEventSendToWorker connect : " + connectorEventSendToWorkerConnection)

	r.connectorEventReceiveFromAggregatorConnection = connectorEventReceiveFromAggregatorConnection
	r.connectorEventReceiveFromAggregator = zmq.NewXSub(connectorEventReceiveFromAggregatorConnection)
	r.connectorEventReceiveFromAggregator.Identity(r.Identity)
	fmt.Printf("connectorEventReceiveFromAggregator connect : " + connectorEventReceiveFromAggregatorConnection)

	r.connectorEventSendToAggregatorConnection = connectorEventSendToAggregatorConnection
	r.connectorEventSendToAggregator = zmq.NewXPub(connectorEventSendToAggregatorConnection)
	r.connectorEventSendToAggregator.Identity(r.Identity)
	fmt.Printf("connectorEventSendToAggregator connect : " + connectorEventSendToAggregatorConnection)

	r.connectorEventReceiveFromWorkerConnection = connectorEventReceiveFromWorkerConnection
	r.connectorEventReceiveFromWorker = zmq.NewXSub(connectorEventReceiveFromWorkerConnection)
	r.connectorEventReceiveFromWorker.Identity(r.Identity)
	fmt.Printf("connectorEventReceiveFromWorker connect : " + connectorEventReceiveFromWorkerConnection)
}

func (r ConnectorEventRoutine) close() err error {
}

func (r ConnectorEventRoutine) reconnectToProxy() err error {

}

func (r ConnectorEventRoutine) run() err error {
	go cleanEventsByTimeout()

	pi := zmq.PollItems{
		zmq.PollItem{Socket: connectorEventSendToWorker, Events: zmq.POLLIN},
		zmq.PollItem{Socket: connectorEventReceiveFromAggregator, Events: zmq.POLLIN},
		zmq.PollItem{Socket: connectorEventSendToAggregator, Events: zmq.POLLIN},
		zmq.PollItem{Socket: connectorEventReceiveFromWorker, Events: zmq.POLLIN}}

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
	go eventMessage.sendEventWith(r.connectorEventReceiveFromAggregator)
}

func (r ConnectorEventRoutine) processEventReceiveA2W(event [][]byte) err error {
	eventMessage := EventMessage.decodeEvent(event[1])
	go eventMessage.sendEventWith(r.connectorEventSendToWorker)
}

func (r ConnectorEventRoutine) processEventSendW2A(event [][]byte) err error {
	eventMessage := EventMessage.decodeEvent(event[1])
	go eventMessage.sendEventWith(r.connectorEventReceiveFromWorker)
}

func (r ConnectorEventRoutine) processEventReceiveW2A(event [][]byte) err error {
	
	if event[0] == Constant.COMMAND_VALIDATION_FUNCTIONS {
		commandFunctions := decodeCommandCommandsEvents(command[2])
		result := r.validationEvents(workerSource, commandFunctions.events)
        if result {
			r.connectorMapWorkerEvents[workerSource] = events
			commandFunctionReply := CommandFunctionReply.New(result)
			go commandFunctionReply.sendCommandFunctionReplyWith(r.connectorCommandSendA2W)
        }
	}
	else {
		eventMessage := EventMessage.decodeEvent(event[1])
		go eventMessage.sendEventWith(r.connectorEventSendToAggregator)
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