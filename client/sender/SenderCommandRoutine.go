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
	MapUUIDCommandStates map[string]string
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

func NewSenderCommandRoutineList(identity string, senderCommandConnections []string) (senderCommandRoutine *SenderCommandRoutine) {
	senderCommandRoutine = new(SenderCommandRoutine)
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

func (r SenderCommandRoutine) SendCommand(context, timeout, uuid, connectorType, commandType, command, payload string) {
	commandMessage := message.NewCommandMessage(context, timeout, uuid, connectorType, commandType, command, payload)
	commandMessage.DestinationAggregator = "aggregator1"
	commandMessage.DestinationConnector = "connector1"
	commandMessage.SendMessageWith(r.SenderCommandSend)
}

func (r SenderCommandRoutine) SendCommandReply(commandMessage message.CommandMessage, reply, payload string) {
	commandMessageReply := new(message.CommandMessageReply)
	commandMessageReply.From(commandMessage, reply, payload)
	commandMessageReply.SendMessageWith(r.SenderCommandSend)
}

func (r SenderCommandRoutine) cleanByTimeout() {

}

func (r SenderCommandRoutine) close() {
}
