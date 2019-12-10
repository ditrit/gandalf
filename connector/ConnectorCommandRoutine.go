package connector

import (
	"fmt"
	"message"
	"container/list"	
	zmq "github.com/zeromq/goczmq"
)

type ConnectorCommandRoutine struct {
	commandZMsgSlice 				     map[string][]CommandMessage					
	commandWorkerCommands 				 map[string][]string				
	connectorCommandSendA2W              zmq.Sock
	connectorCommandSendA2WConnection    string
	connectorCommandReceiveA2W           zmq.Sock
	connectorCommandReceiveA2WConnection string
	connectorCommandSendW2A              zmq.Sock
	connectorCommandSendW2AConnection    string
	connectorCommandReceiveW2A           zmq.Sock
	connectorCommandReceiveW2AConnection string
	identity                              string
}

func (r ConnectorCommandRoutine) New(identity, connectorCommandSendA2WConnection, connectorCommandReceiveA2WConnection, connectorCommandSendW2AConnection, connectorCommandReceiveW2AConnection string) err error {
	r.identity = identity
	r.connectorCommandSendA2WConnection = connectorCommandSendA2WConnection
	r.connectorCommandSendA2W = zmq.NewDealer(r.connectorCommandSendA2WConnection)
	r.connectorCommandSendA2W.Identity(r.identity)
	fmt.Printf("connectorCommandSendA2W connect : " + connectorCommandSendA2WConnection)

	r.connectorCommandReceiveA2WConnection = connectorCommandReceiveA2WConnection
	r.connectorCommandReceiveA2W = zmq.NewRouter(connectorCommandReceiveA2WConnection)
	r.connectorCommandReceiveA2W.Identity(r.identity)
	fmt.Printf("connectorCommandReceiveA2W connect : " + connectorCommandReceiveA2WConnection)

	r.connectorCommandSendW2AConnection = connectorCommandSendW2AConnection
	r.connectorCommandSendW2A = zmq.NewDealer(connectorCommandSendW2AConnection)
	r.connectorCommandSendW2A.Identity(r.identity)
	fmt.Printf("connectorCommandSendW2A connect : " + connectorCommandSendW2AConnection)

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
	pi := zmq.PollItems{
		zmq.PollItem{Socket: connectorCommandSendA2W, Events: zmq.POLLIN},
		zmq.PollItem{Socket: connectorCommandReceiveA2W, Events: zmq.POLLIN},
		zmq.PollItem{Socket: connectorCommandSendW2A, Events: zmq.POLLIN},
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
	commandMessage.sendWith(r.connectorCommandReceiveW2A, commandMessage.sourceConnector)
}

func (r ConnectorCommandRoutine) processCommandReceiveA2W(command [][]byte) err error {
	commandMessage := CommandMessage.decodeCommand(command[1])
	r.commandZMsgSlice.append(r.commandZMsgSlice[currentCommand.command], commandMessage)

}

func (r ConnectorCommandRoutine) processCommandSendW2A(command [][]byte) err error {
	commandMessage := CommandMessage.decodeCommand(command[1])
	commandMessage.sendWith(r.connectorCommandReceiveW2A)
}

func (r ConnectorCommandRoutine) processCommandReceiveW2A(command [][]byte) err error {
    workerSource := command[0]
    if command[1] == Constant.COMMAND_READY {
        //commandReady := decodeCommandReady(command[2])
        commandMessage, err := r.getCommandByWorkerCommands(workerSource)
        if err != nil {
        }
		commandMessage.sendWith(r.connectorCommandSendA2W, workerSource)
	}
	else if command[1] == Constant.COMMAND_VALIDATION_FUNCTIONS {
		commandCommandsEvents := decodeCommandCommandsEvents(command[2])
		result := r.validationCommandsEvents(commandCommandsEvents.commands , commandCommandsEvents.events)
        if result {
			//TODO ADD WORKER
			commandCommandsEventsReply := CommandCommandsEventsReply.New(result)
			commandCommandsEventsReply.sendCommandCommandsEventsReplyWith(r.connectorCommandSendA2W)
        }
	}
    else {
		commandMessage = CommandMessage.decodeCommand(command[1])
		commandMessage.sourceWorker = workerSource
		commandMessage.sendWith(r.connectorCommandSendW2A, workerSource)
    }
}

func (r ConnectorCommandRoutine) getCommandByWorkerCommands(String worker) (commandMessage CommandMessage, err error) {
	
	var maxCommand string
	maxTimestamp := -1
	currentTimestamp := -1
	commandsWorker := r.commandWorkerCommands[worker]
	var commands []string
	
	for i, commandWorker := range commandsWorker {
		if currentCommandWorker, ok := r.commandZMsgSlice[commandWorker]; ok {
			commands[i] = currentCommandWorker
		}
	}
	
	for i, command := range commands {
		if command.timestamp >= currentTimestamp {
			maxTimestamp = command.timestamp
			maxCommand = command
		}
	}
	
	commandMessage = r.commandZMsgSlice(maxCommand)
	append(commandZMsgSlice[:0], commandZMsgSlice[0+1:]...)

	return 
}

func (r ConnectorCommandRoutine) getCommandZMsgSlice(String command) (commandMessage message.CommandMessage, err error) {
    if commandMessage, ok := r.commandZMsgSlice[command]; ok {
		return commandMessage
	}
}

func (r ConnectorCommandRoutine) validationCommandsEvents(commands, events []string) (result bool, err error) {
	//TODO
	result &= r.validationCommands(commands)
	result &= r.validationEvents(events)
	return
}

func (r ConnectorCommandRoutine) validationCommands(commands []string) (result bool, err error) {
	//TODO
	result := true
	return
}

func (r ConnectorCommandRoutine) validationEvents(events []string) (result bool, err error) {
	//TODO
	result := true
	return
}