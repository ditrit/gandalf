package function

import (
	"fmt"
	"gandalf-go/message"
)

type FunctionTest struct {
}

func (ft FunctionTest) ExecuteCommand(commandMessage message.CommandMessage, replys chan message.CommandMessageReply) {
	fmt.Println("COMMAND")
	fmt.Println(commandMessage)
}

func (ft FunctionTest) ExecuteEvent(eventMessage message.EventMessage) {
	fmt.Println("EVENT")
	fmt.Println(eventMessage)
}
