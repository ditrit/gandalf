package test

import (
	"context"
	"fmt"
	pb "garcimore/grpc"
	"shoset/msg"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

//SenderCommandGrpc :
type ClientGrpcTest struct {
	clientGrpcTestConnection string
	Identity                 string
	ClientCommand            pb.ConnectorCommandClient
	ClientEvent              pb.ConnectorEventClient
}

//NewSenderCommandGrpc :
func NewClientGrpcTest(identity, clientGrpcTestConnection string) (clientGrpcTest *ClientGrpcTest) {
	clientGrpcTest = new(ClientGrpcTest)
	clientGrpcTest.Identity = identity
	clientGrpcTest.clientGrpcTestConnection = clientGrpcTestConnection
	conn, _ := grpc.Dial(clientGrpcTest.clientGrpcTestConnection, grpc.WithInsecure())
	// if err != nil {
	// 	// TODO implement erreur
	// }
	clientGrpcTest.ClientCommand = pb.NewConnectorCommandClient(conn)
	fmt.Println("clientGrpcTest connect : " + clientGrpcTest.clientGrpcTestConnection)

	clientGrpcTest.ClientEvent = pb.NewConnectorEventClient(conn)
	fmt.Println("clientGrpcTest connect : " + clientGrpcTest.clientGrpcTestConnection)

	return
}

//SendCommand :
func (r ClientGrpcTest) SendCommand(timeout, connectorType, command, payload string) *pb.CommandMessageUUID {
	commandMessage := new(pb.CommandMessage)
	commandMessage.Timeout = timeout
	commandMessage.UUID = uuid.New().String()
	commandMessage.ConnectorType = connectorType
	commandMessage.Command = command
	commandMessage.Payload = payload

	CommandMessageUUIDGrpc, _ := r.ClientCommand.SendCommandMessage(context.Background(), commandMessage)
	//commandMessageUUID = msg.CommandMessageUUIDFromGrpc(CommandMessageUUIDGrpc)

	return CommandMessageUUIDGrpc
}

//WaitCommand :
func (r ClientGrpcTest) WaitCommand(command, id string) msg.Command {
	commandMessageWait := new(pb.CommandMessageWait)
	commandMessageWait.WorkerSource = r.Identity
	commandMessageWait.Value = command
	commandMessageWait.IteratorId = id
	commandMessageGrpc, _ := r.ClientCommand.WaitCommandMessage(context.Background(), commandMessageWait)
	fmt.Println(commandMessageGrpc)

	for commandMessageGrpc == nil {
		time.Sleep(time.Duration(1) * time.Millisecond)
	}

	return pb.CommandFromGrpc(commandMessageGrpc)
}

//CreateIteratorCommand :
func (r ClientGrpcTest) CreateIteratorCommand() string {
	empty := new(pb.Empty)

	iteratorMessage, _ := r.ClientCommand.CreateIteratorCommand(context.Background(), empty)
	fmt.Println(iteratorMessage)

	for iteratorMessage == nil {
		time.Sleep(time.Duration(1) * time.Millisecond)
	}

	return iteratorMessage.GetId()
}

//SendEvent :
func (r ClientGrpcTest) SendEvent(topic, timeout, event, payload string) *pb.Empty {
	eventMessage := new(pb.EventMessage)
	eventMessage.Topic = topic
	eventMessage.Timeout = timeout
	eventMessage.UUID = uuid.New().String()
	eventMessage.Event = event
	eventMessage.Payload = payload
	empty, _ := r.ClientEvent.SendEventMessage(context.Background(), eventMessage)

	return empty
}

//CreateIteratorEvent :
func (r ClientGrpcTest) CreateIteratorEvent() string {
	empty := new(pb.Empty)

	iteratorMessage, _ := r.ClientEvent.CreateIteratorEvent(context.Background(), empty)
	fmt.Println(iteratorMessage)

	for iteratorMessage == nil {
		time.Sleep(time.Duration(1) * time.Millisecond)
	}

	return iteratorMessage.GetId()
}

//WaitEvent :
func (r ClientGrpcTest) WaitEvent(event, topic, id string) (eventMessage msg.Event) {
	eventMessageWait := new(pb.EventMessageWait)
	eventMessageWait.WorkerSource = r.Identity
	eventMessageWait.Topic = topic
	eventMessageWait.Event = event
	eventMessageWait.IteratorId = id
	eventMessageGrpc, _ := r.ClientEvent.WaitEventMessage(context.Background(), eventMessageWait)

	for eventMessageGrpc == nil {
		time.Sleep(time.Duration(1) * time.Millisecond)
	}

	return pb.EventFromGrpc(eventMessageGrpc)
}

//WaitEvent :
func (r ClientGrpcTest) WaitTopic(topic, id string) (eventMessage msg.Event) {
	topicMessageWait := new(pb.TopicMessageWait)
	topicMessageWait.WorkerSource = r.Identity
	topicMessageWait.Topic = topic
	topicMessageWait.IteratorId = id
	eventMessageGrpc, _ := r.ClientEvent.WaitTopicMessage(context.Background(), topicMessageWait)

	for eventMessageGrpc == nil {
		time.Sleep(time.Duration(1) * time.Millisecond)
	}

	return pb.EventFromGrpc(eventMessageGrpc)
}
