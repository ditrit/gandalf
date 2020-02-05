package waitergrpc

import (
	"context"
	"time"

	pb "gandalf-go/grpc"
	"gandalf-go/message"

	"google.golang.org/grpc"
)

type WaiterEventGrpc struct {
	WaiterEventGrpcConnection string
	Identity                  string
	client                    pb.ConnectorEventClient
}

func NewWaiterEventGrpc(identity, waiterEventGrpcConnection string) (waiterEventGrpc *WaiterEventGrpc) {
	waiterEventGrpc = new(WaiterEventGrpc)

	waiterEventGrpc.Identity = identity
	waiterEventGrpc.WaiterEventGrpcConnection = waiterEventGrpcConnection

	conn, _ := grpc.Dial(waiterEventGrpc.WaiterEventGrpcConnection, grpc.WithInsecure())
	// if err != nil {
	// 	// TODO Handle error
	// }
	waiterEventGrpc.client = pb.NewConnectorEventClient(conn)

	return
}

func (r WaiterEventGrpc) WaitEvent(event, topic string) (eventMessage message.EventMessage) {
	eventMessageWait := new(pb.EventMessageWait)
	eventMessageWait.WorkerSource = r.Identity
	eventMessageWait.Topic = topic
	eventMessageWait.Event = event
	eventMessageGrpc, _ := r.client.WaitEventMessage(context.Background(), eventMessageWait)

	for eventMessageGrpc == nil {
		time.Sleep(time.Duration(1) * time.Millisecond)
	}

	return message.EventMessageFromGrpc(eventMessageGrpc)
}
