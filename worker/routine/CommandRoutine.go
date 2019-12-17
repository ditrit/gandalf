package routine

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

func (cp CommandPrint) executeCommand() {
	fmt.Print("%s", "COMMAND")
}
