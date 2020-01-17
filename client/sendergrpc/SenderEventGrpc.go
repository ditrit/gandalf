package sendergrpc

type SenderEventGrpc struct {
	SenderEventGrpcConnection string
	Identity   string
	client 	   connectorevent.ConnectorEventClient
}

func NewSenderEventGrpc(identity, senderEventGrpcConnection string) (senderEventGrpc *SenderEventGrpc) {
	senderEventGrpc = new(SenderEventGrpc)
	senderEventGrpc.Identity = identity
	senderEventGrpc.SenderEventGrpcConnection = senderEventGrpcConnection

	conn, err := grpc.Dial(*senderEventGrpc.SenderEventGrpcConnection)
	if err != nil {
		...
	}
	defer conn.Close()
	client := connector.NewConnectorEventClient(conn)
	return
}

func (r SenderEventGrpc) SendEvent(topic, timeout, uuid, event, payload string) (*pb.Empty, error) {
	eventMessage = new(pb.EventMessage)
	eventMessage.Topic = topic
	eventMessage.Timeout = timeout
	eventMessage.Uuid = uuid
	eventMessage.Event = event
	eventMessage.Payload = payload

	return r.client.SendEventMessage(context.Background(), eventMessage)
}

