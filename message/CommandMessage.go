package message

import (
	"fmt"

	msgpack "github.com/shamaton/msgpack"
)

type CommandMessage struct {
	sourceAggregator    string
	sourceConnector string
	sourceWorker   string
	destinationAggregator    string
    destinationConnector    string
    destinationWorker string
    tenant   string
    token    string
    context    string
    timeout string
    timestamp   string
    major    string
    minor    string
	uuid string
	connectorType string
    commandType   string
    command    string
    payload    string
}

func (c CommandMessage) New(context, timeout, uuid, connectorType, commandType, command, payload string) err error {
	c.context = context
	c.timeout = timeout
	c.uuid = uuid
	c.connectorType = connectorType
	c.commandType = commandType
	c.command = command
	c.payload = payload
	c.timestamp = time.Now()
}

func (c CommandMessage) sendWith(socket zmq.Sock) err error {
	socket.send(c.encodeCommand(c))
}

func (c CommandMessage) from(command []byte) err error {
	c.sourceAggregator = command[0]
	c.sourceConnector = command[1]
	c.sourceWorker = command[2]
	c.destinationAggregator = command[3]
    c.destinationConnector = command[4]
    c.destinationWorker = command[5]
    c.tenant = command[6]
    c.token = command[7]
    c.context = command[8]
    c.timeout = command[9]
    c.timestamp = command[10]
    c.major = command[11]
    c.minor = command[12]
	c.uuid = command[13]
	c.connectorType = command[14]
    c.commandType = command[15]
    c.command = command[16]
    c.payload = command[17]
}

func (c Command) reply(socket zmq.Sock, reply, payload string) err error {
	commandResponse := CommandResponse.from(c, reply, payload)
	commandResponse.sendWith(socket)
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

type CommandReply struct {
	sourceAggregator    string
	sourceConnector string
	sourceWorker   string
	destinationAggregator    string
    destinationConnector    string
    destinationWorker string
    tenant   string
    token    string
    context    string
    timeout string
    timestamp   string
	uuid string
	reply    string
    payload    string
}

func (cr CommandResponse) sendWith(socket zmq.Sock) {
	socket.send(cr.encodeCommandResponse(cr))
}

func (cr CommandResponse) from(commandMessage CommandMessage, reply, payload string) {
	cr.sourceAggregator = commandMessage.sourceAggregator
	cr.sourceConnector = commandMessage.sourceConnector
	cr.sourceWorker = commandMessage.sourceWorker
	cr.destinationAggregator = commandMessage.destinationAggregator
    cr.destinationConnector = commandMessage.destinationConnector
    cr.destinationWorker = commandMessage.destinationWorker
    cr.tenant = commandMessage.tenant
    cr.token = commandMessage.token
    cr.context = commandMessage.context
    cr.timeout = commandMessage.timeout
    cr.timestamp = commandMessage.timestamp
	cr.uuid = commandMessage.uuid
	cr.reply = reply
    cr.payload = payload
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
