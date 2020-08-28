//Package goclient :
//File ClientGandalf.go
package goclient

import (
	"strings"
	"time"

	"github.com/ditrit/gandalf/libraries/goclient/models"

	pb "github.com/ditrit/gandalf/libraries/goclient/grpc"

	"github.com/ditrit/gandalf/libraries/goclient/client"

	"github.com/ditrit/shoset/msg"
)

var clientIndex = 0
var defaultTimeout = "10000"

//ClientGandalf :
type ClientGandalf struct {
	Identity          string
	ClientConnections []string
	Clients           []*client.ClientGrpc
	ClientStopChannel chan int
	Timeout           string
}

//NewClientGandalf
func NewClientGandalf(identity, timeout string, clientConnections []string) (clientGandalf *ClientGandalf) {
	clientGandalf = new(ClientGandalf)
	clientGandalf.ClientStopChannel = make(chan int)

	clientGandalf.Identity = identity
	clientGandalf.ClientConnections = clientConnections
	if timeout == "" {
		clientGandalf.Timeout = defaultTimeout
	} else {
		clientGandalf.Timeout = timeout
	}

	for _, connection := range clientGandalf.ClientConnections {
		clientGandalf.Clients = append(clientGandalf.Clients, client.NewClientGrpc(clientGandalf.Identity, connection))
	}

	return
}

//SendCommand
func (cg ClientGandalf) SendCommand(request string, options *models.Options) (commandMessageUUID *pb.CommandMessageUUID) {
	var notSend bool
	requestSplit := strings.Split(request, ".")
	timeout := options.GetTimeout()
	if timeout == "" {
		timeout = cg.Timeout
	}

	for stay, timeoutLoop := true, time.After(time.Second); stay; {

		commandMessageUUID = cg.Clients[getClientIndex(cg.Clients, true)].SendCommand(requestSplit[0], requestSplit[1], timeout, options.GetPayload())
		if commandMessageUUID != nil {
			notSend = false
			break
		}

		select {
		case <-timeoutLoop:
			stay = false
			notSend = true
		default:
		}
	}

	if notSend {
		return nil
	}

	return commandMessageUUID
}

//SendEvent
func (cg ClientGandalf) SendEvent(topic, event string, options *models.Options) (empty *pb.Empty) {
	var notSend bool
	timeout := options.GetTimeout()
	if timeout == "" {
		timeout = cg.Timeout
	}

	for stay, timeoutLoop := true, time.After(time.Second); stay; {

		empty = cg.Clients[getClientIndex(cg.Clients, true)].SendEvent(topic, event, "", timeout, options.GetPayload())

		if empty != nil {
			notSend = false
			break
		}

		select {
		case <-timeoutLoop:
			stay = false
			notSend = true
		default:
		}
	}

	if notSend {
		return nil
	}

	return empty
}

//SendEvent
func (cg ClientGandalf) SendReply(topic, event, referenceUUID string, options *models.Options) (empty *pb.Empty) {
	var notSend bool
	timeout := options.GetTimeout()
	if timeout == "" {
		timeout = cg.Timeout
	}

	for stay, timeoutLoop := true, time.After(time.Second); stay; {

		empty = cg.Clients[getClientIndex(cg.Clients, true)].SendEvent(topic, event, referenceUUID, timeout, options.GetPayload())

		if empty != nil {
			notSend = false
			break
		}

		select {
		case <-timeoutLoop:
			stay = false
			notSend = true
		default:
		}
	}

	if notSend {
		return nil
	}

	return empty
}

//SendCommandList
func (cg ClientGandalf) SendCommandList(version int64, commands []string) (empty *pb.Empty) {
	empty = cg.Clients[getClientIndex(cg.Clients, true)].SendCommandList(version, commands)

	return empty
}

//WaitCommand
func (cg ClientGandalf) WaitCommand(command, idIterator string, version int64) (commandMessage msg.Command) {
	return pb.CommandFromGrpc(cg.Clients[getClientIndex(cg.Clients, false)].WaitCommand(command, idIterator, version))
}

//WaitEvent
func (cg ClientGandalf) WaitEvent(topic, event, idIterator string) (eventMessage msg.Event) {
	return pb.EventFromGrpc(cg.Clients[getClientIndex(cg.Clients, false)].WaitEvent(topic, event, "", idIterator))
}

//WaitReplyByEvent
func (cg ClientGandalf) WaitReplyByEvent(topic, event, referenceUUID, idIterator string) (eventMessage msg.Event) {
	return pb.EventFromGrpc(cg.Clients[getClientIndex(cg.Clients, false)].WaitEvent(topic, event, referenceUUID, idIterator))
}

//WaitTopic
func (cg ClientGandalf) WaitTopic(topic, idIterator string) (eventMessage msg.Event) {
	return pb.EventFromGrpc(cg.Clients[getClientIndex(cg.Clients, false)].WaitTopic(topic, "", idIterator))
}

//WaitReplyByTopic
func (cg ClientGandalf) WaitReplyByTopic(topic, referenceUUID, idIterator string) (eventMessage msg.Event) {
	return pb.EventFromGrpc(cg.Clients[getClientIndex(cg.Clients, false)].WaitTopic(topic, referenceUUID, idIterator))
}

//WaitAllReplyByTopic
func (cg ClientGandalf) WaitAllReplyByTopic(topic, referenceUUID, idIterator string) (eventMessages []msg.Event) {
	client := cg.Clients[getClientIndex(cg.Clients, false)]
	for {
		message := pb.EventFromGrpc(client.WaitTopic(topic, referenceUUID, idIterator))
		eventMessages = append(eventMessages, message)

		if message.GetEvent() == "SUCCES" {
			break
		}
	}
	return
	//return pb.EventFromGrpc(cg.Clients[getClientIndex(cg.Clients, false)].WaitTopic(topic, referenceUUID, idIterator))
}

/* //WaitTopicEvent
func (cg ClientGandalf) WaitTopicEvent(topic, event, refUUID, idIterator string) (eventMessage msg.Event) {
	var message *pb.EventMessage
	client := cg.Clients[getClientIndex(cg.Clients, false)]
	for {
		message = client.WaitTopic(topic, refUUID, idIterator)
		if message.GetEvent() == event {
			return pb.EventFromGrpc(message)
		}
	}
	//return pb.EventFromGrpc(cg.Clients[getClientIndex(cg.Clients, false)].WaitTopic(topic, idIterator))
} */

//CreateIteratorCommand
func (cg ClientGandalf) CreateIteratorCommand() string {
	return cg.Clients[getClientIndex(cg.Clients, false)].CreateIteratorCommand().GetId()
}

//CreateIteratorEvent
func (cg ClientGandalf) CreateIteratorEvent() string {
	return cg.Clients[getClientIndex(cg.Clients, false)].CreateIteratorEvent().GetId()
}

func getClientIndex(conns []*client.ClientGrpc, updateIndex bool) int {
	aux := clientIndex

	if updateIndex {
		clientIndex++
		if clientIndex >= len(conns) {
			clientIndex = 0
		}
	}

	return aux
}
