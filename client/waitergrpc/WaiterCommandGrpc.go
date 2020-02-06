package waitergrpc

import (
	"context"
	"fmt"
	"time"

	pb "gandalf-go/grpc"
	"gandalf-go/message"

	"google.golang.org/grpc"
)

//WaiterCommandGrpc :
type WaiterCommandGrpc struct {
	WaiterCommandGrpcConnection string
	Identity                    string
	client                      pb.ConnectorCommandClient
}

//NewWaiterCommandGrpc :
func NewWaiterCommandGrpc(identity, waiterCommandGrpcConnection string) (waiterCommandGrpc *WaiterCommandGrpc) {
	waiterCommandGrpc = new(WaiterCommandGrpc)

	waiterCommandGrpc.Identity = identity
	waiterCommandGrpc.WaiterCommandGrpcConnection = waiterCommandGrpcConnection
	conn, err := grpc.Dial(waiterCommandGrpc.WaiterCommandGrpcConnection, grpc.WithInsecure())

	if err != nil {
		//TODO : Handle this case
		return
	}

	waiterCommandGrpc.client = pb.NewConnectorCommandClient(conn)

	return
}

//WaitCommand :
func (r WaiterCommandGrpc) WaitCommand(command string) message.CommandMessage {
	commandMessageWait := new(pb.CommandMessageWait)
	commandMessageWait.WorkerSource = r.Identity
	commandMessageWait.Value = command
	commandMessageGrpc, _ := r.client.WaitCommandMessage(context.Background(), commandMessageWait)
	fmt.Println(commandMessageGrpc)

	for commandMessageGrpc == nil {
		time.Sleep(time.Duration(1) * time.Millisecond)
	}

	return message.CommandMessageFromGrpc(commandMessageGrpc)
}

//WaitCommandReply :
func (r WaiterCommandGrpc) WaitCommandReply(uuid string) message.CommandMessageReply {
	commandMessageWait := new(pb.CommandMessageWait)
	commandMessageWait.WorkerSource = r.Identity
	commandMessageWait.Value = uuid
	commandMessageReplyGrpc, _ := r.client.WaitCommandMessageReply(context.Background(), commandMessageWait)

	for commandMessageReplyGrpc == nil {
		time.Sleep(time.Duration(1) * time.Millisecond)
	}

	return message.CommandMessageReplyFromGrpc(commandMessageReplyGrpc)
}
