package library

type LibraryGandalf struct {
	clientGandalf   ClientGandalf
	listenerGandalf ListenerGandalf
}

func (lg LibraryGandalf) main() {
	//identity, workerCommandReceiveC2WConnection, workerEventReceiveC2WConnection string, topics *string
	//LOAD CONF
	cg.clientGandalf = ClientGandalf.new()
	cg.listenerGandalf = ListenerGandalf.new()
}
