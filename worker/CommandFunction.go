package worker

import (
	"fmt"
)

//TODO CHANNEL ?
type CommandFunction interface {
	//executeCommand(command [][]byte, commandStates *CommandStates, referenceState *ReferenceState) string
	executeCommand() string
}

type CommandPrint struct {
	print string
}

func (cp CommandPrint) new() {

}

func (cp CommandPrint) executeCommand() string {
	fmt.Print(cp.print)
	return print
}
