package routine

import (
	"gandalf-go/message"
)

type EventRoutine interface {
	//executeEvent(event [][]byte)
	ExecuteEvent(eventMessage message.EventMessage) string
}
/* 
type EventPrint struct {
	print string
}

func (ep EventPrint) ExecuteEvent(eventMessage message.EventMessage) {
	fmt.Print("%s", "EVENT")
}
 */