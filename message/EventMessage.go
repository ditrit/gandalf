package message

import (
	"fmt"
	"message"
	msgpack "github.com/shamaton/msgpack"
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

func (e EventMessage) New(topic, timeout, event, payload string) err error {
	e.topic = topic
	e.timeout = timeout
	e.event = event
	e.payload = payload
	e.timestamp = time.Now()
}

func (e EventMessage) sendWith(socket zmq.Sock, header string) isSend bool {
	validation = e.sendHeaderWith(socket, header)
	validation += e.sendEventWith(socket)
	isSend := validation > 0 ? true : false
	return
} 

func (e EventMessage) sendHeaderWith(socket zmq.Sock, header string) isSend bool {
	validation = zmq_send(socket, header, ZMQ_SNDMORE);
	isSend := validation > 0 ? true : false
	return
}

func (e EventMessage) sendEventWith(socket zmq.Sock) isSend bool {
	validation = zmq_send(socket, e.topic, ZMQ_SNDMORE);
	validation += zmq_send(socket, e.encodeEvent(e), 0);
	isSend := validation > 0 ? true : false
	return
}

func (e EventMessage) from(event []byte) err error {
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

func (e EventMessage) decodeEvent(bytesContent []byte) (event Event, commandError error) {
	err := msgpack.Decode(bytesContent, event)
	if err != nil {
		commandError = fmt.Errorf("Event %s", err)
		return
	}
	return
}
