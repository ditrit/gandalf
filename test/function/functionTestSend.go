//Package function :
//File functionTestSend.go
package function

import (
	"fmt"
	"gandalf-go/commons/client"
	"gandalf-go/commons/message"
)

//FunctionTestSend :
type FunctionTestSend struct {
	Replys        chan message.CommandMessageReply
	ClientGandalf *client.ClientGandalfGrpc
}

//NewFunctionTestSend :
func NewFunctionTestSend(clientGandalf client.ClientGandalfGrpc, replys chan message.CommandMessageReply) {
	fmt.Println("COMMAND")
}

//ExecuteCommand :
func (fts FunctionTestSend) ExecuteCommand() {
	fmt.Println("Send")

	// Uncompilable unable to restore
	// fts.ClientGandalf.SenderGandalfGrpc.SenderCommandRoutine.SendCommandSync("context", "timeout", "uuid", "connectorType", "commandType", "send", "payload")
	fmt.Println("End Send")
}
