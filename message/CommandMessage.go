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
    Tenant   string
    Token    string
    Context    string
    Timeout string
    Timestamp   string
    Major    string
    Minor    string
	Uuid string
	ConnectorType string
    CommandType   string
    Command    string
    Payload    string
}

func NewCommandMessage(context, timeout, uuid, connectorType, commandType, command, payload string) (commandMessage *CommandMessage) {
	commandMessage = new(CommandMessage)

	commandMessage.Context = context
	commandMessage.Timeout = timeout
	commandMessage.Uuid = uuid
	commandMessage.ConnectorType = connectorType
	commandMessage.CommandType = commandType
	commandMessage.Command = command
	commandMessage.Payload = payload
	commandMessage.Timestamp = time.Now().String()

	return
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
    c.Tenant = command[6]
    c.Token = command[7]
    c.Context = command[8]
    c.Timeout = command[9]
    c.Timestamp = command[10]
    c.Major = command[11]
    c.Minor = command[12]
	c.Uuid = command[13]
	c.ConnectorType = command[14]
    c.CommandType = command[15]
    c.Command = command[16]
    c.Payload = command[17]
}

//

type CommandMessageReply struct {
	SourceAggregator    string
	SourceConnector string
	SourceWorker   string
	DestinationAggregator    string
    DestinationConnector    string
    DestinationWorker string
    Tenant   string
    Token    string
    Context    string
    Timeout string
    Timestamp   string
	Uuid string
	Reply    string
    Payload    string
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
    cr.Tenant = commandMessage.Tenant
    cr.Token = commandMessage.Token
    cr.Context = commandMessage.Context
    cr.Timeout = commandMessage.Timeout
    cr.Timestamp = commandMessage.Timestamp
	cr.Uuid = commandMessage.Uuid
	cr.Reply = reply
    cr.Payload = payload
}

//

type CommandFunction struct {
	Functions    []string
}

func NewCommandFunction(functions []string) (commandFunction *CommandFunction) {
	commandFunction = new(CommandFunction)
	commandFunction.Functions = functions

	return
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
	Validation bool
}

func NewCommandFunctionReply(validation bool) (commandFunctionReply *CommandFunctionReply) {
	commandFunctionReply = new(CommandFunctionReply) 
	commandFunctionReply.Validation = validation

	return
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

func NewCommandMessageReady() (commandMessageReady *CommandMessageReady) {
	commandMessageReady = new(CommandMessageReady)

	return
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
		commandError = fmt.Errorf("command %s", err)
		return
	}
	return
}

func EncodeCommandMessageReply(commandMessageReply CommandMessageReply) (bytesContent []byte, commandError error) {
	bytesContent, err := msgpack.Encode(commandMessageReply)
	if err != nil {
		commandError = fmt.Errorf("command %s", err)
		return
	}
	return
}

func EncodeCommandMessageReady(commandMessageReady CommandMessageReady) (bytesContent []byte, commandError error) {
	bytesContent, err := msgpack.Encode(commandMessageReady)
	if err != nil {
		commandError = fmt.Errorf("command %s", err)
		return
	}
	return
}

func EncodeCommandFunction(commandFunction CommandFunction) (bytesContent []byte, commandError error) {
	bytesContent, err := msgpack.Encode(commandFunction)
	if err != nil {
		commandError = fmt.Errorf("command %s", err)
		return
	}
	return
}

func EncodeCommandFunctionReply(commandFunctionReply CommandFunctionReply) (bytesContent []byte, commandError error) {
	bytesContent, err := msgpack.Encode(commandFunctionReply)
	if err != nil {
		commandError = fmt.Errorf("command %s", err)
		return
	}
	return
}

func DecodeCommandMessage(bytesContent []byte) (commandMessage CommandMessage, commandError error) {
	err := msgpack.Decode(bytesContent, commandMessage)
	if err != nil {
		commandError = fmt.Errorf("command %s", err)
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