package listener

import (
	"fmt"
	"message"
	zmq "github.com/zeromq/goczmq"
)

type ListenerEventRoutine struct {
	listenerEventReceive           zmq.Sock
	listenerEventReceiveConnection string
	identity                       string
	events []message.EventMessage
}

func (r ListenerEventRoutine) New(identity, listenerEventReceiveConnection string) err error {
	r.identity = identity

	r.listenerEventReceiveConnection = listenerEventReceiveConnection
	r.listenerEventReceive = zmq.NewSub(listenerEventReceiveConnection)
	r.listenerEventReceive.Identity(r.identity)
	fmt.Printf("listenerEventReceive connect : " + listenerEventReceiveConnection)
}

func (r ListenerEventRoutine) close() err error {
	r.listenerEventReceive.close()
	r.Context.close()
}

func (r ListenerEventRoutine) run() {
	pi := zmq.PollItems{
		zmq.PollItem{Socket: listenerEventReceive, Events: zmq.POLLIN}

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