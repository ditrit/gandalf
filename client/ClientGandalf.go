package client

import (
	"gandalf-go/client/receiver"
	"gandalf-go/client/sender"
	"gandalf-go/message"
	"gandalf-go/worker/routine"
)

type ClientGandalf struct {
	Identity                  string
	SenderCommandConnection   string
	SenderEventConnection     string
	ReceiverCommandConnection string
	ReceiverEventConnection   string
	Replys                    chan message.CommandMessageReply
	CommandsRoutine           map[string][]routine.CommandRoutine
	EventsRoutine             map[string][]routine.EventRoutine
	SenderGandalf             *sender.SenderGandalf
	ReceiverGandalf           *receiver.ReceiverGandalf
}

func NewClientGandalf(identty, senderCommandConnection, senderEventConnection, receiverCommandConnection, receiverEventConnection, commandsRoutine map[string][]routine.CommandRoutine, eventsRoutine map[string][]routine.EventRoutine, replys chan message.CommandMessageReply) (clientGandalf *ClientGandalf) {
	clientGandalf = new(ClientGandalf)
	clientGandalf.Identity = identity
	clientGandalf.SenderCommandConnection = senderCommandConnection
	clientGandalf.SenderEventConnection = senderEventConnection
	clientGandalf.ReceiverCommandConnection = receiverCommandcgConnection
	clientGandalf.ReceiverEventConnection = receiverEventConnection
	clientGandalf.CommandsRoutine = commandsRoutine
	clientGandalf.EventsRoutine = eventsRoutine
	//TODO USELESS ??
	clientGandalf.Replys = replys

	clientGandalf.SenderGandalf = NewSenderGandalf(clientGandalf.identity, clientGandalf.senderCommandConnection, clientGandalf.senderEventConnection)
	clientGandalf.ReceiverGandalf = NewReceiverGandalf(clientGandalf.identity, clientGandalf.receiverCommandConnection, clientGandalf.receiverEventConnection, clientGandalf.commandsRoutine, clientGandalf.eventsRoutine, clientGandalf.results)
}

func (cg ClientGandalf) run() {
	go cg.ReceiverGandalf.run()
}
