package waitergrpc

import (
	"context"
)

type WaiterEventGrpc struct {
	WaiterEventGrpcConnection string
	Identity                  string
	client 	   connectorevent.ConnectorEventClient
}

func NewWaiterEventGrpc(identity, waiterEventGrpcConnection string) (waiterEventGrpc *WaiterEventGrpc) {
	waiterEventRoutine = new(WaiterEventRoutine)

	waiterEventRoutine.Identity = identity
	waiterEventRoutine.WaiterEventGrpcConnection = waiterEventGrpcConnection

	conn, err := grpc.Dial(*waiterEventRoutine.WaiterEventGrpcConnection)
	if err != nil {
		...
	}
	defer conn.Close()
	client := connector.NewConnectorEventClient(conn)
	return
}

func (r WaiterEventGrpc) WaitEvent(event, topic string) (*pb.EventMessage, error) {
	eventMessageWait = new(pb.EventMessageWait)
	eventMessageWait.WorkerSource = r.Identity
	eventMessageWait.Topic = topic
	eventMessageWait.Event = event
	return r.client.WaitEventMessage(context.Background(), eventMessageRequest)
}
