package message

import (
	"fmt"

	msgpack "github.com/shamaton/msgpack"
)

type Event struct {
	topic string
	uuid  string
	acces string
	info  string
}

func (e Event) New(topic, uuid, acces, info string) err error {
	e.topic = topic
	e.uuid = uuid
	e.acces = acces
	e.info = info
}

func (e Event) sendWith() err error {

}

func (e Event) from() err error {

}

func (e Event) encodeEvent() (bytesContent []byte, commandError error) {
	bytesContent, err := msgpack.Encode(e)
	if err != nil {
		commandError = fmt.Errorf("Event %s", err)
		return
	}
	return
}

func (e Event) decodeEvent(bytesContent []byte) (event Event, commandError error) {
	err := msgpack.Decode(bytesContent, event)
	if err != nil {
		commandError = fmt.Errorf("Event %s", err)
		return
	}
	return
}
