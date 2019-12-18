package connector

import (
	"fmt"
	"errors"
	"gandalfgo/message"
	"gandalfgo/constant"
	"github.com/pebbe/zmq4"
)

type ConnectorEventRoutine struct {
	Context											*zmq4.Context
	ConnectorMapUUIDEventMessage	   				map[string][]message.EventMessage					
	ConnectorMapWorkerEvents 		   				map[string][]string	
	ConnectorEventSendToWorker              		*zmq4.Socket
	ConnectorEventSendToWorkerConnection    		string
	ConnectorEventReceiveFromAggregator           	*zmq4.Socket
	ConnectorEventReceiveFromAggregatorConnection 	string
	ConnectorEventSendToAggregator              	*zmq4.Socket
	ConnectorEventSendToAggregatorConnection    	string
	ConnectorEventReceiveFromWorker           		*zmq4.Socket
	ConnectorEventReceiveFromWorkerConnection 		string
	Identity                            			string
}

func NewConnectorEventRoutine(identity, connectorEventSendToWorkerConnection, connectorEventReceiveFromAggregatorConnection, connectorEventSendToAggregatorConnection, connectorEventReceiveFromWorkerConnection string) (connectorEventRoutine *ConnectorEventRoutine) {
	connectorEventRoutine = new(ConnectorEventRoutine)
	
	connectorEventRoutine.Identity = identity

	connectorEventRoutine.Context, _ = zmq4.NewContext()
	connectorEventRoutine.ConnectorEventSendToWorkerConnection = connectorEventSendToWorkerConnection
	connectorEventRoutine.ConnectorEventSendToWorker, _ = connectorEventRoutine.Context.NewSocket(zmq4.XPUB)
	connectorEventRoutine.ConnectorEventSendToWorker.SetIdentity(connectorEventRoutine.Identity)
	connectorEventRoutine.ConnectorEventSendToWorker.Bind(connectorEventRoutine.ConnectorEventSendToWorkerConnection)
	fmt.Printf("connectorEventSendToWorker bind : " + connectorEventSendToWorkerConnection)

	connectorEventRoutine.ConnectorEventReceiveFromAggregatorConnection = connectorEventReceiveFromAggregatorConnection
	connectorEventRoutine.ConnectorEventReceiveFromAggregator, _ = connectorEventRoutine.Context.NewSocket(zmq4.XSUB)
	connectorEventRoutine.ConnectorEventReceiveFromAggregator.SetIdentity(connectorEventRoutine.Identity)
	connectorEventRoutine.ConnectorEventReceiveFromAggregator.Connect(connectorEventRoutine.ConnectorEventReceiveFromAggregatorConnection)
	fmt.Printf("connectorEventReceiveFromAggregator connect : " + connectorEventReceiveFromAggregatorConnection)

	connectorEventRoutine.ConnectorEventSendToAggregatorConnection = connectorEventSendToAggregatorConnection
	connectorEventRoutine.ConnectorEventSendToAggregator, _ = connectorEventRoutine.Context.NewSocket(zmq4.XPUB)
	connectorEventRoutine.ConnectorEventSendToAggregator.SetIdentity(connectorEventRoutine.Identity)
	connectorEventRoutine.ConnectorEventSendToAggregator.Connect(connectorEventRoutine.ConnectorEventSendToAggregatorConnection)
	fmt.Printf("connectorEventSendToAggregator connect : " + connectorEventSendToAggregatorConnection)

	connectorEventRoutine.ConnectorEventReceiveFromWorkerConnection = connectorEventReceiveFromWorkerConnection
	connectorEventRoutine.ConnectorEventReceiveFromWorker, _ = connectorEventRoutine.Context.NewSocket(zmq4.XSUB)
	connectorEventRoutine.ConnectorEventReceiveFromWorker.SetIdentity(connectorEventRoutine.Identity)
	connectorEventRoutine.ConnectorEventReceiveFromWorker.Bind(connectorEventRoutine.ConnectorEventReceiveFromWorkerConnection)
	fmt.Printf("connectorEventReceiveFromWorker bind : " + connectorEventReceiveFromWorkerConnection)
	
	return
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
	//go r.cleanEventsByTimeout()

	poller := zmq4.NewPoller()
	poller.Add(r.ConnectorEventSendToWorker, zmq4.POLLIN)
	poller.Add(r.ConnectorEventReceiveFromAggregator, zmq4.POLLIN)
	poller.Add(r.ConnectorEventSendToAggregator, zmq4.POLLIN)
	poller.Add(r.ConnectorEventReceiveFromWorker, zmq4.POLLIN)

	event := [][]byte{}
	err := errors.New("")

	for {
		sockets, _ := poller.Poll(-1)
		for _, socket := range sockets {

			switch currentSocket := socket.Socket; currentSocket {
			case r.ConnectorEventSendToWorker:

				event, err = currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}
				err = r.processEventSendToWorker(event)
				if err != nil {
					panic(err)
				}

			case r.ConnectorEventReceiveFromAggregator:

				event, err = currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}
				err = r.processEventReceiveFromAggregator(event)
				if err != nil {
					panic(err)
				}

			case r.ConnectorEventSendToAggregator:

				event, err = currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}
				err = r.processEventSendToAggregator(event)
				if err != nil {
					panic(err)
				}

			case r.ConnectorEventReceiveFromWorker:

				event, err = currentSocket.RecvMessageBytes(0)
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

func (r ConnectorEventRoutine) processEventSendToWorker(event [][]byte) (err error) {
	eventMessage, _ := message.DecodeEventMessage(event[1])
	//r.addEvents(eventMessage)
	go eventMessage.SendEventWith(r.ConnectorEventReceiveFromAggregator)

	return
}

func (r ConnectorEventRoutine) processEventReceiveFromAggregator(event [][]byte) (err error) {
	eventMessage, _ := message.DecodeEventMessage(event[1])
	go eventMessage.SendEventWith(r.ConnectorEventSendToWorker)

	return
}

func (r ConnectorEventRoutine) processEventSendToAggregator(event [][]byte) (err error) {
	eventMessage, _ := message.DecodeEventMessage(event[1])
	go eventMessage.SendEventWith(r.ConnectorEventReceiveFromWorker)
	
	return
}

func (r ConnectorEventRoutine) processEventReceiveFromWorker(event [][]byte) (err error) {
	if  string(event[0]) == constant.EVENT_VALIDATION_TOPIC && string(event[1]) == constant.COMMAND_VALIDATION_FUNCTIONS {
		eventFunctions, _ := message.DecodeEventFunction(event[2])
		result, _ := r.validationEvents(eventFunctions.Worker, eventFunctions.Functions)
        if result {
			r.ConnectorMapWorkerEvents[eventFunctions.Worker] = eventFunctions.Functions
			eventFunctionReply := message.NewEventFunctionReply(result)
			go eventFunctionReply.SendEventFunctionReplyWith(r.ConnectorEventReceiveFromAggregator)
        }
	} else {
		eventMessage, _ := message.DecodeEventMessage(event[2])
		go eventMessage.SendEventWith(r.ConnectorEventSendToAggregator)
	}

	return
}

func (r ConnectorEventRoutine) validationEvents(workerSource string, events []string) (result bool, err error) {
	//TODO
	result = true
	return
}

/*func (r ConnectorEventRoutine) addEvents(eventMessage message.EventMessage) {
	if val, ok := r.ConnectorMapUUIDEventMessage[eventMessage.Uuid]; ok {
		if !ok {
			r.ConnectorMapUUIDEventMessage[eventMessage.Uuid] = eventMessage
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
}*/