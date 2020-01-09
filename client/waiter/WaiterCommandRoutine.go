package waiter

import (
	"errors"
	"fmt"
	"gandalf-go/message"
	"gandalf-go/worker/routine"

	"github.com/pebbe/zmq4"
)

type WaiterCommandRoutine struct {
	Context                      *zmq4.Context
	WaiterCommandReceive         *zmq4.Socket
	WaiterCommandConnection    string
	WaiterEventReceive           *zmq4.Socket
	WaiterEventReceiveConnection string
	Identity                     string
}

func NewWaiterCommandRoutine(identity, waiterCommandConnection string, commandsRoutine map[string][]routine.CommandRoutine, results chan message.CommandMessageReply) (waiterCommandRoutine *WaiterCommandRoutine) {
	waiterCommandRoutine = new(WaiterCommandRoutine)

	waiterCommandRoutine.Identity = identity
	waiterCommandRoutine.WaiterCommandConnection = waiterCommandConnection
	waiterCommandRoutine.CommandsRoutine = commandsRoutine
	waiterCommandRoutine.Replys = make(chan message.CommandMessageReply)

	waiterCommandRoutine.Context, _ = zmq4.NewContext()
	waiterCommandRoutine.WaiterCommandReceive, _ = waiterCommandRoutine.Context.NewSocket(zmq4.DEALER)
	waiterCommandRoutine.WaiterCommandReceive.SetIdentity(waiterCommandRoutine.Identity)
	waiterCommandRoutine.WaiterCommandReceive.Connect(waiterCommandRoutine.WaiterCommandConnection)
	fmt.Println("workerCommandReceive connect : " + waiterCommandConnection)

	return
}

func (r WaiterCommandRoutine) WaitCommand(string uuid) (commandMessage CommandMessage) {
	//SEND REQUEST
	for {
		command, err := r.WaiterCommandReceive.RecvMessageBytes(0)
		if err != nil {
			panic(err)
		}
		commandMessage, _ := message.DecodeCommandMessage(command[1])
		return
	}
}

func (r WaiterCommandRoutine) WaitCommandReply(uuid string) (commandMessageReply CommandMessageReply) {
	//SEND REQUEST
	for {
		command, err := r.WaiterCommandReceive.RecvMessageBytes(0)
		if err != nil {
			panic(err)
		}
		commandMessageReply, _ := message.DecodeCommandMessageReply(command[1])
		return
	}
}
