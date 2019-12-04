package listener

import (
	"fmt"

	zmq "github.com/zeromq/goczmq"
)

type ListenerEventRoutine struct {
	listenerEventReceive           zmq.Sock
	listenerEventReceiveConnection string
	identity                       string
	events [][]byte{}
}

func (r ListenerEventRoutine) new(identity, listenerEventReceiveConnection string) {
	r.identity = identity

	r.listenerEventReceiveConnection = listenerEventReceiveConnection
	r.listenerEventReceive = zmq.NewSub(listenerEventReceiveConnection)
	r.listenerEventReceive.Identity(r.identity)
	fmt.Printf("listenerEventReceive connect : " + listenerEventReceiveConnection)
}

func (r ListenerEventRoutine) close() {
	r.listenerEventReceive.close()
	r.Context.close()
}

func (r ListenerEventRoutine) run() {
	pi := zmq.PollItems{
		zmq.PollItem{Socket: listenerEventReceive, Events: zmq.POLLIN}

	var command = [][]byte{}

	for {
		_, _ = zmq.Poll(pi, -1)

		switch {
		case pi[0].REvents&zmq.POLLIN != 0:

			command, err := pi[0].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			//STORE IN events
			err = routerSock.SendMessage(msg)
			if err != nil {
				panic(err)
			}
		}
	}

	fmt.Println("done")

}
