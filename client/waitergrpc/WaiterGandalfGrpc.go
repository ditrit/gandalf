package waitergrpc

import (
	pb "gandalf-go/grpc"
)

type WaiterGandalfGrpc struct {
	Identity                    string
	WaiterCommandGrpcConnection string
	WaiterEventGrpcConnection   string
	WaiterCommandGrpc           *WaiterCommandGrpc
	WaiterEventGrpc             *WaiterEventGrpc
}

func NewWaiterGandalfGrpc(identity, waiterCommandGrpcConnection, waiterEventGrpcConnection string) (waiterGandalfGrpc *WaiterGandalfGrpc) {
	waiterGandalfGrpc = new(WaiterGandalfGrpc)

	waiterGandalfGrpc.Identity = identity
	waiterGandalfGrpc.WaiterCommandGrpcConnection = waiterCommandGrpcConnection
	waiterGandalfGrpc.WaiterEventGrpcConnection = waiterEventGrpcConnection

	waiterGandalfGrpc.WaiterCommandGrpc = NewWaiterCommandGrpc(waiterGandalfGrpc.Identity, waiterGandalfGrpc.WaiterCommandGrpcConnection)
	waiterGandalfGrpc.WaiterEventGrpc = NewWaiterEventGrpc(waiterGandalfGrpc.Identity, waiterGandalfGrpc.WaiterEventGrpcConnection)

	return
}

func (wg WaiterGandalfGrpc) WaitEvent(event, topic string) (eventMessage *pb.EventMessage) {
	eventMessage = wg.WaiterEventGrpc.WaitEvent(event, topic)
	return
}

func (wg WaiterGandalfGrpc) WaitCommand(uuid string) (commandMessage *pb.CommandMessage) {
	return wg.WaiterCommandGrpc.WaitCommand(uuid)
}

func (wg WaiterGandalfGrpc) WaitCommandReply(uuid string) (commandMessageReply *pb.CommandMessageReply) {
	return wg.WaiterCommandGrpc.WaitCommandReply(uuid)
}
