package message

import (
	"fmt"

	msgpack "github.com/shamaton/msgpack"
)

type CommandMessage struct {
	sourceAggregator    string
	sourceConnector string
	sourceWorker   string
	targetAggregator    string
    targetConnector    string
    targetWorker string
    tenant   string
    token    string
    context    string
    timeout string
    timestamp   string
    major    string
    minor    string
    uuid string
    commandType   string
    command    string
    payload    string
}

func (c CommandMessage) New(context, timeout, uuid, commandType, command, payload string) err error {
	c.context = context
	c.timeout = timeout
	c.uuid = uuid
	c.commandType = commandType
	c.command = command
	c.payload = payload
	c.timestamp = time.Now()
}

func (c CommandMessage) sendWith() err error {

}

func (c CommandMessage) from() err error {

}

func (c Command) response() err error {

}

func (c CommandMessage) encodeCommand() (bytesContent []byte, commandError error) {
	bytesContent, err := msgpack.Encode(c)
	if err != nil {
		commandError = fmt.Errorf("Command %s", err)
		return
	}
	return
}

func (c CommandMessage) decodeCommand(bytesContent []byte) (command Command, commandError error) {
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
