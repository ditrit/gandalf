package client

import "gandalf-go/client/sender"

type LibraryGandalf struct {
	SenderGandalf   sender.SenderGandalf
	ListenerGandalf listener.ListenerGandalf
}

func (lg LibraryGandalf) NewLibraryGandalf(path string) (libraryGandalf LibraryGandalf) {
	//identity, workerCommandReceiveC2WConnection, workerEventReceiveC2WConnection string, topics *string
	//LOAD CONF
	libraryGandalf = new(LibraryGandalf)

	libraryGandalf.SenderGandalf = NewSenderGandalf()
	libraryGandalf.ListenerGandalf = NewListenerGandalf()
}
