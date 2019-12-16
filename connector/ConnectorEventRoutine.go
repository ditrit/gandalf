package connector

import (
	"fmt"
	"gandalfgo/message"
	zmq4 "github.com/pebbe/zmq4"
)

type ConnectorEventRoutine struct {
	context											zmq4.Context
	connectorMapUUIDEventMessage	   				map[string][]EventMessage					
	connectorMapWorkerEvents 		   				map[string][]string	
	connectorEventSendToWorker              		zmq4.Socket
	connectorEventSendToWorkerConnection    		string
	connectorEventReceiveFromAggregator           	zmq4.Socket
	connectorEventReceiveFromAggregatorConnection 	string
	connectorEventSendToAggregator              	zmq4.Socket
	connectorEventSendToAggregatorConnection    	string
	connectorEventReceiveFromWorker           		zmq4.Socket
	connectorEventReceiveFromWorkerConnection 		string
	identity                            			string
}

func (r ConnectorEventRoutine) New(identity, connectorEventSendToWorkerConnection, connectorEventReceiveFromAggregatorConnection, connectorEventSendToAggregatorConnection, connectorEventReceiveFromWorkerConnection string) err error {
	r.identity = identity

	r.context, _ := zmq4.NewContext()
	r.connectorEventSendToWorkerConnection = connectorEventSendToWorkerConnection
	r.connectorEventSendToWorker = r.context.NewXPub(connectorEventSendToWorkerConnection)
	r.connectorEventSendToWorker.Identity(r.Identity)
	fmt.Printf("connectorEventSendToWorker connect : " + connectorEventSendToWorkerConnection)

	r.connectorEventReceiveFromAggregatorConnection = connectorEventReceiveFromAggregatorConnection
	r.connectorEventReceiveFromAggregator = r.context.NewXSub(connectorEventReceiveFromAggregatorConnection)
	r.connectorEventReceiveFromAggregator.Identity(r.Identity)
	fmt.Printf("connectorEventReceiveFromAggregator connect : " + connectorEventReceiveFromAggregatorConnection)

	r.connectorEventSendToAggregatorConnection = connectorEventSendToAggregatorConnection
	r.connectorEventSendToAggregator = r.context.NewXPub(connectorEventSendToAggregatorConnection)
	r.connectorEventSendToAggregator.Identity(r.Identity)
	fmt.Printf("connectorEventSendToAggregator connect : " + connectorEventSendToAggregatorConnection)

	r.connectorEventReceiveFromWorkerConnection = connectorEventReceiveFromWorkerConnection
	r.connectorEventReceiveFromWorker = r.context.NewXSub(connectorEventReceiveFromWorkerConnection)
	r.connectorEventReceiveFromWorker.Identity(r.Identity)
	fmt.Printf("connectorEventReceiveFromWorker connect : " + connectorEventReceiveFromWorkerConnection)
}

func (r ConnectorEventRoutine) close() {
}

func (r ConnectorEventRoutine) reconnectToProxy() {

}

func (r ConnectorEventRoutine) run() {
	go cleanEventsByTimeout()

	pi := zmq4.PollItems{
		zmq4.PollItem{Socket: connectorEventSendToWorker, Events: zmq4.POLLIN},
		zmq4.PollItem{Socket: connectorEventReceiveFromAggregator, Events: zmq4.POLLIN},
		zmq4.PollItem{Socket: connectorEventSendToAggregator, Events: zmq4.POLLIN},
		zmq4.PollItem{Socket: connectorEventReceiveFromWorker, Events: zmq4.POLLIN}}

	var event = [][]byte{}

	for {

		_, _ = zmq4.Poll(pi, -1)

		switch {
		case pi[0].REvents&zmq4.POLLIN != 0:

			event, err := pi[0].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			err = r.processEventSendToWorker(event)
			if err != nil {
				panic(err)
			}

		case pi[1].REvents&zmq4.POLLIN != 0:

			event, err := pi[1].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			err = r.processEventReceiveFromAggregator(event)
			if err != nil {
				panic(err)
			}

		case pi[2].REvents&zmq4.POLLIN != 0:

			event, err := pi[2].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			err = r.processEventSendToAggregator(event)
			if err != nil {
				panic(err)
			}

		case pi[3].REvents&zmq4.POLLIN != 0:

			event, err := pi[3].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			err = r.processEventReceiveFromWorker(event)
			if err != nil {
				panic(err)
			}
		}
	}
}

func (r ConnectorEventRoutine) processEventSendToWorker(event [][]byte) {
	eventMessage := EventMessage.decodeEvent(event[1])
	r.addEvents(eventMessage)
	go eventMessage.sendEventWith(r.connectorEventReceiveFromAggregator)
}

func (r ConnectorEventRoutine) processEventReceiveFromAggregator(event [][]byte) {
	eventMessage := EventMessage.decodeEvent(event[1])
	go eventMessage.sendEventWith(r.connectorEventSendToWorker)
}

func (r ConnectorEventRoutine) processEventSendToAggregator(event [][]byte) {
	eventMessage := EventMessage.decodeEvent(event[1])
	go eventMessage.sendEventWith(r.connectorEventReceiveFromWorker)
}

func (r ConnectorEventRoutine) processEventReceiveFromWorker(event [][]byte) {
	
	if event[0] == Constant.COMMAND_VALIDATION_FUNCTIONS {
		commandFunctions := decodeCommandCommandsEvents(command[2])
		result := r.validationEvents(workerSource, commandFunctions.events)
        if result {
			r.connectorMapWorkerEvents[workerSource] = events
			commandFunctionReply := CommandFunctionReply.New(result)
			go commandFunctionReply.sendCommandFunctionReplyWith(r.connectorCommandSendToWorker)
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