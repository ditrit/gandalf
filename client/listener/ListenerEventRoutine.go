package listener

import (
	"errors"
	"fmt"
	"gandalf-go/message"

	"github.com/pebbe/zmq4"
)

type ListenerEventRoutine struct {
	Context                        *zmq4.Context
	ListenerEventReceive           *zmq4.Socket
	ListenerEventReceiveConnection string
	Identity                       string
	Events                         []message.EventMessage
}

func NewListenerEventRoutine(identity, listenerEventReceiveConnection string) (listenerEventRoutine *ListenerEventRoutine) {
	listenerEventRoutine = new(ListenerEventRoutine)
	listenerEventRoutine.Identity = identity

	listenerEventRoutine.Context, _ = zmq4.NewContext()
	listenerEventRoutine.ListenerEventReceiveConnection = listenerEventReceiveConnection
	listenerEventRoutine.ListenerEventReceive, _ = listenerEventRoutine.Context.NewSocket(zmq4.SUB)
	listenerEventRoutine.ListenerEventReceive.SetIdentity(listenerEventRoutine.Identity)
	listenerEventRoutine.ListenerEventReceive.Connect(listenerEventRoutine.ListenerEventReceiveConnection)
	fmt.Printf("listenerEventReceive connect : " + listenerEventReceiveConnection)

	return
}

func (r ListenerEventRoutine) close() {
	r.ListenerEventReceive.Close()
	r.Context.Term()
}

func (r ListenerEventRoutine) run() {

	poller := zmq4.NewPoller()
	poller.Add(r.ListenerEventReceive, zmq4.POLLIN)

	event := [][]byte{}
	err := errors.New("")

	for {

		sockets, _ := poller.Poll(-1)
		for _, socket := range sockets {

			switch currentSocket := socket.Socket; currentSocket {
			case r.ListenerEventReceive:

				event, err = currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}
				r.processEventReceive(event)
			}
		}
	}
	fmt.Println("done")
}

func (r ListenerEventRoutine) processEventReceive(event [][]byte) {
	eventMessage, _ := message.DecodeEventMessage(event[1])
	r.Events = append(r.Events, eventMessage)
}

func (r ListenerEventRoutine) getEvents() (lastEvent message.EventMessage, err error) {
	lastEvent = r.Events[0]
	r.Events = append(r.Events[:0], r.Events[0+1])
	return
}
