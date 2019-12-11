package client

type ClientGandalf struct {
	identity string
	senderCommandConnection string
	senderEventConnection string
	receiverCommandConnection string 
	receiverEventConnection string
	results chan ResponseMessage
	commandsRoutine map[string][]CommandFunction
	eventsRoutine map[string][]EventFunction										

	senderGandalf   SenderGandalf
	receiverGandalf ReceiverGandalf
}

func (cg ClientGandalf) New(identity, senderCommandConnection, senderEventConnection, receiverCommandConnection, receiverEventConnection, commandsRoutine map[string][]CommandFunction, eventsRoutine map[string][]EventFunction , 	results chan ResponseMessage) {
	cg.identity = identity
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
