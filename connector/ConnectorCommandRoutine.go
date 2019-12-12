package connector

import (
	"fmt"
	"message"
	"container/list"	
	zmq "github.com/zeromq/goczmq"
)

type ConnectorCommandRoutine struct {
	connectorMapUUIDCommandMessage		 map[string][]CommandMessage					
	connectorMapWorkerCommands 			 map[string][]string				
	connectorCommandSendToWorker              zmq.Sock
	connectorCommandSendToWorkerConnection    string
	connectorCommandReceiveFromAggregator           zmq.Sock
	connectorCommandReceiveFromAggregatorConnection string
	connectorCommandSendToAggregator              zmq.Sock
	connectorCommandSendToAggregatorConnection    string
	connectorCommandReceiveW2A           zmq.Sock
	connectorCommandReceiveW2AConnection string
	identity                              string
}

func (r ConnectorCommandRoutine) New(identity, connectorCommandSendToWorkerConnection, connectorCommandReceiveFromAggregatorConnection, connectorCommandSendToAggregatorConnection, connectorCommandReceiveW2AConnection string) err error {
	r.identity = identity
	r.connectorCommandSendToWorkerConnection = connectorCommandSendToWorkerConnection
	r.connectorCommandSendToWorker = zmq.NewDealer(r.connectorCommandSendToWorkerConnection)
	r.connectorCommandSendToWorker.Identity(r.identity)
	fmt.Printf("connectorCommandSendToWorker connect : " + connectorCommandSendToWorkerConnection)

	r.connectorCommandReceiveFromAggregatorConnection = connectorCommandReceiveFromAggregatorConnection
	r.connectorCommandReceiveFromAggregator = zmq.NewRouter(connectorCommandReceiveFromAggregatorConnection)
	r.connectorCommandReceiveFromAggregator.Identity(r.identity)
	fmt.Printf("connectorCommandReceiveFromAggregator connect : " + connectorCommandReceiveFromAggregatorConnection)

	r.connectorCommandSendToAggregatorConnection = connectorCommandSendToAggregatorConnection
	r.connectorCommandSendToAggregator = zmq.NewDealer(connectorCommandSendToAggregatorConnection)
	r.connectorCommandSendToAggregator.Identity(r.identity)
	fmt.Printf("connectorCommandSendToAggregator connect : " + connectorCommandSendToAggregatorConnection)

	r.connectorCommandReceiveW2AConnection = connectorCommandReceiveW2AConnection
	r.connectorCommandReceiveW2A = zmq.NewRouter(connectorCommandReceiveW2AConnection)
	r.connectorCommandReceiveW2A.Identity(r.identity)
	fmt.Printf("connectorCommandReceiveW2A connect : " + connectorCommandReceiveW2AConnection)
}

func (r ConnectorCommandRoutine) close() err error {
}

func (r ConnectorCommandRoutine) reconnectToProxy() err error {

}

func (r ConnectorCommandRoutine) run() err error {
	go cleanCommandsByTimeout()

	pi := zmq.PollItems{
		zmq.PollItem{Socket: connectorCommandSendToWorker, Events: zmq.POLLIN},
		zmq.PollItem{Socket: connectorCommandReceiveFromAggregator, Events: zmq.POLLIN},
		zmq.PollItem{Socket: connectorCommandSendToAggregator, Events: zmq.POLLIN},
		zmq.PollItem{Socket: connectorCommandReceiveW2A, Events: zmq.POLLIN},

		var command = [][]byte{}

	for {

		_, _ = zmq.Poll(pi, -1)

		switch {
		case pi[0].REvents&zmq.POLLIN != 0:

			command, err := pi[0].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}

			err = r.processCommandSendA2W(command)
			if err != nil {
				panic(err)
			}

		case pi[1].REvents&zmq.POLLIN != 0:

			command, err := pi[1].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			err = r.processCommandReceiveA2W(command)
			if err != nil {
				panic(err)
			}

		case pi[2].REvents&zmq.POLLIN != 0:

			command, err := pi[2].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			err = r.processCommandSendW2A(command)
			if err != nil {
				panic(err)
			}

		case pi[3].REvents&zmq.POLLIN != 0:

			command, err := pi[3].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			err = r.processCommandReceiveW2A(command)
			if err != nil {
				panic(err)
			}
		}
	}
}

func (r ConnectorCommandRoutine) processCommandSendA2W(command [][]byte) err error {
	commandMessage := CommandMessage.decodeCommand(command[1])
	r.addCommands(commandMessage)
	go commandMessage.sendWith(r.connectorCommandReceiveW2A, commandMessage.sourceConnector)
}

func (r ConnectorCommandRoutine) processCommandReceiveA2W(command [][]byte) err error {
	commandMessage := CommandMessage.decodeCommand(command[1])
	r.connectorMapUUIDCommandMessage.append(r.connectorMapUUIDCommandMessage[currentCommand.command], commandMessage)
}

func (r ConnectorCommandRoutine) processCommandSendW2A(command [][]byte) err error {
	commandMessage := CommandMessage.decodeCommand(command[1])
	go commandMessage.sendWith(r.connectorCommandReceiveW2A)
}

func (r ConnectorCommandRoutine) processCommandReceiveW2A(command [][]byte) err error {
    workerSource := command[0]
    if command[1] == Constant.COMMAND_READY {
        //commandReady := decodeCommandReady(command[2])
        commandMessage, err := r.getCommandByWorkerCommands(workerSource)
        if err != nil {
        }
		go commandMessage.sendWith(r.connectorCommandSendToWorker, workerSource)
	}
	else if command[1] == Constant.COMMAND_VALIDATION_FUNCTIONS {
		commandFunction := decodeCommandFunction(command[2])
		result := r.validationCommands(workerSource, commandFunction.functions)
        if result {
			r.connectorMapWorkerCommands[workerSource] = commands 
			commandFunctionReply := CommandFunctionReply.New(result)
			go commandFunctionReply.sendCommandFunctionReplyWith(r.connectorCommandSendToWorker)

        }
	}
    else {
		commandMessage = CommandMessage.decodeCommand(command[1])
		commandMessage.sourceWorker = workerSource
		go commandMessage.sendWith(r.connectorCommandSendToAggregator, workerSource)
    }
}

func (r ConnectorCommandRoutine) getCommandByWorkerCommands(String worker) (commandMessage CommandMessage, err error) {
	
	var maxCommand string
	maxTimestamp := -1
	currentTimestamp := -1
	commandsWorker := r.commandWorkerCommands[worker]
	var commands []string
	
	for i, commandWorker := range commandsWorker {
		if currentCommandWorker, ok := r.connectorMapUUIDCommandMessage[commandWorker]; ok {
			commands[i] = currentCommandWorker
		}
	}
	
	for i, command := range commands {
		if command.timestamp >= currentTimestamp {
			maxTimestamp = command.timestamp
			maxCommand = command
		}
	}
	
	commandMessage = r.connectorMapUUIDCommandMessage[maxCommand]
	append(connectorMapUUIDCommandMessage[:0], connectorMapUUIDCommandMessage[0+1:]...)

	return 
}

func (r ConnectorCommandRoutine) getConnectorMapUUIDCommandMessage(String command) (commandMessage message.CommandMessage, err error) {
    if commandMessage, ok := r.connectorMapUUIDCommandMessage[command]; ok {
		if ok {
			return commandMessage
		}
	}
}

func (r ConnectorCommandRoutine) validationCommands(workerSource string, commands []string) (result bool, err error) {
	//TODO
	result := true

	return
}

func (r ConnectorCommandRoutine) addCommands(commandMessage CommandMessage) {
	if val, ok := r.connectorMapUUIDCommandMessage[commandMessage.uuid]; ok {
		if !ok {
			r.connectorMapUUIDCommandMessage[commandMessage.uuid] = commandMessage
		}
	}
}

func (r ConnectorCommandRoutine) cleanCommandsByTimeout() {
	maxTimeout = 0
	for {
		for uuid, commandMessage := range r.connectorMapUUIDCommandMessage { 
			if commandMessage.timestamp - commandMessage.timeout == 0 {
				delete(r.connectorMapUUIDCommandMessage, uuid) 	
			}
			else {
				if commandMessage.timeout >= maxTimeout {
					maxTimeout = commandMessage.timeout
				}
			}
		}
		time.Sleep(maxTimeout * time.Millisecond)
	}
}