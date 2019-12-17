package message

import (
	"fmt"
	"time"
	"github.com/shamaton/msgpack"
	zmq4 "github.com/pebbe/zmq4"
)

type EventMessage struct {
	Tenant string
	Token  string
	Topic string
	Timeout  string
	Timestamp  string
	Event  string
	Payload  string
}

func (e EventMessage) New(topic, timeout, event, payload string) {
	e.Topic = topic
	e.Timeout = timeout
	e.Event = event
	e.Payload = payload
	e.Timestamp = time.Now().String()
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
	e.Event = event[5]
	e.Payload = event[6]
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
