package sender

type SenderGandalf struct {
	identity 			 	string
	senderCommandConnection string
	senderEventConnection 	string
	senderCommandRoutine 	SenderCommandRoutine
	senderEventRoutine   	SenderEventRoutine
}

func (sg SenderGandalf) New(identity, senderCommandConnection, senderEventConnection) {
	sg.identity = identity
	sg.senderCommandConnection = senderCommandConnection
	sg.senderEventConnection = senderEventConnection
	sg.senderCommandRoutine = SenderCommandRoutine.New(identity, senderCommandConnection)
	sg.senderEventRoutine = SenderEventRoutine.New(identity, senderEventConnection)
}
