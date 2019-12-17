package receiver

import(
	"gandalfgo/worker/routine"
	"gandalfgo/message"
)  

type ReceiverGandalf struct {
	Identity 					string   
	ReceiverCommandConnection 	string
	ReceiverEventConnection 	string
	ReceiverCommandRoutine 		*ReceiverCommandRoutine
	ReceiverEventRoutine   		*ReceiverEventRoutine
	Replys 						chan message.CommandMessageReply
	CommandsRoutine 			map[string][]routine.CommandRoutine
	EventsRoutine 				map[string][]routine.EventRoutine
}

func NewReceiverGandalf(identity, receiverCommandConnection, receiverEventConnection string, commandsRoutine map[string][]routine.CommandRoutine, eventsRoutine map[string][]routine.EventRoutine, replys chan message.CommandMessageReply) (receiverGandalf ReceiverGandalf) {
	receiverGandalf = new(ReceiverGandalf)

	receiverGandalf.Identity = identity
	receiverGandalf.ReceiverCommandConection = receiverCommandConnection
	receiverGandalf.ReceiverEventRoutine = receiverEventRoutine
	receiverGandalf.CommandsRoutine = commandsRoutine
	receiverGandalf.EventsRoutine = eventsRoutine
	receiverGandalf.Replys = replys
	receiverGandalf.ReceiverCommandRoutine = ReceiverCommandRoutine.New(receiverGandalf.Identity, receiverGandalf.ReceiverCommandConnection, receiverGandalf.CommandsRoutine , receiverGandalf.Results)
	receiverGandalf.ReceiverEventRoutine = ReceiverEventRoutine.New(receiverGandalf.Identity, receiverGandalf.ReceiverEventConnection, receiverGandalf.EventsRoutine)
}
