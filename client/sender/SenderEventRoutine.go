package sender

import (
	"fmt"
	"message"
	zmq "github.com/zeromq/goczmq"
)

type SenderEventRoutine struct {
	senderEventSend            zmq.Sock
	senderEventConnection  string
	senderEventConnections *string
	Identity                   string
	Responses                  *zmq.Message
}

func (r SenderEventRoutine) New(identity, senderEventConnection string) err error {
	r.identity = identity
	r.senderEventConnection = senderEventConnection
	r.senderEventSend = zmq.NewDealer(senderEventConnection)
	r.senderEventSend.Identity(r.identity)
	fmt.Printf("senderEventSend connect : " + senderEventConnection)
}

func (r SenderEventRoutine) NewList(identity string, senderEventConnections *string) err error {
	r.identity = identity
	r.senderEventConnections = senderEventConnections
	r.senderEventSend = zmq.NewDealer(senderEventConnections)
	r.senderEventSend.Identity(r.identity)
	fmt.Printf("senderEventSend connect : " + senderEventConnections)
}

func (r SenderEventRoutine) sendEvent(topic, timeout, event, payload string) err error {
	eventMessage := eventMessage.New(topic, timeout, event, payload)
	if err != nil {
		panic(err)
	}
	eventMessage.sendWith(senderEventSend)
}

func (r SenderEventRoutine) close() err error {
}
