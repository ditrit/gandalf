package client

import (
	"fmt"
	"message"
	"constant"
	zmq "github.com/zeromq/goczmq"
)

type ClientCommandRoutine struct {
	clientCommandSend            	zmq.Sock
	clientCommandSendConnections 	*string
	clientCommandSendConnection  	string
	identity                 		string
	result               			chan
	mapUUIDCommandStates            map[string]State
}

func (r ClientCommandRoutine) New(identity, sendClientConnection string) err error {
	result := make(chan Result)
	r.identity = identity
	r.sendClientConnection = sendClientConnection
	r.clientCommandSend = zmq.NewDealer(sendClientConnection)
	r.clientCommandSend.Identity(r.identity)
	fmt.Printf("clientCommandSend connect : " + sendClientConnection)
}

func (r ClientCommandRoutine) NewList(identity string, clientCommandSendConnections *string) err error {
	result := make(chan Result)
	r.identity = identity
	r.clientCommandSendConnections = clientCommandSendConnections
	r.clientCommandSend = zmq.NewDealer(clientCommandSendConnections)
	r.clientCommandSend.Identity(r.identity)
	fmt.Printf("clientCommandSend connect : " + clientCommandSendConnections)
}

func (r ClientCommandRoutine) sendCommandSync(context, timeout, uuid, connectorType, commandtype, command, payload string) (commandResponse CommandResponse, err error) {
	commandMessage := CommandMessage.New(context, timeout, uuid, connectorType, commandType, command, payload)
	if err != nil {
		panic(err)
	}
	commandMessage.sendWith(clientCommandSend)
	commandResponse, err := getCommandResultSync(commandMessage.uuid)
	if err != nil {
		panic(err)
	}
	return commandResponse
}

//TODO UTILISATION MAP
func (r ClientCommandRoutine) getCommandResultSync(uuid string) (commandResponse CommandResponse, err error) {
	for {
		command, err := r.clientCommandSend.RecvMessage()
        if err != nil {
			panic(err)
		}
		commandResponse := CommandResponse.decode(command)
		return
    }
}

func (r ClientCommandRoutine) sendCommandAsync(context, timeout, uuid, connectorType, commandtype, command, payload string) (zmq.Message, err error) {
	commandMessage := CommandMessage.New(context, timeout, uuid, connectorType, commandType, command, payload)
	if err != nil {
		panic(err)
	}
	commandMessage.sendWith(clientCommandSend)
	go getCommandResultAsync(commandMessage)
}

func (r ClientCommandRoutine) getCommandResultAsync(commandMessage string) (err error) {
	select {
		case command, err := r.clientCommandSend.RecvMessage():
			if err != nil {
				panic(err)
			}
			r.result <- command
			return
		case <-time.After(commandMessage.timeout):
			fmt.Println("timeout")
	}	
}

func (r ClientCommandRoutine) cleanByTimeout() err error {
}

func (r ClientCommandRoutine) close() err error {
}
