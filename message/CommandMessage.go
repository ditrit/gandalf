package message

import (
	"fmt"

	msgpack "github.com/shamaton/msgpack"
)

type Command struct {
	uuid    string
	routing string
	acces   string
	info    string
}

func (c Command) new(uuid, routing, acces, info string) {
	c.uuid = uuid
	c.routing = routing
	c.acces = acces
	c.info = info
}

func (c Command) sendWith() {

}

func (c Command) from() {

}

func (c Command) response() {

}

func (c Command) encodeCommand() (bytesContent []byte, commandError error) {
	bytesContent, err := msgpack.Encode(c)
	if err != nil {
		commandError = fmt.Errorf("Command %s", err)
		return
	}
	return
}

func (c Command) decodeCommand(bytesContent []byte) (command Command, commandError error) {
	err := msgpack.Decode(bytesContent, command)
	if err != nil {
		commandError = fmt.Errorf("Command %s", err)
		return
	}
	return
}

type CommandResponse struct {
	command Command
}

func (cr CommandResponse) new(uuid, routing, acces, info string) {
	cr.command.uuid = uuid
	cr.command.routing = routing
	cr.command.acces = acces
	cr.command.info = info
}

func (cr CommandResponse) sendWith() {
}

func (cr CommandResponse) from() {

}

func (cr CommandResponse) encode() {

}

func (cr CommandResponse) decode() {

}

func (cr CommandResponse) encodeCommandResponse() (bytesContent []byte, commandError error) {
	bytesContent, err := msgpack.Encode(cr)
	if err != nil {
		commandError = fmt.Errorf("CommandResponse %s", err)
		return
	}
	return
}

func (cr CommandResponse) decodeCommandResponse(bytesContent []byte) (commandResponse CommandResponse, commandError error) {
	err := msgpack.Decode(bytesContent, commandResponse)
	if err != nil {
		commandError = fmt.Errorf("CommandResponse %s", err)
		return
	}
	return
}
