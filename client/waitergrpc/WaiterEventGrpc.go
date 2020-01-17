package waitergrpc

import (
	"context"

	pb "gandalf-go/grpc"

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

	conn, err := grpc.Dial(waiterEventGrpc.WaiterEventGrpcConnection)
	if err != nil {
	}
	defer conn.Close()
	waiterEventGrpc.client = pb.NewConnectorEventClient(conn)
	return
}

func (r WaiterEventGrpc) WaitEvent(event, topic string) (eventMessage *pb.EventMessage) {
	eventMessageWait := new(pb.EventMessageWait)
	eventMessageWait.WorkerSource = r.Identity
	eventMessageWait.Topic = topic
	eventMessageWait.Event = event
	eventMessage, _ = r.client.WaitEventMessage(context.Background(), eventMessageWait)
	return
}
