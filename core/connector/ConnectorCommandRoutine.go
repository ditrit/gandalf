//Package connector :
//File ConnectorCommandRoutine.go
package connector

import (
	"context"
	"fmt"
	"gandalf-go/commons/message"
	constant "gandalf-go/core"
	"log"
	"net"
	"time"

	pb "gandalf-go/commons/grpc"

	"github.com/pebbe/zmq4"
	"google.golang.org/grpc"
)

//ConnectorCommandRoutine :
type ConnectorCommandRoutine struct {
	Context                                          *zmq4.Context
	ConnectorCommandWorkerConnection                 string
	ConnectorCommandReceiveFromAggregator            *zmq4.Socket
	ConnectorCommandReceiveFromAggregatorConnections []string
	ConnectorCommandSendToAggregator                 *zmq4.Socket
	ConnectorCommandSendToAggregatorConnections      []string
	Identity                                         string
	ConnectorMapCommandNameCommandMessage            *Queue
	ConnectorMapUUIDCommandMessageReply              *Queue
	ConnectorMapWorkerCommands                       map[string][]string
	ConnectorMapWorkerIterators                      map[string][]*Iterator
	ConnectorCommandGrpcServer                       *grpc.Server
	ConnectorCommandChannel                          chan message.CommandMessage
	ConnectorCommandReplyChannel                     chan message.CommandMessageReply
}

//NewConnectorCommandRoutine :
func NewConnectorCommandRoutine(
	identity string,
	connectorCommandWorkerConnection string,
	connectorCommandReceiveFromAggregatorConnections []string,
	connectorCommandSendToAggregatorConnections []string) *ConnectorCommandRoutine {
	connectorCommandRoutine := new(ConnectorCommandRoutine)
	connectorCommandRoutine.Identity = identity
	connectorCommandRoutine.ConnectorMapWorkerIterators = make(map[string][]*Iterator)
	connectorCommandRoutine.ConnectorCommandChannel = make(chan message.CommandMessage)
	connectorCommandRoutine.ConnectorCommandReplyChannel = make(chan message.CommandMessageReply)

	connectorCommandRoutine.ConnectorMapCommandNameCommandMessage = NewQueue()
	connectorCommandRoutine.ConnectorMapUUIDCommandMessageReply = NewQueue()

	connectorCommandRoutine.Context, _ = zmq4.NewContext()
	connectorCommandRoutine.ConnectorCommandReceiveFromAggregatorConnections = connectorCommandReceiveFromAggregatorConnections
	connectorCommandRoutine.ConnectorCommandReceiveFromAggregator, _ = connectorCommandRoutine.Context.NewSocket(zmq4.ROUTER)
	connectorCommandRoutine.ConnectorCommandReceiveFromAggregator.SetIdentity(connectorCommandRoutine.Identity)

	for _, connection := range connectorCommandRoutine.ConnectorCommandReceiveFromAggregatorConnections {
		connectorCommandRoutine.ConnectorCommandReceiveFromAggregator.Connect(connection)
		fmt.Println("connectorCommandReceiveFromAggregator connect : " + connection)
	}

	connectorCommandRoutine.ConnectorCommandSendToAggregatorConnections = connectorCommandSendToAggregatorConnections
	connectorCommandRoutine.ConnectorCommandSendToAggregator, _ = connectorCommandRoutine.Context.NewSocket(zmq4.DEALER)
	connectorCommandRoutine.ConnectorCommandSendToAggregator.SetIdentity(connectorCommandRoutine.Identity)

	for _, connection := range connectorCommandRoutine.ConnectorCommandSendToAggregatorConnections {
		connectorCommandRoutine.ConnectorCommandSendToAggregator.Connect(connection)
		fmt.Println("connectorCommandSendToAggregator connect : " + connection)
	}

	connectorCommandRoutine.ConnectorCommandWorkerConnection = connectorCommandWorkerConnection
	//go connectorCommandRoutine.StartGrpcServer(connectorCommandRoutine.ConnectorCommandWorkerConnection)
	fmt.Println("ConnectorCommandWorkerConnection connect : " + connectorCommandRoutine.ConnectorCommandWorkerConnection)

	return connectorCommandRoutine
}

//close :
func (r ConnectorCommandRoutine) close() {
	r.ConnectorCommandReceiveFromAggregator.Close()
	r.ConnectorCommandSendToAggregator.Close()
	r.Context.Term()
}

// TODO : implement
// func (r ConnectorCommandRoutine) reconnectToProxy() {

// }

//run :
func (r ConnectorCommandRoutine) run() {
	//go r.cleanCommandsByTimeout()
	poller := zmq4.NewPoller()

	poller.Add(r.ConnectorCommandReceiveFromAggregator, zmq4.POLLIN)

	for {
		fmt.Println("Running ConnectorCommandRoutine")

		sockets, _ := poller.Poll(-1)

		for _, socket := range sockets {
			currentSocket := socket.Socket

			if currentSocket == r.ConnectorCommandReceiveFromAggregator {
				fmt.Println("Receive Aggregator")

				command, err := currentSocket.RecvMessageBytes(0)

				if err != nil {
					panic(err)
				}

				r.processCommandReceiveFromAggregator(command)
			}
		}
	}
}

//processCommandReceiveFromAggregator :
func (r ConnectorCommandRoutine) processCommandReceiveFromAggregator(command [][]byte) {
	commandType := string(command[2])
	if commandType == constant.COMMAND_MESSAGE {
		commandMessage, _ := message.DecodeCommandMessage(command[3])
		r.ConnectorMapCommandNameCommandMessage.Push(commandMessage)
	} else {
		r.ConnectorMapUUIDCommandMessageReply.Print()
		commandMessageReply, _ := message.DecodeCommandMessageReply(command[3])
		r.ConnectorMapUUIDCommandMessageReply.Push(commandMessageReply)
		r.ConnectorMapUUIDCommandMessageReply.Print()
	}
}

//validationCommands :
func (r ConnectorCommandRoutine) validationCommands(workerSource string, commands []string) (result bool) {
	//TODO
	result = true

	return
}

//addCommands :
func (r ConnectorCommandRoutine) addCommands(commandMessage message.CommandMessage) {
	r.ConnectorMapCommandNameCommandMessage.Push(commandMessage)
}

//runIteratorCommandMessage :
//TODO
func (r ConnectorCommandRoutine) runIteratorCommandMessage(target, value string, iterator *Iterator, channel chan message.CommandMessage) {
	notfound := true
	for notfound {
		iterator.PrintQueue()
		messageIterator := iterator.Get()

		if messageIterator != nil {
			commandMessage := (*messageIterator).(message.CommandMessage)

			if value == commandMessage.Command {
				channel <- commandMessage

				notfound = false
			}
		}

		time.Sleep(time.Duration(2000) * time.Millisecond)
	}
	delete(r.ConnectorMapWorkerIterators, target)
}

//runIteratorCommandMessageReply :
//TODO
func (r ConnectorCommandRoutine) runIteratorCommandMessageReply(target, value string, iterator *Iterator, channel chan message.CommandMessageReply) {
	notfound := true
	for notfound {
		iterator.PrintQueue()

		messageIterator := iterator.Get()

		if messageIterator != nil {
			commandMessageReply := (*messageIterator).(message.CommandMessageReply)
			if value == commandMessageReply.UUID {
				channel <- commandMessageReply

				notfound = false
			}
		}

		time.Sleep(time.Duration(2000) * time.Millisecond)
	}
	delete(r.ConnectorMapWorkerIterators, target)
}

//startGrpcServer :
//GRPC
func (r ConnectorCommandRoutine) startGrpcServer() {
	lis, err := net.Listen("tcp", r.ConnectorCommandWorkerConnection)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	r.ConnectorCommandGrpcServer = grpc.NewServer()

	pb.RegisterConnectorCommandServer(r.ConnectorCommandGrpcServer, &r)
	r.ConnectorCommandGrpcServer.Serve(lis)
}

//SendCommandMessage :
func (r ConnectorCommandRoutine) SendCommandMessage(ctx context.Context, in *pb.CommandMessage) (*pb.CommandMessageUUID, error) {
	commandMessage := message.CommandMessageFromGrpc(in)

	go commandMessage.SendMessageWith(r.ConnectorCommandSendToAggregator)

	return &pb.CommandMessageUUID{UUID: commandMessage.UUID}, nil
}

//SendCommandMessageReply :
func (r ConnectorCommandRoutine) SendCommandMessageReply(ctx context.Context, in *pb.CommandMessageReply) (*pb.Empty, error) {
	commandMessageReply := message.CommandMessageReplyFromGrpc(in)

	go commandMessageReply.SendMessageWith(r.ConnectorCommandSendToAggregator)

	return &pb.Empty{}, nil
}

//WaitCommandMessage :
func (r ConnectorCommandRoutine) WaitCommandMessage(ctx context.Context, in *pb.CommandMessageWait) (commandMessage *pb.CommandMessage, err error) {
	target := in.GetWorkerSource()
	iterator := NewIterator(r.ConnectorMapCommandNameCommandMessage)

	r.ConnectorMapWorkerIterators[target] = append(r.ConnectorMapWorkerIterators[target], iterator)

	go r.runIteratorCommandMessage(target, in.GetValue(), iterator, r.ConnectorCommandChannel)
	messageChannel := <-r.ConnectorCommandChannel
	commandMessage = message.CommandMessageToGrpc(messageChannel)

	return
}

//WaitCommandMessageReply :
func (r ConnectorCommandRoutine) WaitCommandMessageReply(ctx context.Context, in *pb.CommandMessageWait) (commandMessageReply *pb.CommandMessageReply, err error) {
	target := in.GetWorkerSource()
	iterator := NewIterator(r.ConnectorMapUUIDCommandMessageReply)

	r.ConnectorMapWorkerIterators[target] = append(r.ConnectorMapWorkerIterators[target], iterator)

	go r.runIteratorCommandMessageReply(target, in.GetValue(), iterator, r.ConnectorCommandReplyChannel)
	messageReplyChannel := <-r.ConnectorCommandReplyChannel
	commandMessageReply = message.CommandMessageReplyToGrpc(messageReplyChannel)

	return
}
