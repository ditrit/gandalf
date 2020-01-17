package grpc

type SenderEventGrpc struct {
	serverAddr string
	Identity   string
	client 	   connectorevent.ConnectorEventClient
}

func NewSenderEventGrpc(identity, serverAddr string) (senderEventGrpc *SenderEventGrpc) {
	senderEventGrpc = new(SenderEventGrpc)
	senderEventGrpc.Identity = identity
	senderEventGrpc.serverAddr = serverAddr

	conn, err := grpc.Dial(*serverAddr)
	if err != nil {
		...
	}
	defer conn.Close()
	client := connector.NewConnectorEventClient(conn)
	return
}

func (r NewSenderEventGrpc) SendEventMessage(eventMessage *EventMessage) (*Empty, error) {
	return r.client.SendEventMessage(context.Background(), eventMessage)
}

func (r NewSenderEventGrpc) WaitEventMessage(eventMessageRequest *EventMessageRequest) (*EventMessage, error) {
	return r.client.WaitEventMessage(context.Background(), eventMessageRequest)
}