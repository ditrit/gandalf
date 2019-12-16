package client

import(
	"gandalfgo/client/sender"
	"gandalfgo/client/listener"
)

type LibraryGandalf struct {
	senderGandalf   SenderGandalf
	listenerGandalf ListenerGandalf
}

func (lg LibraryGandalf) New(path string) {
	//identity, workerCommandReceiveC2WConnection, workerEventReceiveC2WConnection string, topics *string
	//LOAD CONF
	cg.senderGandalf = SenderGandalf.new()
	cg.listenerGandalf = ListenerGandalf.new()
}
