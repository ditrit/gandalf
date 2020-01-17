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
	ConnectorCommandGrpcServer                       grpc.Server
}

func NewConnectorCommandRoutine(identity, connectorCommandSendToWorkerConnection, connectorCommandReceiveFromWorkerConnection string, connectorCommandReceiveFromAggregatorConnections, connectorCommandSendToAggregatorConnections []string) (connectorCommandRoutine *ConnectorCommandRoutine) {
	connectorCommandRoutine = new(ConnectorCommandRoutine)
	connectorCommandRoutine.Identity = identity
	connectorCommandRoutine.ConnectorMapWorkerIterators = make(map[string][]*Iterator)

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

	connectorCommandRoutine.StartGrpcServer(ConnectorCommandWorkerConnection)
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

func (r ConnectorCommandRoutine) processCommandSendToWorker(command [][]byte) {
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
}

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

func (r ConnectorCommandRoutine) processCommandReceiveFromWorker(command [][]byte) {
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
}

func (r ConnectorCommandRoutine) validationCommands(workerSource string, commands []string) (result bool) {
	//TODO
	result = true

	return
}

func (r ConnectorCommandRoutine) addCommands(commandMessage message.CommandMessage) {
	r.ConnectorMapCommandNameCommandMessage.Push(commandMessage)
}

//TODO
func (r ConnectorCommandRoutine) runIterator(target, commandType, value string, iterator *Iterator) (interface{}) {

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
			if commandType == constant.COMMAND_MESSAGE {
				commandMessage := (*messageIterator).(message.CommandMessage)
				if value == commandMessage.Command {
					return commandMessage
				}
			} else {
				commandMessageReply := (*messageIterator).(message.CommandMessageReply)
				if value == commandMessageReply.Uuid {
					return commandMessageReply
				}
			}
		}
		time.Sleep(time.Duration(2000 * time.Millisecond))

	}
	delete(r.ConnectorMapWorkerIterators, target)
}

//GRPC
func (r ConnectorCommandRoutine) StartGrpcServer(port string) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	r.ConnectorCommandGrpcServer = grpc.NewServer()
	pb.RegisterConnectorCommandServer(r.ConnectorCommandGrpcServer, &connectorCommandServer{})
	grpcServer.Serve(lis)
}

func (ccg *ConnectorCommandGrpc) SendCommandMessage(ctx context.Context, in *pb.CommandMessage) (*CommandMessageUUID, error) {
	commandMessage = new(message.CommandMessage)
	commandMessage.FromGrpc(in)
	go commandMessage.SendMessageWith(r.ConnectorCommandSendToAggregator)
	//TODO REPLACE TOTO BY UUID
	return &pb.CommandMessageUUID{uuid: commandMessage.Uuid}, nil
}

func (ccg *ConnectorCommandGrpc) SendCommandMessageReply(ctx context.Context, in *pb.CommandMessageReply) (*Empty, error) {
	commandMessageReply = new(message.CommandMessageReply)
	commandMessageReply.FromGrpc(in)
	go commandMessageReply.SendMessageWith(r.ConnectorCommandSendToAggregator)
	//TODO REPLACE TOTO BY UUID
	return &pb.Empty{}, nil
}

func (ccg *ConnectorCommandGrpc) WaitCommandMessage(ctx context.Context, in *pb.CommandMessageWait) (*pb.CommandMessage, error) {

	target := commandMessageWait.WorkerSource
	fmt.Println("QUEUE")
	fmt.Println(r.ConnectorMapCommandNameCommandMessage)
	iterator = NewIterator(r.ConnectorMapCommandNameCommandMessage)
	
	r.ConnectorMapWorkerIterators[target] = append(r.ConnectorMapWorkerIterators[target], iterator)

	//TODO REVOIR
	commandMessage := go r.runIterator(target, in.GetCommandType(), in.GetValue(), iterator).(message.CommandMessage)
	commandMessageReply.ToGrpc(message)
	return 
}

func (ccg *ConnectorCommandGrpc) WaitCommandMessageReply(ctx context.Context, in *pb.CommandMessageWait) (commandMessageReply *pb.CommandMessageReply, error) {

	target := in.GetWorkerSource()
	fmt.Println("QUEUE2")
	fmt.Println(r.ConnectorMapUUIDCommandMessageReply)
	iterator = NewIterator(r.ConnectorMapUUIDCommandMessageReply)

	r.ConnectorMapWorkerIterators[target] = append(r.ConnectorMapWorkerIterators[target], iterator)

		//TODO REVOIR
	commandMessageReply := go r.runIterator(target, in.GetCommandType(), in.GetValue(), iterator).(message.CommandMessageReply)
	commandMessageReply.ToGrpc(messageReply)
	return 
}
