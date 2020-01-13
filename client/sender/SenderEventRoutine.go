package sender

import (
	"fmt"
	"gandalf-go/message"

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
	senderEventRoutine.SenderEventSend, _ = senderEventRoutine.Context.NewSocket(zmq4.PUB)
	senderEventRoutine.SenderEventSend.SetIdentity(senderEventRoutine.Identity)
	senderEventRoutine.SenderEventSend.Connect(senderEventRoutine.SenderEventConnection)
	fmt.Println("senderEventSend connect : " + senderEventConnection)

	return
}

func NewSenderEventRoutineList(identity string, senderEventConnections []string) (senderEventRoutine *SenderEventRoutine) {
	senderEventRoutine = new(SenderEventRoutine)

	senderEventRoutine.Identity = identity

	senderEventRoutine.Context, _ = zmq4.NewContext()
	senderEventRoutine.SenderEventConnections = senderEventConnections
	senderEventRoutine.SenderEventSend, _ = senderEventRoutine.Context.NewSocket(zmq4.PUB)
	senderEventRoutine.SenderEventSend.SetIdentity(senderEventRoutine.Identity)

	for _, connection := range senderEventRoutine.SenderEventConnections {
		senderEventRoutine.SenderEventSend.Connect(connection)
		fmt.Println("senderEventSend connect : " + connection)
	}
	return
}

func (r SenderEventRoutine) SendEvent(topic, timeout, uuid, event, payload string) {
	eventMessage := message.NewEventMessage(topic, timeout, uuid, event, payload)
	go eventMessage.SendMessageWith(r.SenderEventSend)
}

func (r SenderEventRoutine) close() {
}
