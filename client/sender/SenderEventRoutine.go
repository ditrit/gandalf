package sender

import (
	"fmt"
	"gandalfgo/message"
	"github.com/pebbe/zmq4"
)

type SenderEventRoutine struct {
	context						*zmq4.Context
	senderEventSend            	*zmq4.Socket
	senderEventConnection  		string
	senderEventConnections 		*string
	Identity                   	string
	Responses                  	*zmq4.Message
}

func (r SenderEventRoutine) New(identity, senderEventConnection string) err error {
	r.identity = identity

	r.context, _ := zmq4.NewContext()
	r.senderEventConnection = senderEventConnection
	r.senderEventSend = r.context.NewDealer(senderEventConnection)
	r.senderEventSend.Identity(r.identity)
	fmt.Printf("senderEventSend connect : " + senderEventConnection)
}

func (r SenderEventRoutine) NewList(identity string, senderEventConnections *string) err error {
	r.identity = identity

	context, _ := zmq4.NewContext()
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
