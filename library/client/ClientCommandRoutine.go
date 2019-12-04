package client

import (
	"fmt"

	zmq "github.com/zeromq/goczmq"
)

type ClientCommandRoutine struct {
	clientCommandSend            zmq.Sock
	clientCommandSendConnections *string
	clientCommandSendConnection  string
	identity                 string
	responses                *zmq.Message
	mapUUIDCommandStates              map[string]string
}

func (r ClientCommandRoutine) new(identity, sendClientConnection string) {
	cc.identity = identity
	cc.sendClientConnection = sendClientConnection
	cc.clientCommandSend = zmq.NewDealer(sendClientConnection)
	cc.clientCommandSend.Identity(cc.identity)
	fmt.Printf("clientCommandSend connect : " + sendClientConnection)
}

func (r ClientCommandRoutine) new(identity string, clientCommandSendConnections *string) {
	cc.identity = identity
	cc.clientCommandSendConnections = clientCommandSendConnections
	cc.clientCommandSend = zmq.NewDealer(clientCommandSendConnections)
	cc.clientCommandSend.Identity(cc.identity)
	fmt.Printf("clientCommandSend connect : " + clientCommandSendConnections)
}

func (r ClientCommandRoutine) sendCommandSync(context, timeout, uuid, commandtype, command, payload string) zmq.Message {
	 //command = message.CommandMessage.new(type)
	 commandMessage, err := msgpack.Marshal(&command)
	if err != nil {
		panic(err)
	}
	result := make(chan Result)
	cc.sendClient.SendMessage(commandMessage, result)
	go getCommandResultSync(commandMessage, result)

	return result
}

//TODO UTILISATION MAP
func (r ClientCommandRoutine) getCommandResultSync(commandMessage string, result chan) chan {
	cc.sendClient.SendMessage(commandMessage)
	select {
		case event, err := cc.sendClient.RecvMessage():
			if err != nil {
				panic(err)
			}
			result <- event
		case <-time.After(3 * time.Second):
			fmt.Println("timeout 2")
	}	
}

//TODO UTILISATION MAP //REVOIR
func (r ClientCommandRoutine) getCommandResultAsync() mangos.Message {
	cc.sendClient.SendMessage(commandMessage)
	select {
		case event, err := cc.sendClient.RecvMessage(): //APPEL ROUTINE
			if err != nil {
				panic(err)
			}
			result <- event
		case <-time.After(3 * time.Second):
			fmt.Println("timeout 2")
    }	
}

func (r ClientCommandRoutine) close() {
}
