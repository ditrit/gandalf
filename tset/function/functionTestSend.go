package function

import (
	"fmt"
	"gandalf-go/client"
	"gandalf-go/message"
)

type FunctionTestSend struct {
	Replys        chan message.CommandMessageReply
	ClientGandalf *client.ClientGandalf
}

func NewFunctionTestSend(clientGandalf client.ClientGandalf, Replys chan message.CommandMessageReply) {
	fmt.Println("COMMAND")
}

func (fts FunctionTestSend) ExecuteCommand() {
	fmt.Println("Send")
	fts.ClientGandalf.SenderGandalf.SenderCommandRoutine.SendCommandSync("context", "timeout", "uuid", "connectorType", "commandType", "send", "payload")
}
