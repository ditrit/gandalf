//Package message :
//File EventMessage.go
package message

import (
	"fmt"
	"gandalf-go/constant"
	pb "gandalf-go/grpc"
	"time"

	"github.com/pebbe/zmq4"
	"github.com/shamaton/msgpack"
)

//EventMessage :
type EventMessage struct {
	Tenant    string
	Token     string
	Topic     string
	Timeout   string
	Timestamp string
	UUID      string
	Event     string
	Payload   string
}

//NewEventMessage :
func NewEventMessage(topic, timeout, uuid, event, payload string) (eventMessage *EventMessage) {
	eventMessage = new(EventMessage)
	eventMessage.Topic = topic
	eventMessage.Timeout = timeout
	eventMessage.Timestamp = time.Now().String()
	eventMessage.UUID = uuid
	eventMessage.Event = event
	eventMessage.Payload = payload

	return
}

//GetUUID :
func (e EventMessage) GetUUID() string {
	return e.UUID
}

//GetTimeout :
func (e EventMessage) GetTimeout() string {
	return e.Timeout
}

//SendWith :
func (e EventMessage) SendWith(socket *zmq4.Socket, header string) (isSend bool) {
	for {
		isSend = e.SendHeaderWith(socket, header)

		isSend = isSend && e.SendMessageWith(socket)
		if isSend {
			return
		}

		time.Sleep(2 * time.Second)
	}
}

//SendHeaderWith :
func (e EventMessage) SendHeaderWith(socket *zmq4.Socket, header string) (isSend bool) {
	for {
		_, err := socket.Send(header, zmq4.SNDMORE)
		if err == nil {
			isSend = true
			return
		}

		time.Sleep(2 * time.Second)
	}
}

//SendMessageWith :
func (e EventMessage) SendMessageWith(socket *zmq4.Socket) (isSend bool) {
	for {
		_, err := socket.SendBytes([]byte(e.Topic), zmq4.SNDMORE)
		if err == nil {
			encoded, _ := EncodeEventMessage(e)
			fmt.Println(encoded)
			_, err = socket.SendBytes(encoded, 0)

			if err == nil {
				isSend = true
				return
			}
		}

		time.Sleep(2 * time.Second)
	}
}

//From :
func (e EventMessage) From(event []string) {
	e.Tenant = event[0]
	e.Token = event[1]
	e.Topic = event[2]
	e.Timeout = event[3]
	e.Timestamp = event[4]
	e.UUID = event[5]
	e.Event = event[6]
	e.Payload = event[7]
}

//EventMessageFromGrpc :
func EventMessageFromGrpc(eventMessage *pb.EventMessage) (e EventMessage) {
	e.Tenant = eventMessage.GetTenant()
	e.Token = eventMessage.GetToken()
	e.Timeout = eventMessage.GetTimeout()
	e.Timestamp = eventMessage.GetTimestamp()
	e.UUID = eventMessage.GetUuid()
	e.Topic = eventMessage.GetTopic()
	e.Event = eventMessage.GetEvent()
	e.Payload = eventMessage.GetPayload()

	return
}

//EventMessageToGrpc :
func EventMessageToGrpc(e EventMessage) (eventMessage *pb.EventMessage) {
	eventMessage = new(pb.EventMessage)
	eventMessage.Tenant = e.Tenant
	eventMessage.Token = e.Token
	eventMessage.Timeout = e.Timeout
	eventMessage.Timestamp = e.Timestamp
	eventMessage.UUID = e.UUID
	eventMessage.Topic = e.Topic
	eventMessage.Event = e.Event
	eventMessage.Payload = e.Payload

	return
}

//EventFunction :
type EventFunction struct {
	Worker    string
	Functions []string
}

//NewEventFunction :
func NewEventFunction(worker string, functions []string) (eventFunction *EventFunction) {
	eventFunction = new(EventFunction)
	eventFunction.Functions = functions

	return
}

//SendWith :
func (cf EventFunction) SendWith(socket *zmq4.Socket) (isSend bool) {
	for {
		_, err := socket.Send(constant.EVENT_VALIDATION_FUNCTIONS, zmq4.SNDMORE)
		if err == nil {
			encoded, _ := EncodeEventFunction(cf)
			_, err = socket.SendBytes(encoded, 0)

			if err == nil {
				isSend = true
				return
			}
		}

		time.Sleep(2 * time.Second)
	}
}

//EventFunctionReply :
type EventFunctionReply struct {
	Validation bool
}

//NewEventFunctionReply :
func NewEventFunctionReply(validation bool) (eventFunctionReply *EventFunctionReply) {
	eventFunctionReply = new(EventFunctionReply)
	eventFunctionReply.Validation = validation

	return
}

//SendWith :
func (cfr EventFunctionReply) SendWith(socket *zmq4.Socket, header string) (isSend bool) {
	for {
		isSend = cfr.SendHeaderWith(socket, header)
		isSend = isSend && cfr.SendMessageWith(socket)

		if isSend {
			return
		}

		time.Sleep(2 * time.Second)
	}
}

//SendHeaderWith :
func (cfr EventFunctionReply) SendHeaderWith(socket *zmq4.Socket, header string) (isSend bool) {
	for {
		_, err := socket.Send(header, zmq4.SNDMORE)
		if err == nil {
			isSend = true
			return
		}

		time.Sleep(2 * time.Second)
	}
}

//SendMessageWith :
func (cfr EventFunctionReply) SendMessageWith(socket *zmq4.Socket) (isSend bool) {
	for {
		_, err := socket.Send(constant.COMMAND_VALIDATION_FUNCTIONS_REPLY, zmq4.SNDMORE)
		if err == nil {
			encoded, _ := EncodeEventFunctionReply(cfr)
			_, err = socket.SendBytes(encoded, 0)

			if err == nil {
				isSend = true
				return
			}
		}

		time.Sleep(2 * time.Second)
	}
}

//SendSubscribeTopic :
func SendSubscribeTopic(socket *zmq4.Socket, topic []byte) (isSend bool) {
	for {
		_, err := socket.SendBytes(topic, 0)
		if err == nil {
			isSend = true
			return
		}

		time.Sleep(2 * time.Second)
	}
}

//EventMessageWait :
type EventMessageWait struct {
	WorkerSource string
	Event        string
	Topic        string
}

//NewEventMessageWait :
func NewEventMessageWait(workerSource, event, topic string) (eventMessageWait *EventMessageWait) {
	eventMessageWait = new(EventMessageWait)
	eventMessageWait.WorkerSource = workerSource
	eventMessageWait.Event = event
	eventMessageWait.Topic = topic

	return
}

//EventMessageWaitFromGrpc :
func EventMessageWaitFromGrpc(eventMessageWait pb.EventMessageWait) (emw EventMessageWait) {
	emw.WorkerSource = eventMessageWait.GetWorkerSource()
	emw.Event = eventMessageWait.GetEvent()
	emw.Topic = eventMessageWait.GetTopic()

	return
}

//EventMessageWaitToGrpc :
func EventMessageWaitToGrpc(emw EventMessageWait) (eventMessageWait pb.EventMessageWait) {
	eventMessageWait.WorkerSource = emw.WorkerSource
	eventMessageWait.Event = emw.Event
	eventMessageWait.Topic = emw.Topic

	return
}

//SendWith :
func (emw EventMessageWait) SendWith(socket *zmq4.Socket) (isSend bool) {
	for {
		_, err := socket.Send(constant.EVENT_WAIT, zmq4.SNDMORE)
		if err == nil {
			encoded, _ := EncodeEventMessageWait(emw)
			_, err = socket.SendBytes(encoded, 0)

			if err == nil {
				isSend = true
				return
			}
		}

		time.Sleep(2 * time.Second)
	}
}

//EncodeEventMessage :
func EncodeEventMessage(eventMessage EventMessage) (bytesContent []byte, commandError error) {
	bytesContent, err := msgpack.Encode(eventMessage)
	if err != nil {
		commandError = fmt.Errorf("event %s", err)
	}

	return
}

//DecodeEventMessage :
func DecodeEventMessage(bytesContent []byte) (eventMessage EventMessage, commandError error) {
	err := msgpack.Decode(bytesContent, &eventMessage)
	if err != nil {
		commandError = fmt.Errorf("event %s", err)
	}

	return
}

//EncodeEventFunction :
func EncodeEventFunction(eventFunction EventFunction) (bytesContent []byte, commandError error) {
	bytesContent, err := msgpack.Encode(eventFunction)
	if err != nil {
		commandError = fmt.Errorf("event %s", err)
	}

	return
}

//DecodeEventFunction :
func DecodeEventFunction(bytesContent []byte) (eventFunction EventFunction, commandError error) {
	err := msgpack.Decode(bytesContent, &eventFunction)
	if err != nil {
		commandError = fmt.Errorf("event %s", err)
	}

	return
}

//EncodeEventFunctionReply :
func EncodeEventFunctionReply(eventFunctionReply EventFunctionReply) (bytesContent []byte, commandError error) {
	bytesContent, err := msgpack.Encode(eventFunctionReply)
	if err != nil {
		commandError = fmt.Errorf("event %s", err)
	}

	return
}

//DecodeEventFunctionReply :
func DecodeEventFunctionReply(bytesContent []byte) (eventFunctionReply EventFunctionReply, commandError error) {
	err := msgpack.Decode(bytesContent, &eventFunctionReply)
	if err != nil {
		commandError = fmt.Errorf("event %s", err)
	}

	return
}

//EncodeEventMessageWait :
func EncodeEventMessageWait(eventMessageWait EventMessageWait) (bytesContent []byte, commandError error) {
	bytesContent, err := msgpack.Encode(eventMessageWait)
	if err != nil {
		commandError = fmt.Errorf("event %s", err)
	}

	return
}

//DecodeEventMessageWait :
func DecodeEventMessageWait(bytesContent []byte) (eventMessageWait EventMessageWait, commandError error) {
	err := msgpack.Decode(bytesContent, &eventMessageWait)
	if err != nil {
		commandError = fmt.Errorf("event %s", err)
	}

	return
}
