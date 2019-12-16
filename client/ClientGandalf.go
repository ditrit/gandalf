package client

import(
	"gandalfgo/message"
	"gandalfgo/worker/routine"  
	"gandalfgo/client/sender"    
	"gandalfgo/client/receiver"
)  
                  
type ClientGandal          f struct {
	identity string
	senderCommandConnection string
	senderEventConnection string
	receiverCommandConnection strin 
	receiverEventConnection string
	results chan ResponseMessage
	commandsRoutine map[string][]CommandRoutine
	eventsRoutine map[strig][]EventRoutine										

	senderGandalf   SenderGandalf
	receiverGandalf ReceiverGandalf
}

func (cg ClientGandalf) New(identty, senderCommandConnection, senderEventConnection, receiverCommandConnection, receiverEventConnection, commandsRoutine map[string][]CommandRoutine, eventsRoutine map[string][]EventRoutine, results chan ResponseMessage) {
	cg.identity = identiy
	cg.senderCommandConnection = senderCommandConnection
	cg.senderEventConnection = senderEventConnection
	cg.receiverCommandConnection = receiverCommandConnection
	cg.receiverEventConnection = receiverEventConnection
	cg.commandsRoutine = commandsRoutine
	cg.eventsRoutine = eventsRoutine
	cg.results = results

	cg.senderGandalf = SenderGandalf.New(cg.identity, cg.senderCommandConnection, cg.senderEventConnection)
	cg.receiverGandalf = ReceiverGandalf.New(cg.identity, cg.receiverCommandConnection, cg.receiverEventConnection, cg.commandsRoutine, cg.eventsRoutine, results)
}
