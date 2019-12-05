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

func (r ListenerCommandRoutine) New(identity, listenerCommandReceiveConnection string) err error {
	r.Identity = identity

	r.listenerCommandReceiveConnection = listenerCommandReceiveConnection
	r.listenerCommandReceive = zmq.NewDealer(listenerCommandReceiveConnection)
	r.listenerCommandReceive.Identity(r.identity)
	fmt.Printf("listenerCommandReceive connect : " + listenerCommandReceiveConnection)
}

func (r ListenerCommandRoutine) close() err error {
	r.listenerCommandReceive.close()
	r.Context.close()
}

func (r ListenerCommandRoutine) run() err error {
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
			err = r.processCommandReceive(command)
			if err != nil {
				panic(err)
			}
		}
	}
	fmt.Println("done")
}

func (r ListenerCommandRoutine) processCommandReceive(event [][]byte) err error {
}