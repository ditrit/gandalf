package sender

import (
	"fmt"
	"gandalfgo/message"
	"gandalfgo/constant"
	"github.com/pebbe/zmq4"
)

type SenderCommandRoutine struct {
	context							zmq4.Context
	senderCommandSend            	zmq4.Socket
	senderCommandConnections 		[]string
	senderCommandConnection  		string
	identity                 		string
	result               			chan
	mapUUIDCommandStates            map[string]State
}

func (r SenderCommandRoutine) New(identity, senderCommandConnection string) {
	result := make(chan Result)
	r.identity = identity

	r.context, _ := zmq4.NewContext()
	r.sendSenderConnection = sendSenderConnection
	r.senderCommandSend = r.context.NewSocket(zmq4.DEALER)
	r.senderCommandSend.SetIdentity(r.identity)
	r.senderCommandSend.Connect(r.senderCommandConnection)
	fmt.Printf("senderCommandSend connect : " + senderCommandConnection)
}

func (r SenderCommandRoutine) NewList(identity string, senderCommandConnections *string) {
	result := make(chan Result)
	r.identity = identity

	r.context, _ := zmq4.NewContext()
	r.senderCommandConnections = senderCommandConnections
	r.senderCommandSend = r.context.NewSocket(zmq4.DEALER)
	r.senderCommandSend.SetIdentity(r.identity)
	for _, connection := range r.aggregatorCommandReceiveFromClusterConnections {
		r.senderCommandSend.Connect(r.connection)
		fmt.Printf("senderCommandSend connect : " + connection)
	}

}

func (r SenderCommandRoutine) sendCommandSync(context, timeout, uuid, connectorType, commandtype, command, payload string) (commandResponse CommandResponse, err error) {
	commandMessage := CommandMessage.New(context, timeout, uuid, connectorType, commandType, command, payload)
	if err != nil {
		panic(err)
	}
	go commandMessage.sendWith(senderCommandSend)

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

func (r SenderCommandRoutine) sendCommandAsync(context, timeout, uuid, connectorType, commandtype, command, payload string) (zmq4.Message, err error) {
	commandMessage := CommandMessage.New(context, timeout, uuid, connectorType, commandType, command, payload)
	if err != nil {
		panic(err)
	}
	go commandMessage.sendWith(senderCommandSend)

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
