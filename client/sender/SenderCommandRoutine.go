package sender

import (
	"fmt"
	"message"
	"constant"
	zmq "github.com/zeromq/goczmq"
)

type SenderCommandRoutine struct {
	senderCommandSend            	zmq.Sock
	senderCommandConnections 	*string
	senderCommandConnection  	string
	identity                 		string
	result               			chan
	mapUUIDCommandStates            map[string]State
}

func (r SenderCommandRoutine) New(identity, sendSenderConnection string) err error {
	result := make(chan Result)
	r.identity = identity
	r.sendSenderConnection = sendSenderConnection
	r.senderCommandSend = zmq.NewDealer(sendSenderConnection)
	r.senderCommandSend.Identity(r.identity)
	fmt.Printf("senderCommandSend connect : " + sendSenderConnection)
}

func (r SenderCommandRoutine) NewList(identity string, senderCommandConnections *string) err error {
	result := make(chan Result)
	r.identity = identity
	r.senderCommandConnections = senderCommandConnections
	r.senderCommandSend = zmq.NewDealer(senderCommandConnections)
	r.senderCommandSend.Identity(r.identity)
	fmt.Printf("senderCommandSend connect : " + senderCommandConnections)
}

func (r SenderCommandRoutine) sendCommandSync(context, timeout, uuid, connectorType, commandtype, command, payload string) (commandResponse CommandResponse, err error) {
	commandMessage := CommandMessage.New(context, timeout, uuid, connectorType, commandType, command, payload)
	if err != nil {
		panic(err)
	}
	for {
		isSend = commandMessage.sendWith(senderCommandSend)
		if isSend {
			break
		}
		time.Sleep(2 * time.Second)
	}

	commandResponse, err := getCommandResultSync(commandMessage.uuid)
	if err != nil {
		panic(err)
	}
	return commandResponse
}

//TODO UTILISATION MAP
func (r SenderCommandRoutine) getCommandResultSync(uuid string) (commandResponse CommandResponse, err error) {
	for {
		command, err := r.senderCommandSend.RecvMessage()
        if err != nil {
			panic(err)
		}
		commandResponse := CommandResponse.decode(command)
		return
    }
}

func (r SenderCommandRoutine) sendCommandAsync(context, timeout, uuid, connectorType, commandtype, command, payload string) (zmq.Message, err error) {
	commandMessage := CommandMessage.New(context, timeout, uuid, connectorType, commandType, command, payload)
	if err != nil {
		panic(err)
	}
	for {
		isSend = commandMessage.sendWith(senderCommandSend)
		if isSend {
			break
		}
		time.Sleep(2 * time.Second)
	}
	go getCommandResultAsync(commandMessage)
}

func (r SenderCommandRoutine) getCommandResultAsync(commandMessage string) (err error) {
	select {
		case command, err := r.senderCommandSend.RecvMessage():
			if err != nil {
				panic(err)
			}
			r.result <- command
			return
		case <-time.After(commandMessage.timeout):
			fmt.Println("timeout")
	}	
}

func (r SenderCommandRoutine) cleanByTimeout() err error {

}

func (r SenderCommandRoutine) close() err error {
}
