package waitergrpc

import (
	"context"
	"fmt"

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

	conn, err := grpc.Dial(waiterEventGrpc.WaiterEventGrpcConnection, grpc.WithInsecure())
	if err != nil {
		fmt.Println("ERROR EVENT")
	}
	fmt.Println("CONNN WAITER EVENT")
	fmt.Println(conn)
	waiterEventGrpc.client = pb.NewConnectorEventClient(conn)
	fmt.Println("waiterEventGrpc connect : " + waiterEventGrpc.WaiterEventGrpcConnection)
	return
}

func (r WaiterEventGrpc) WaitEvent(event, topic string) (eventMessage message.EventMessage) {
	eventMessageWait := new(pb.EventMessageWait)
	eventMessageWait.WorkerSource = r.Identity
	eventMessageWait.Topic = topic
	eventMessageWait.Event = event
	eventMessageGrpc, _ := r.client.WaitEventMessage(context.Background(), eventMessageWait)
	eventMessage = message.EventMessageFromGrpc(eventMessageGrpc)
	return
}
