package connector

import (
	"fmt"
	"gandalfgo/message"
	"container/list"	
	"github.com/pebbe/zmq4"
)

type ConnectorCommandRoutine struct {
	context												zmq4.Context
	connectorMapUUIDCommandMessage		 				map[string][]CommandMessage					
	connectorMapWorkerCommands 			 				map[string][]string				
	connectorCommandSendToWorker              			zmq4.Socket
	connectorCommandSendToWorkerConnection    			string
	connectorCommandReceiveFromAggregator           	zmq4.Socket
	connectorCommandReceiveFromAggregatorConnections 	[]string
	connectorCommandSendToAggregator              		zmq4.Socket
	connectorCommandSendToAggregatorConnections    		[]string
	connectorCommandReceiveFromWorker           		zmq4.Socket
	connectorCommandReceiveFromWorkerConnection 		string
	identity                              				string
}

func (r ConnectorCommandRoutine) New(identity, connectorCommandSendToWorkerConnection, connectorCommandReceiveFromAggregatorConnections, connectorCommandSendToAggregatorConnections, connectorCommandReceiveFromWorkerConnection string) err error {
	r.identity = identity

	r.context, _ := zmq4.NewContext()
	r.connectorCommandSendToWorkerConnection = connectorCommandSendToWorkerConnection
	r.connectorCommandSendToWorker = r.context.NewRouter(r.connectorCommandSendToWorkerConnection)
	r.connectorCommandSendToWorker.Identity(r.identity)
	fmt.Printf("connectorCommandSendToWorker connect : " + connectorCommandSendToWorkerConnection)

	r.connectorCommandReceiveFromAggregatorConnections = connectorCommandReceiveFromAggregatorConnections
	r.connectorCommandReceiveFromAggregator = r.context.NewDealer(connectorCommandReceiveFromAggregatorConnections)
	r.connectorCommandReceiveFromAggregator.Identity(r.identity)
	fmt.Printf("connectorCommandReceiveFromAggregator connect : " + connectorCommandReceiveFromAggregatorConnections)

	r.connectorCommandSendToAggregatorConnections = connectorCommandSendToAggregatorConnections
	r.connectorCommandSendToAggregator = r.context.NewRouter(connectorCommandSendToAggregatorConnections)
	r.connectorCommandSendToAggregator.Identity(r.identity)
	fmt.Printf("connectorCommandSendToAggregator connect : " + connectorCommandSendToAggregatorConnections)

	r.connectorCommandReceiveFromWorkerConnection = connectorCommandReceiveFromWorkerConnection
	r.connectorCommandReceiveFromWorker = r.context.NewDealer(connectorCommandReceiveFromWorkerConnection)
	r.connectorCommandReceiveFromWorker.Identity(r.identity)
	fmt.Printf("connectorCommandReceiveFromWorker connect : " + connectorCommandReceiveFromWorkerConnection)
}

func (r ConnectorCommandRoutine) close() {
}

func (r ConnectorCommandRoutine) reconnectToProxy() {

}

func (r ConnectorCommandRoutine) run() {
	go cleanCommandsByTimeout()

	pi := zmq4.PollItems{
		zmq4.PollItem{Socket: connectorCommandSendToWorker, Events: zmq4.POLLIN},
		zmq4.PollItem{Socket: connectorCommandReceiveFromAggregator, Events: zmq4.POLLIN},
		zmq4.PollItem{Socket: connectorCommandSendToAggregator, Events: zmq4.POLLIN},
		zmq4.PollItem{Socket: connectorCommandReceiveFromWorker, Events: zmq4.POLLIN},

		var command = [][]byte{}

	for {

		_, _ = zmq4.Poll(pi, -1)

		switch {
		case pi[0].REvents&zmq4.POLLIN != 0:

			command, err := pi[0].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}

			err = r.processCommandSendToWorker(command)
			if err != nil {
				panic(err)
			}

		case pi[1].REvents&zmq4.POLLIN != 0:

			command, err := pi[1].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			err = r.processCommandReceiveFromAggregator(command)
			if err != nil {
				panic(err)
			}

		case pi[2].REvents&zmq4.POLLIN != 0:

			command, err := pi[2].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			err = r.processCommandSendAggregator(command)
			if err != nil {
				panic(err)
			}

		case pi[3].REvents&zmq4.POLLIN != 0:

			command, err := pi[3].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			err = r.processCommandReceiveFromWorker(command)
			if err != nil {
				panic(err)
			}
		}
	}
}

func (r ConnectorCommandRoutine) processCommandSendToWorker(command [][]byte) {
	commandMessage := CommandMessage.decodeCommand(command[1])
	r.addCommands(commandMessage)
	go commandMessage.sendWith(r.connectorCommandReceiveFromCluster, commandMessage.sourceConnector)
}

func (r ConnectorCommandRoutine) processCommandReceiveFromAggregator(command [][]byte) {
	commandMessage := CommandMessage.decodeCommand(command[1])
	r.connectorMapUUIDCommandMessage.append(r.connectorMapUUIDCommandMessage[currentCommand.command], commandMessage)
}

func (r ConnectorCommandRoutine) processCommandSendAggregator(command [][]byte) {
	commandMessage := CommandMessage.decodeCommand(command[1])
	go commandMessage.sendWith(r.connectorCommandReceiveFromWorker)
}

func (r ConnectorCommandRoutine) processCommandReceiveFromWorker(command [][]byte) {
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