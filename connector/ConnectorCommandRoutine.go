package connector

import (
	"errors"
	"fmt"
	"gandalf-go/constant"
	"gandalf-go/message"

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
	ConnectorMapCommandNameCommandMessage            *Queue
	ConnectorMapUUIDCommandMessageReply              *Queue
	ConnectorMapWorkerCommands                       map[string][]string
	ConnectorMapWorkerIterators                      map[string][]*Iterator
}

func NewConnectorCommandRoutine(identity, connectorCommandSendToWorkerConnection, connectorCommandReceiveFromWorkerConnection string, connectorCommandReceiveFromAggregatorConnections, connectorCommandSendToAggregatorConnections []string) (connectorCommandRoutine *ConnectorCommandRoutine) {
	connectorCommandRoutine = new(ConnectorCommandRoutine)
	connectorCommandRoutine.Identity = identity
	connectorCommandRoutine.ConnectorMapWorkerIterators = make(map[string][]*Iterator)

	connectorCommandRoutine.ConnectorMapCommandNameCommandMessage = NewQueue()
	connectorCommandRoutine.ConnectorMapCommandNameCommandMessage.Init()
	connectorCommandRoutine.ConnectorMapUUIDCommandMessageReply = NewQueue()
	connectorCommandRoutine.ConnectorMapUUIDCommandMessageReply.Init()

	connectorCommandRoutine.Context, _ = zmq4.NewContext()
	connectorCommandRoutine.ConnectorCommandSendToWorkerConnection = connectorCommandSendToWorkerConnection
	connectorCommandRoutine.ConnectorCommandSendToWorker, _ = connectorCommandRoutine.Context.NewSocket(zmq4.DEALER)
	connectorCommandRoutine.ConnectorCommandSendToWorker.SetIdentity(connectorCommandRoutine.Identity)
	connectorCommandRoutine.ConnectorCommandSendToWorker.Bind(connectorCommandRoutine.ConnectorCommandSendToWorkerConnection)
	fmt.Println("connectorCommandSendToWorker bind : " + connectorCommandSendToWorkerConnection)

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
	poller.Add(r.ConnectorCommandSendToWorker, zmq4.POLLIN)
	poller.Add(r.ConnectorCommandReceiveFromAggregator, zmq4.POLLIN)
	poller.Add(r.ConnectorCommandReceiveFromWorker, zmq4.POLLIN)

	command := [][]byte{}
	err := errors.New("")

	for {
		fmt.Println("Running ConnectorCommandRoutine")
		sockets, _ := poller.Poll(-1)
		for _, socket := range sockets {
			fmt.Println("Running ConnectorCommandRoutine2")

			switch currentSocket := socket.Socket; currentSocket {
			case r.ConnectorCommandSendToWorker:
				fmt.Println("Connector send worker")

				command, err = currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}
				r.processCommandSendToWorker(command)

			case r.ConnectorCommandReceiveFromAggregator:
				fmt.Println("Connector receive aggregator")

				command, err = currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}
				r.processCommandReceiveFromAggregator(command)
			case r.ConnectorCommandReceiveFromWorker:
				fmt.Println("Connector receive worker")

				command, err = currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}
				r.processCommandReceiveFromWorker(command)
			}
		}
	}
}

func (r ConnectorCommandRoutine) processCommandSendToWorker(command [][]byte) {
	target := string(command[0])
	commandType := string(command[1])
	if commandType == constant.COMMAND_WAIT {
		commandMessageWait, _ := message.DecodeCommandMessageWait(command[2])
		var iterator *Iterator
		if commandMessageWait.TypeCommand == constant.COMMAND_MESSAGE {
			iterator = NewIterator(r.ConnectorMapUUIDCommandMessage)
		} else {
			iterator = NewIterator(r.ConnectorMapCommandNameCommandMessageReply)
		}
		r.ConnectorMapWorkerIterators[target] = append(r.ConnectorMapWorkerIterators[target], iterator)

		go r.runIterator(target, commandType, commandMessageWait.Value, iterator)
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
		fmt.Println("MESSAGE")
		fmt.Println(commandMessage)
		r.ConnectorMapCommandNameCommandMessage.Push(commandMessage)
	} else {
		commandMessageReply, _ := message.DecodeCommandMessageReply(command[3])
		r.ConnectorMapUUIDCommandMessageReply.Push(commandMessageReply)
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

func (r ConnectorCommandRoutine) runIterator(target, commandType, value string, iterator *Iterator) {
	notfound := true
	for notfound {
		message := iterator.Get()
		if message != nil {
			if commandType == constant.COMMAND_MESSAGE {
				commandMessage := (*message).(message.CommandMessage)
				if value == message.Command {
					commandMessage.SendMessageWith(r.ConnectorCommandSendToWorker, target)
				}
			} else {
				commandMessage := (*message).(message.CommandMessageReply)
				if value == message.Uuid {
					commandMessage.SendMessageWith(r.ConnectorCommandSendToWorker, target)
				}
			}
		}
	}
	delete(r.ConnectorMapWorkerIterators, "target")
}

/* func (r ConnectorCommandRoutine) cleanCommandsByTimeout() {
	maxTimeout := 0
	currentTimestamp := -1
	currentTimeout := -1
	for {
		for uuid, commandMessageSlice := range r.ConnectorMapCommandNameCommandMessage {
			for _, commandMessage := range commandMessageSlice {
				currentTimestamp, _ = strconv.Atoi(commandMessage.Timestamp)
				currentTimeout, _ = strconv.Atoi(commandMessage.Timeout)

				if currentTimestamp-currentTimeout == 0 {
					delete(r.ConnectorMapCommandNameCommandMessage, uuid)
				} else {
					if currentTimeout >= maxTimeout {
						maxTimeout = currentTimeout
					}
				}
			}
		}
		time.Sleep(time.Duration(maxTimeout) * time.Millisecond)
	}
} */
