package receiver

  
type ReceiverGandalf struct {
	Identity 					string   
	ReceiverCommandConnection 	string
	ReceiverEventConnection 	string
	ReceiverCommandRoutine 		*ReceiverCommandoutine
	ReceiverEventRoutine   		*ReceiverEventRoutine
	Results 					chan ResponseMessage
	CommandsRoutine 			map[string][]CommandFunction
	EventsRoutine 				map[string][]EventFunction
}

func NewReceiverGandalf(identity, receiverCommandConnection, receiverEventConnection string, commandsRoutine map[string][]CommandFunction, eventsRoutine map[string][]EventFunction, results chan ResponseMessage) (receiverGandalf ReceiverGandalf) {
	receiverGandalf = new(ReceiverGandalf)

	receiverGandalf.Identity = identity
	receiverGandalf.ReceiverCommandConection = receiverCommandConnection
	receiverGandalf.ReceiverEventRoutine = receiverEventRoutine
	receiverGandalf.CommandsRoutine = commandsRoutine
	receiverGandalf.EventsRoutine = eventsRoutine
	receiverGandalf.Results = results
	receiverGandalf.ReceiverCommandRoutine = ReceiverCommandRoutine.New(receiverGandalf.Identity, receiverGandalf.ReceiverCommandConnection, receiverGandalf.CommandsRoutine , receiverGandalf.Results)
	receiverGandalf.ReceiverEventRoutine = ReceiverEventRoutine.New(receiverGandalf.Identity, receiverGandalf.ReceiverEventConnection, receiverGandalf.EventsRoutine)
}
