//Package client :
//File ClientGrpc.go
package client

import (
	pb "github.com/ditrit/gandalf/libraries/gogrpc"

	"github.com/ditrit/gandalf/libraries/goclient/command"
	"github.com/ditrit/gandalf/libraries/goclient/event"

	"github.com/ditrit/gandalf/libraries/goclient/base"
)

//ClientGrpc :
type ClientGrpc struct {
	Identity          string
	ClientConnection  string
	ClientBase        *base.ClientBase
	ClientCommand     *command.ClientCommand
	ClientEvent       *event.ClientEvent
	ClientStopChannel chan int
}

//NewClientGandalf
func NewClientGrpc(identity, clientConnection string) (clientGrpc *ClientGrpc) {
	clientGrpc = new(ClientGrpc)
	clientGrpc.ClientStopChannel = make(chan int)

	clientGrpc.Identity = identity
	clientGrpc.ClientConnection = clientConnection

	clientGrpc.ClientBase = base.NewClientBase(clientGrpc.Identity, clientGrpc.ClientConnection)
	clientGrpc.ClientCommand = command.NewClientCommand(clientGrpc.Identity, clientGrpc.ClientConnection)
	clientGrpc.ClientEvent = event.NewClientEvent(clientGrpc.Identity, clientGrpc.ClientConnection)

	return
}

//SendCommandList
func (cg ClientGrpc) SendCommandList(major, minor int64, commands []string) *pb.Validate {
	return cg.ClientBase.SendCommandList(major, minor, commands)
}

//SendStop
func (cg ClientGrpc) SendStop(major, minor int64) *pb.Validate {
	return cg.ClientBase.SendStop(major, minor)
}

//SendCommand
func (cg ClientGrpc) SendCommand(connectorType, command, timeout, payload string) *pb.CommandMessageUUID {
	return cg.ClientCommand.SendCommand(connectorType, command, timeout, payload)
}

//SendEvent
func (cg ClientGrpc) SendEvent(topic, event, referenceUUID, timeout, payload string) *pb.Empty {
	return cg.ClientEvent.SendEvent(topic, event, referenceUUID, timeout, payload)
}

//WaitCommand
func (cg ClientGrpc) WaitCommand(command, idIterator string, version int64) (commandMessage *pb.CommandMessage) {
	return cg.ClientCommand.WaitCommand(command, idIterator, version)
}

//WaitEvent
func (cg ClientGrpc) WaitEvent(topic, event, referenceUUID, idIterator string) (eventMessage *pb.EventMessage) {
	return cg.ClientEvent.WaitEvent(topic, event, referenceUUID, idIterator)
}

//WaitEvent
func (cg ClientGrpc) WaitTopic(topic, referenceUUID, idIterator string) (eventMessage *pb.EventMessage) {
	return cg.ClientEvent.WaitTopic(topic, referenceUUID, idIterator)
}

//CreateIteratorCommand
func (cg ClientGrpc) CreateIteratorCommand() (iteratorMessage *pb.IteratorMessage) {
	return cg.ClientCommand.CreateIteratorCommand()
}

//CreateIteratorEvent
func (cg ClientGrpc) CreateIteratorEvent() (iteratorMessage *pb.IteratorMessage) {
	return cg.ClientEvent.CreateIteratorEvent()
}
