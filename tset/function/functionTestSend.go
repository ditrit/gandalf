package function

import (
	"fmt"
	"gandalf-go/client"
	"gandalf-go/message"
)

type FunctionTestSend struct {
	Replys        chan message.CommandMessageReply
	ClientGandalf *client.ClientGandalfGrpc
}

func NewFunctionTestSend(clientGandalf client.ClientGandalfGrpc, replys chan message.CommandMessageReply) {
	fmt.Println("COMMAND")
}

func (fts FunctionTestSend) ExecuteCommand() {
	fmt.Println("Send")

	// Uncompilable unable to restore
	// fts.ClientGandalf.SenderGandalfGrpc.SenderCommandRoutine.SendCommandSync("context", "timeout", "uuid", "connectorType", "commandType", "send", "payload")
	fmt.Println("End Send")
}
