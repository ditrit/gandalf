package message

import (
	"fmt"
	"time"
	"github.com/shamaton/msgpack"
	zmq4 "github.com/pebbe/zmq4"
)

type EventMessage struct {
	tenant string
	token  string
	topic string
	timeout  string
	timestamp  string
	event  string
	payload  string
}

func (e EventMessage) New(topic, timeout, event, payload string) {
	e.topic = topic
	e.timeout = timeout
	e.event = event
	e.payload = payload
	e.timestamp = time.Now().String()
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
		_, err := socket.Send(e.topic, zmq4.SNDMORE)
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
	e.tenant = event[0]
	e.token = event[1]
	e.topic = event[2]
	e.timeout = event[3]
	e.timestamp = event[4]
	e.event = event[5]
	e.payload = event[6]
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
