package message

import (
	"fmt"
	"time"
	"gandalfgo/constant"
	"github.com/shamaton/msgpack"
	zmq4 "github.com/pebbe/zmq4"
)

type EventMessage struct {
	Tenant 		string
	Token  		string
	Topic 		string
	Timeout  	string
	Timestamp  	string
	Uuid 		string
	Event  		string
	Payload  	string
}

func NewEventMessage(topic, timeout, uuid, event, payload string) (eventMessage *EventMessage) {
	eventMessage = new(EventMessage)
	eventMessage.Topic = topic
	eventMessage.Timeout = timeout
	eventMessage.Timestamp = time.Now().String()
	eventMessage.Uuid = uuid
	eventMessage.Event = event
	eventMessage.Payload = payload
	
	return
}

func (e EventMessage) SendWith(socket *zmq4.Socket, header string) (isSend bool) {
	for {
		isSend = e.SendHeaderWith(socket, header)
		isSend = isSend && isSend && e.SendEventWith(socket)
		if isSend {
			return
		}
		time.Sleep(2 * time.Second)
	}
} 

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

func (e EventMessage) SendEventWith(socket *zmq4.Socket) (isSend bool) {
	for {
		_, err := socket.Send(e.Topic, zmq4.SNDMORE)
		if err == nil {
			encoded, _ := EncodeEventMessage(e)
			_, err = socket.SendBytes(encoded, 0)
			if err == nil {
				isSend = true
				return
			}
		}
		time.Sleep(2 * time.Second)
	}
}

func (e EventMessage) From(event []string) {
	e.Tenant = event[0]
	e.Token = event[1]
	e.Topic = event[2]
	e.Timeout = event[3]
	e.Timestamp = event[4]
	e.Uuid = event[5]
	e.Event = event[6]
	e.Payload = event[7]
}

type EventFunction struct {
	Worker		 string
	Functions    []string
}

func NewEventFunction(worker string, functions []string) (eventFunction *EventFunction) {
	eventFunction = new(EventFunction)
	eventFunction.Functions = functions

	return
}

func (cf EventFunction) SendWith(socket *zmq4.Socket) (isSend bool) {
	for {
		_, err := socket.Send(constant.EVENT_VALIDATION_TOPIC, zmq4.SNDMORE);
		_, err = socket.Send(constant.EVENT_VALIDATION_FUNCTIONS, zmq4.SNDMORE);
		if err == nil {
			encoded, _ := EncodeEventFunction(cf)
			_, err = socket.SendBytes(encoded, 0);
			if err == nil {
				isSend = true
				return
			}
		}
		time.Sleep(2 * time.Second)
	}
}

//

type EventFunctionReply struct {
	Validation bool
}

func NewEventFunctionReply(validation bool) (eventFunctionReply *EventFunctionReply) {
	eventFunctionReply = new(EventFunctionReply) 
	eventFunctionReply.Validation = validation

	return
}

func (cfr EventFunctionReply) SendWith(socket *zmq4.Socket, header string) (isSend bool) {
	for {
		isSend = cfr.SendHeaderWith(socket, header)
		isSend = isSend && cfr.SendEventFunctionReplyWith(socket)
		if isSend {
			return
		}
		time.Sleep(2 * time.Second)
	}
}

func (cfr EventFunctionReply) SendHeaderWith(socket *zmq4.Socket, header string) (isSend bool) {
	for {
		_, err := socket.Send(header, zmq4.SNDMORE);
		if err == nil {
			isSend = true
			return
		}
		time.Sleep(2 * time.Second)
	}
}

func (cfr EventFunctionReply) SendEventFunctionReplyWith(socket *zmq4.Socket) (isSend bool) {
	for {
		_, err := socket.Send(constant.COMMAND_VALIDATION_FUNCTIONS_REPLY, zmq4.SNDMORE);
		if err == nil {
			encoded, _ := EncodeEventFunctionReply(cfr)
			_, err = socket.SendBytes(encoded, 0);
			if err == nil {
				isSend = true
				return
			}
		}
		
		time.Sleep(2 * time.Second)
	}
}

func EncodeEventMessage(eventMessage EventMessage) (bytesContent []byte, commandError error) {
	bytesContent, err := msgpack.Encode(eventMessage)
	if err != nil {
		commandError = fmt.Errorf("Event %s", err)
		return
	}
	return
}

func DecodeEventMessage(bytesContent []byte) (eventMessage EventMessage, commandError error) {
	err := msgpack.Decode(bytesContent, eventMessage)
	if err != nil {
		commandError = fmt.Errorf("Event %s", err)
		return
	}
	return
}

func EncodeEventFunction(eventFunction EventFunction) (bytesContent []byte, commandError error) {
	bytesContent, err := msgpack.Encode(eventFunction)
	if err != nil {
		commandError = fmt.Errorf("Event %s", err)
		return
	}
	return
}

func DecodeEventFunction(bytesContent []byte) (eventFunction EventFunction, commandError error) {
	err := msgpack.Decode(bytesContent, eventFunction)
	if err != nil {
		commandError = fmt.Errorf("Event %s", err)
		return
	}
	return
}

func EncodeEventFunctionReply(eventFunctionReply EventFunctionReply) (bytesContent []byte, commandError error) {
	bytesContent, err := msgpack.Encode(eventFunctionReply)
	if err != nil {
		commandError = fmt.Errorf("Event %s", err)
		return
	}
	return
}

func DecodeEventFunctionReply(bytesContent []byte) (eventFunctionReply EventFunctionReply, commandError error) {
	err := msgpack.Decode(bytesContent, eventFunctionReply)
	if err != nil {
		commandError = fmt.Errorf("Event %s", err)
		return
	}
	return
}