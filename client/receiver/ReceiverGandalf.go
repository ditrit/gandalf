package receiver

import(
	"gandalfg                 eiver"
)
  
type ReceiverGandalf str   t {
	identity 					string   
	receiver                  CommandConnection 	sring
	receiverEventCon          nection 	string
	receiverComman            dRoutine 		ReceiverCommandoutine
	eceiverEventRoutine   		ReceiverEventRoutine
results chan ResponseMessage
	commandsRoutine map[string][]CommandFunction
	eventsRoutine map[strig][]EventFunction
}

func (rg ReceiverGandalf) New(identit, receiverCommandConnection, receiverEventConnection string, 	commandsRoutine map[string][]CommandFunction, eventsRoutine map[string][]EventFunction, results chan ResponseMessage) {
	rg.identity = identity
	rg.receiverCommandConection = receiverCommandConnection
	rg.receiverEventRoutine = receiverEventRoutine
	rg.commandsRoutine = commandsRoutine
	g.eventsRoutine = eventsRoutine
	rg.results = results
	rg.receiverCommandRoutine = ReceiverCommandRoutine.New(rg.identity, rg.receiverCommandConnection, rg.commandsRoutine , rg.results)
	rg.receiverEventRoutine = ReceiverEventRoutine.New(rg.idenitty, rg.receiverEventConnection, rg.eventsRoutine)
}
