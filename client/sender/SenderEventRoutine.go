package sender

import (
	"fmt"
	"gandalfgo/message"
	"github.com/zeromq/goczmq"
)

type SenderEventRoutine struct {
	context							*zmq.Context
	senderEventSend            zmq.Sock
	senderEventConnection  string
	senderEventConnections *string
	Identity                   string
	Responses                  *zmq.Message
}

func (r SenderEventRoutine) New(identity, senderEventConnection string) err error {
	r.identity = identity

	context, _ := zmq.NewContext()
	r.senderEventConnection = senderEventConnection
	r.senderEventSend = context.NewDealer(senderEventConnection)
	r.senderEventSend.Identity(r.identity)
	fmt.Printf("senderEventSend connect : " + senderEventConnection)
}

func (r SenderEventRoutine) NewList(identity string, senderEventConnections *string) err error {
	r.identity = identity

	context, _ := zmq.NewContext()
	r.senderEventConnections = senderEventConnections
	r.senderEventSend = context.NewDealer(senderEventConnections)
	r.senderEventSend.Identity(r.identity)
	fmt.Printf("senderEventSend connect : " + senderEventConnections)
}

func (r SenderEventRoutine) sendEvent(topic, timeout, event, payload string) err error {
	eventMessage := eventMessage.New(topic, timeout, event, payload)
	if err != nil {
		panic(err)
	}
	go eventMessage.sendWith(senderEventSend)
}

func (r SenderEventRoutine) close() err error {
}
