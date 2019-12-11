package client

type ClientGandalf struct {
	senderGandalf   SenderGandalf
	receiverGandalf ReceiverGandalf
}

func (cg ClientGandalf) main() {
	clientConfiguration := ClientConfiguration.loadConfiguration(path)

	cg.senderGandalf = SenderGandalf.New(clientConfiguration.identity, clientConfiguration.senderCommandConnection, clientConfiguration.senderEventConnection)
	cg.receiverGandalf = ReceiverGandalf.New(clientConfiguration.identity, clientConfiguration.receiverCommandConnection, clientConfiguration.receiverEventConnection)
}
