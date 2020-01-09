package client

import (
	"fmt"
	"gandalf-go/client/waiter"
	"gandalf-go/client/sender"
	"gandalf-go/message"
	"gandalf-go/worker/routine"
)

type ClientGandalf struct {
	Identity                  string
	SenderCommandConnection   string
	SenderEventConnection     string
	WaiterCommandConnection string
	WaiterEventConnection   string
	SenderGandalf             *sender.SenderGandalf
	WaiterGandalf             *waiter.WaiterGandalf
}

func NewClientGandalf(identity, senderCommandConnection, senderEventConnection, waiterCommandConnection, waiterEventConnection string) (clientGandalf *ClientGandalf) {
	clientGandalf = new(ClientGandalf)
	clientGandalf.ClientStopChannel = make(chan int)

	clientGandalf.Identity = identity
	clientGandalf.SenderCommandConnection = senderCommandConnection
	clientGandalf.SenderEventConnection = senderEventConnection
	clientGandalf.WaiterCommandConnection = waiterCommandConnection
	clientGandalf.WaiterEventConnection = waiterEventConnection

	clientGandalf.SenderGandalf = sender.NewSenderGandalf(clientGandalf.Identity, clientGandalf.SenderCommandConnection, clientGandalf.SenderEventConnection)
	clientGandalf.WaiterGandalf = waiter.NewWaiterGandalf(clientGandalf.Identity, clientGandalf.WaiterCommandConnection, clientGandalf.WaiterEventConnection)

	return
}
