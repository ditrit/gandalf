package waitergrpc

import (
	"context"

	pb "gandalf-go/grpc"

	"google.golang.org/grpc"
)

type WaiterCommandGrpc struct {
	WaiterCommandGrpcConnection string
	Identity                    string
	client                      pb.ConnectorCommandClient
}

func NewWaiterCommandGrpc(identity, waiterCommandGrpcConnection string) (waiterCommandGrpc *WaiterCommandGrpc) {
	waiterCommandGrpc = new(WaiterCommandGrpc)

	waiterCommandGrpc.Identity = identity
	waiterCommandGrpc.WaiterCommandGrpcConnection = waiterCommandGrpcConnection

	conn, err := grpc.Dial(waiterCommandGrpc.WaiterCommandGrpcConnection)
	if err != nil {

	}
	defer conn.Close()
	waiterCommandGrpc.client = pb.NewConnectorCommandClient(conn)
	return
}

func (r WaiterCommandGrpc) WaitCommand(command string) *pb.CommandMessage {
	commandMessageWait := new(pb.CommandMessageWait)
	commandMessageWait.WorkerSource = r.Identity
	commandMessageWait.Value = command
	commandMessage, _ := r.client.WaitCommandMessage(context.Background(), commandMessageWait)
	return commandMessage

}

func (r WaiterCommandGrpc) WaitCommandReply(uuid string) *pb.CommandMessageReply {
	commandMessageWait := new(pb.CommandMessageWait)
	commandMessageWait.WorkerSource = r.Identity
	commandMessageWait.Value = uuid
	commandMessageReply, _ := r.client.WaitCommandMessageReply(context.Background(), commandMessageWait)
	return commandMessageReply

}
