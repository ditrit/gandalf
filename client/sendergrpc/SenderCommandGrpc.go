package sendergrpc

type SenderCommandGrpc struct {
	SenderCommandGrpcConnection string
	Identity   string
	client 	   connectorcommand.ConnectorCommandClient
}

func NewSenderCommandGrpc(identity, senderCommandGrpcConnection string) (senderCommandGrpc *SenderCommandGrpc) {
	senderCommandGrpc = new(SenderCommandGrpc)
	senderCommandGrpc.Identity = identity
	senderCommandGrpc.SenderCommandGrpcConnection = senderCommandGrpcConnection

	conn, err := grpc.Dial(*senderCommandGrpc.SenderCommandGrpcConnection)
	if err != nil {
		...
	}
	defer conn.Close()
	client := connector.NewConnectorCommandClient(conn)
	return
}

func (r SenderCommandGrpc) SendCommand(context, timeout, uuid, connectorType, commandType, command, payload string) (*pb.CommandMessageUUID, error) {
	commandMessage = new(pb.CommandMessage)
	commandMessage.Context = context
	commandMessage.Timeout = timeout
	commandMessage.Uuid = uuid
	commandMessage.ConnectorType = connectorType
	commandMessage.CommandType = command
	commandMessage.Command = command
	commandMessage.Payload = payload

	return r.client.SendCommandMessage(context.Background(), commandMessage)
}

func (r SenderCommandGrpc) SendCommandReply(commandMessage message.CommandMessage, reply, payload string) (*pb.Empty, error) {
	commandMessageReply = new(pb.CommandMessageReply)
	SourceAggregator = commandMessage.SourceAggregator   
	SourceConnector = commandMessage.SourceConnector
	SourceWorker = commandMessage.SourceWorker 
	DestinationAggregator = commandMessage.DestinationAggregator
	DestinationConnector = commandMessage.DestinationConnector 
	DestinationWorker = commandMessage.DestinationWorker    
	Tenant = commandMessage.Tenant               
	Token = commandMessage.Token                
	Context = commandMessage.Context              
	Timeout = commandMessage.Timeout             
	Timestamp = commandMessage.Timestamp            
	Uuid = commandMessage.Uuid                  
	Reply = reply                 
	Payload = payload                 
	return r.client.SendCommandMessageReply(context.Background(), commandMessageReply)

}

