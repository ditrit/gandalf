package receiver

type ReceiverGandalf struct {
	identity 					string
	receiverCommandConnection 	string
	receiverEventConnection 	string
	receiverCommandRoutine 		ReceiverCommandRoutine
	receiverEventRoutine   		ReceiverEventRoutine
	results chan ResponseMessage
	commandsRoutine map[string][]CommandFunction
	eventsRoutine map[string][]EventFunction
}

func (rg ReceiverGandalf) New(identity, receiverCommandConnection, receiverEventConnection string, 	commandsRoutine map[string][]CommandFunction, eventsRoutine map[string][]EventFunction, results chan ResponseMessage) {
	rg.identity = identity
	rg.receiverCommandConnection = receiverCommandConnection
	rg.receiverEventRoutine = receiverEventRoutine
	rg.commandsRoutine = commandsRoutine
	rg.eventsRoutine = eventsRoutine
	rg.results = results
	rg.receiverCommandRoutine = ReceiverCommandRoutine.New(rg.identity, rg.receiverCommandConnection, rg.commandsRoutine , rg.results)
	rg.receiverEventRoutine = ReceiverEventRoutine.New(rg.idenitty, rg.receiverEventConnection, rg.eventsRoutine)
}
