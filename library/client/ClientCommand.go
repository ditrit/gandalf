package client

import (
	message "gandalfgo/message"
	zmq "github.com/zeromq/goczmq"
)

type ClientCommand struct {
	sendClient            zmq.Sock
	sendClientConnections *string
	sendClientConnection  string
	identity                 string
	responses                *zmq.Message
	mapUUIDCommandStates              map[string]string
}

func (cc ClientCommand) new(identity, sendClientConnection string) {
	cc.sendClientConnection = sendClientConnection
	cc.sendClient = zmq.NewDealer(sendClientConnection)
	cc.sendClient.Identity(cc.identity)
	fmt.Printf("sendClient connect : " + sendClientConnection)
}

func (cc ClientCommand) new(identity string, sendClientConnections *string) {
	cc.sendClientConnections = sendClientConnections
	cc.sendClient = zmq.NewDealer(sendClientConnections)
	cc.sendClient.Identity(cc.identity)
	fmt.Printf("sendClient connect : " + sendClientConnections)
}

func (cc ClientCommand sendCommandSync(context, timeout, uuid, commandtype, command, payload string) zmq.Message {
	 command = message.CommandMessage.new(type)
	 command 
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
go func (cc ClientCommand getCommandResultSync(commandMessage string, result chan) mangos.Message {
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
go func (cc ClientCommand getCommandResultAsync() mangos.Message {
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

func (cc ClientCommand close() {
}
