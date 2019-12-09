package client

import (
	"fmt"
	"message"
	zmq "github.com/zeromq/goczmq"
)

type ClientEventRoutine struct {
	clientEventSend            zmq.Sock
	clientEventSendConnection  string
	clientEventSendConnections *string
	Identity                   string
	Responses                  *zmq.Message
}

func (r ClientEventRoutine) New(identity, clientEventSendConnection string) err error {
	r.identity = identity
	r.clientEventSendConnection = clientEventSendConnection
	r.clientEventSend = zmq.NewDealer(clientEventSendConnection)
	r.clientEventSend.Identity(r.identity)
	fmt.Printf("clientEventSend connect : " + clientEventSendConnection)
}

func (r ClientEventRoutine) NewList(identity string, clientEventSendConnections *string) err error {
	r.identity = identity
	r.clientEventSendConnections = clientEventSendConnections
	r.clientEventSend = zmq.NewDealer(clientEventSendConnections)
	r.clientEventSend.Identity(r.identity)
	fmt.Printf("clientEventSend connect : " + clientEventSendConnections)
}

func (r ClientEventRoutine) sendEvent(topic, timeout, event, payload string) err error {
	eventMessage := eventMessage.New(topic, timeout, event, payload)
	if err != nil {
		panic(err)
	}
	eventMessage.sendWith(clientEventSend)
}

func (r ClientEventRoutine) close() err error {
}
