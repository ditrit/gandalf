package waitergrpc

import "context"

type WaiterCommandGrpc struct {
	WaiterCommandGrpcConnection string
	Identity                    string
	client 	   connectorcommand.ConnectorCommandClient
}

func NewWaiterCommandGrpc(identity, waiterCommandGrpcConnection string) (waiterCommandGrpc *WaiterCommandGrpc) {
	waiterCommandGrpc = new(WaiterCommandRoutine)

	waiterCommandGrpc.Identity = identity
	waiterCommandGrpc.WaiterCommandGrpcConnection = waiterCommandGrpcConnection

	conn, err := grpc.Dial(*waiterCommandGrpc.WaiterCommandGrpcConnection)
	if err != nil {
		...
	}
	defer conn.Close()
	client := connector.NewConnectorCommandClient(conn)
	return
}

func (r WaiterCommandGrpc) WaitCommand(command string) (*CommandMessage, error) {
	commandMessageWait = new(pb.CommandMessageWait)
	commandMessageWait.WorkerSource = r.Identity
	commandMessageWait.Value = command
	return r.client.WaitCommandMessage(context.Background(), commandMessageWait)

}

func (r WaiterCommandGrpc) WaitCommandReply(uuid string) (*CommandMessageReply, error) {
	commandMessageUUID = new(pb.CommandMessageUUID)
	commandMessageUUID.Uuid = uuid
	return r.client.WaitCommandMessageReply(context.Background(), commandMessageUUID)

}
