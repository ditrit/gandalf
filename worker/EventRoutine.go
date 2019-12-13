package worker

import "fmt"

type EventRoutine interface {
	//executeEvent(event [][]byte)
	executeEvent()
}

type EventPrint struct {
	print string
}

func (ep EventPrint) executeEvent() err error {
	fmt.Print(ep.print)
}
