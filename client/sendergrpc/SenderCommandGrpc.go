package sendergrpc

import (
	"context"
	pb "gandalf-go/grpc"
	"gandalf-go/message"

	"google.golang.org/grpc"
)

type SenderCommandGrpc struct {
	SenderCommandGrpcConnection string
	Identity                    string
	client                      pb.ConnectorCommandClient
}

func NewSenderCommandGrpc(identity, senderCommandGrpcConnection string) (senderCommandGrpc *SenderCommandGrpc) {
	senderCommandGrpc = new(SenderCommandGrpc)
	senderCommandGrpc.Identity = identity
	senderCommandGrpc.SenderCommandGrpcConnection = senderCommandGrpcConnection

	conn, err := grpc.Dial(senderCommandGrpc.SenderCommandGrpcConnection)
	if err != nil {
	}
	defer conn.Close()
	senderCommandGrpc.client = pb.NewConnectorCommandClient(conn)
	return
}

func (r SenderCommandGrpc) SendCommand(context, timeout, uuid, connectorType, commandType, command, payload string) *pb.CommandMessageUUID {
	commandMessage := new(pb.CommandMessage)
	commandMessage.Context = context
	commandMessage.Timeout = timeout
	commandMessage.Uuid = uuid
	commandMessage.ConnectorType = connectorType
	commandMessage.CommandType = command
	commandMessage.Command = command
	commandMessage.Payload = payload

	CommandMessageUUID, _ := r.client.SendCommandMessage(context.Background(), commandMessage)
	return CommandMessageUUID

}

func (r SenderCommandGrpc) SendCommandReply(commandMessage message.CommandMessage, reply, payload string) *pb.Empty {
	commandMessageReply := new(pb.CommandMessageReply)
	commandMessageReply.SourceAggregator = commandMessage.SourceAggregator
	commandMessageReply.SourceConnector = commandMessage.SourceConnector
	commandMessageReply.SourceWorker = commandMessage.SourceWorker
	commandMessageReply.DestinationAggregator = commandMessage.DestinationAggregator
	commandMessageReply.DestinationConnector = commandMessage.DestinationConnector
	commandMessageReply.DestinationWorker = commandMessage.DestinationWorker
	commandMessageReply.Tenant = commandMessage.Tenant
	commandMessageReply.Token = commandMessage.Token
	commandMessageReply.Context = commandMessage.Context
	commandMessageReply.Timeout = commandMessage.Timeout
	commandMessageReply.Timestamp = commandMessage.Timestamp
	commandMessageReply.Uuid = commandMessage.Uuid
	commandMessageReply.Reply = reply
	commandMessageReply.Payload = payload
	empty, _ := r.client.SendCommandMessageReply(context.Background(), commandMessageReply)
	return empty
}
