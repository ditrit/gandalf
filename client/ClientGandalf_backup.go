package client

import (
	"fmt"
	"gandalf-go/client/receiver"
	"gandalf-go/client/sender"
	"gandalf-go/message"
	"gandalf-go/worker/routine"
)

type ClientGandalf struct {
	Identity                  string
	SenderCommandConnection   string
	SenderEventConnection     string
	ReceiverCommandConnection string
	ReceiverEventConnection   string
	Replys                    chan message.CommandMessageReply
	CommandsRoutine           map[string][]routine.CommandRoutine
	EventsRoutine             map[string][]routine.EventRoutine
	SenderGandalf             *sender.SenderGandalf
	ReceiverGandalf           *receiver.ReceiverGandalf
	ClientStopChannel         chan int
}

func NewClientGandalf(identity, senderCommandConnection, senderEventConnection, receiverCommandConnection, receiverEventConnection string, commandsRoutine map[string][]routine.CommandRoutine, eventsRoutine map[string][]routine.EventRoutine, replys chan message.CommandMessageReply) (clientGandalf *ClientGandalf) {
	clientGandalf = new(ClientGandalf)
	clientGandalf.ClientStopChannel = make(chan int)

	clientGandalf.Identity = identity
	clientGandalf.SenderCommandConnection = senderCommandConnection
	clientGandalf.SenderEventConnection = senderEventConnection
	clientGandalf.ReceiverCommandConnection = receiverCommandConnection
	clientGandalf.ReceiverEventConnection = receiverEventConnection
	clientGandalf.CommandsRoutine = commandsRoutine
	clientGandalf.EventsRoutine = eventsRoutine
	//TODO USELESS ??
	clientGandalf.Replys = replys

	clientGandalf.SenderGandalf = sender.NewSenderGandalf(clientGandalf.Identity, clientGandalf.SenderCommandConnection, clientGandalf.SenderEventConnection)
	clientGandalf.ReceiverGandalf = receiver.NewReceiverGandalf(clientGandalf.Identity, clientGandalf.ReceiverCommandConnection, clientGandalf.ReceiverEventConnection, clientGandalf.CommandsRoutine, clientGandalf.EventsRoutine, clientGandalf.Replys)

	return
}

func (cg ClientGandalf) Run() {
	go cg.ReceiverGandalf.Run()
	for {
		select {
		case <-cg.ClientStopChannel:
			fmt.Println("quit")
			break
		}
	}
}

func (cg ClientGandalf) Stop() {
	cg.ClientStopChannel <- 0
}
