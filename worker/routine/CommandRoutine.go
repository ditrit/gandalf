package routine

import (
	"fmt"
	"gandalf-go/message"
)

//TODO CHANNEL ?
type CommandRoutine interface {
	//executeCommand(command [][]byte, commandStates *CommandStates, referenceState *ReferenceState) string
	ExecuteCommand(commandMessage message.CommandMessage, Replys chan message.CommandMessageReply) string
}

type CommandPrint struct {
	print string
}

func (cp CommandPrint) New() {

}

func (cp CommandPrint) ExecuteCommand(commandMessage message.CommandMessage, Replys chan message.CommandMessageReply) {
	fmt.Print("%s", "COMMAND")
	fmt.Print("%s", commandMessage)
}
