package client

import (
	"gandalf-go/client/sender"
	"gandalf-go/client/waiter"
	"gandalf-go/message"
)

type ClientGandalf struct {
	Identity                string
	SenderCommandConnection string
	SenderEventConnection   string
	WaiterCommandConnection string
	WaiterEventConnection   string
	SenderGandalf           *sender.SenderGandalf
	WaiterGandalf           *waiter.WaiterGandalf
	ClientStopChannel       chan int
}

func NewClientGandalf(identity, senderCommandConnection, senderEventConnection, waiterCommandConnection, waiterEventConnection string) (clientGandalf *ClientGandalf) {
	clientGandalf = new(ClientGandalf)
	clientGandalf.ClientStopChannel = make(chan int)

	clientGandalf.Identity = identity
	clientGandalf.SenderCommandConnection = senderCommandConnection
	clientGandalf.SenderEventConnection = senderEventConnection
	clientGandalf.WaiterCommandConnection = waiterCommandConnection
	clientGandalf.WaiterEventConnection = waiterEventConnection

	clientGandalf.SenderGandalf = sender.NewSenderGandalf(clientGandalf.Identity, clientGandalf.SenderCommandConnection, clientGandalf.SenderEventConnection)
	clientGandalf.WaiterGandalf = waiter.NewWaiterGandalf(clientGandalf.Identity, clientGandalf.WaiterCommandConnection, clientGandalf.WaiterEventConnection)

	return
}

func (cg ClientGandalf) SendCommand(context, timeout, uuid, connectorType, commandType, command, payload string) {
	cg.SenderGandalf.SendCommand(context, timeout, uuid, connectorType, commandType, command, payload)
}

func (cg ClientGandalf) SendCommandReply(commandMessage message.CommandMessage, reply, payload string) {
	cg.SenderGandalf.SendCommandReply(commandMessage, reply, payload)
}

func (cg ClientGandalf) SendEvent(topic, timeout, uuid, event, payload string) {
	cg.SenderGandalf.SendEvent(topic, timeout, uuid, event, payload)
}

func (cg ClientGandalf) WaitCommand(uuid string) (commandMessage message.CommandMessage) {
	//SEND WAIT
	return cg.WaiterGandalf.WaitCommand(uuid)
}

func (cg ClientGandalf) WaitCommandReply(uuid string) (commandMessageReply message.CommandMessageReply) {
	//SEND WAIT
	return cg.WaiterGandalf.WaitCommandReply(uuid)
}

func (cg ClientGandalf) WaitEvent(event string) (eventMessage message.EventMessage) {
	//SEND WAIT
	return cg.WaiterGandalf.WaitEvent(event)
}
