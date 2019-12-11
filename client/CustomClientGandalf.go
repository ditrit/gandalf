package library

type LibraryGandalf struct {
	senderGandalf   SenderGandalf
	listenerGandalf ListenerGandalf
}

func (lg LibraryGandalf) main() {
	//identity, workerCommandReceiveC2WConnection, workerEventReceiveC2WConnection string, topics *string
	//LOAD CONF
	cg.senderGandalf = SenderGandalf.new()
	cg.listenerGandalf = ListenerGandalf.new()
}
