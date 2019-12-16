package listener

import (
	"fmt"
	"gandalfgo/message"
	"github.com/pebbe/zmq4"
)

type ListenerCommandRoutine struct {
	context								zmq4.Context
	listenerCommandReceive              zmq4.Socket
	listenerCommandReceiveConnection   	string
	identity                            string
	commands []CommandMessage
}

func (r ListenerCommandRoutine) New(identity, listenerCommandReceiveConnection string) {
	r.Identity = identity

	r.context, _ = zmq4.NewContext()
	r.listenerCommandReceiveConnection = listenerCommandReceiveConnection
	r.listenerCommandReceive = r.context.NewSocket(zmq4.DEALER)
	r.listenerCommandReceive.SetIdentity(r.identity)
	r.listenerCommandReceive.Connect(r.listenerCommandReceiveConnection)
	fmt.Printf("listenerCommandReceive connect : " + listenerCommandReceiveConnection)
}

func (r ListenerCommandRoutine) close() {
	r.listenerCommandReceive.close()
	r.Context.close()
}

func (r ListenerCommandRoutine) run() {

	poller := zmq4.NewPoller()
	poller.Add(r.listenerCommandReceive, zmq4.POLLIN)

	command := [][]byte{}

	for {
		r.sendReadyCommand()

		sockets, _ := poller.Poll(-1)
		for _, socket := range sockets {

			switch currentSocket := socket.Socket; currentSocket {
			case listenerCommandReceive:

				command, err := currentSocket.RecvMessage()
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
	r.commands.append(CommandMessage.decodeCommand(command))
}

func (r ListenerCommandRoutine) GetCommand() (lastCommand CommandMessage, err error) {
	lastCommand = r.commands[0]
	r.commands = append(r.commands[:0][], s[0+1][]...)
	return
}