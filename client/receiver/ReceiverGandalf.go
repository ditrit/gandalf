package receiver

import (
	"gandalf-go/message"
	"gandalf-go/worker/routine"
)

type ReceiverGandalf struct {
	Identity                  string
	ReceiverCommandConnection string
	ReceiverEventConnection   string
	ReceiverCommandRoutine    *ReceiverCommandRoutine
	ReceiverEventRoutine      *ReceiverEventRoutine
	Replys                    chan message.CommandMessageReply
	CommandsRoutine           map[string][]routine.CommandRoutine
	EventsRoutine             map[string][]routine.EventRoutine
}

func NewReceiverGandalf(identity, receiverCommandConnection, receiverEventConnection string, commandsRoutine map[string][]routine.CommandRoutine, eventsRoutine map[string][]routine.EventRoutine, replys chan message.CommandMessageReply) (receiverGandalf *ReceiverGandalf) {
	receiverGandalf = new(ReceiverGandalf)

	receiverGandalf.Identity = identity
	receiverGandalf.ReceiverCommandConnection = receiverCommandConnection
	receiverGandalf.ReceiverEventConnection = receiverEventConnection
	receiverGandalf.CommandsRoutine = commandsRoutine
	receiverGandalf.EventsRoutine = eventsRoutine
	receiverGandalf.Replys = replys

	receiverGandalf.ReceiverCommandRoutine = NewReceiverCommandRoutine(receiverGandalf.Identity, receiverGandalf.ReceiverCommandConnection, receiverGandalf.CommandsRoutine, receiverGandalf.Replys)
	receiverGandalf.ReceiverEventRoutine = NewReceiverEventRoutine(receiverGandalf.Identity, receiverGandalf.ReceiverEventConnection, receiverGandalf.EventsRoutine)

	return
}

func (rg ReceiverGandalf) Run() {
	go rg.ReceiverCommandRoutine.run()
	go rg.ReceiverEventRoutine.run()
}
