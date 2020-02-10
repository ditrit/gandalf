//Package client :
//File ClientGandalfGrpc.go
package client

import (
	"gandalf-go/commons/client/sendergrpc"
	"gandalf-go/commons/client/waitergrpc"
	pb "gandalf-go/commons/grpc"

	"gandalf-go/commons/message"
)

//ClientGandalfGrpc :
type ClientGandalfGrpc struct {
	Identity                    string
	SenderCommandGrpcConnection string
	SenderEventGrpcConnection   string
	WaiterCommandGrpcConnection string
	WaiterEventGrpcConnection   string
	SenderGandalfGrpc           *sendergrpc.SenderGandalfGrpc
	WaiterGandalfGrpc           *waitergrpc.WaiterGandalfGrpc
	ClientStopChannel           chan int
}

//NewClientGandalfGrpc :
func NewClientGandalfGrpc(
	identity string,
	senderCommandGrpcConnection string,
	senderEventGrpcConnection string,
	waiterCommandGrpcConnection string,
	waiterEventGrpcConnection string) (clientGandalfGrpc *ClientGandalfGrpc) {
	clientGandalfGrpc = new(ClientGandalfGrpc)
	clientGandalfGrpc.ClientStopChannel = make(chan int)

	clientGandalfGrpc.Identity = identity
	clientGandalfGrpc.SenderCommandGrpcConnection = senderCommandGrpcConnection
	clientGandalfGrpc.SenderEventGrpcConnection = senderEventGrpcConnection
	clientGandalfGrpc.WaiterCommandGrpcConnection = waiterCommandGrpcConnection
	clientGandalfGrpc.WaiterEventGrpcConnection = waiterEventGrpcConnection

	clientGandalfGrpc.SenderGandalfGrpc = sendergrpc.NewSenderGandalfGrpc(
		clientGandalfGrpc.Identity,
		clientGandalfGrpc.SenderCommandGrpcConnection,
		clientGandalfGrpc.SenderEventGrpcConnection)
	clientGandalfGrpc.WaiterGandalfGrpc = waitergrpc.NewWaiterGandalfGrpc(
		clientGandalfGrpc.Identity,
		clientGandalfGrpc.WaiterCommandGrpcConnection,
		clientGandalfGrpc.WaiterEventGrpcConnection)

	return
}

//SendCommand :
//nolint: lll
func (cg ClientGandalfGrpc) SendCommand(context, timeout, uuid, connectorType, commandType, command, payload string) message.CommandMessageUUID {
	return cg.SenderGandalfGrpc.SendCommand(context, timeout, uuid, connectorType, commandType, command, payload)
}

//SendCommandReply :
//nolint: lll
func (cg ClientGandalfGrpc) SendCommandReply(commandMessage message.CommandMessage, reply, payload string) *pb.Empty {
	return cg.SenderGandalfGrpc.SendCommandReply(commandMessage, reply, payload)
}

//SendEvent :
func (cg ClientGandalfGrpc) SendEvent(topic, timeout, uuid, event, payload string) *pb.Empty {
	return cg.SenderGandalfGrpc.SendEvent(topic, timeout, uuid, event, payload)
}

//WaitCommand :
func (cg ClientGandalfGrpc) WaitCommand(command string) (commandMessage message.CommandMessage) {
	return cg.WaiterGandalfGrpc.WaitCommand(command)
}

//WaitCommandReply :
func (cg ClientGandalfGrpc) WaitCommandReply(uuid string) (commandMessageReply message.CommandMessageReply) {
	return cg.WaiterGandalfGrpc.WaitCommandReply(uuid)
}

//WaitEvent :
func (cg ClientGandalfGrpc) WaitEvent(event, topic string) (eventMessage message.EventMessage) {
	return cg.WaiterGandalfGrpc.WaitEvent(event, topic)
}
