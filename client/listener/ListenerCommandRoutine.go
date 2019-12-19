package listener

import (
	"errors"
	"fmt"
	"gandalf-go/message"

	"github.com/pebbe/zmq4"
)

type ListenerCommandRoutine struct {
	Context                          *zmq4.Context
	ListenerCommandReceive           *zmq4.Socket
	ListenerCommandReceiveConnection string
	Identity                         string
	Commands                         []message.CommandMessage
}

func NewListenerCommandRoutine(identity, listenerCommandReceiveConnection string) (listenerCommandRoutine *ListenerCommandRoutine) {
	listenerCommandRoutine = new(ListenerCommandRoutine)

	listenerCommandRoutine.Identity = identity

	listenerCommandRoutine.Context, _ = zmq4.NewContext()
	listenerCommandRoutine.ListenerCommandReceiveConnection = listenerCommandReceiveConnection
	listenerCommandRoutine.ListenerCommandReceive, _ = listenerCommandRoutine.Context.NewSocket(zmq4.DEALER)
	listenerCommandRoutine.ListenerCommandReceive.SetIdentity(listenerCommandRoutine.Identity)
	listenerCommandRoutine.ListenerCommandReceive.Connect(listenerCommandRoutine.ListenerCommandReceiveConnection)
	fmt.Printf("listenerCommandReceive connect : " + listenerCommandReceiveConnection)

	return
}

func (r ListenerCommandRoutine) close() {
	r.ListenerCommandReceive.Close()
	r.Context.Term()
}

func (r ListenerCommandRoutine) run() {

	poller := zmq4.NewPoller()
	poller.Add(r.ListenerCommandReceive, zmq4.POLLIN)

	command := [][]byte{}
	err := errors.New("")

	for {

		sockets, _ := poller.Poll(-1)
		for _, socket := range sockets {

			switch currentSocket := socket.Socket; currentSocket {
			case r.ListenerCommandReceive:

				command, err = currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}
				r.processCommandReceive(command)
			}
		}
	}
	fmt.Println("done")
}

func (r ListenerCommandRoutine) processCommandReceive(command [][]byte) {
	commandMessage, _ := message.DecodeCommandMessage(command[1])
	r.Commands = append(r.Commands, commandMessage)
}

func (r ListenerCommandRoutine) GetCommand() (lastCommand message.CommandMessage, err error) {
	lastCommand = r.Commands[0]
	//TODO REVOIR
	r.Commands = append(r.Commands[:0], r.Commands[0+1])
	return
}
