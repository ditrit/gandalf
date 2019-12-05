package connector

import (
	"fmt"

	zmq "github.com/zeromq/goczmq"
)

type ConnectorCommandRoutine struct {
    commandMessage                        CommandMessage
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
	command = r.updateHeaderCommandSendA2W(command)
	r.connectorCommandReceiveW2A.SendMessage(command)
}

func (r ConnectorCommandRoutine) updateHeaderCommandSendA2W(command [][]byte) err error {
    currentCommand, err := r.commandMessage.decodeCommand(command[1])
    if err != nil {
        //RESPONSE WORKER
    }
    command[0] = currentCommand.sourceConnector
}

func (r ConnectorCommandRoutine) processCommandReceiveA2W(command [][]byte) err error {
	command = r.updateHeaderCommandReceiveA2W(command)
	r.connectorCommandSendA2W.SendMessage(command)
}

func (r ConnectorCommandRoutine) updateHeaderCommandReceiveA2W(command [][]byte) err error {
    currentCommand, err := r.commandMessage.decodeCommand(command[1])
    if err != nil {
        //STOCK COMMAND
    }
    command[0] = []byte(currentCommand.targetConnector)
}

func (r ConnectorCommandRoutine) processCommandSendW2A(command [][]byte) err error {
	command = r.updateHeaderCommandSendW2A(command)
	r.connectorCommandReceiveW2A.SendMessage(command)
	//SEND
}

func (r ConnectorCommandRoutine) updateHeaderCommandSendW2A(command [][]byte) err error {
    //TODO NOTHING
}

func (r ConnectorCommandRoutine) processCommandReceiveW2A(command [][]byte) err error {
    //READY && RESULT
    if command[1] == Constant.COMMAND_READY {
        commands := command[2]
        workerCommand, err := r.getCommandByWorkerCommands()
        if err != nil {
        }
        workerCommand = r.updateIdentityCommandReadyMessage(workerCommand)
        workerCommand = r.updateHeaderCommandReceiveReadyMessage(workerCommand)
    	r.connectorCommandSendW2A.SendMessage(command)

    }
    else {
        command = r.updateHeaderCommandReceiveW2A(command)
    	r.connectorCommandSendW2A.SendMessage(command)
    }
}

func (r ConnectorCommandRoutine) updateIdentityCommandReadyMessage(command [][]byte) (command [][]byte, err error) {
    //TODO
}

func (r ConnectorCommandRoutine) updateHeaderCommandReceiveReadyMessage(command [][]byte) (command [][]byte, err error) {
    //TODO
}

func (r ConnectorCommandRoutine) updateHeaderCommandReceiveW2A(command [][]byte) (command [][]byte, err error) {
       command = append(command[:i][], s[i+1][]...)
       return command
}

func (r ConnectorCommandRoutine) getCommandByWorkerCommands() (command [][]byte, err error) {
    //TODO
}