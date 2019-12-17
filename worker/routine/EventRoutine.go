package routine

import (
	"fmt"
	"gandalfgo/message"
)

type EventRoutine interface {
	//executeEvent(event [][]byte)
	ExecuteEvent(eventMessage message.EventMessage)
}

type EventPrint struct {
	print string
}

func (ep EventPrint) ExecuteEvent(eventMessage message.EventMessage) {
	fmt.Print("%s", "EVENT")
}
