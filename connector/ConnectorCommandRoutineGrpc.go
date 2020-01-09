package connector

import (
	"errors"
	"fmt"
	"gandalf-go/constant"
	"gandalf-go/message"
	"strconv"
	"time"

	"github.com/pebbe/zmq4"
)

type ConnectorCommandRoutine struct {
	Context                                          *zmq4.Context
	ConnectorCommandSendToWorker                     *zmq4.Socket
	ConnectorCommandSendToWorkerConnection           string
	ConnectorCommandReceiveFromAggregator            *zmq4.Socket
	ConnectorCommandReceiveFromAggregatorConnections []string
	ConnectorCommandSendToAggregator                 *zmq4.Socket
	ConnectorCommandSendToAggregatorConnections      []string
	ConnectorCommandReceiveFromWorker                *zmq4.Socket
	ConnectorCommandReceiveFromWorkerConnection      string
	Identity                                         string
	ConnectorMapUUIDCommandMessage                   map[string][]message.CommandMessage
	ConnectorMapUUIDCommandMessageReply              map[string][]message.CommandMessageReply
	ConnectorMapWorkerCommands                       map[string][]string
}

func NewConnectorCommandRoutine(identity, connectorCommandSendToWorkerConnection, connectorCommandReceiveFromWorkerConnection string, connectorCommandReceiveFromAggregatorConnections, connectorCommandSendToAggregatorConnections []string) (connectorCommandRoutine *ConnectorCommandRoutine) {
	connectorCommandRoutine = new(ConnectorCommandRoutine)
	connectorCommandRoutine.Identity = identity
	connectorCommandRoutine.ConnectorMapUUIDCommandMessage = make(map[string][]message.CommandMessage)
	connectorCommandRoutine.ConnectorMapWorkerCommands = make(map[string][]string)

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

	connectorCommandRoutine.Context, _ = zmq4.NewContext()
	connectorCommandRoutine.ConnectorCommandSendToWorkerConnection = connectorCommandSendToWorkerConnection
	connectorCommandRoutine.ConnectorCommandSendToWorker, _ = connectorCommandRoutine.Context.NewSocket(zmq4.DEALER)
	connectorCommandRoutine.ConnectorCommandSendToWorker.SetIdentity(connectorCommandRoutine.Identity)
	connectorCommandRoutine.ConnectorCommandSendToWorker.Bind(connectorCommandRoutine.ConnectorCommandSendToWorkerConnection)
	fmt.Println("connectorCommandSendToWorker bind : " + connectorCommandSendToWorkerConnection)

	connectorCommandRoutine.ConnectorCommandReceiveFromWorkerConnection = connectorCommandReceiveFromWorkerConnection
	connectorCommandRoutine.ConnectorCommandReceiveFromWorker, _ = connectorCommandRoutine.Context.NewSocket(zmq4.ROUTER)
	connectorCommandRoutine.ConnectorCommandReceiveFromWorker.SetIdentity(connectorCommandRoutine.Identity)
	connectorCommandRoutine.ConnectorCommandReceiveFromWorker.Bind(connectorCommandRoutine.ConnectorCommandReceiveFromWorkerConnection)
	fmt.Println("connectorCommandReceiveFromWorker bind : " + connectorCommandReceiveFromWorkerConnection)

	return
}

func (r ConnectorCommandRoutine) close() {
	r.ConnectorCommandSendToWorker.Close()
	r.ConnectorCommandReceiveFromAggregator.Close()
	r.ConnectorCommandSendToAggregator.Close()
	r.ConnectorCommandReceiveFromWorker.Close()
	r.Context.Term()
}

func (r ConnectorCommandRoutine) reconnectToProxy() {

}

func (r ConnectorCommandRoutine) run() {
	//go r.cleanCommandsByTimeout()

	poller := zmq4.NewPoller()
	poller.Add(r.ConnectorCommandReceiveFromAggregator, zmq4.POLLIN)
	poller.Add(r.ConnectorCommandSendToAggregator, zmq4.POLLIN)
	//poller.Add(r.ConnectorCommandSendToWorker, zmq4.POLLIN)
	//poller.Add(r.ConnectorCommandReceiveFromWorker, zmq4.POLLIN)

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

			case r.ConnectorCommandSendToAggregator:
				fmt.Println("Connector send aggregator")

				command, err = currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}
				r.processCommandSendAggregator(command)

/* 			case r.ConnectorCommandReceiveFromWorker:
				fmt.Println("Connector receive worker")
				command, err = currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}
				r.processCommandReceiveFromWorker(command)
			case r.ConnectorCommandSendToWorker:
				fmt.Println("Connector send worker")

				command, err = currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}
				r.processCommandSendToWorker(command) */
			}
		}
	}
}

func startCommandGRPCServer(address string) error {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	s := grpc.ConnectorServer{}

	grpcServer := grpc.NewServer(opts...)

	grpc.RegisterConnectorCommandServer(grpcServer, &s)

	log.Printf("starting HTTP/2 gRPC server on %s", address)
	if err := grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %s", err)
	}

	return nil
}



func (r ConnectorCommandRoutine) processCommandSendToWorker(command [][]byte) {
	commandMessage, _ := message.DecodeCommandMessage(command[1])
	r.addCommands(commandMessage)
	go commandMessage.SendWith(r.ConnectorCommandReceiveFromAggregator, commandMessage.SourceConnector)
}

func (r ConnectorCommandRoutine) processCommandReceiveFromAggregator(command [][]byte) {
	fmt.Println("CMD")
	fmt.Println(command)
	fmt.Println(string(command[0]))
	fmt.Println(string(command[1]))
	commandMessage, _ := message.DecodeCommandMessage(command[2])
	fmt.Println("BEFORE")
	fmt.Println(r.ConnectorMapUUIDCommandMessage[commandMessage.Command])
	r.ConnectorMapUUIDCommandMessage[commandMessage.Command] = append(r.ConnectorMapUUIDCommandMessage[commandMessage.Command], commandMessage)
	fmt.Println("AFTER")
	fmt.Println(r.ConnectorMapUUIDCommandMessage[commandMessage.Command])
}

func (r ConnectorCommandRoutine) processCommandSendAggregator(command [][]byte) {
	commandMessage, _ := message.DecodeCommandMessage(command[1])
	go commandMessage.SendCommandWith(r.ConnectorCommandReceiveFromWorker)
}

func (r ConnectorCommandRoutine) processCommandReceiveFromWorker(command [][]byte) {
	workerSource := string(command[0])
	commandHeader := string(command[1])

	if commandHeader == constant.COMMAND_READY {
		//commandReady := decodeCommandReady(command[2])
		commandMessage, err := r.getCommandByWorkerCommands(workerSource)
		if err != nil {
		}
		go commandMessage.SendWith(r.ConnectorCommandSendToWorker, workerSource)
	} else if commandHeader == constant.COMMAND_VALIDATION_FUNCTIONS {
		commandFunction, _ := message.DecodeCommandFunction(command[2])
		result := r.validationCommands(workerSource, commandFunction.Functions)
		if result {
			r.ConnectorMapWorkerCommands[workerSource] = commandFunction.Functions
			commandFunctionReply := message.NewCommandFunctionReply(result)
			go commandFunctionReply.SendCommandFunctionReplyWith(r.ConnectorCommandSendToWorker)
		}
	} else {
		commandMessage, _ := message.DecodeCommandMessage(command[1])
		commandMessage.SourceWorker = workerSource
		go commandMessage.SendWith(r.ConnectorCommandSendToAggregator, workerSource)
	}
}

func (r ConnectorCommandRoutine) getCommandByWorkerCommands(worker string) (commandMessage message.CommandMessage, err error) {

	var maxCommand string
	maxTimestamp := -1
	currentTimestamp := -1
	commandsWorker := r.ConnectorMapWorkerCommands[worker]
	var commands []message.CommandMessage

	for i, commandWorker := range commandsWorker {
		if currentCommandWorker, ok := r.ConnectorMapUUIDCommandMessage[commandWorker]; ok {
			commands[i] = currentCommandWorker[0]
		}
	}

	for _, command := range commands {
		currentTimestamp, _ = strconv.Atoi(command.Timestamp)
		if currentTimestamp >= maxTimestamp {
			maxTimestamp, _ = strconv.Atoi(command.Timestamp)
			maxCommand = command.Command
		}
	}

	commandMessage = r.ConnectorMapUUIDCommandMessage[maxCommand][0]
	delete(r.ConnectorMapUUIDCommandMessage, maxCommand)

	return
}

func (r ConnectorCommandRoutine) getConnectorMapUUIDCommandMessage(command string) (commandMessage []message.CommandMessage) {
	if commandMessage, ok := r.ConnectorMapUUIDCommandMessage[command]; ok {
		if ok {
			return commandMessage
		}
	}
	return
}

func (r ConnectorCommandRoutine) validationCommands(workerSource string, commands []string) (result bool) {
	//TODO
	result = true

	return
}

func (r ConnectorCommandRoutine) addCommands(commandMessage message.CommandMessage) {
	if _, ok := r.ConnectorMapUUIDCommandMessage[commandMessage.Uuid]; ok {
		if !ok {
			r.ConnectorMapUUIDCommandMessage[commandMessage.Uuid] = append(r.ConnectorMapUUIDCommandMessage[commandMessage.Uuid], commandMessage)
		}
	}
}

func (r ConnectorCommandRoutine) cleanCommandsByTimeout() {
	maxTimeout := 0
	currentTimestamp := -1
	currentTimeout := -1
	for {
		for uuid, commandMessageSlice := range r.ConnectorMapUUIDCommandMessage {
			for _, commandMessage := range commandMessageSlice {
				currentTimestamp, _ = strconv.Atoi(commandMessage.Timestamp)
				currentTimeout, _ = strconv.Atoi(commandMessage.Timeout)

				if currentTimestamp-currentTimeout == 0 {
					delete(r.ConnectorMapUUIDCommandMessage, uuid)
				} else {
					if currentTimeout >= maxTimeout {
						maxTimeout = currentTimeout
					}
				}
			}
		}
		time.Sleep(time.Duration(maxTimeout) * time.Millisecond)
	}
}
