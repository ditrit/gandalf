//Package event :
//File ClientEvent.go
package event

import (
	"context"
	"fmt"
	"time"

	pb "github.com/ditrit/gandalf/libraries/goclient/grpc"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

//ClientEvent :
type ClientEvent struct {
	ClientEventConnection string
	Identity              string
	client                pb.ConnectorEventClient
}

//NewClientEvent :
func NewClientEvent(identity, clientEventConnection string) (clientEvent *ClientEvent) {
	clientEvent = new(ClientEvent)
	clientEvent.Identity = identity
	clientEvent.ClientEventConnection = clientEventConnection

	conn, err := grpc.Dial(clientEvent.ClientEventConnection, grpc.WithInsecure())

	if err != nil {
		fmt.Println("clientEvent failed dial")
	}

	clientEvent.client = pb.NewConnectorEventClient(conn)
	fmt.Println("clientEvent connect : " + clientEvent.ClientEventConnection)

	return
}

//SendEvent :
func (ce ClientEvent) SendEvent(topic, event, referenceUUID, timeout, payload string) *pb.Empty {
	eventMessage := new(pb.EventMessage)
	eventMessage.Topic = topic
	eventMessage.Timeout = timeout
	eventMessage.UUID = uuid.New().String()
	eventMessage.Event = event
	eventMessage.Payload = payload
	eventMessage.ReferenceUUID = referenceUUID
	empty, _ := ce.client.SendEventMessage(context.Background(), eventMessage)

	return empty
}

//WaitEvent :
func (ce ClientEvent) WaitEvent(topic, event, referenceUUID, idIterator string) *pb.EventMessage {
	eventMessageWait := new(pb.EventMessageWait)
	eventMessageWait.WorkerSource = ce.Identity
	eventMessageWait.Topic = topic
	eventMessageWait.Event = event
	eventMessageWait.IteratorId = idIterator
	eventMessageWait.ReferenceUUID = referenceUUID
	eventMessage, _ := ce.client.WaitEventMessage(context.Background(), eventMessageWait)

	for eventMessage == nil {
		time.Sleep(time.Duration(1) * time.Second)
	}

	return eventMessage
}

//WaitEvent :
func (ce ClientEvent) WaitTopic(topic, referenceUUID, idIterator string) *pb.EventMessage {
	topicMessageWait := new(pb.TopicMessageWait)
	topicMessageWait.WorkerSource = ce.Identity
	topicMessageWait.Topic = topic
	topicMessageWait.IteratorId = idIterator
	topicMessageWait.ReferenceUUID = referenceUUID
	eventMessage, _ := ce.client.WaitTopicMessage(context.Background(), topicMessageWait)

	for eventMessage == nil {
		time.Sleep(time.Duration(1) * time.Second)
	}

	return eventMessage
}

//CreateIteratorEvent :
func (ce ClientEvent) CreateIteratorEvent() *pb.IteratorMessage {
	iteratorMessage, _ := ce.client.CreateIteratorEvent(context.Background(), new(pb.Empty))
	return iteratorMessage
}
