package message

import (
	"fmt"
	"time"
	"github.com/shamaton/msgpack"
	"github.com/zeromq/goczmq"
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
	e.timestamp = time.Now()
}

func (e EventMessage) sendWith(socket Socket, header string) {
	for {
		isSend := e.sendHeaderWith(socket, header)
		isSend += e.sendEventWith(socket)
		if isSend > 0 {
			break
		}
		time.Sleep(2 * time.Second)
	}
} 

func (e EventMessage) sendHeaderWith(socket Socket, header string) {
	for {
		isSend := socket.Send(header, FlagMore);
		if isSend > 0 {
			break
		}
		time.Sleep(2 * time.Second)
	}
}

func (e EventMessage) sendEventWith(socket Socket) {
	for {
		isSend := socket.Send(e.topic, FlagMore);
		isSend += socket.SendBytes(e.encodeEvent(e), 0);
		if isSend > 0 {
			break
		}
		time.Sleep(2 * time.Second)
	}
}

func (e EventMessage) from(event []byte) {
	e.tenant = event[0]
	e.token = event[1]
	e.topic = event[2]
	e.timeout = event[3]
	e.timestamp = event[4]
	e.event = event[5]
	e.payload = event[6]
}

func (e EventMessage) encodeEvent() (bytesContent []byte, commandError error) {
	bytesContent, err := msgpack.Encode(e)
	if err != nil {
		commandError = fmt.Errorf("Event %s", err)
		return
	}
	return
}

func (e EventMessage) decodeEvent(bytesContent []byte) (eventMessage EventMessage, commandError error) {
	err := msgpack.Decode(bytesContent, eventMessage)
	if err != nil {
		commandError = fmt.Errorf("Event %s", err)
		return
	}
	return
}
