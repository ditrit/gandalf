package sender

import (
	"fmt"
	"gandalf-go/message"

	"github.com/pebbe/zmq4"
)

type SenderCommandRoutine struct {
	Context                  *zmq4.Context
	SenderCommandSend        *zmq4.Socket
	SenderCommandConnections []string
	SenderCommandConnection  string
	Identity                 string
	//Replys                   chan message.CommandMessageReply
	MapUUIDCommandStates     map[string]string
}

func NewSenderCommandRoutine(identity, senderCommandConnection string) (senderCommandRoutine *SenderCommandRoutine) {
	senderCommandRoutine = new(SenderCommandRoutine)

	senderCommandRoutine.MapUUIDCommandStates = make(map[string]string)
	//senderCommandRoutine.Replys = make(chan message.CommandMessageReply)
	senderCommandRoutine.Identity = identity

	senderCommandRoutine.Context, _ = zmq4.NewContext()
	senderCommandRoutine.SenderCommandConnection = senderCommandConnection
	senderCommandRoutine.SenderCommandSend, _ = senderCommandRoutine.Context.NewSocket(zmq4.DEALER)
	senderCommandRoutine.SenderCommandSend.SetIdentity(senderCommandRoutine.Identity)
	senderCommandRoutine.SenderCommandSend.Connect(senderCommandRoutine.SenderCommandConnection)
	fmt.Println("senderCommandSend connect : " + senderCommandConnection)

	return
}

func NewLenderCommandRoutine(identity string, senderCommandConnections []string) (senderCommandRoutine *SenderCommandRoutine) {
	senderCommandRoutine = new(SenderCommandRoutine)

	//senderCommandRoutine.Replys = make(chan message.CommandMessageReply)
	senderCommandRoutine.Identity = identity

	senderCommandRoutine.Context, _ = zmq4.NewContext()
	senderCommandRoutine.SenderCommandConnections = senderCommandConnections
	senderCommandRoutine.SenderCommandSend, _ = senderCommandRoutine.Context.NewSocket(zmq4.DEALER)
	senderCommandRoutine.SenderCommandSend.SetIdentity(senderCommandRoutine.Identity)

	for _, connection := range senderCommandRoutine.SenderCommandConnections {
		senderCommandRoutine.SenderCommandSend.Connect(connection)
		fmt.Println("senderCommandSend connect : " + connection)
	}

	return
}

/* func (r SenderCommandRoutine) SendCommandSync(context, timeout, uuid, connectorType, commandType, command, payload string) (commandMessageReply message.CommandMessageReply) {
	commandMessage := message.NewCommandMessage(context, timeout, uuid, connectorType, commandType, command, payload)

	go commandMessage.SendCommandWith(r.SenderCommandSend)
	commandMessageReply = r.getCommandResultSync(commandMessage.Uuid)

	return
} */

func (r SenderCommandRoutine) SendCommand(context, timeout, uuid, connectorType, commandType, command, payload string)  {
	commandMessage := message.NewCommandMessage(context, timeout, uuid, connectorType, commandType, command, payload)

	go commandMessage.SendMessageWith(r.SenderCommandSend)
}

func (r SenderCommandRoutine) SendCommandReply(commandMessage message.CommandMessage, reply, payload string) {
	commandMessageReply := new(message.CommandMessageReply)
	commandMessageReply.From(commandMessage, reply, payload)

	go commandMessageReply.SendMessageWith(r.SenderCommandSend)
	return
}

/* //TEST
func (r SenderCommandRoutine) SendCommandSyncTEST(context, timeout, uuid, connectorType, commandType, command, payload string) (commandMessageReply message.CommandMessageReply) {
	commandMessage := message.NewCommandMessage(context, timeout, uuid, connectorType, commandType, command, payload)
	commandMessage.DestinationAggregator = "aggregator2"
	commandMessage.DestinationConnector = "connector2"
	commandMessage.DestinationWorker = "worker2"
	go commandMessage.SendCommandWith(r.SenderCommandSend)
	commandMessageReply = r.getCommandResultSync(commandMessage.Uuid)

	return
}

//FIN TEST

//TODO UTILISATION MAP
func (r SenderCommandRoutine) getCommandResultSync(uuid string) (commandMessageReply message.CommandMessageReply) {
	for {
		command, err := r.SenderCommandSend.RecvMessageBytes(0)
		if err != nil {
			panic(err)
		}
		commandMessageReply, _ = message.DecodeCommandMessageReply(command[1])
		return
	}
}

func (r SenderCommandRoutine) sendCommandAsync(context, timeout, uuid, connectorType, commandType, command, payload string) {
	commandMessage := message.NewCommandMessage(context, timeout, uuid, connectorType, commandType, command, payload)

	go commandMessage.SendCommandWith(r.SenderCommandSend)
	go r.getCommandResultAsync()
}

func (r SenderCommandRoutine) getCommandResultAsync() {
	for {
		command, err := r.SenderCommandSend.RecvMessageBytes(0)
		if err != nil {
			panic(err)
		}
		commandMessage, _ := message.DecodeCommandMessageReply(command[1])
		r.Replys <- commandMessage

		return
	}
} */

func (r SenderCommandRoutine) cleanByTimeout() {

}

func (r SenderCommandRoutine) close() {
}
