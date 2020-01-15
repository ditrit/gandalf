package connector

import (
	"errors"
	"fmt"
	"gandalf-go/constant"
	"gandalf-go/message"

	"github.com/pebbe/zmq4"
)

type ConnectorEventRoutine struct {
	Context                                        *zmq4.Context
	ConnectorMapEventNameEventMessage              *Queue
	ConnectorMapWorkerEvents                       map[string][]string
	ConnectorMapWorkerIterators                    map[string][]*Iterator
	ConnectorEventSendToWorker                     *zmq4.Socket
	ConnectorEventSendToWorkerConnection           string
	ConnectorEventReceiveFromAggregator            *zmq4.Socket
	ConnectorEventReceiveFromAggregatorConnections []string
	ConnectorEventSendToAggregator                 *zmq4.Socket
	ConnectorEventSendToAggregatorConnections      []string
	ConnectorEventReceiveFromWorker                *zmq4.Socket
	ConnectorEventReceiveFromWorkerConnection      string
	Identity                                       string
}

func NewConnectorEventRoutine(identity, connectorEventSendToWorkerConnection, connectorEventReceiveFromWorkerConnection string, connectorEventReceiveFromAggregatorConnections, connectorEventSendToAggregatorConnections []string) (connectorEventRoutine *ConnectorEventRoutine) {
	connectorEventRoutine = new(ConnectorEventRoutine)
	connectorEventRoutine.Identity = identity
	connectorEventRoutine.ConnectorMapWorkerIterators = make(map[string][]*Iterator)
	connectorEventRoutine.ConnectorMapEventNameEventMessage = NewQueue()

	connectorEventRoutine.Context, _ = zmq4.NewContext()
	connectorEventRoutine.ConnectorEventSendToWorkerConnection = connectorEventSendToWorkerConnection
	connectorEventRoutine.ConnectorEventSendToWorker, _ = connectorEventRoutine.Context.NewSocket(zmq4.ROUTER)
	connectorEventRoutine.ConnectorEventSendToWorker.SetIdentity(connectorEventRoutine.Identity)
	connectorEventRoutine.ConnectorEventSendToWorker.Bind(connectorEventRoutine.ConnectorEventSendToWorkerConnection)
	fmt.Println("connectorEventSendToWorker bind : " + connectorEventSendToWorkerConnection)

	connectorEventRoutine.ConnectorEventReceiveFromWorkerConnection = connectorEventReceiveFromWorkerConnection
	connectorEventRoutine.ConnectorEventReceiveFromWorker, _ = connectorEventRoutine.Context.NewSocket(zmq4.XSUB)
	connectorEventRoutine.ConnectorEventReceiveFromWorker.SetIdentity(connectorEventRoutine.Identity)
	connectorEventRoutine.ConnectorEventReceiveFromWorker.Bind(connectorEventRoutine.ConnectorEventReceiveFromWorkerConnection)
	fmt.Println("connectorEventReceiveFromWorker bind : " + connectorEventReceiveFromWorkerConnection)
	connectorEventRoutine.ConnectorEventReceiveFromWorker.SendBytes([]byte{0x01}, 0) //SUBSCRIBE ALL

	connectorEventRoutine.ConnectorEventReceiveFromAggregatorConnections = connectorEventReceiveFromAggregatorConnections
	connectorEventRoutine.ConnectorEventReceiveFromAggregator, _ = connectorEventRoutine.Context.NewSocket(zmq4.XSUB)
	connectorEventRoutine.ConnectorEventReceiveFromAggregator.SetIdentity(connectorEventRoutine.Identity)
	//connectorEventRoutine.ConnectorEventReceiveFromAggregator.Connect(connectorEventRoutine.ConnectorEventReceiveFromAggregatorConnections)
	//fmt.Println("connectorEventReceiveFromAggregator connect : " + connectorEventReceiveFromAggregatorConnection)
	for _, connection := range connectorEventRoutine.ConnectorEventReceiveFromAggregatorConnections {
		connectorEventRoutine.ConnectorEventReceiveFromAggregator.Connect(connection)
		fmt.Println("connectorEventReceiveFromAggregatorConnections connect : " + connection)
	}
	connectorEventRoutine.ConnectorEventReceiveFromAggregator.SendBytes([]byte{0x01}, 0) //SUBSCRIBE ALL

	connectorEventRoutine.ConnectorEventSendToAggregatorConnections = connectorEventSendToAggregatorConnections
	connectorEventRoutine.ConnectorEventSendToAggregator, _ = connectorEventRoutine.Context.NewSocket(zmq4.XPUB)
	connectorEventRoutine.ConnectorEventSendToAggregator.SetIdentity(connectorEventRoutine.Identity)
	//connectorEventRoutine.ConnectorEventSendToAggregator.Connect(connectorEventRoutine.ConnectorEventSendToAggregatorConnection)
	//fmt.Println("connectorEventSendToAggregator connect : " + connectorEventSendToAggregatorConnection)
	for _, connection := range connectorEventRoutine.ConnectorEventSendToAggregatorConnections {
		connectorEventRoutine.ConnectorEventSendToAggregator.Connect(connection)
		fmt.Println("connectorEventSendToAggregator connect : " + connection)
	}

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
		fmt.Println("Running ConnectorEventRoutine")
		sockets, _ := poller.Poll(-1)
		for _, socket := range sockets {
			switch currentSocket := socket.Socket; currentSocket {
			case r.ConnectorEventSendToWorker:
				fmt.Println("SEND WORKER")
				event, err = currentSocket.RecvMessageBytes(0)
				fmt.Println(event)
				if err != nil {
					panic(err)
				}
				r.processEventSendToWorker(event)

			case r.ConnectorEventReceiveFromAggregator:
				fmt.Println("RECEIVER AGG")
				event, err = currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}
				r.processEventReceiveFromAggregator(event)
			case r.ConnectorEventSendToAggregator:
				fmt.Println("SEND AGG")
				event, err = currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}
				r.processEventSendToAggregator(event)
			case r.ConnectorEventReceiveFromWorker:
				fmt.Println("RECEIVER WORKER")
				event, err = currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}
				r.processEventReceiveFromWorker(event)

			}
		}
	}
}

func (r ConnectorEventRoutine) processEventSendToWorker(event [][]byte) {
	fmt.Println("processEventSendToWorker")
	fmt.Println(event)
	eventType := string(event[1])
	if eventType == constant.EVENT_WAIT {
		fmt.Println(string(event[0]))
		fmt.Println(string(event[1]))
		fmt.Println(string(event[2]))
		eventMessageWait, _ := message.DecodeEventMessageWait(event[2])
		target := eventMessageWait.WorkerSource
		iterator := NewIterator(r.ConnectorMapEventNameEventMessage)
		r.ConnectorMapWorkerIterators[eventMessageWait.Event] = append(r.ConnectorMapWorkerIterators[eventMessageWait.Event], iterator)

		fmt.Println("SUB")
		fmt.Println("ConnectorEventReceiveFromAggregator")
		fmt.Println(eventMessageWait.Topic)
		//go message.SendSubscribeTopic(r.ConnectorEventReceiveFromAggregator, []byte(eventMessageWait.Topic))

		go r.runIterator(target, eventMessageWait.Event, iterator)
	}
}

func (r ConnectorEventRoutine) processEventReceiveFromAggregator(event [][]byte) {
	eventMessage, _ := message.DecodeEventMessage(event[1])
	r.ConnectorMapEventNameEventMessage.Push(eventMessage)
	//go eventMessage.SendEventWith(r.ConnectorEventSendToWorker)
}

func (r ConnectorEventRoutine) processEventSendToAggregator(event [][]byte) {
	fmt.Println("processEventSendToAggregator")
	fmt.Println(event)
	if len(event) == 1 {
		topic := event[0]
		fmt.Println("SUB")
		fmt.Println("ConnectorEventReceiveFromWorker")
		fmt.Println(topic)
		fmt.Println(string(topic))

		//go message.SendSubscribeTopic(r.ConnectorEventReceiveFromWorker, topic)
	}
	//eventMessage, _ := message.DecodeEventMessage(event[0])
	//go eventMessage.SendEventWith(r.ConnectorEventReceiveFromWorker)
}

func (r ConnectorEventRoutine) processEventReceiveFromWorker(event [][]byte) {
	//TODO REVOIR EVENT IF
	fmt.Println(event[0])
	fmt.Println(event[1])
	if string(event[0]) == constant.EVENT_VALIDATION_FUNCTIONS {
		eventFunctions, _ := message.DecodeEventFunction(event[0])
		result, _ := r.validationEvents(eventFunctions.Worker, eventFunctions.Functions)
		if result {
			r.ConnectorMapWorkerEvents[eventFunctions.Worker] = eventFunctions.Functions
			eventFunctionReply := message.NewEventFunctionReply(result)
			go eventFunctionReply.SendMessageWith(r.ConnectorEventReceiveFromAggregator)
		}
	} else {
		eventMessage, _ := message.DecodeEventMessage(event[1])
		go eventMessage.SendMessageWith(r.ConnectorEventSendToAggregator)
	}
}

func (r ConnectorEventRoutine) validationEvents(workerSource string, events []string) (result bool, err error) {
	//TODO
	result = true
	return
}

func (r ConnectorEventRoutine) runIterator(target, value string, iterator *Iterator) {
	notfound := true
	for notfound {
		messageIterator := iterator.Get()
		if messageIterator != nil {
			eventMessage := (*messageIterator).(message.EventMessage)
			if value == eventMessage.Event {
				fmt.Println("TOTO")
				eventMessage.SendWith(r.ConnectorEventSendToWorker, target)
				notfound = false
			}
		}
	}
	delete(r.ConnectorMapWorkerIterators, "target")
}

/*func (r ConnectorEventRoutine) addEvents(eventMessage message.EventMessage) {
	if val, ok := r.ConnectorMapEventNameEventMessage[eventMessage.Uuid]; ok {
		if !ok {
			r.ConnectorMapEventNameEventMessage[eventMessage.Uuid] = eventMessage
		}
	}
}

func (r ConnectorEventRoutine) cleanEventsByTimeout() {
	maxTimeout = 0
	for {
		for uuid, eventMessage := range r.ConnectorMapEventNameEventMessage {
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
