package waiter

import (
	"fmt"
	"gandalf-go/message"

	"github.com/pebbe/zmq4"
)

type WaiterEventRoutine struct {
	Context                      *zmq4.Context
	WaiterEventReceive           *zmq4.Socket
	WaiterEventConnection        string
	WaiterEventReceiveConnection string
	Identity                     string
}

func NewWaiterEventRoutine(identity, waiterEventConnection string) (waiterEventRoutine *WaiterEventRoutine) {
	waiterEventRoutine = new(WaiterEventRoutine)

	waiterEventRoutine.Identity = identity
	waiterEventRoutine.WaiterEventConnection = waiterEventConnection

	waiterEventRoutine.Context, _ = zmq4.NewContext()
	waiterEventRoutine.WaiterEventReceive, _ = waiterEventRoutine.Context.NewSocket(zmq4.DEALER)
	waiterEventRoutine.WaiterEventReceive.SetIdentity(waiterEventRoutine.Identity)
	waiterEventRoutine.WaiterEventReceive.Connect(waiterEventRoutine.WaiterEventConnection)
	fmt.Println("workerEventReceive connect : " + waiterEventConnection)

	return
}

func (r WaiterEventRoutine) WaitEvent(event string) (eventMessage message.EventMessage) {
	eventMessageWait := message.NewEventMessageWait(event)
	eventMessageWait.SendWith(r.WaiterEventReceive)
	for {
		event, err := r.WaiterEventReceive.RecvMessageBytes(0)
		if err != nil {
			panic(err)
		}
		eventMessage, _ = message.DecodeEventMessage(event[1])
		break
	}
	return
}
