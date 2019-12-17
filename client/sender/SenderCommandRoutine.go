package sender

import (
	"fmt"
	"gandalfgo/message"
	"gandalfgo/constant"
	"github.com/pebbe/zmq4"
)

type SenderCommandRoutine struct {
	Context							*zmq4.Context
	SenderCommandSend            	*zmq4.Socket
	SenderCommandConnections 		[]string
	SenderCommandConnection  		string
	Identity                 		string
	Replys               			chan CommandMessageReply
	MapUUIDCommandStates            map[string]State
}

func (r SenderCommandRoutine) NewSenderCommandRoutine(identity, senderCommandConnection string) (senderCommandRoutine *SenderCommandRoutine) {
	senderCommandRoutine = new(SenderCommandRoutine)

	senderCommandRoutine.Replys := make(chan CommandMessageReply)
	senderCommandRoutine.Identity = Identity

	senderCommandRoutine.Context, _ := zmq4.NewContext()
	senderCommandRoutine.SendSenderConnection = sendSenderConnection
	senderCommandRoutine.SenderCommandSend = senderCommandRoutine.Context.NewSocket(zmq4.DEALER)
	senderCommandRoutine.SenderCommandSend.SetIdentity(senderCommandRoutine.Identity)
	senderCommandRoutine.SenderCommandSend.Connect(senderCommandRoutine.SenderCommandConnection)
	fmt.Printf("senderCommandSend connect : " + senderCommandConnection)
}

func (r SenderCommandRoutine) NewLenderCommandRoutine(identity string, senderCommandConnections []string) {
	senderCommandRoutine = new(SenderCommandRoutine)

	senderCommandRoutine.Replys := make(chan CommandMessageReply)
	senderCommandRoutine.Identity = Identity

	senderCommandRoutine.Context, _ := zmq4.NewContext()
	senderCommandRoutine.SenderCommandConnections = senderCommandConnections
	senderCommandRoutine.SenderCommandSend = senderCommandRoutine.Context.NewSocket(zmq4.DEALER)
	senderCommandRoutine.SenderCommandSend.SetIdentity(senderCommandRoutine.Identity)

	for _, connection := range senderCommandRoutine.senderCommandConnections {
		senderCommandRoutine.SenderCommandSend.Connect(connection)
		fmt.Printf("senderCommandSend connect : " + connection)
	}

}

func (r SenderCommandRoutine) sendCommandSync(context, timeout, uuid, connectorType, commandtype, command, payload string) (commandResponse CommandResponse, err error) {
	commandMessage := message.NewCommandMessage(context, timeout, uuid, connectorType, commandType, command, payload)
	if err != nil {
		panic(err)
	}
	go commandMessage.sendWith(r.SenderCommandSend)

	commandResponse, err := getCommandResultSync(commandMessage.Uuid)
	if err != nil {
		panic(err)
	}
	return commandResponse
}

//TODO UTILISATION MAP
func (r SenderCommandRoutine) getCommandResultSync(uuid string) (commandResponse CommandResponse, err error) {
	for {
		command, err := r.SenderCommandSend.RecvMessageBytes(0)
        if err != nil {
			panic(err)
		}
		commandResponse := message.decodeCommandResponse(command)
		return
    }
}

func (r SenderCommandRoutine) sendCommandAsync(context, timeout, uuid, connectorType, commandtype, command, payload string) (zmq4.Message, err error) {
	commandMessage := message.NewCommandMessage(context, timeout, uuid, connectorType, commandType, command, payload)
	if err != nil {
		panic(err)
	}
	go commandMessage.sendWith(r.SenderCommandSend)

	go getCommandResultAsync(commandMessage)
}

func (r SenderCommandRoutine) getCommandResultAsync(commandMessage string) (err error) {
	select {
		case command, err := r.SenderCommandSend.RecvMessageBytes(0):
			if err != nil {
				panic(err)
			}
			r.Replys <- command
			return
		case <-time.After(commandMessage.timeout):
			fmt.Println("timeout")
	}	
}

func (r SenderCommandRoutine) cleanByTimeout() err error {

}

func (r SenderCommandRoutine) close() err error {
}
