package listener

import (
	"fmt"
	"gandalfgo/message"
	"github.com/pebbe/zmq4"
)

type ListenerEventRoutine struct {
	context							zmq4.Context
	listenerEventReceive           	zmq4.Sock
	listenerEventReceiveConnection 	string
	identity                       	string
	events []EventMessage
}

func (r ListenerEventRoutine) New(identity, listenerEventReceiveConnection string) {
	r.identity = identity

	r.context, _ := zmq4.NewContext()	
	r.listenerEventReceiveConnection = listenerEventReceiveConnection
	r.listenerEventReceive = r.context.NewSocket(zmq4.SUB)
	r.listenerEventReceive.SetIdentity(r.identity)
	r.listenerEventReceive.Connect(r.listenerEventReceiveConnection)
	fmt.Printf("listenerEventReceive connect : " + listenerEventReceiveConnection)
}

func (r ListenerEventRoutine) close() {
	r.listenerEventReceive.close()
	r.Context.close()
}

func (r ListenerEventRoutine) run() {

	poller := zmq4.NewPoller()
	poller.Add(r.listenerEventReceive, zmq4.POLLIN)

	event := [][]byte{}

	for {

		sockets, _ := poller.Poll(-1)
		for _, socket := range sockets {

			switch currentSocket := socket.Socket; currentSocket {
			case listenerEventReceive:

				event, err := pi[0].currentSocket.RecvMessage()
				if err != nil {
					panic(err)
				}
				err = r.processEventReceive(event)
				if err != nil {
					panic(err)
				}
			}
		}
	}
	fmt.Println("done")
}

func (r ListenerEventRoutine) processEventReceive(event [][]byte) {
	r.events.append(message.EventMessage.decodeEvent(event))
}

func (r ListenerEventRoutine) getEvents() (lastEvent EventMessage, err error) {
	lastEvent := r.events[0]
	r.events = append(r.events[:0][], s[0+1][]...)
	return
}