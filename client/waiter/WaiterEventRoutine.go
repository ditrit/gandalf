package waiter

import (
	"fmt"
	"gandalf-go/message"
	"gandalf-go/worker/routine"

	"github.com/pebbe/zmq4"
)

type WaiterEventRoutine struct {
	Context                      *zmq4.Context
	WaiterEventReceive         	 *zmq4.Socket
	WaiterEventConnection    	 string
	WaiterEventReceiveConnection string
	Identity                     string
}

func NewWaiterEventRoutine(identity, waiterEventConnection string, commandsRoutine map[string][]routine.EventRoutine) (waiterEventRoutine *WaiterEventRoutine) {
	waiterEventRoutine = new(WaiterEventRoutine)

	waiterEventRoutine.Identity = identity
	waiterEventRoutine.WaiterEventConnection = waiterEventConnection
	waiterEventRoutine.EventsRoutine = commandsRoutine

	waiterEventRoutine.Context, _ = zmq4.NewContext()
	waiterEventRoutine.WaiterEventReceive, _ = waiterEventRoutine.Context.NewSocket(zmq4.DEALER)
	waiterEventRoutine.WaiterEventReceive.SetIdentity(waiterEventRoutine.Identity)
	waiterEventRoutine.WaiterEventReceive.Connect(waiterEventRoutine.WaiterEventConnection)
	fmt.Println("workerEventReceive connect : " + waiterEventConnection)

	return
}

func (r WaiterEventRoutine) WaitEvent(event string) (eventMessage EventMessage) {
	eventMessageWait := EventMessageWait.NewEventMessageWait(uuid)
	go eventMessageWait.SendWith(r.WaiterEventReceive)
	for {
		event, err := r.WaiterEventReceive.RecvMessageBytes(0)
		if err != nil {
			panic(err)
		}
		eventMessage, _ := message.DecodeEventMessage(command[1])

		return
	}
}

