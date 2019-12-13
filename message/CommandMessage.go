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
	c.timestamp = time.Now().String()
}

func (c CommandMessage) sendWith(socket goczmq.Sock, header string) (isSend bool) {
	for {
		isSend := c.sendHeaderWith(socket, header)
		isSend = isSend && c.sendCommandWith(socket)
		if isSend {
			return
		}
		time.Sleep(2 * time.Second)
	}
}

func (c CommandMessage) sendHeaderWith(socket goczmq.Sock, header string) (isSend bool) {
	for {
		err := socket.SendFrame([]byte(header), goczmq.FlagMore);
		if err == nil {
			isSend = true
			return
		}
		time.Sleep(2 * time.Second)
	}
}

func (c CommandMessage) sendCommandWith(socket goczmq.Sock) (isSend bool) {
	for {
		err := socket.SendFrame(encodeCommandMessage(c), 0);
		if err == nil {
			isSend = true
			return
		}
		time.Sleep(2 * time.Second)
	}
}

func (c CommandMessage) from(command []string) {
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

func (cr CommandMessageReply) sendWith(socket goczmq.Sock, header string) (isSend bool) {
	for {
		isSend = cr.sendHeaderWith(socket, header)
		isSend = isSend && cr.sendCommandReplyWith(socket)
		if isSend {
			return
		}
		time.Sleep(2 * time.Second)
	}
}

func (cr CommandMessageReply) sendHeaderWith(socket goczmq.Sock, header string) (isSend bool) {
	for {
		err := socket.SendFrame([]byte(header), goczmq.FlagMore);
		if err == nil {
			isSend = true
			return
		}
		time.Sleep(2 * time.Second)
	}
}

func (cr CommandMessageReply) sendCommandReplyWith(socket goczmq.Sock) (isSend bool) {
	for {
		err := socket.SendFrame(encodeCommandMessageReply(cr), 0);
		if err == nil {
			isSend = true
			return
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

func (cf CommandFunction) sendWith(socket goczmq.Sock) (isSend bool) {
	for {
		err := socket.SendFrame([]byte(constant.COMMAND_VALIDATION_FUNCTIONS), goczmq.FlagMore);
		if err == nil {
			err = socket.SendFrame(encodeCommandFunction(cf), 0);
			if err == nil {
				isSend = true
				return
			}
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

func (cfr CommandFunctionReply) sendWith(socket goczmq.Sock, header string) (isSend bool) {
	for {
		isSend = cfr.sendHeaderWith(socket, header)
		isSend = isSend && cfr.sendCommandFunctionReplyWith(socket)
		if isSend {
			return
		}
		time.Sleep(2 * time.Second)
	}
}

func (cfr CommandFunctionReply) sendHeaderWith(socket goczmq.Sock, header string) (isSend bool) {
	for {
		err := socket.SendFrame([]byte(header), goczmq.FlagMore);
		if err == nil {
			isSend = true
			return
		}
		time.Sleep(2 * time.Second)
	}
}

func (cfr CommandFunctionReply) sendCommandFunctionReplyWith(socket goczmq.Sock) (isSend bool) {
	for {
		err := socket.SendFrame([]byte(constant.COMMAND_VALIDATION_FUNCTIONS_REPLY), goczmq.FlagMore);
		if err == nil {
			err = socket.SendFrame(encodeCommandFunctionReply(cfr), 0);
			if err == nil {
				isSend = true
				return
			}
		}
		
		time.Sleep(2 * time.Second)
	}
}

//

type CommandMessageReady struct {
	// ???
}

func (cry CommandMessageReady) New() {
}

func (cry CommandMessageReady) sendWith(socket goczmq.Sock) (isSend bool) {
	for {
		err := socket.SendFrame([]byte(constant.COMMAND_READY), goczmq.FlagMore);
		if err == nil {
			err = socket.SendFrame(encodeCommandMessageReady(cry), 0);
			if err == nil {
				isSend = true
				return
			}
		}
		time.Sleep(2 * time.Second)
	}
}

//

func encodeCommandMessage(commandMessage CommandMessage) (bytesContent []byte, commandError error) {
	bytesContent, err := msgpack.Encode(commandMessage)
	if err != nil {
		commandError = fmt.Errorf("Command %s", err)
		return
	}
	return
}

func encodeCommandMessageReply(commandMessageReply CommandMessageReply) (bytesContent []byte, commandError error) {
	bytesContent, err := msgpack.Encode(commandMessageReply)
	if err != nil {
		commandError = fmt.Errorf("Command %s", err)
		return
	}
	return
}

func encodeCommandMessageReady(commandMessageReady CommandMessageReady) (bytesContent []byte, commandError error) {
	bytesContent, err := msgpack.Encode(commandMessageReady)
	if err != nil {
		commandError = fmt.Errorf("Command %s", err)
		return
	}
	return
}

func encodeCommandFunction(commandFunction CommandFunction) (bytesContent []byte, commandError error) {
	bytesContent, err := msgpack.Encode(commandFunction)
	if err != nil {
		commandError = fmt.Errorf("Command %s", err)
		return
	}
	return
}

func encodeCommandFunctionReply(commandFunctionReply CommandFunctionReply) (bytesContent []byte, commandError error) {
	bytesContent, err := msgpack.Encode(commandFunctionReply)
	if err != nil {
		commandError = fmt.Errorf("Command %s", err)
		return
	}
	return
}

func decodeCommandMessage(bytesContent []byte) (commandMessage CommandMessage, commandError error) {
	err := msgpack.Decode(bytesContent, commandMessage)
	if err != nil {
		commandError = fmt.Errorf("Command %s", err)
		return
	}
	return
}	

func decodeCommandMessageReply(bytesContent []byte) (commandMessageReply CommandMessageReply, commandError error) {
	err := msgpack.Decode(bytesContent, commandMessageReply)
	if err != nil {
		commandError = fmt.Errorf("CommandResponse %s", err)
		return
	}
	return
}

func decodeCommandMessageReady(bytesContent []byte) (commandMessageReady CommandMessageReady, commandError error) {
	err := msgpack.Decode(bytesContent, commandMessageReady)
	if err != nil {
		commandError = fmt.Errorf("CommandResponse %s", err)
		return
	}
	return
}

func decodeCommandFunction(bytesContent []byte) (commandFunction CommandFunction, commandError error) {
	err := msgpack.Decode(bytesContent, commandFunction)
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