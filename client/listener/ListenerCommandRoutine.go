package listener

import (
	"fmt"
	"gandalfgo/message"
	"github.com/pebbe/zmq4"
)

type ListenerCommandRoutine struct {
	context								*zmq4.Context
	listenerCommandReceive              zmq4.Socket
	listenerCommandReceiveConnection   	string
	identity                            string
	commands []CommandMessage
}

func (r ListenerCommandRoutine) New(identity, listenerCommandReceiveConnection string) err error {
	r.Identity = identity

	r.context, _ := zmq4.NewContext()
	r.listenerCommandReceiveConnection = listenerCommandReceiveConnection
	r.listenerCommandReceive = r.context.NewDealer(listenerCommandReceiveConnection)
	r.listenerCommandReceive.Identity(r.identity)
	fmt.Printf("listenerCommandReceive connect : " + listenerCommandReceiveConnection)
}

func (r ListenerCommandRoutine) close() err error {
	r.listenerCommandReceive.close()
	r.Context.close()
}

func (r ListenerCommandRoutine) run() err error {
	pi := zmq4.PollItems{
		zmq4.PollItem{Socket: listenerCommandReceive, Events: zmq4.POLLIN}

	var command = [][]byte{}

	for {
		r.sendReadyCommand()

		_, _ = zmq4.Poll(pi, -1)

		switch {
		case pi[0].REvents&zmq4.POLLIN != 0:

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