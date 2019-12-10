package message

import (
	"fmt"
	"constant"
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

func (c CommandMessage) sendWith(socket zmq.Sock, header string) {
	cr.sendHeaderWith(socket, header)
	cr.sendCommandWith(socket)
}

func (c CommandMessage) sendHeaderWith(socket zmq.Sock, header string) {
	zmq_send(socket, header, ZMQ_SNDMORE);
}

func (c CommandMessage) sendCommandWith(socket zmq.Sock) {
	zmq_send(socket, encode(c), 0);
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

//

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

func (cr CommandReply) sendWith(socket zmq.Sock, header string) {
	cr.sendHeaderWith(socket, header)
	cr.sendCommandReplyWith(socket)
}

func (cr CommandReply) sendHeaderCommandReplyWith(socket zmq.Sock, header string) {
	zmq_send(socket, header, ZMQ_SNDMORE);
}

func (cr CommandReply) sendCommandReplyWith(socket zmq.Sock) {
	zmq_send(socket, encode(cr), 0);
}

func (cr CommandReply) from(commandMessage CommandMessage, reply, payload string) {
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

//

type CommandCommandsEvents struct {
	commands    []string
	events    	[]string
}

func (cce CommandCommandsEvents) New(commands, events []string) err error {
	cce.commands = commands
	cce.events = events
}

func (cce CommandCommandsEvents) sendWith(socket zmq.Sock) {
	zmq_send(socket, constant.COMMAND_VALIDATION_FUNCTIONS, ZMQ_SNDMORE);
	zmq_send(socket, encode(cce), 0);
}

//

type CommandCommandsEventsReply struct {
	validation bool
}

func (ccer CommandCommandsEventsReply) New(validation bool) err error {
	ccer.validation = validation
}

func (ccer CommandCommandsEventsReply) sendWith(socket zmq.Sock, header string) {
	ccer.sendHeaderWith(socket, header)
	ccer.sendCommandCommandsEventsReplyWith(socket)
}

func (ccer CommandCommandsEventsReply) sendHeaderWith(socket zmq.Sock, header string) {
	zmq_send(socket, header, ZMQ_SNDMORE);
}

func (ccer CommandCommandsEventsReply) sendCommandCommandsEventsReplyWith(socket zmq.Sock) {
	zmq_send(socket, constant.COMMAND_VALIDATION_FUNCTIONS_REPLY, ZMQ_SNDMORE);
	zmq_send(socket, encode(ccer), 0);
}

//

type CommandReady struct {
	// ???
}

func (cry CommandReady) New() err error {
}

func (cry CommandReady) sendWith(socket zmq.Sock) {
	zmq_send(socket, constant.COMMAND_READY, ZMQ_SNDMORE);
	zmq_send(socket, encode(cry), 0);
}

//

func encode() (bytesContent []byte, commandError error) {
	bytesContent, err := msgpack.Encode(c)
	if err != nil {
		commandError = fmt.Errorf("Command %s", err)
		return
	}
	return
}

func decodeCommand(bytesContent []byte) (command Command, commandError error) {
	err := msgpack.Decode(bytesContent, command)
	if err != nil {
		commandError = fmt.Errorf("Command %s", err)
		return
	}
	return
}	

func decodeCommandReply(bytesContent []byte) (commandReply CommandReply, commandError error) {
	err := msgpack.Decode(bytesContent, commandReply)
	if err != nil {
		commandError = fmt.Errorf("CommandResponse %s", err)
		return
	}
	return
}

func decodeCommandReady(bytesContent []byte) (commandReady CommandReady, commandError error) {
	err := msgpack.Decode(bytesContent, CommandReady)
	if err != nil {
		commandError = fmt.Errorf("CommandResponse %s", err)
		return
	}
	return
}

func decodeCommandCommandsEvents(bytesContent []byte) (commandCommandsEvents CommandCommandsEvents, commandError error) {
	err := msgpack.Decode(bytesContent, commandCommandsEvents)
	if err != nil {
		commandError = fmt.Errorf("CommandResponse %s", err)
		return
	}
	return
}

func decodeCommandCommandsEventsReply(bytesContent []byte) (commandCommandsEventsReply CommandCommandsEventsReply, commandError error) {
	err := msgpack.Decode(bytesContent, commandCommandsEventsReply)
	if err != nil {
		commandError = fmt.Errorf("CommandResponse %s", err)
		return
	}
	return
}