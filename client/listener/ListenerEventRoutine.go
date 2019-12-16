package listener

import (
	"fmt"
	"gandalfgo/message"
	"github.com/pebbe/zmq4"
)

type ListenerEventRoutine struct {
	context							*zmq4.Context
	listenerEventReceive           	zmq4.Sock
	listenerEventReceiveConnection 	string
	identity                       	string
	events []EventMessage
}

func (r ListenerEventRoutine) New(identity, listenerEventReceiveConnection string) err error {
	r.identity = identity

	r.context, _ := zmq4.NewContext()	
	r.listenerEventReceiveConnection = listenerEventReceiveConnection
	r.listenerEventReceive = r.context.NewSub(listenerEventReceiveConnection)
	r.listenerEventReceive.Identity(r.identity)
	fmt.Printf("listenerEventReceive connect : " + listenerEventReceiveConnection)
}

func (r ListenerEventRoutine) close() err error {
	r.listenerEventReceive.close()
	r.Context.close()
}

func (r ListenerEventRoutine) run() {
	pi := zmq4.PollItems{
		zmq4.PollItem{Socket: listenerEventReceive, Events: zmq4.POLLIN}

	var event = [][]byte{}

	for {
		_, _ = zmq4.Poll(pi, -1)

		switch {
		case pi[0].REvents&zmq4.POLLIN != 0:

			event, err := pi[0].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			err = r.processEventReceive(event)
			if err != nil {
				panic(err)
			}
		}
	}
	fmt.Println("done")
}

func (r ListenerEventRoutine) processEventReceive(event [][]byte) err error {
	r.events.append(message.EventMessage.decodeEvent(event))
}

func (r ListenerEventRoutine) getEvents() (lastEvent EventMessage, err error) {
	lastEvent := r.events[0]
	r.events = append(r.events[:0][], s[0+1][]...)
	return
}