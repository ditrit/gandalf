package waiter

import (
	"gandalf-go/message"
)

type WaiterGandalf struct {
	Identity                string
	WaiterCommandConnection string
	WaiterEventConnection   string
	WaiterCommandRoutine    *WaiterCommandRoutine
	WaiterEventRoutine      *WaiterEventRoutine
	WaiterStopChannel       chan int
}

func NewWaiterGandalf(identity, waiterCommandConnection, waiterEventConnection string) (waiterGandalf *WaiterGandalf) {
	waiterGandalf = new(WaiterGandalf)

	waiterGandalf.Identity = identity
	waiterGandalf.WaiterCommandConnection = waiterCommandConnection
	waiterGandalf.WaiterEventConnection = waiterEventConnection

	waiterGandalf.WaiterCommandRoutine = NewWaiterCommandRoutine(waiterGandalf.Identity, waiterGandalf.WaiterCommandConnection)
	waiterGandalf.WaiterEventRoutine = NewWaiterEventRoutine(waiterGandalf.Identity, waiterGandalf.WaiterEventConnection)

	return
}

func (wg WaiterGandalf) WaitEvent(event string) (eventMessage message.EventMessage) {
	return wg.WaiterEventRoutine.WaitEvent(event)
}

func (wg WaiterGandalf) WaitCommand(uuid string) (commandMessage message.CommandMessage) {
	return wg.WaiterCommandRoutine.WaitCommand(uuid)
}

func (wg WaiterGandalf) WaitCommandReply(uuid string) (commandMessageReply message.CommandMessageReply) {
	return wg.WaiterCommandRoutine.WaitCommandReply(uuid)
}

func (wg WaiterGandalf) Stop() {
	wg.WaiterStopChannel <- 0
}
