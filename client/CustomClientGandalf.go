package client

import(
	"gandalfgo/client/sender"
	"gandalfgo/client/listener"
)

type LibraryGandalf struct {
	SenderGandalf   SenderGandalf
	ListenerGandalf ListenerGandalf
}

func (lg LibraryGandalf) NewLibraryGandalf(path string) (libraryGandalf LibraryGandalf) {
	//identity, workerCommandReceiveC2WConnection, workerEventReceiveC2WConnection string, topics *string
	//LOAD CONF
	libraryGandalf = new(LibraryGandalf)

	libraryGandalf.SenderGandalf = NewSenderGandalf()
	libraryGandalf.ListenerGandalf = NewListenerGandalf()
}
