package sender

import (
	"fmt"
	"gandalfgo/message"
	"github.com/pebbe/zmq4"
)

type SenderEventRoutine struct {
	Context						*zmq4.Context
	SenderEventSend            	*zmq4.Socket
	SenderEventConnection  		string
	SenderEventConnections 		string
	Identity                   	string
}

func NewSenderEventRoutine(identity, senderEventConnection string) (senderEventRoutine *SenderEventRoutine) {
	senderEventRoutine = new(SenderEventRoutine)

	senderEventRoutine.Identity = identity

	senderEventRoutine.Context, _ := zmq4.NewContext()
	senderEventRoutine.SenderEventConnection = senderEventConnection
	senderEventRoutine.SenderEventSend = senderEventRoutine.Context.NewSocket(zmq.DEALER)
	senderEventRoutine.SenderEventSend.SetIdentity(senderEventRoutine.Identity)
	senderEventRoutine.SenderEventSend.Connect(senderEventRoutine.SenderEventConnection)
	fmt.Printf("senderEventSend connect : " + senderEventConnection)
}

func (r SenderEventRoutine) NewList(identity string, senderEventConnections *string) {
	senderEventRoutine = newSenderEventRoutine)

	senderEventRoutine.Identity = identity

	senderEventRoutine.Context, _ := zmq4.NewContext()
	senderEventRoutine.SenderEventConnections = senderEventConnections
	senderEventRoutine.SenderEventSend = senderEventRoutine.Context.NewSocket(zmq.DEALER)
	senderEventRoutine.SenderEventSend.SetIdentity(senderEventRoutine.Identity)

	fmt.Printf("senderEventSend connect : " + senderEventConnection)
	for _, connection := range senderEventRoutine.senderEventConnections {
		senderEventRoutine.senderEventSend.Connect(connection)
		fmt.Printf("senderEventSend connect : " + senderEventConnections)
	}
}

func (r SenderEventRoutine) sendEvent(topic, timeout, event, payload string) {
	eventMessage := message.NewEventMessage(topic, timeout, event, payload)
	if err != nil {
		panic(err)
	}
	go eventMessage.sendWith(r.SenderEventSend)
}

func (r SenderEventRoutine) close() {
}
