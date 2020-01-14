package waiter

import (
	"fmt"
	"gandalf-go/constant"
	"gandalf-go/message"

	"github.com/pebbe/zmq4"
)

type WaiterCommandRoutine struct {
	Context                      *zmq4.Context
	WaiterCommandReceive         *zmq4.Socket
	WaiterCommandConnection      string
	WaiterEventReceive           *zmq4.Socket
	WaiterEventReceiveConnection string
	Identity                     string
}

func NewWaiterCommandRoutine(identity, waiterCommandConnection string) (waiterCommandRoutine *WaiterCommandRoutine) {
	waiterCommandRoutine = new(WaiterCommandRoutine)

	waiterCommandRoutine.Identity = identity
	waiterCommandRoutine.WaiterCommandConnection = waiterCommandConnection

	waiterCommandRoutine.Context, _ = zmq4.NewContext()
	waiterCommandRoutine.WaiterCommandReceive, _ = waiterCommandRoutine.Context.NewSocket(zmq4.DEALER)
	waiterCommandRoutine.WaiterCommandReceive.SetIdentity(waiterCommandRoutine.Identity)
	waiterCommandRoutine.WaiterCommandReceive.Connect(waiterCommandRoutine.WaiterCommandConnection)
	fmt.Println("workerCommandReceive connect : " + waiterCommandConnection)

	return
}

func (r WaiterCommandRoutine) WaitCommand(uuid string) (commandMessage message.CommandMessage) {
	commandMessageWait := message.NewCommandMessageWait(uuid, constant.COMMAND_MESSAGE)
	commandMessageWait.SendWith(r.WaiterCommandReceive)
	for {
		command, err := r.WaiterCommandReceive.RecvMessageBytes(0)
		if err != nil {
			panic(err)
		}
		commandMessage, _ = message.DecodeCommandMessage(command[1])
		break
	}
	return
}

func (r WaiterCommandRoutine) WaitCommandReply(uuid string) (commandMessageReply message.CommandMessageReply) {
	commandMessageWait := message.NewCommandMessageWait(uuid, constant.COMMAND_MESSAGE_REPLY)
	commandMessageWait.SendWith(r.WaiterCommandReceive)
	for {
		command, err := r.WaiterCommandReceive.RecvMessageBytes(0)
		if err != nil {
			panic(err)
		}
		commandMessageReply, _ = message.DecodeCommandMessageReply(command[1])
		break
	}
	return
}
