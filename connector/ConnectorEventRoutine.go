package connector

import (
	"fmt"
	"gandalfgo/message"
	zmq4 "github.com/pebbe/zmq4"
)

type ConnectorEventRoutine struct {
	Context											zmq4.Context
	ConnectorMapUUIDEventMessage	   				map[string][]message.EventMessage					
	ConnectorMapWorkerEvents 		   				map[string][]string	
	ConnectorEventSendToWorker              		zmq4.Socket
	ConnectorEventSendToWorkerConnection    		string
	ConnectorEventReceiveFromAggregator           	zmq4.Socket
	ConnectorEventReceiveFromAggregatorConnection 	string
	ConnectorEventSendToAggregator              	zmq4.Socket
	ConnectorEventSendToAggregatorConnection    	string
	ConnectorEventReceiveFromWorker           		zmq4.Socket
	ConnectorEventReceiveFromWorkerConnection 		string
	Identity                            			string
}

func NewConnectorEventRoutine(identity, connectorEventSendToWorkerConnection, connectorEventReceiveFromAggregatorConnection, connectorEventSendToAggregatorConnection, connectorEventReceiveFromWorkerConnection string) (connectorEventRoutine *ConnectorEventRoutine) {
	connectorEventRoutine = new(ConnectorEventRoutine)
	
	connectorEventRoutine.identity = identity

	connectorEventRoutine.Context, _ = zmq4.NewContext()
	connectorEventRoutine.ConnectorEventSendToWorkerConnection = connectorEventSendToWorkerConnection
	connectorEventRoutine.ConnectorEventSendToWorker = connectorEventRoutine.Context.NewSocket(zmq4.XPUB)
	connectorEventRoutine.ConnectorEventSendToWorker.Identity(connectorEventRoutine.Identity)
	connectorEventRoutineConnectorEventSendToWorker.Bind(connectorEventRoutine.ConnectorEventSendToWorkerConnection)
	fmt.Printf("connectorEventSendToWorker bind : " + connectorEventSendToWorkerConnection)

	connectorEventRoutine.ConnectorEventReceiveFromAggregatorConnection = connectorEventReceiveFromAggregatorConnection
	connectorEventRoutine.ConnectorEventReceiveFromAggregator = connectorEventRoutine.Context.NewSocket(zmq4.XSUB)
	connectorEventRoutine.ConnectorEventReceiveFromAggregator.Identity(connectorEventRoutine.Identity)
	connectorEventRoutine.ConnectorEventReceiveFromAggregator.Connect(connectorEventRoutine.ConnectorEventReceiveFromAggregatorConnection)
	fmt.Printf("connectorEventReceiveFromAggregator connect : " + connectorEventReceiveFromAggregatorConnection)

	connectorEventRoutine.ConnectorEventSendToAggregatorConnection = connectorEventSendToAggregatorConnection
	connectorEventRoutine.ConnectorEventSendToAggregator = connectorEventRoutine.Context.NewSocket(zmq4.XPUB)
	connectorEventRoutine.ConnectorEventSendToAggregator.Identity(connectorEventRoutine.Identity)
	connectorEventRoutine.ConnectorEventSendToAggregator.Connect(connectorEventRoutine.ConnectorEventSendToAggregatorConnection)
	fmt.Printf("connectorEventSendToAggregator connect : " + connectorEventSendToAggregatorConnection)

	connectorEventRoutine.ConnectorEventReceiveFromWorkerConnection = connectorEventReceiveFromWorkerConnection
	connectorEventRoutine.ConnectorEventReceiveFromWorker = connectorEventRoutine.Context.NewSocket(zmq4.XSUB)
	connectorEventRoutine.ConnectorEventReceiveFromWorker.Identity(connectorEventRoutine.Identity)
	connectorEventRoutine.ConnectorEventReceiveFromWorker.Bind(connectorEventRoutine.ConnectorEventReceiveFromWorkerConnection)
	fmt.Printf("connectorEventReceiveFromWorker bind : " + connectorEventReceiveFromWorkerConnection)
}

func (r ConnectorEventRoutine) close() {
	r.ConnectorEventSendToWorker.Close()
	r.ConnectorEventReceiveFromAggregator.Close()
	r.ConnectorEventSendToAggregator.Close()
	r.ConnectorEventReceiveFromWorker.Close()
	r.Context.Term()
}

func (r ConnectorEventRoutine) reconnectToProxy() {

}

func (r ConnectorEventRoutine) run() {
	go cleanEventsByTimeout()

	poller := zmq4.NewPoller()
	poller.Add(r.ConnectorEventSendToWorker, zmq4.POLLIN)
	poller.Add(r.ConnectorEventReceiveFromAggregator, zmq4.POLLIN)
	poller.Add(r.ConnectorEventSendToAggregator, zmq4.POLLIN)
	poller.Add(r.ConnectorEventReceiveFromWorker, zmq4.POLLIN)

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
	eventMessage := message.DecodeEvent(event[1])
	r.addEvents(eventMessage)
	go eventMessage.sendEventWith(r.connectorEventReceiveFromAggregator)
}

func (r ConnectorEventRoutine) processEventReceiveFromAggregator(event [][]byte) {
	eventMessage := message.DecodeEvent(event[1])
	go eventMessage.sendEventWith(r.connectorEventSendToWorker)
}

func (r ConnectorEventRoutine) processEventSendToAggregator(event [][]byte) {
	eventMessage := message.DecodeEvent(event[1])
	go eventMessage.SendEventWith(r.connectorEventReceiveFromWorker)
}

func (r ConnectorEventRoutine) processEventReceiveFromWorker(event [][]byte) {
	
	if event[0] == Constant.COMMAND_VALIDATION_FUNCTIONS {
		commandFunctions := decodeCommandCommandsEvents(command[2])
		result := r.validationEvents(workerSource, commandFunctions.events)
        if result {
			r.ConnectorMapWorkerEvents[workerSource] = events
			commandFunctionReply := CommandFunctionReply.New(result)
			go commandFunctionReply.SendCommandFunctionReplyWith(r.connectorCommandSendToWorker)
        }
	} else {
		eventMessage := EventMessage.decodeEvent(event[1])
		go eventMessage.SendEventWith(r.connectorEventSendToAggregator)
	}
}

func (r ConnectorEventRoutine) validationEvents(workerSource string, events []string) (result bool, err error) {
	//TODO
	result = true
	return
}

func (r ConnectorEventRoutine) addEvents(eventMessage message.EventMessage) {
	if val, ok := r.connectorMapUUIDEventMessage[eventMessage.Uuid]; ok {
		if !ok {
			r.connectorMapUUIDEventMessage[eventMessage.Uuid] = eventMessage
		}
	}
}

func (r ConnectorEventRoutine) cleanEventsByTimeout() {
	maxTimeout = 0
	for {
		for uuid, eventMessage := range r.connectorMapUUIDEventMessage { 
			if commandMessage.timestamp - commandMessage.timeout == 0 {
				delete(r.commandUUIDCommandMessage, uuid) 	
			} else {
				if commandMessage.timeout >= maxTimeout {
					maxTimeout = commandMessage.timeout
				}
			}
		}
		time.Sleep(maxTimeout * time.Millisecond)
	}
}