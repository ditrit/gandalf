package sender

type SenderGandalf struct {
	Identity 			 	string
	SenderCommandConnection string
	SenderEventConnection 	string
	SenderCommandRoutine 	*SenderCommandRoutine
	SenderEventRoutine   	*SenderEventRoutine
}

func NewSenderGandalf(identity, senderCommandConnection, senderEventConnection) (senderGandalf SenderGandalf) {
	senderGandalf.Identity = identity
	senderGandalf.SenderCommandConnection = senderCommandConnection
	senderGandalf.SenderEventConnection = senderEventConnection
	senderGandalf.SenderCommandRoutine = NewSenderCommandRoutine(identity, senderCommandConnection)
	senderGandalf.SenderEventRoutine = NewSenderEventRoutine(identity, senderEventConnection)
}
