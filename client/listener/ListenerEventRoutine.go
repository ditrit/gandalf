package listener

import (
	"fmt"
	"gandalfgo/message"
	"github.com/alecthomas/gozmq"
)

type ListenerEventRoutine struct {
	context							*gozmq.Context
	listenerEventReceive           	gozmq.Sock
	listenerEventReceiveConnection 	string
	identity                       	string
	events []EventMessage
}

func (r ListenerEventRoutine) New(identity, listenerEventReceiveConnection string) err error {
	r.identity = identity

	r.context, _ := gozmq.NewContext()	
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
	pi := gozmq.PollItems{
		gozmq.PollItem{Socket: listenerEventReceive, Events: gozmq.POLLIN}

	var event = [][]byte{}

	for {
		_, _ = zmq.Poll(pi, -1)

		switch {
		case pi[0].REvents&zmq.POLLIN != 0:

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