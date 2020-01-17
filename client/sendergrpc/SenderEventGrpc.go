package sendergrpc

import (
	"context"
	pb "gandalf-go/grpc"

	"google.golang.org/grpc"
)

type SenderEventGrpc struct {
	SenderEventGrpcConnection string
	Identity                  string
	client                    pb.ConnectorEventClient
}

func NewSenderEventGrpc(identity, senderEventGrpcConnection string) (senderEventGrpc *SenderEventGrpc) {
	senderEventGrpc = new(SenderEventGrpc)
	senderEventGrpc.Identity = identity
	senderEventGrpc.SenderEventGrpcConnection = senderEventGrpcConnection

	conn, err := grpc.Dial(senderEventGrpc.SenderEventGrpcConnection)
	if err != nil {
	}
	defer conn.Close()
	senderEventGrpc.client = pb.NewConnectorEventClient(conn)
	return
}

func (r SenderEventGrpc) SendEvent(topic, timeout, uuid, event, payload string) *pb.Empty {
	eventMessage := new(pb.EventMessage)
	eventMessage.Topic = topic
	eventMessage.Timeout = timeout
	eventMessage.Uuid = uuid
	eventMessage.Event = event
	eventMessage.Payload = payload
	empty, _ := r.client.SendEventMessage(context.Background(), eventMessage)
	return empty
}
