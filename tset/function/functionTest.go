package function

import (
	"fmt"
	"gandalf-go/client"
	"gandalf-go/message"
)

type FunctionTest struct {
	Replys              chan message.CommandMessageReply
	ClientGandalf       *client.ClientGandalf
}

func NewFunctionTest(clientGandalf client.ClientGandalf, Replys chan message.CommandMessageReply) {
	fmt.Print("%s", "COMMAND")
}

func (ft FunctionTest) ExecuteCommand(commandMessage message.CommandMessage, Replys chan message.CommandMessageReply) {
	fmt.Print("%s", "COMMAND")
	fmt.Print("%s", commandMessage)
}
