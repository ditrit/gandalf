package grpc

type SenderCommandGrpc struct {
	serverAddr string
	Identity   string
	client 	   connectorcommand.ConnectorCommandClient
}

func NewSenderCommandGrpc(identity, serverAddr string) (senderCommandGrpc *SenderCommandGrpc) {
	senderCommandGrpc = new(SenderCommandGrpc)
	senderCommandGrpc.Identity = identity
	senderCommandGrpc.serverAddr = serverAddr

	conn, err := grpc.Dial(*serverAddr)
	if err != nil {
		...
	}
	defer conn.Close()
	client := connector.NewConnectorCommandClient(conn)
	return
}

func (r SenderCommandGrpc) SendCommandMessage(commandMessage *CommandMessage) (*CommandMessageUUID, error) {
	return r.client.SendCommandMessage(context.Background(), commandMessage)
}

func (r SenderCommandGrpc) SendCommandMessageReply(commandMessageReply *CommandMessageReply) (*Empty, error) {
	return r.client.SendCommandMessageReply(context.Background(), commandMessageReply)

}

func (r SenderCommandGrpc) WaitCommandMessage(commandMessageRequest *CommandMessageRequest) (*CommandMessage, error) {
	return r.client.WaitCommandMessage(context.Background(), commandMessageRequest)

}

func (r SenderCommandGrpc) WaitCommandMessageReply(commandMessageUUID *CommandMessageUUID) (*CommandMessageReply, error) {
	return r.client.WaitCommandMessageReply(context.Background(), commandMessageUUID)

}
