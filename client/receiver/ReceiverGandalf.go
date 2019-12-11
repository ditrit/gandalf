package receiver

type ReceiverGandalf struct {
	identity 					string
	receiverCommandConnection 	string
	receiverEventConnection 	string
	receiverCommandRoutine 		ReceiverCommandRoutine
	receiverEventRoutine   		ReceiverEventRoutine
}

func (rg ReceiverGandalf) New(identity, receiverCommandConnection, receiverEventConnection string) {
	rg.identity = identity
	rg.receiverCommandConnection = receiverCommandConnection
	rg.receiverEventRoutine = receiverEventRoutine
	rg.receiverCommandRoutine = ReceiverCommandRoutine.New(identity, receiverCommandConnection)
	rg.receiverEventRoutine = ReceiverEventRoutine.New(idenitty, receiverEventConnection)
}
