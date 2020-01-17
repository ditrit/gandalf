package connector

import (
	"context"
	"errors"
	"fmt"
	"gandalf-go/constant"
	"gandalf-go/message"
	"log"
	"net"
	"time"

	pb "gandalf-go/grpc"

	"github.com/pebbe/zmq4"
	"google.golang.org/grpc"
)

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

func NewConnectorCommandRoutine(identity, connectorCommandWorkerConnection string, connectorCommandReceiveFromAggregatorConnections, connectorCommandSendToAggregatorConnections []string) (connectorCommandRoutine *ConnectorCommandRoutine) {
	connectorCommandRoutine = new(ConnectorCommandRoutine)
	connectorCommandRoutine.Identity = identity
	connectorCommandRoutine.ConnectorMapWorkerIterators = make(map[string][]*Iterator)
	connectorCommandRoutine.ConnectorCommandChannel = make(chan message.CommandMessage)
	connectorCommandRoutine.ConnectorCommandReplyChannel = make(chan message.CommandMessageReply)

	connectorCommandRoutine.ConnectorMapCommandNameCommandMessage = NewQueue()
	connectorCommandRoutine.ConnectorMapUUIDCommandMessageReply = NewQueue()

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
	connectorCommandRoutine.StartGrpcServer(connectorCommandRoutine.ConnectorCommandWorkerConnection)
	return
}

func (r ConnectorCommandRoutine) close() {
	r.ConnectorCommandReceiveFromAggregator.Close()
	r.ConnectorCommandSendToAggregator.Close()
	r.Context.Term()
}

func (r ConnectorCommandRoutine) reconnectToProxy() {

}

func (r ConnectorCommandRoutine) run() {
	//go r.cleanCommandsByTimeout()

	poller := zmq4.NewPoller()
	poller.Add(r.ConnectorCommandReceiveFromAggregator, zmq4.POLLIN)

	command := [][]byte{}
	err := errors.New("")

	for {
		fmt.Println("Running ConnectorCommandRoutine")
		sockets, _ := poller.Poll(-1)
		for _, socket := range sockets {
			fmt.Println("Running ConnectorCommandRoutine2")

			switch currentSocket := socket.Socket; currentSocket {
			case r.ConnectorCommandReceiveFromAggregator:
				fmt.Println("Connector receive aggregator")

				command, err = currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}
				r.processCommandReceiveFromAggregator(command)
			}
		}
	}
}

/* func (r ConnectorCommandRoutine) processCommandSendToWorker(command [][]byte) {
	fmt.Println("WAIIIITTTT")
	commandType := string(command[0])
	fmt.Println(commandType)
	if commandType == constant.COMMAND_WAIT {
		commandMessageWait, _ := message.DecodeCommandMessageWait(command[1])
		target := commandMessageWait.WorkerSource
		var iterator *Iterator
		if commandMessageWait.CommandType == constant.COMMAND_MESSAGE {
			fmt.Println("QUEUE")
			fmt.Println(r.ConnectorMapCommandNameCommandMessage)
			iterator = NewIterator(r.ConnectorMapCommandNameCommandMessage)
		} else {
			fmt.Println("QUEUE2")
			fmt.Println(r.ConnectorMapUUIDCommandMessageReply)
			iterator = NewIterator(r.ConnectorMapUUIDCommandMessageReply)
		}
		r.ConnectorMapWorkerIterators[target] = append(r.ConnectorMapWorkerIterators[target], iterator)

		go r.runIterator(target, commandMessageWait.CommandType, commandMessageWait.Value, iterator)
	}
} */

func (r ConnectorCommandRoutine) processCommandReceiveFromAggregator(command [][]byte) {
	fmt.Println("CMD")
	fmt.Println(command)
	fmt.Println(string(command[0]))
	fmt.Println(string(command[1]))
	fmt.Println(string(command[2]))
	commandType := string(command[2])
	if commandType == constant.COMMAND_MESSAGE {
		commandMessage, _ := message.DecodeCommandMessage(command[3])
		fmt.Println("QUEUE CMD")
		toto := r.ConnectorMapCommandNameCommandMessage
		fmt.Println(&toto)
		r.ConnectorMapCommandNameCommandMessage.Print()
		r.ConnectorMapCommandNameCommandMessage.Push(commandMessage)
		fmt.Println("QUEUE CMD2")
		toto = r.ConnectorMapCommandNameCommandMessage
		fmt.Println(&toto)
	} else {
		fmt.Println("QUEUE REPLY")
		r.ConnectorMapUUIDCommandMessageReply.Print()
		commandMessageReply, _ := message.DecodeCommandMessageReply(command[3])
		r.ConnectorMapUUIDCommandMessageReply.Push(commandMessageReply)
		fmt.Println("QUEUE REPLY")
		r.ConnectorMapUUIDCommandMessageReply.Print()
	}
}

/* func (r ConnectorCommandRoutine) processCommandReceiveFromWorker(command [][]byte) {
	workerSource := string(command[0])
	commandHeader := string(command[1])
	fmt.Println(workerSource)
	fmt.Println(commandHeader)
	if commandHeader == constant.COMMAND_VALIDATION_FUNCTIONS {
		commandFunction, _ := message.DecodeCommandFunction(command[2])
		result := r.validationCommands(workerSource, commandFunction.Functions)
		if result {
			r.ConnectorMapWorkerCommands[workerSource] = commandFunction.Functions
			commandFunctionReply := message.NewCommandFunctionReply(result)
			go commandFunctionReply.SendMessageWith(r.ConnectorCommandSendToWorker)
		}
	} else if commandHeader == constant.COMMAND_MESSAGE {
		commandMessage, _ := message.DecodeCommandMessage(command[2])
		commandMessage.SourceWorker = workerSource
		fmt.Println("MESSAGE")
		fmt.Println(commandMessage)
		go commandMessage.SendMessageWith(r.ConnectorCommandSendToAggregator)
	} else if commandHeader == constant.COMMAND_MESSAGE_REPLY {
		commandMessageReply, _ := message.DecodeCommandMessageReply(command[2])
		commandMessageReply.SourceWorker = workerSource
		go commandMessageReply.SendMessageWith(r.ConnectorCommandSendToAggregator)
	} else {
		//COMMAND WAIT
	}
} */

func (r ConnectorCommandRoutine) validationCommands(workerSource string, commands []string) (result bool) {
	//TODO
	result = true

	return
}

func (r ConnectorCommandRoutine) addCommands(commandMessage message.CommandMessage) {
	r.ConnectorMapCommandNameCommandMessage.Push(commandMessage)
}

//TODO
func (r ConnectorCommandRoutine) runIteratorCommandMessage(target, value string, iterator *Iterator, channel chan message.CommandMessage) {

	notfound := true
	for notfound {
		fmt.Println("ITERATOR PRINT QUEUE")
		toto := iterator.GetQueue()
		fmt.Println(&toto)
		fmt.Println(iterator.GetQueue())
		iterator.PrintQueue()

		messageIterator := iterator.Get()
		fmt.Println("GET")
		fmt.Println(messageIterator)
		if messageIterator != nil {
			commandMessage := (*messageIterator).(message.CommandMessage)
			if value == commandMessage.Command {
				channel <- commandMessage
				notfound = false
			}
		}
		time.Sleep(time.Duration(2000 * time.Millisecond))

	}
	delete(r.ConnectorMapWorkerIterators, target)
}

//TODO
func (r ConnectorCommandRoutine) runIteratorCommandMessageReply(target, value string, iterator *Iterator, channel chan message.CommandMessageReply) {

	notfound := true
	for notfound {
		fmt.Println("ITERATOR PRINT QUEUE")
		toto := iterator.GetQueue()
		fmt.Println(&toto)
		fmt.Println(iterator.GetQueue())
		iterator.PrintQueue()

		messageIterator := iterator.Get()
		fmt.Println("GET")
		fmt.Println(messageIterator)
		if messageIterator != nil {
			commandMessageReply := (*messageIterator).(message.CommandMessageReply)
			if value == commandMessageReply.Uuid {
				channel <- commandMessageReply
				notfound = false
			}

		}
		time.Sleep(time.Duration(2000 * time.Millisecond))

	}
	delete(r.ConnectorMapWorkerIterators, target)
}

//GRPC
func (r ConnectorCommandRoutine) StartGrpcServer(port string) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	r.ConnectorCommandGrpcServer = grpc.NewServer()
	pb.RegisterConnectorCommandServer(r.ConnectorCommandGrpcServer, &r)
	r.ConnectorCommandGrpcServer.Serve(lis)
}

func (r ConnectorCommandRoutine) SendCommandMessage(ctx context.Context, in *pb.CommandMessage) (*pb.CommandMessageUUID, error) {
	commandMessage := new(message.CommandMessage)
	commandMessage.FromGrpc(in)
	go commandMessage.SendMessageWith(r.ConnectorCommandSendToAggregator)
	return &pb.CommandMessageUUID{Uuid: commandMessage.Uuid}, nil
}

func (r ConnectorCommandRoutine) SendCommandMessageReply(ctx context.Context, in *pb.CommandMessageReply) (*pb.Empty, error) {
	commandMessageReply := new(message.CommandMessageReply)
	commandMessageReply.FromGrpc(in)
	go commandMessageReply.SendMessageWith(r.ConnectorCommandSendToAggregator)
	return &pb.Empty{}, nil
}

func (r ConnectorCommandRoutine) WaitCommandMessage(ctx context.Context, in *pb.CommandMessageWait) (commandMessage *pb.CommandMessage, err error) {

	target := in.GetWorkerSource()
	fmt.Println("QUEUE")
	fmt.Println(r.ConnectorMapCommandNameCommandMessage)
	iterator := NewIterator(r.ConnectorMapCommandNameCommandMessage)

	r.ConnectorMapWorkerIterators[target] = append(r.ConnectorMapWorkerIterators[target], iterator)

	go r.runIteratorCommandMessage(target, in.GetValue(), iterator, r.ConnectorCommandChannel)
	select {
	case message := <-r.ConnectorCommandChannel:
		fmt.Println("command")
		commandMessage = message.ToGrpc()
		return
	default:
		fmt.Println("nope")
	}
	return
}

func (r ConnectorCommandRoutine) WaitCommandMessageReply(ctx context.Context, in *pb.CommandMessageWait) (commandMessageReply *pb.CommandMessageReply, err error) {

	target := in.GetWorkerSource()
	fmt.Println("QUEUE2")
	fmt.Println(r.ConnectorMapUUIDCommandMessageReply)
	iterator := NewIterator(r.ConnectorMapUUIDCommandMessageReply)

	r.ConnectorMapWorkerIterators[target] = append(r.ConnectorMapWorkerIterators[target], iterator)

	go r.runIteratorCommandMessageReply(target, in.GetValue(), iterator, r.ConnectorCommandReplyChannel)
	select {
	case messageReply := <-r.ConnectorCommandReplyChannel:
		fmt.Println("commandReply")
		commandMessageReply = messageReply.ToGrpc()
		return
	default:
		fmt.Println("nope")
	}
	return
}
