package listener

import (
	"fmt"

	zmq "github.com/zeromq/goczmq"
)

type ListenerCommandRoutine struct {
	listenerCommandReceive              zmq.Sock
	listenerCommandReceiveConnection   string
	identity                               string
	commands [][]byte{}
}

func (r ListenerCommandRoutine) new(identity, listenerCommandReceiveConnection string) {
	r.Identity = identity

	r.listenerCommandReceiveConnection = listenerCommandReceiveConnection
	r.listenerCommandReceive = zmq.NewDealer(listenerCommandReceiveConnection)
	r.listenerCommandReceive.Identity(r.identity)
	fmt.Printf("listenerCommandReceive connect : " + listenerCommandReceiveConnection)
}

func (r ListenerCommandRoutine) close() {
	r.listenerCommandReceive.close()
	r.Context.close()
}

func (r ListenerCommandRoutine) run() {
	pi := zmq.PollItems{
		zmq.PollItem{Socket: listenerCommandReceive, Events: zmq.POLLIN}

	var command = [][]byte{}

	for {
		r.sendReadyCommand()

		_, _ = zmq.Poll(pi, -1)

		switch {
		case pi[0].REvents&zmq.POLLIN != 0:

			command, err := pi[0].Socket.RecvMessage()
			if err != nil {
				panic(err)
			}
			//STORE IN COMMANDS
			err = routerSock.SendMessage(msg)
			if err != nil {
				panic(err)
			}
		}
	}

	fmt.Println("done")

}
