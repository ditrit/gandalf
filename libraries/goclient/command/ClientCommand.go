//Package command :
//File ClientCommand.go
package command

import (
	"context"
	"fmt"
	"time"

	pb "github.com/ditrit/gandalf/libraries/gogrpc"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

//ClientCommand :
type ClientCommand struct {
	ClientCommandConnection string
	Identity                string
	client                  pb.ConnectorCommandClient
}

//NewClientCommand :
func NewClientCommand(identity, clientCommandConnection string) (clientCommand *ClientCommand) {
	clientCommand = new(ClientCommand)
	clientCommand.Identity = identity
	clientCommand.ClientCommandConnection = clientCommandConnection
	conn, _ := grpc.Dial(clientCommand.ClientCommandConnection, grpc.WithInsecure())
	// if err != nil {
	// 	// TODO implement erreur
	// }
	clientCommand.client = pb.NewConnectorCommandClient(conn)
	fmt.Println("clientCommand connect : " + clientCommand.ClientCommandConnection)

	return
}

//SendCommand :
func (cc ClientCommand) SendCommand(connectorType, command, timeout, payload string) (commandMessageUUID *pb.CommandMessageUUID) {
	commandMessage := new(pb.CommandMessage)
	commandMessage.Timeout = timeout
	commandMessage.UUID = uuid.New().String()
	commandMessage.ConnectorType = connectorType
	commandMessage.Command = command
	commandMessage.Payload = payload

	commandMessageUUID, _ = cc.client.SendCommandMessage(context.Background(), commandMessage)

	return commandMessageUUID
}

//WaitCommand :
func (cc ClientCommand) WaitCommand(command, idIterator string, major int64) *pb.CommandMessage {
	commandMessageWait := new(pb.CommandMessageWait)
	commandMessageWait.WorkerSource = cc.Identity
	commandMessageWait.Value = command
	commandMessageWait.IteratorId = idIterator
	commandMessageWait.Major = major
	commandMessage, _ := cc.client.WaitCommandMessage(context.Background(), commandMessageWait)
	fmt.Println(commandMessage)

	for commandMessage == nil {
		time.Sleep(time.Duration(1) * time.Second)
	}

	return commandMessage
}

//CreateIteratorCommand :
func (cc ClientCommand) CreateIteratorCommand() *pb.IteratorMessage {
	iteratorMessage, _ := cc.client.CreateIteratorCommand(context.Background(), new(pb.Empty))
	return iteratorMessage
}
