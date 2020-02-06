package sendergrpc

import (
	"context"
	"fmt"
	pb "gandalf-go/grpc"

	"google.golang.org/grpc"
)

//SenderEventGrpc :
type SenderEventGrpc struct {
	SenderEventGrpcConnection string
	Identity                  string
	client                    pb.ConnectorEventClient
}

//NewSenderEventGrpc :
func NewSenderEventGrpc(identity, senderEventGrpcConnection string) (senderEventGrpc *SenderEventGrpc) {
	senderEventGrpc = new(SenderEventGrpc)
	senderEventGrpc.Identity = identity
	senderEventGrpc.SenderEventGrpcConnection = senderEventGrpcConnection

	conn, err := grpc.Dial(senderEventGrpc.SenderEventGrpcConnection, grpc.WithInsecure())

	if err != nil {
		fmt.Println("senderEventGrpc failed dial")
	}

	senderEventGrpc.client = pb.NewConnectorEventClient(conn)
	fmt.Println("senderEventGrpc connect : " + senderEventGrpc.SenderEventGrpcConnection)

	return
}

//SendEvent :
func (r SenderEventGrpc) SendEvent(topic, timeout, uuid, event, payload string) *pb.Empty {
	eventMessage := new(pb.EventMessage)
	eventMessage.Topic = topic
	eventMessage.Timeout = timeout
	eventMessage.UUID = uuid
	eventMessage.Event = event
	eventMessage.Payload = payload
	empty, _ := r.client.SendEventMessage(context.Background(), eventMessage)

	return empty
}
