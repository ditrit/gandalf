package routine

import "fmt"

type EventRoutine interface {
	//executeEvent(event [][]byte)
	executeEvent()
}

type EventPrint struct {
	print string
}

func (ep EventPrint) executeEvent() {
	fmt.Print("%s", "EVENT")
}
