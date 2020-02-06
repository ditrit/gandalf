//Package function :
//File functionTest.go
package function

import (
	"fmt"
	"gandalf-go/message"
)

//FunctionTest :
type FunctionTest struct {
}

//ExecuteCommand :
func (ft FunctionTest) ExecuteCommand(commandMessage message.CommandMessage, replys chan message.CommandMessageReply) {
	fmt.Println("COMMAND")
	fmt.Println(commandMessage)
}

//ExecuteEvent :
func (ft FunctionTest) ExecuteEvent(eventMessage message.EventMessage) {
	fmt.Println("EVENT")
	fmt.Println(eventMessage)
}
