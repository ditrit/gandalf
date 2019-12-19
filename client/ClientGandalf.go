package client

type ClientGandalf struct {
	Identity                  string
	SenderCommandConnection   string
	SenderEventConnection     string
	ReceiverCommandConnection string
	ReceiverEventConnection   string
	Replys                    chan CommandMessageReply
	CommandsRoutine           map[string][]CommandRoutine
	EventsRoutine             map[string][]EventRoutine
	SenderGandalf             *SenderGandalf
	ReceiverGandalf           *ReceiverGandalf
}

func NewClientGandalf(identty, senderCommandConnection, senderEventConnection, receiverCommandConnection, receiverEventConnection, commandsRoutine map[string][]CommandRoutine, eventsRoutine map[string][]EventRoutine, replys chan ResponseMessage) (clientGandalf *ClientGandalf) {
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
