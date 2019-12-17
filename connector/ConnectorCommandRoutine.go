package connector

import (
	"fmt"
	"errors"
	"gandalfgo/message"
	"gandalfgo/constant"
	"github.com/pebbe/zmq4"
)

type ConnectorCommandRoutine struct {
	context												*zmq4.Context		
	connectorCommandSendToWorker              			*zmq4.Socket
	connectorCommandSendToWorkerConnection    			string
	connectorCommandReceiveFromAggregator           	*zmq4.Socket
	connectorCommandReceiveFromAggregatorConnections 	[]string
	connectorCommandSendToAggregator              		*zmq4.Socket
	connectorCommandSendToAggregatorConnections    		[]string
	connectorCommandReceiveFromWorker           		*zmq4.Socket
	connectorCommandReceiveFromWorkerConnection 		string
	identity                              				string
	connectorMapUUIDCommandMessage		 				map[string][]message.CommandMessage					
	connectorMapWorkerCommands 			 				map[string][]string		
}

func (r ConnectorCommandRoutine) New(identity, connectorCommandSendToWorkerConnection, connectorCommandReceiveFromWorkerConnection string, connectorCommandReceiveFromAggregatorConnections, connectorCommandSendToAggregatorConnections []string) {
	r.identity = identity

	r.context, _ = zmq4.NewContext()
	r.connectorCommandSendToWorkerConnection = connectorCommandSendToWorkerConnection
	r.connectorCommandSendToWorker, _ = r.context.NewSocket(zmq4.ROUTER)
	r.connectorCommandSendToWorker.SetIdentity(r.identity)
	r.connectorCommandSendToWorker.Bind(r.connectorCommandSendToWorkerConnection)
	fmt.Printf("connectorCommandSendToWorker bind : " + connectorCommandSendToWorkerConnection)

	r.connectorCommandReceiveFromAggregatorConnections = connectorCommandReceiveFromAggregatorConnections
	r.connectorCommandReceiveFromAggregator, _ = r.context.NewSocket(zmq4.DEALER)
	r.connectorCommandReceiveFromAggregator.SetIdentity(r.identity)
	for _, connection := range r.connectorCommandReceiveFromAggregatorConnections {
		r.connectorCommandReceiveFromAggregator.Connect(connection)
		fmt.Printf("connectorCommandReceiveFromAggregator connect : " + connection)
	}

	r.connectorCommandSendToAggregatorConnections = connectorCommandSendToAggregatorConnections
	r.connectorCommandSendToAggregator, _ = r.context.NewSocket(zmq4.ROUTER)
	r.connectorCommandSendToAggregator.SetIdentity(r.identity)
	for _, connection := range r.connectorCommandSendToAggregatorConnections {
		r.connectorCommandSendToAggregator.Connect(connection)
		fmt.Printf("connectorCommandSendToAggregator connect : " + connection)
	}


	r.connectorCommandReceiveFromWorkerConnection = connectorCommandReceiveFromWorkerConnection
	r.connectorCommandReceiveFromWorker, _ = r.context.NewSocket(zmq4.DEALER)
	r.connectorCommandReceiveFromWorker.SetIdentity(r.identity)
	r.connectorCommandReceiveFromWorker.Bind(r.connectorCommandReceiveFromWorkerConnection)
	fmt.Printf("connectorCommandReceiveFromWorker bind : " + connectorCommandReceiveFromWorkerConnection)
}

func (r ConnectorCommandRoutine) close() {	
	r.connectorCommandSendToWorker.Close()
	r.connectorCommandReceiveFromAggregator.Close()
	r.connectorCommandSendToAggregator.Close()
	r.connectorCommandReceiveFromWorker.Close()
	r.context.Term()
}

func (r ConnectorCommandRoutine) reconnectToProxy() {

}

func (r ConnectorCommandRoutine) run() {
	go r.cleanCommandsByTimeout()

	poller := zmq4.NewPoller()
	poller.Add(r.connectorCommandSendToWorker, zmq4.POLLIN)
	poller.Add(r.connectorCommandReceiveFromAggregator, zmq4.POLLIN)
	poller.Add(r.connectorCommandSendToAggregator, zmq4.POLLIN)
	poller.Add(r.connectorCommandReceiveFromWorker, zmq4.POLLIN)

	command := [][]byte{}
	err := errors.New("")

	for {
	
		sockets, _ := poller.Poll(-1)
		for _, socket := range sockets {

			switch currentSocket := socket.Socket; currentSocket {
			case r.connectorCommandSendToWorker:

				command, err = currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}

				err = r.processCommandSendToWorker(command)
				if err != nil {
					panic(err)
				}

			case r.connectorCommandReceiveFromAggregator:

				command, err = currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}
				err = r.processCommandReceiveFromAggregator(command)
				if err != nil {
					panic(err)
				}

			case r.connectorCommandSendToAggregator:

				command, err = currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}
				err = r.processCommandSendAggregator(command)
				if err != nil {
					panic(err)
				}

			case r.connectorCommandReceiveFromWorker:

				command, err = currentSocket.RecvMessageBytes(0)
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
}

func (r ConnectorCommandRoutine) processCommandSendToWorker(command [][]byte) (err error) {
	commandMessage, err := message.DecodeCommandMessage(command[1])
	r.addCommands(commandMessage)
	go commandMessage.SendWith(r.connectorCommandReceiveFromAggregator, commandMessage.SourceConnector)
	return
}

func (r ConnectorCommandRoutine) processCommandReceiveFromAggregator(command [][]byte) (err error) {
	commandMessage, err := message.DecodeCommandMessage(command[1])
	r.connectorMapUUIDCommandMessage.append(r.connectorMapUUIDCommandMessage[commandMessage.Command], commandMessage)
	return
}

func (r ConnectorCommandRoutine) processCommandSendAggregator(command [][]byte) (err error) {
	commandMessage, err := message.DecodeCommandMessage(command[1])
	go commandMessage.SendCommandWith(r.connectorCommandReceiveFromWorker)
	return
}

func (r ConnectorCommandRoutine) processCommandReceiveFromWorker(command [][]byte) (err error) {
	workerSource := string(command[0])
	commandHeader := string(command[1])
    if commandHeader == constant.COMMAND_READY {
        //commandReady := decodeCommandReady(command[2])
        commandMessage, err := r.getCommandByWorkerCommands(workerSource)
        if err != nil {
        }
		go commandMessage.SendWith(r.connectorCommandSendToWorker, workerSource)
	} else if commandHeader == constant.COMMAND_VALIDATION_FUNCTIONS {
		commandFunction, err := message.DecodeCommandFunction(command[2])
		result, _ := r.validationCommands(workerSource, commandFunction.Functions)
        if result {
			r.connectorMapWorkerCommands[workerSource] := commandFunction 
			commandFunctionReply := message.CommandFunctionReply.New(result)
			go commandFunctionReply.SendCommandFunctionReplyWith(r.connectorCommandSendToWorker)
        }
	} else {
		commandMessage, err := message.DecodeCommandMessage(command[1])
		commandMessage.SourceWorker = workerSource
		go commandMessage.SendWith(r.connectorCommandSendToAggregator, workerSource)
	}
	return
}

func (r ConnectorCommandRoutine) getCommandByWorkerCommands(worker string) (commandMessage message.CommandMessage, err error) {
	
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
		if command.Timestamp >= currentTimestamp {
			maxTimestamp = command.Timestamp
			maxCommand = command
		}
	}
	
	commandMessage = r.connectorMapUUIDCommandMessage[maxCommand][0]
	append(r.connectorMapUUIDCommandMessage[:0], connectorMapUUIDCommandMessage[0+1:]...)

	return 
}

func (r ConnectorCommandRoutine) getConnectorMapUUIDCommandMessage(command string) (commandMessage message.CommandMessage, err error) {
    if commandMessage, ok := r.connectorMapUUIDCommandMessage[command]; ok {
		if ok {
			return commandMessage
		}
	}
}

func (r ConnectorCommandRoutine) validationCommands(workerSource string, commands []string) (result bool, err error) {
	//TODO
	result = true

	return
}

func (r ConnectorCommandRoutine) addCommands(commandMessage message.CommandMessage) {
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
			} else {
				if commandMessage.timeout >= maxTimeout {
					maxTimeout = commandMessage.timeout
				}
			}
		}
		time.Sleep(maxTimeout * time.Millisecond)
	}
}