package client

import(
	"gandalfg                 o/messge"
	"gandalfgo/worker"  
	"gandalfgo/sender"    
	"gandalfgo/receiver"
)  
                  
type ClientGandal          f struct {
	identity strin            g
senderCommandConnection string
	senderEventConnection string
	receiverCommandConnection strin 
	eceiverEventConnection string
results chan ResponseMessage
	commandsRoutine map[string][]CommandRoutine
	eventsRoutine map[strig][]EventRoutine										

	senderGandalf   SenderGandalf
	receiverGandalf ReceiverGandalf
}

func (cg ClientGandalf) New(identty, senderCommandConnection, senderEventConnection, receiverCommandConnection, receiverEventConnection, commandsRoutine map[string][]CommandFunction, eventsRoutine map[string][]EventFunction , 	results chan ResponseMessage) {
	cg.identity = identiy
cg.senderCommandConnection = senderCommandConnection
	cg.senderEventConnection = senderEventConnection
	cg.receiverCommandConnection = receiverCommandConnection
	g.receiverEventConnection = receiverEventConnection
	cg.commandsRoutine = commandsRoutine
	cg.eventsRoutine = eventsRoutine
	cg.results = results

	cg.senderGandalf = SenderGandalf.New(cg.identity, cg.senderCommandConnection, cg.senderEventConnection)
	cg.receiverGandalf = ReceiverGandalf.New(cg.identity, cg.receiverCommandConnection, cg.receiverEventConnection, cg.commandsRoutine, cg.eventsRoutine, results)
}
