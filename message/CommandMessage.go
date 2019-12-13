package message

import (
	"fmt"
	"time"
	"gandalfgo/constant"
	"github.com/shamaton/msgpack"
	"github.com/zeromq/goczmq"
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

func (c CommandMessage) New(context, timeout, uuid, connectorType, commandType, command, payload string) {
	c.context = context
	c.timeout = timeout
	c.uuid = uuid
	c.connectorType = connectorType
	c.commandType = commandType
	c.command = command
	c.payload = payload
	c.timestamp = time.Now()
}

func (c CommandMessage) sendWith(socket Socket, header string) {
	for {
		isSend := c.sendHeaderWith(socket, header)
		isSend += c.sendCommandWith(socket)
		if isSend > 0 {
			break
		}
		time.Sleep(2 * time.Second)
	}
}

func (c CommandMessage) sendHeaderWith(socket Socket, header string) {
	for {
		isSend := socket.Send(header, FlagMore);
		if isSend > 0 {
			break
		}
		time.Sleep(2 * time.Second)
	}
}

func (c CommandMessage) sendCommandWith(socket Socket) {
	for {
		isSend := socket.SendBytes(encode(c), 0);
		if isSend > 0 {
			break
		}
		time.Sleep(2 * time.Second)
	}
}

func (c CommandMessage) from(command []byte) {
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

type CommandMessageReply struct {
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

func (cr CommandMessageReply) sendWith(socket Socket, header string) {
	for {
		isSend := cr.sendHeaderWith(socket, header)
		isSend += cr.sendCommandReplyWith(socket)
		if isSend > 0 {
			break
		}
		time.Sleep(2 * time.Second)
	}
}

func (cr CommandMessageReply) sendHeaderCommandReplyWith(socket Socket, header string) {
	for {
		isSend := socket.Send(header, FlagMore);
		if isSend > 0 {
			break
		}
		time.Sleep(2 * time.Second)
	}
}

func (cr CommandMessageReply) sendCommandReplyWith(socket Socket) {
	for {
		isSend := socket.SendBytes(encode(cr), 0);
		if isSend > 0 {
			break
		}
		time.Sleep(2 * time.Second)
	}
}

func (cr CommandMessageReply) from(commandMessage CommandMessage, reply, payload string) {
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

type CommandFunction struct {
	functions    []string
}

func (cf CommandFunction) New(functions []string) {
	cf.functions = functions
}

func (cf CommandFunction) sendWith(socket Socket) {
	for {
		isSend := socket.Send(constant.COMMAND_VALIDATION_FUNCTIONS, FlagMore);
		isSend += socket.SendBytes(encode(cf), 0);
		if isSend > 0 {
			break
		}
		time.Sleep(2 * time.Second)
	}
}

//

type CommandFunctionReply struct {
	validation bool
}

func (cfr CommandFunctionReply) New(validation bool) {
	cfr.validation = validation
}

func (cfr CommandFunctionReply) sendWith(socket Socket, header string) {
	for {
		isSend := cfr.sendHeaderWith(socket, header)
		isSend += cfr.sendCommandCommandsEventsReplyWith(socket)
		if isSend > 0 {
			break
		}
		time.Sleep(2 * time.Second)
	}
}

func (cfr CommandFunctionReply) sendHeaderWith(socket Socket, header string) {
	for {
		isSend := socket.Send(header, FlagMore);
		if isSend > 0 {
			break
		}
		time.Sleep(2 * time.Second)
	}
}

func (cfr CommandFunctionReply) sendCommandFunctionReplyWith(socket Socket) {
	for {
		isSend := socket.Send(constant.COMMAND_VALIDATION_FUNCTIONS_REPLY, FlagMore);
		isSend += socket.SendBytes(encode(ccer), 0);
		if isSend > 0 {
			break
		}
		time.Sleep(2 * time.Second)
	}
}

//

type CommandReady struct {
	// ???
}

func (cry CommandReady) New() {
}

func (cry CommandReady) sendWith(socket Socket) {
	for {
		isSend := socket.Send(constant.COMMAND_READY, FlagMore);
		isSend += socket.SendBytes(encode(cry), 0);
		if isSend > 0 {
			break
		}
		time.Sleep(2 * time.Second)
	}
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

func decodeCommand(bytesContent []byte) (commandMessage CommandMessage, commandError error) {
	err := msgpack.Decode(bytesContent, command)
	if err != nil {
		commandError = fmt.Errorf("Command %s", err)
		return
	}
	return
}	

func decodeCommandReply(bytesContent []byte) (commandReply CommandMessageReply, commandError error) {
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

func decodeCommandFunction(bytesContent []byte) (commandFunction CommandFunction, commandError error) {
	err := msgpack.Decode(bytesContent, commandFunctions)
	if err != nil {
		commandError = fmt.Errorf("CommandResponse %s", err)
		return
	}
	return
}

func decodeCommandFunctionReply(bytesContent []byte) (commandFunctionReply CommandFunctionReply, commandError error) {
	err := msgpack.Decode(bytesContent, commandFunctionReply)
	if err != nil {
		commandError = fmt.Errorf("CommandResponse %s", err)
		return
	}
	return
}