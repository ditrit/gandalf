package sender

import (
	"fmt"
	"gandalfgo/message"

	"github.com/pebbe/zmq4"
)

type SenderEventRoutine struct {
	Context                *zmq4.Context
	SenderEventSend        *zmq4.Socket
	SenderEventConnection  string
	SenderEventConnections []string
	Identity               string
}

func NewSenderEventRoutine(identity, senderEventConnection string) (senderEventRoutine *SenderEventRoutine) {
	senderEventRoutine = new(SenderEventRoutine)

	senderEventRoutine.Identity = identity

	senderEventRoutine.Context, _ = zmq4.NewContext()
	senderEventRoutine.SenderEventConnection = senderEventConnection
	senderEventRoutine.SenderEventSend, _ = senderEventRoutine.Context.NewSocket(zmq4.DEALER)
	senderEventRoutine.SenderEventSend.SetIdentity(senderEventRoutine.Identity)
	senderEventRoutine.SenderEventSend.Connect(senderEventRoutine.SenderEventConnection)
	fmt.Printf("senderEventSend connect : " + senderEventConnection)

	return
}

func NewSenderEventRoutineList(identity string, senderEventConnections []string) (senderEventRoutine *SenderEventRoutine) {
	senderEventRoutine = new(SenderEventRoutine)

	senderEventRoutine.Identity = identity

	senderEventRoutine.Context, _ = zmq4.NewContext()
	senderEventRoutine.SenderEventConnections = senderEventConnections
	senderEventRoutine.SenderEventSend, _ = senderEventRoutine.Context.NewSocket(zmq4.DEALER)
	senderEventRoutine.SenderEventSend.SetIdentity(senderEventRoutine.Identity)

	for _, connection := range senderEventRoutine.SenderEventConnections {
		senderEventRoutine.SenderEventSend.Connect(connection)
		fmt.Printf("senderEventSend connect : " + connection)
	}
	return
}

func (r SenderEventRoutine) sendEvent(topic, timeout, uuid, event, payload string) {
	eventMessage := message.NewEventMessage(topic, timeout, uuid, event, payload)
	if err != nil {
		panic(err)
	}
	go eventMessage.SendEventWith(r.SenderEventSend)
}

func (r SenderEventRoutine) close() {
}
