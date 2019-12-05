package message

import (
	"fmt"

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

func (e EventMessage) sendWith() err error {

}

func (e EventMessage) from() err error {

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
