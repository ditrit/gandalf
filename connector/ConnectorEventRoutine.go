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
	r.connectorEventSendToWorker = r.context.NewSocket(zmq4.XPUB)
	r.connectorEventSendToWorker.Identity(r.Identity)
	r.connectorEventSendToWorker.Bind(r.connectorEventSendToWorkerConnection)
	fmt.Printf("connectorEventSendToWorker bind : " + connectorEventSendToWorkerConnection)

	r.connectorEventReceiveFromAggregatorConnection = connectorEventReceiveFromAggregatorConnection
	r.connectorEventReceiveFromAggregator = r.context.NewSocket(zmq4.XSUB)
	r.connectorEventReceiveFromAggregator.Identity(r.Identity)
	r.connectorEventReceiveFromAggregator.Connect(r.connectorEventReceiveFromAggregatorConnection)
	fmt.Printf("connectorEventReceiveFromAggregator connect : " + connectorEventReceiveFromAggregatorConnection)

	r.connectorEventSendToAggregatorConnection = connectorEventSendToAggregatorConnection
	r.connectorEventSendToAggregator = r.context.NewSocket(zmq4.XPUB)
	r.connectorEventSendToAggregator.Identity(r.Identity)
	r.connectorEventSendToAggregator.Connect(r.connectorEventSendToAggregatorConnection)
	fmt.Printf("connectorEventSendToAggregator connect : " + connectorEventSendToAggregatorConnection)

	r.connectorEventReceiveFromWorkerConnection = connectorEventReceiveFromWorkerConnection
	r.connectorEventReceiveFromWorker = r.context.NewSocket(zmq4.XSUB)
	r.connectorEventReceiveFromWorker.Identity(r.Identity)
	r.connectorEventReceiveFromWorker.Bind(r.connectorEventReceiveFromWorkerConnection)
	fmt.Printf("connectorEventReceiveFromWorker bind : " + connectorEventReceiveFromWorkerConnection)
}

func (r ConnectorEventRoutine) close() {
}

func (r ConnectorEventRoutine) reconnectToProxy() {

}

func (r ConnectorEventRoutine) run() {
	go cleanEventsByTimeout()

	poller := zmq4.NewPoller()
	poller.Add(r.connectorEventSendToWorker, zmq4.POLLIN)
	poller.Add(r.connectorEventReceiveFromAggregator, zmq4.POLLIN)
	poller.Add(r.connectorEventSendToAggregator, zmq4.POLLIN)
	poller.Add(r.connectorEventReceiveFromWorker, zmq4.POLLIN)

	event := [][]byte{}

	for {
		r.sendReadyCommand()
		sockets, _ := poller.Poll(-1)
		for _, socket := range sockets {

			switch currentSocket := socket.Socket; currentSocket {
			case connectorEventSendToWorker:

				event, err := currentSocket.RecvMessage()
				if err != nil {
					panic(err)
				}
				err = r.processEventSendToWorker(event)
				if err != nil {
					panic(err)
				}

			case connectorEventReceiveFromAggregator:

				event, err := currentSocket.RecvMessage()
				if err != nil {
					panic(err)
				}
				err = r.processEventReceiveFromAggregator(event)
				if err != nil {
					panic(err)
				}

			case connectorEventSendToAggregator:

				event, err := currentSocket.RecvMessage()
				if err != nil {
					panic(err)
				}
				err = r.processEventSendToAggregator(event)
				if err != nil {
					panic(err)
				}

			case connectorEventReceiveFromWorker:

				event, err := currentSocket.RecvMessage()
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