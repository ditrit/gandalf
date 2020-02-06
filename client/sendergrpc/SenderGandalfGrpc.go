//Package sendergrpc :
//File SenderGandalfGrpc.go
package sendergrpc

import (
	pb "gandalf-go/grpc"
	"gandalf-go/message"
)

//SenderGandalfGrpc :
type SenderGandalfGrpc struct {
	Identity                string
	SenderCommandConnection string
	SenderEventConnection   string
	SenderCommandGrpc       *SenderCommandGrpc
	SenderEventGrpc         *SenderEventGrpc
}

//NewSenderGandalfGrpc :
func NewSenderGandalfGrpc(identity, senderCommandConnection, senderEventConnection string) (senderGandalfGrpc *SenderGandalfGrpc) {
	senderGandalfGrpc = new(SenderGandalfGrpc)
	senderGandalfGrpc.Identity = identity
	senderGandalfGrpc.SenderCommandConnection = senderCommandConnection
	senderGandalfGrpc.SenderEventConnection = senderEventConnection
	senderGandalfGrpc.SenderCommandGrpc = NewSenderCommandGrpc(identity, senderCommandConnection)
	senderGandalfGrpc.SenderEventGrpc = NewSenderEventGrpc(identity, senderEventConnection)

	return
}

//SendEvent :
func (sg SenderGandalfGrpc) SendEvent(topic, timeout, uuid, event, payload string) *pb.Empty {
	return sg.SenderEventGrpc.SendEvent(topic, timeout, uuid, event, payload)
}

//SendCommand :
func (sg SenderGandalfGrpc) SendCommand(context, timeout, uuid, connectorType, commandType, command, payload string) message.CommandMessageUUID {
	return sg.SenderCommandGrpc.SendCommand(context, timeout, uuid, connectorType, commandType, command, payload)
}

//SendCommandReply :
func (sg SenderGandalfGrpc) SendCommandReply(commandMessage message.CommandMessage, reply, payload string) *pb.Empty {
	return sg.SenderCommandGrpc.SendCommandReply(commandMessage, reply, payload)
}
