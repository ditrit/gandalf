package message

import (
	"fmt"
	"time"
	"gandalfgo/constant"
	"github.com/shamaton/msgpack"
	"github.com/pebbe/zmq4"
)

type CommandMessage struct {
	SourceAggregator    string
	SourceConnector string
	SourceWorker   string
	DestinationAggregator    string
    DestinationConnector    string
    DestinationWorker string
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

func (c CommandMessage) SendWith(socket *zmq4.Socket, header string) (isSend bool) {
	for {
		isSend = c.SendHeaderWith(socket, header)
		isSend = isSend && c.SendCommandWith(socket)
		if isSend {
			return
		}
		time.Sleep(2 * time.Second)
	}
}

func (c CommandMessage) SendHeaderWith(socket *zmq4.Socket, header string) (isSend bool) {
	for {
		_, err := socket.Send(header, zmq4.SNDMORE);
		if err == nil {
			isSend = true
			return
		}
		time.Sleep(2 * time.Second)
	}
}

func (c CommandMessage) SendCommandWith(socket *zmq4.Socket) (isSend bool) {
	for {
		encoded, _ := EncodeCommandMessage(c)
		_, err := socket.SendBytes(encoded, 0);
		if err == nil {
			isSend = true
			return
		}
		time.Sleep(2 * time.Second)
	}
}

func (c CommandMessage) From(command []string) {
	c.SourceAggregator = command[0]
	c.SourceConnector = command[1]
	c.SourceWorker = command[2]
	c.DestinationAggregator = command[3]
    c.DestinationConnector = command[4]
    c.DestinationWorker = command[5]
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
	SourceAggregator    string
	SourceConnector string
	SourceWorker   string
	DestinationAggregator    string
    DestinationConnector    string
    DestinationWorker string
    tenant   string
    token    string
    context    string
    timeout string
    timestamp   string
	uuid string
	reply    string
    payload    string
}

func (cr CommandMessageReply) SendWith(socket *zmq4.Socket, header string) (isSend bool) {
	for {
		isSend = cr.SendHeaderWith(socket, header)
		isSend = isSend && cr.SendCommandReplyWith(socket)
		if isSend {
			return
		}
		time.Sleep(2 * time.Second)
	}
}

func (cr CommandMessageReply) SendHeaderWith(socket *zmq4.Socket, header string) (isSend bool) {
	for {
		_, err := socket.Send(header, zmq4.SNDMORE);
		if err == nil {
			isSend = true
			return
		}
		time.Sleep(2 * time.Second)
	}
}

func (cr CommandMessageReply) SendCommandReplyWith(socket *zmq4.Socket) (isSend bool) {
	for {
		encoded, _ := EncodeCommandMessageReply(cr)
		_, err := socket.SendBytes(encoded, 0);
		if err == nil {
			isSend = true
			return
		}
		time.Sleep(2 * time.Second)
	}
}

func (cr CommandMessageReply) From(commandMessage CommandMessage, reply, payload string) {
	cr.SourceAggregator = commandMessage.SourceAggregator
	cr.SourceConnector = commandMessage.SourceConnector
	cr.SourceWorker = commandMessage.SourceWorker
	cr.DestinationAggregator = commandMessage.DestinationAggregator
    cr.DestinationConnector = commandMessage.DestinationConnector
    cr.DestinationWorker = commandMessage.DestinationWorker
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

func (cf CommandFunction) SendWith(socket *zmq4.Socket) (isSend bool) {
	for {
		_, err := socket.Send(constant.COMMAND_VALIDATION_FUNCTIONS, zmq4.SNDMORE);
		if err == nil {
			encoded, _ := EncodeCommandFunction(cf)
			_, err = socket.SendBytes(encoded, 0);
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

func (cfr CommandFunctionReply) SendWith(socket *zmq4.Socket, header string) (isSend bool) {
	for {
		isSend = cfr.SendHeaderWith(socket, header)
		isSend = isSend && cfr.SendCommandFunctionReplyWith(socket)
		if isSend {
			return
		}
		time.Sleep(2 * time.Second)
	}
}

func (cfr CommandFunctionReply) SendHeaderWith(socket *zmq4.Socket, header string) (isSend bool) {
	for {
		_, err := socket.Send(header, zmq4.SNDMORE);
		if err == nil {
			isSend = true
			return
		}
		time.Sleep(2 * time.Second)
	}
}

func (cfr CommandFunctionReply) SendCommandFunctionReplyWith(socket *zmq4.Socket) (isSend bool) {
	for {
		_, err := socket.Send(constant.COMMAND_VALIDATION_FUNCTIONS_REPLY, zmq4.SNDMORE);
		if err == nil {
			encoded, _ := EncodeCommandFunctionReply(cfr)
			_, err = socket.SendBytes(encoded, 0);
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

func (cry CommandMessageReady) SendWith(socket *zmq4.Socket) (isSend bool) {
	for {
		_, err := socket.Send(constant.COMMAND_READY, zmq4.SNDMORE);
		if err == nil {
			encoded, _ := EncodeCommandMessageReady(cry)
			_, err = socket.SendBytes(encoded, 0);
			if err == nil {
				isSend = true
				return
			}
		}
		time.Sleep(2 * time.Second)
	}
}

//

func EncodeCommandMessage(commandMessage CommandMessage) (bytesContent []byte, commandError error) {
	bytesContent, err := msgpack.Encode(commandMessage)
	if err != nil {
		commandError = fmt.Errorf("Command %s", err)
		return
	}
	return
}

func EncodeCommandMessageReply(commandMessageReply CommandMessageReply) (bytesContent []byte, commandError error) {
	bytesContent, err := msgpack.Encode(commandMessageReply)
	if err != nil {
		commandError = fmt.Errorf("Command %s", err)
		return
	}
	return
}

func EncodeCommandMessageReady(commandMessageReady CommandMessageReady) (bytesContent []byte, commandError error) {
	bytesContent, err := msgpack.Encode(commandMessageReady)
	if err != nil {
		commandError = fmt.Errorf("Command %s", err)
		return
	}
	return
}

func EncodeCommandFunction(commandFunction CommandFunction) (bytesContent []byte, commandError error) {
	bytesContent, err := msgpack.Encode(commandFunction)
	if err != nil {
		commandError = fmt.Errorf("Command %s", err)
		return
	}
	return
}

func EncodeCommandFunctionReply(commandFunctionReply CommandFunctionReply) (bytesContent []byte, commandError error) {
	bytesContent, err := msgpack.Encode(commandFunctionReply)
	if err != nil {
		commandError = fmt.Errorf("Command %s", err)
		return
	}
	return
}

func DecodeCommandMessage(bytesContent []byte) (commandMessage CommandMessage, commandError error) {
	err := msgpack.Decode(bytesContent, commandMessage)
	if err != nil {
		commandError = fmt.Errorf("Command %s", err)
		return
	}
	return
}	

func DecodeCommandMessageReply(bytesContent []byte) (commandMessageReply CommandMessageReply, commandError error) {
	err := msgpack.Decode(bytesContent, commandMessageReply)
	if err != nil {
		commandError = fmt.Errorf("CommandResponse %s", err)
		return
	}
	return
}

func DecodeCommandMessageReady(bytesContent []byte) (commandMessageReady CommandMessageReady, commandError error) {
	err := msgpack.Decode(bytesContent, commandMessageReady)
	if err != nil {
		commandError = fmt.Errorf("CommandResponse %s", err)
		return
	}
	return
}

func DecodeCommandFunction(bytesContent []byte) (commandFunction CommandFunction, commandError error) {
	err := msgpack.Decode(bytesContent, commandFunction)
	if err != nil {
		commandError = fmt.Errorf("CommandResponse %s", err)
		return
	}
	return
}

func DecodeCommandFunctionReply(bytesContent []byte) (commandFunctionReply CommandFunctionReply, commandError error) {
	err := msgpack.Decode(bytesContent, commandFunctionReply)
	if err != nil {
		commandError = fmt.Errorf("CommandResponse %s", err)
		return
	}
	return
}