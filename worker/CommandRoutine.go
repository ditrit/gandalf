package worker

import (
	"fmt"
)

//TODO CHANNEL ?
type CommandRoutine interface {
	//executeCommand(command [][]byte, commandStates *CommandStates, referenceState *ReferenceState) string
	executeCommand() string
}

type CommandPrint struct {
	print string
}

func (cp CommandPrint) New() {

}

func (cp CommandPrint) executeCommand() (result string, err error) {
	fmt.Print(cp.print)
	return print
}
