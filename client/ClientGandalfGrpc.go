package client

import (
	"gandalf-go/client/sendergrpc"
	"gandalf-go/client/waitergrpc"
	pb "gandalf-go/grpc"

	"gandalf-go/message"
)

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

func NewClientGandalfGrpc(identity, senderCommandGrpcConnection, senderEventGrpcConnection, waiterCommandGrpcConnection, waiterEventGrpcConnection string) (clientGandalfGrpc *ClientGandalfGrpc) {
	clientGandalfGrpc = new(ClientGandalfGrpc)
	clientGandalfGrpc.ClientStopChannel = make(chan int)

	clientGandalfGrpc.Identity = identity
	clientGandalfGrpc.SenderCommandGrpcConnection = senderCommandGrpcConnection
	clientGandalfGrpc.SenderEventGrpcConnection = senderEventGrpcConnection
	clientGandalfGrpc.WaiterCommandGrpcConnection = waiterCommandGrpcConnection
	clientGandalfGrpc.WaiterEventGrpcConnection = waiterEventGrpcConnection

	clientGandalfGrpc.SenderGandalfGrpc = sendergrpc.NewSenderGandalfGrpc(clientGandalfGrpc.Identity, clientGandalfGrpc.SenderCommandGrpcConnection, clientGandalfGrpc.SenderEventGrpcConnection)
	clientGandalfGrpc.WaiterGandalfGrpc = waitergrpc.NewWaiterGandalfGrpc(clientGandalfGrpc.Identity, clientGandalfGrpc.WaiterCommandGrpcConnection, clientGandalfGrpc.WaiterEventGrpcConnection)

	return
}

func (cg ClientGandalfGrpc) SendCommand(context, timeout, uuid, connectorType, commandType, command, payload string) message.CommandMessageUUID {
	return cg.SenderGandalfGrpc.SendCommand(context, timeout, uuid, connectorType, commandType, command, payload)
}

func (cg ClientGandalfGrpc) SendCommandReply(commandMessage message.CommandMessage, reply, payload string) *pb.Empty {
	return cg.SenderGandalfGrpc.SendCommandReply(commandMessage, reply, payload)
}

func (cg ClientGandalfGrpc) SendEvent(topic, timeout, uuid, event, payload string) *pb.Empty {
	return cg.SenderGandalfGrpc.SendEvent(topic, timeout, uuid, event, payload)
}

func (cg ClientGandalfGrpc) WaitCommand(command string) (commandMessage message.CommandMessage) {
	return cg.WaiterGandalfGrpc.WaitCommand(command)
}

func (cg ClientGandalfGrpc) WaitCommandReply(uuid string) (commandMessageReply message.CommandMessageReply) {
	return cg.WaiterGandalfGrpc.WaitCommandReply(uuid)
}

func (cg ClientGandalfGrpc) WaitEvent(event, topic string) (eventMessage message.EventMessage) {
	return cg.WaiterGandalfGrpc.WaitEvent(event, topic)
}
