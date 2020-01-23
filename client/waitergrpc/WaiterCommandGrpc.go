package waitergrpc

import (
	"context"
	"fmt"

	pb "gandalf-go/grpc"
	"gandalf-go/message"

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
	conn, err := grpc.Dial(waiterCommandGrpc.WaiterCommandGrpcConnection, grpc.WithInsecure())
	if err != nil {
	}
	waiterCommandGrpc.client = pb.NewConnectorCommandClient(conn)
	return
}

func (r WaiterCommandGrpc) WaitCommand(command string) (commandMessage message.CommandMessage) {
	commandMessageWait := new(pb.CommandMessageWait)
	commandMessageWait.WorkerSource = r.Identity
	commandMessageWait.Value = command
	commandMessageGrpc, _ := r.client.WaitCommandMessage(context.Background(), commandMessageWait)
	fmt.Println(commandMessageGrpc)
	for {
		if commandMessageGrpc != nil {
			commandMessage = message.CommandMessageFromGrpc(commandMessageGrpc)
			break
		}
	}
	return

}

func (r WaiterCommandGrpc) WaitCommandReply(uuid string) (commandMessageReply message.CommandMessageReply) {
	commandMessageWait := new(pb.CommandMessageWait)
	commandMessageWait.WorkerSource = r.Identity
	commandMessageWait.Value = uuid
	commandMessageReplyGrpc, _ := r.client.WaitCommandMessageReply(context.Background(), commandMessageWait)
	for {
		if commandMessageReplyGrpc != nil {
			commandMessageReply = message.CommandMessageReplyFromGrpc(commandMessageReplyGrpc)
			break
		}
	}
	return

}
