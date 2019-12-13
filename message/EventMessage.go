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
	e.timestamp = time.Now().String()
}

func (e EventMessage) sendWith(socket goczmq.Sock, header string) (isSend bool) {
	for {
		isSend = e.sendHeaderWith(socket, header)
		isSend = isSend && isSend && e.sendEventWith(socket)
		if isSend {
			return
		}
		time.Sleep(2 * time.Second)
	}
} 

func (e EventMessage) sendHeaderWith(socket goczmq.Sock, header string) (isSend bool) {
	for {
		err := socket.SendFrame([]byte(header), goczmq.FlagMore)
		if err == nil {
			isSend = true
			return
		}
		time.Sleep(2 * time.Second)
	}
}

func (e EventMessage) sendEventWith(socket goczmq.Sock) (isSend bool) {
	for {
		err := socket.SendFrame([]byte(e.topic), goczmq.FlagMore)
		if err == nil {
			encoded, _ := encodeEvent(e)
			err = socket.SendFrame(encoded, 0)
			if err == nil {
				isSend = true
				return
			}
		}
		time.Sleep(2 * time.Second)
	}
}

func (e EventMessage) from(event []string) {
	e.tenant = event[0]
	e.token = event[1]
	e.topic = event[2]
	e.timeout = event[3]
	e.timestamp = event[4]
	e.event = event[5]
	e.payload = event[6]
}

func encodeEvent(eventMessage EventMessage) (bytesContent []byte, commandError error) {
	bytesContent, err := msgpack.Encode(eventMessage)
	if err != nil {
		commandError = fmt.Errorf("Event %s", err)
		return
	}
	return
}

func decodeEvent(bytesContent []byte) (eventMessage EventMessage, commandError error) {
	err := msgpack.Decode(bytesContent, eventMessage)
	if err != nil {
		commandError = fmt.Errorf("Event %s", err)
		return
	}
	return
}
