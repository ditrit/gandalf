package listener

import (
	"fmt"
	"gandalfgo/message"
	"github.com/pebbe/zmq4"
)

type ListenerEventRoutine struct {
	Context							*zmq4.Context
	ListenerEventReceive           	*zmq4.Socket
	ListenerEventReceiveConnection 	string
	Identity                       	string
	Events 							[]EventMessage
}

func NewListenerEventRoutine(identity, listenerEventReceiveConnection string) (listenerEventRoutine *ListenerEventRoutine) {
	listenerEventRoutine = new(ListenerEventRoutine)
	listenerEventRoutine.identity = identity

	listenerEventRoutine.Context, _ := zmq4.NewContext()	
	listenerEventRoutine.ListenerEventReceiveConnection = listenerEventReceiveConnection
	listenerEventRoutine.ListenerEventReceive = listenerEventRoutine.Context.NewSocket(zmq4.SUB)
	listenerEventRoutine.ListenerEventReceive.SetIdentity(listenerEventRoutine.Identity)
	listenerEventRoutine.ListenerEventReceive.Connect(rlistenerEventRoutine.ListenerEventReceiveConnection)
	fmt.Printf("listenerEventReceive connect : " + listenerEventReceiveConnection)
}

func (r ListenerEventRoutine) close() {
	r.ListenerEventReceive.Close()
	r.Context.Term()
}

func (r ListenerEventRoutine) run() {

	poller := zmq4.NewPoller()
	poller.Add(r.ListenerEventReceive, zmq4.POLLIN)

	event := [][]byte{}

	for {

		sockets, _ := poller.Poll(-1)
		for _, socket := range sockets {

			switch currentSocket := socket.Socket; currentSocket {
			case r.ListenerEventReceive:

				event, err := pi[0].currentSocket.RecvMessageBytes(0)
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
	r.Events.append(message.decodeEventMessage(event))
}

func (r ListenerEventRoutine) getEvents() (lastEvent EventMessage, err error) {
	lastEvent := r.Events[0]
	r.events = append(r.events[:0][], s[0+1][]...)
	return
}