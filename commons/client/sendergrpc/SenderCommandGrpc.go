//Package sendergrpc :
//File SenderCommandGrpc.go
package sendergrpc

import (
	"context"
	"fmt"
	pb "gandalf-go/commons/grpc"
	"gandalf-go/commons/message"

	"google.golang.org/grpc"
)

//SenderCommandGrpc :
type SenderCommandGrpc struct {
	SenderCommandGrpcConnection string
	Identity                    string
	client                      pb.ConnectorCommandClient
}

//NewSenderCommandGrpc :
func NewSenderCommandGrpc(
	identity string,
	senderCommandGrpcConnection string) (senderCommandGrpc *SenderCommandGrpc) {
	senderCommandGrpc = new(SenderCommandGrpc)
	senderCommandGrpc.Identity = identity
	senderCommandGrpc.SenderCommandGrpcConnection = senderCommandGrpcConnection
	conn, _ := grpc.Dial(senderCommandGrpc.SenderCommandGrpcConnection, grpc.WithInsecure())
	// if err != nil {
	// 	// TODO implement erreur
	// }
	senderCommandGrpc.client = pb.NewConnectorCommandClient(conn)
	fmt.Println("senderCommandGrpc connect : " + senderCommandGrpc.SenderCommandGrpcConnection)

	return
}

//SendCommand :
func (r SenderCommandGrpc) SendCommand(
	contextCommand string,
	timeout string,
	uuid string,
	connectorType string,
	commandType string,
	command string,
	payload string) (commandMessageUUID message.CommandMessageUUID) {
	commandMessage := new(pb.CommandMessage)
	commandMessage.Context = contextCommand
	commandMessage.Timeout = timeout
	commandMessage.UUID = uuid
	commandMessage.ConnectorType = connectorType
	commandMessage.CommandType = command
	commandMessage.Command = command
	commandMessage.Payload = payload

	CommandMessageUUIDGrpc, _ := r.client.SendCommandMessage(context.Background(), commandMessage)
	commandMessageUUID = message.CommandMessageUUIDFromGrpc(CommandMessageUUIDGrpc)

	return
}

//SendCommandReply :
func (r SenderCommandGrpc) SendCommandReply(
	commandMessage message.CommandMessage,
	reply string,
	payload string) *pb.Empty {
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
	commandMessageReply.UUID = commandMessage.UUID
	commandMessageReply.Reply = reply
	commandMessageReply.Payload = payload

	empty, _ := r.client.SendCommandMessageReply(context.Background(), commandMessageReply)

	return empty
}
