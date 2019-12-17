package listener

import (
	"fmt"
	"gandalfgo/message"
	"github.com/pebbe/zmq4"
)

type ListenerCommandRoutine struct {
	Context								*zmq4.Context
	ListenerCommandReceive              *zmq4.Socket
	ListenerCommandReceiveConnection   	string
	Identity                            string
	Commands 							[]CommandMessage
}

func (r ListenerCommandRoutine) NewListenerCommandRoutine(identity, listenerCommandReceiveConnection string) (listenerCommandRoutine *ListenerCommandRoutine) {
	listenerCommandRoutine = new(ListenerCommandRoutine)

	listenerCommandRoutine.Identity = identity

	listenerCommandRoutine.context, _ = zmq4.NewContext()
	listenerCommandRoutine.ListenerCommandReceiveConnection = listenerCommandReceiveConnection
	listenerCommandRoutine.ListenerCommandReceive = listenerCommandRoutine.Context.NewSocket(zmq4.DEALER)
	listenerCommandRoutine.ListenerCommandReceive.SetIdentity(listenerCommandRoutine.Identity)
	listenerCommandRoutine.ListenerCommandReceive.Connect(listenerCommandRoutine.ListenerCommandReceiveConnection)
	fmt.Printf("listenerCommandReceive connect : " + listenerCommandReceiveConnection)
}

func (r ListenerCommandRoutine) close() {
	r.ListenerCommandReceive.Close()
	r.Context.Term()
}

func (r ListenerCommandRoutine) run() {

	poller := zmq4.NewPoller()
	poller.Add(r.ListenerCommandReceive, zmq4.POLLIN)

	command := [][]byte{}

	for {
		r.sendReadyCommand()

		sockets, _ := poller.Poll(-1)
		for _, socket := range sockets {

			switch currentSocket := socket.Socket; currentSocket {
			case r.ListenerCommandReceive:

				command, err := currentSocket.RecvMessageBytes(0)
				if err != nil {
					panic(err)
				}
				err = r.processCommandReceive(command)
				if err != nil {
					panic(err)
				}
			}
		}
	}
	fmt.Println("done")
}

func (r ListenerCommandRoutine) processCommandReceive(command [][]byte) {
	r.Commands.append(message.decodeCommandMessage(command))
}

func (r ListenerCommandRoutine) GetCommand() (lastCommand CommandMessage, err error) {
	lastCommand = r.Commands[0]
	r.commands = append(r.Commands[:0][], s[0+1][]...)
	return
}