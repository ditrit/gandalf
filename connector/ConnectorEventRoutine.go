package connector

import (
	"context"
	"errors"
	"fmt"
	"gandalf-go/message"
	"log"
	"net"

	pb "gandalf-go/grpc"

	"github.com/pebbe/zmq4"
	"google.golang.org/grpc"
)

type ConnectorEventRoutine struct {
	Context                                        *zmq4.Context
	ConnectorMapEventNameEventMessage              *Queue
	ConnectorMapWorkerEvents                       map[string][]string
	ConnectorMapWorkerIterators                    map[string][]*Iterator
	ConnectorEventWorkerConnection                 string
	ConnectorEventReceiveFromAggregator            *zmq4.Socket
	ConnectorEventReceiveFromAggregatorConnections []string
	ConnectorEventSendToAggregator                 *zmq4.Socket
	ConnectorEventSendToAggregatorConnections      []string
	ConnectorEventChannel                          chan message.EventMessage
	Identity                                       string
	ConnectorEventGrpcServer                       *grpc.Server
}

func NewConnectorEventRoutine(identity, connectorEventWorkerConnection string, connectorEventReceiveFromAggregatorConnections, connectorEventSendToAggregatorConnections []string) (connectorEventRoutine *ConnectorEventRoutine) {
	connectorEventRoutine = new(ConnectorEventRoutine)
	connectorEventRoutine.Identity = identity
	connectorEventRoutine.ConnectorMapWorkerIterators = make(map[string][]*Iterator)
	connectorEventRoutine.ConnectorEventChannel = make(chan message.EventMessage)
	connectorEventRoutine.ConnectorMapEventNameEventMessage = NewQueue()

	connectorEventRoutine.Context, _ = zmq4.NewContext()
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
	connectorEventRoutine.ConnectorEventWorkerConnection = connectorEventWorkerConnection
	connectorEventRoutine.StartGrpcServer(connectorEventRoutine.ConnectorEventWorkerConnection)
	fmt.Println("ConnectorEventWorkerConnection connect : " + connectorEventRoutine.ConnectorEventWorkerConnection)

	return
}

func (r ConnectorEventRoutine) close() {
	r.ConnectorEventReceiveFromAggregator.Close()
	r.ConnectorEventSendToAggregator.Close()
	r.Context.Term()
}

func (r ConnectorEventRoutine) reconnectToProxy() {

}

func (r ConnectorEventRoutine) run() {
	//go r.cleanEventsByTimeout()

	poller := zmq4.NewPoller()
	poller.Add(r.ConnectorEventReceiveFromAggregator, zmq4.POLLIN)
	poller.Add(r.ConnectorEventSendToAggregator, zmq4.POLLIN)

	event := [][]byte{}
	err := errors.New("")

	for {
		fmt.Println("Running ConnectorEventRoutine")
		sockets, _ := poller.Poll(-1)
		for _, socket := range sockets {
			fmt.Println("Running ConnectorEventRoutine2")
			switch currentSocket := socket.Socket; currentSocket {
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
			}
		}
	}
}

/* func (r ConnectorEventRoutine) processEventSendToWorker(event [][]byte) {
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
} */

func (r ConnectorEventRoutine) processEventReceiveFromAggregator(event [][]byte) {
	fmt.Println("processEventReceiveFromAggregator")
	fmt.Println(event)
	fmt.Println(len(event))
	fmt.Println(event[1])
	eventMessage, _ := message.DecodeEventMessage(event[1])
	fmt.Println(eventMessage)
	r.ConnectorMapEventNameEventMessage.Push(eventMessage)

	//go eventMessage.SendEventWith(r.ConnectorEventSendToWorker)
}

func (r ConnectorEventRoutine) processEventSendToAggregator(event [][]byte) {
	/* 	fmt.Println("processEventSendToAggregator")
	   	fmt.Println(event)
	   	if len(event) == 1 {
	   		topic := event[0]
	   		fmt.Println("SUB")
	   		fmt.Println("ConnectorEventReceiveFromWorker")
	   		fmt.Println(topic)
	   		fmt.Println(string(topic))

	   		//go message.SendSubscribeTopic(r.ConnectorEventReceiveFromWorker, topic)
	   	} */
	//eventMessage, _ := message.DecodeEventMessage(event[0])
	//go eventMessage.SendEventWith(r.ConnectorEventReceiveFromWorker)
}

/* func (r ConnectorEventRoutine) processEventReceiveFromWorker(event [][]byte) {
	//TODO REVOIR EVENT IF
	fmt.Println(event)
	fmt.Println(event[0])
	fmt.Println(event[1])
	if string(event[0]) == constant.EVENT_VALIDATION_FUNCTIONS {
		eventFunctions, _ := message.DecodeEventFunction(event[1])
		result, _ := r.validationEvents(eventFunctions.Worker, eventFunctions.Functions)
		if result {
			r.ConnectorMapWorkerEvents[eventFunctions.Worker] = eventFunctions.Functions
			eventFunctionReply := message.NewEventFunctionReply(result)
			go eventFunctionReply.SendMessageWith(r.ConnectorEventReceiveFromAggregator)
		}
	} else {
		fmt.Println("BLIP")
		fmt.Println(event[1])
		eventMessage, _ := message.DecodeEventMessage(event[1])
		fmt.Println(eventMessage)
		go eventMessage.SendMessageWith(r.ConnectorEventSendToAggregator)
	}
} */

func (r ConnectorEventRoutine) validationEvents(workerSource string, events []string) (result bool, err error) {
	//TODO
	result = true
	return
}

func (r ConnectorEventRoutine) runIterator(target, value string, iterator *Iterator, channel chan message.EventMessage) {
	notfound := true
	for notfound {
		messageIterator := iterator.Get()
		if messageIterator != nil {
			eventMessage := (*messageIterator).(message.EventMessage)
			if value == eventMessage.Event {
				fmt.Println("TOTO")
				channel <- eventMessage
				notfound = false
			}
		}
	}
	delete(r.ConnectorMapWorkerIterators, "target")
}

//GRPC
func (r ConnectorEventRoutine) StartGrpcServer(port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	r.ConnectorEventGrpcServer = grpc.NewServer()
	pb.RegisterConnectorEventServer(r.ConnectorEventGrpcServer, &r)
	go r.ConnectorEventGrpcServer.Serve(lis)
}

//TODO REVOIR SERVICE
func (r ConnectorEventRoutine) SendEventMessage(ctx context.Context, in *pb.EventMessage) (*pb.Empty, error) {

	eventMessage := new(message.EventMessage)
	eventMessage.FromGrpc(in)
	fmt.Println(eventMessage)
	go eventMessage.SendMessageWith(r.ConnectorEventSendToAggregator)
	return &pb.Empty{}, nil
}

func (r ConnectorEventRoutine) WaitEventMessage(ctx context.Context, in *pb.EventMessageWait) (messageEvent *pb.EventMessage, err error) {

	target := in.GetWorkerSource()
	iterator := NewIterator(r.ConnectorMapEventNameEventMessage)
	r.ConnectorMapWorkerIterators[in.GetEvent()] = append(r.ConnectorMapWorkerIterators[in.GetEvent()], iterator)

	fmt.Println("SUB")
	fmt.Println("ConnectorEventReceiveFromAggregator")

	go r.runIterator(target, in.GetEvent(), iterator, r.ConnectorEventChannel)
	select {
	case message := <-r.ConnectorEventChannel:
		fmt.Println("commandReply")
		message.ToGrpc(messageEvent)
		return
	default:
		fmt.Println("nope")
	}
	return
}
