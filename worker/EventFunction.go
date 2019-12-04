package worker

import "fmt"

type EventFunction interface {
	//executeEvent(event [][]byte)
	executeEvent()
}

type EventPrint struct {
	print string
}

func (ep EventPrint) executeEvent() {
	fmt.Print(ep.print)
}
