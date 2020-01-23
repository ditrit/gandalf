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
	//go connectorEventRoutine.StartGrpcServer(connectorEventRoutine.ConnectorEventWorkerConnection)
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
	//poller.Add(r.ConnectorEventSendToAggregator, zmq4.POLLIN)

	event := [][]byte{}
	err := errors.New("")

	for {
		fmt.Println("Running ConnectorEventRoutine")
		sockets, _ := poller.Poll(-1)
		for _, socket := range sockets {
			switch currentSocket := socket.Socket; currentSocket {
			case r.ConnectorEventReceiveFromAggregator:
				fmt.Println("Receive Aggregator")
				event, err = currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}
				r.processEventReceiveFromAggregator(event)
				/* 	case r.ConnectorEventSendToAggregator:
					fmt.Println("Send Aggregator")
					event, err = currentSocket.RecvMessageBytes(0)
					if err != nil {
						panic(err)
					}
					r.processEventSendToAggregator(event)
				} */
			}
		}
	}
}

func (r ConnectorEventRoutine) processEventReceiveFromAggregator(event [][]byte) {
	eventMessage, _ := message.DecodeEventMessage(event[1])
	r.ConnectorMapEventNameEventMessage.Push(eventMessage)
}

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
				channel <- eventMessage
				notfound = false
			}
		}
	}
	delete(r.ConnectorMapWorkerIterators, "target")
}

//GRPC
func (r ConnectorEventRoutine) startGrpcServer() {
	lis, err := net.Listen("tcp", r.ConnectorEventWorkerConnection)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	r.ConnectorEventGrpcServer = grpc.NewServer()
	pb.RegisterConnectorEventServer(r.ConnectorEventGrpcServer, &r)
	r.ConnectorEventGrpcServer.Serve(lis)
}

//TODO REVOIR SERVICE
func (r ConnectorEventRoutine) SendEventMessage(ctx context.Context, in *pb.EventMessage) (*pb.Empty, error) {

	eventMessage := message.EventMessageFromGrpc(in)
	fmt.Println(eventMessage)
	go eventMessage.SendMessageWith(r.ConnectorEventSendToAggregator)
	return &pb.Empty{}, nil
}

func (r ConnectorEventRoutine) WaitEventMessage(ctx context.Context, in *pb.EventMessageWait) (messageEvent *pb.EventMessage, err error) {
	target := in.GetWorkerSource()
	iterator := NewIterator(r.ConnectorMapEventNameEventMessage)

	r.ConnectorMapWorkerIterators[in.GetEvent()] = append(r.ConnectorMapWorkerIterators[in.GetEvent()], iterator)

	go r.runIterator(target, in.GetEvent(), iterator, r.ConnectorEventChannel)
	select {
	case messageChannel := <-r.ConnectorEventChannel:
		messageEvent = message.EventMessageToGrpc(messageChannel)
		return
	}
	return
}
