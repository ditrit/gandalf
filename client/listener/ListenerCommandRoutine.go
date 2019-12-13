package listener

import (
	"fmt"
	"gandalfgo/message"
	zmq "github.com/zeromq/goczmq"
)

type ListenerCommandRoutine struct {
	context							*zmq.Context
	listenerCommandReceive              zmq.Sock
	listenerCommandReceiveConnection   string
	identity                               string
	commands []CommandMessage
}

func (r ListenerCommandRoutine) New(identity, listenerCommandReceiveConnection string) err error {
	r.Identity = identity

	context, _ := zmq.NewContext()
	r.listenerCommandReceiveConnection = listenerCommandReceiveConnection
	r.listenerCommandReceive = context.NewDealer(listenerCommandReceiveConnection)
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
			err = r.processCommandReceive(command)
			if err != nil {
				panic(err)
			}
		}
	}
	fmt.Println("done")
}

func (r ListenerCommandRoutine) processCommandReceive(command [][]byte) err error {
	r.commands.append(CommandMessage.decodeCommand(command))
}

func (r ListenerCommandRoutine) GetCommand() (lastCommand CommandMessage, err error) {
	lastCommand := r.commands[0]
	r.commands = append(r.commands[:0][], s[0+1][]...)
	return
}