package message

import (
	"fmt"
	"gandalf-go/constant"
	"time"

	pb "gandalf-go/grpc"

	"github.com/pebbe/zmq4"
	"github.com/shamaton/msgpack"
)

type CommandMessage struct {
	SourceAggregator      string
	SourceConnector       string
	SourceWorker          string
	DestinationAggregator string
	DestinationConnector  string
	DestinationWorker     string
	Tenant                string
	Token                 string
	Context               string
	Timeout               string
	Timestamp             string
	Major                 string
	Minor                 string
	Uuid                  string
	ConnectorType         string
	CommandType           string
	Command               string
	Payload               string
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

func (c CommandMessage) GetUUID() string {
	return c.Uuid
}

func (c CommandMessage) GetTimeout() string {
	return c.Timeout
}

func (c CommandMessage) SendWith(socket *zmq4.Socket, header string) (isSend bool) {
	for {
		isSend = c.SendHeaderWith(socket, header)
		isSend = isSend && c.SendMessageWith(socket)
		fmt.Println(isSend)
		if isSend {
			return
		}
		time.Sleep(2 * time.Second)
	}
}

func (c CommandMessage) SendHeaderWith(socket *zmq4.Socket, header string) (isSend bool) {
	for {
		_, err := socket.Send(header, zmq4.SNDMORE)
		if err == nil {
			isSend = true
			return
		}
		time.Sleep(2 * time.Second)
	}
}

func (c CommandMessage) SendMessageWith(socket *zmq4.Socket) (isSend bool) {
	for {
		_, err := socket.Send(constant.COMMAND_MESSAGE, zmq4.SNDMORE)
		encoded, _ := EncodeCommandMessage(c)
		_, err = socket.SendBytes(encoded, 0)
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

func (c CommandMessage) FromGrpc(commandMessage pb.CommandMessage) {
	c.SourceAggregator = commandMessage.GetSourceAggregator()
	c.SourceConnector = commandMessage.GetSourceConnector()
	c.SourceWorker = commandMessage.GetSourceWorker()
	c.DestinationAggregator = commandMessage.GetDestinationAggregator()
	c.DestinationConnector = commandMessage.GetDestinationConnector()
	c.DestinationWorker = commandMessage.GetDestinationWorker()
	c.Tenant = commandMessage.GetTenant()
	c.Token = commandMessage.GetToken()
	c.Context = commandMessage.GetContext()
	c.Timeout = commandMessage.GetTimeout()
	c.Timestamp = commandMessage.GetTimestamp()
	c.Major = commandMessage.GetMajor()
	c.Minor = commandMessage.GetMinor()
	c.Uuid = commandMessage.GetUuid()
	c.ConnectorType = commandMessage.GetConnectorType()
	c.CommandType = commandMessage.GetCommandType()
	c.Command = commandMessage.GetCommand()
	c.Payload = commandMessage.GetPayload()
}

func (c CommandMessage) ToGrpc() (commandMessage pb.CommandMessage) {
	commandMessage.SourceAggregator = c.SourceAggregator
	commandMessage.SourceConnector = c.SourceConnector
	commandMessage.SourceWorker = c.SourceWorker
	commandMessage.DestinationAggregator = c.DestinationAggregator
	commandMessage.DestinationConnector = c.DestinationConnector
	commandMessage.DestinationWorker = c.DestinationWorker
	commandMessage.Tenant = c.Tenant
	commandMessage.Token = c.Token
	commandMessage.Context = c.Context
	commandMessage.Timeout = c.Timeout
	commandMessage.Timestamp = c.Timestamp
	commandMessage.Major = c.Major
	commandMessage.Minor = c.Minor
	commandMessage.Uuid = c.Uuid
	commandMessage.ConnectorType = c.ConnectorType
	commandMessage.CommandType = c.CommandType
	commandMessage.Command = c.Command
	commandMessage.Payload = c.Payload

	return
}

//

type CommandMessageReply struct {
	SourceAggregator      string
	SourceConnector       string
	SourceWorker          string
	DestinationAggregator string
	DestinationConnector  string
	DestinationWorker     string
	Tenant                string
	Token                 string
	Context               string
	Timeout               string
	Timestamp             string
	Uuid                  string
	Reply                 string
	Payload               string
}

func (cr CommandMessageReply) GetUUID() string {
	return cr.Uuid
}

func (cr CommandMessageReply) GetTimeout() string {
	return cr.Timeout
}

func (cr CommandMessageReply) SendWith(socket *zmq4.Socket, header string) (isSend bool) {
	for {
		isSend = cr.SendHeaderWith(socket, header)
		isSend = isSend && cr.SendMessageWith(socket)
		if isSend {
			return
		}
		time.Sleep(2 * time.Second)
	}
}

func (cr CommandMessageReply) SendHeaderWith(socket *zmq4.Socket, header string) (isSend bool) {
	for {
		_, err := socket.Send(header, zmq4.SNDMORE)
		if err == nil {
			isSend = true
			return
		}
		time.Sleep(2 * time.Second)
	}
}

func (cr CommandMessageReply) SendMessageWith(socket *zmq4.Socket) (isSend bool) {
	for {
		_, err := socket.Send(constant.COMMAND_MESSAGE_REPLY, zmq4.SNDMORE)
		encoded, _ := EncodeCommandMessageReply(cr)
		_, err = socket.SendBytes(encoded, 0)
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

func (cr CommandMessageReply) FromGrpc(commandMessageReply pb.CommandMessageReply) {
	cr.SourceAggregator = commandMessageReply.GetSourceAggregator()
	cr.SourceConnector = commandMessageReply.GetSourceConnector()
	cr.SourceWorker = commandMessageReply.GetSourceWorker()
	cr.DestinationAggregator = commandMessageReply.GetDestinationAggregator()
	cr.DestinationConnector = commandMessageReply.GetDestinationConnector()
	cr.DestinationWorker = commandMessageReply.GetDestinationWorker()
	cr.Tenant = commandMessageReply.GetTenant()
	cr.Token = commandMessageReply.GetToken()
	cr.Context = commandMessageReply.GetContext()
	cr.Timeout = commandMessageReply.GetTimeout()
	cr.Timestamp = commandMessageReply.GetTimestamp()
	cr.Uuid = commandMessageReply.GetUuid()
	cr.Reply = commandMessageReply.GetReply()
	cr.Payload = commandMessageReply.GetPayload()
}

func (cr CommandMessageReply) ToGrpc() (commandMessageReply pb.CommandMessageReply) {
	commandMessageReply.SourceAggregator = cr.SourceAggregator
	commandMessageReply.SourceConnector = cr.SourceConnector
	commandMessageReply.SourceWorker = cr.SourceWorker
	commandMessageReply.DestinationAggregator = cr.DestinationAggregator
	commandMessageReply.DestinationConnector = cr.DestinationConnector
	commandMessageReply.DestinationWorker = cr.DestinationWorker
	commandMessageReply.Tenant = cr.Tenant
	commandMessageReply.Token = cr.Token
	commandMessageReply.Context = cr.Context
	commandMessageReply.Timeout = cr.Timeout
	commandMessageReply.Timestamp = cr.Timestamp
	commandMessageReply.Uuid = cr.Uuid
	commandMessageReply.Reply = cr.Reply
	commandMessageReply.Payload = cr.Payload
	return
}

//

type CommandFunction struct {
	Functions []string
}

func NewCommandFunction(functions []string) (commandFunction *CommandFunction) {
	commandFunction = new(CommandFunction)
	commandFunction.Functions = functions

	return
}

func (cf CommandFunction) SendWith(socket *zmq4.Socket) (isSend bool) {
	for {
		_, err := socket.Send(constant.COMMAND_VALIDATION_FUNCTIONS, zmq4.SNDMORE)
		if err == nil {
			encoded, _ := EncodeCommandFunction(cf)
			_, err = socket.SendBytes(encoded, 0)
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
		isSend = isSend && cfr.SendMessageWith(socket)
		if isSend {
			return
		}
		time.Sleep(2 * time.Second)
	}
}

func (cfr CommandFunctionReply) SendHeaderWith(socket *zmq4.Socket, header string) (isSend bool) {
	for {
		_, err := socket.Send(header, zmq4.SNDMORE)
		if err == nil {
			isSend = true
			return
		}
		time.Sleep(2 * time.Second)
	}
}

func (cfr CommandFunctionReply) SendMessageWith(socket *zmq4.Socket) (isSend bool) {
	for {
		_, err := socket.Send(constant.COMMAND_VALIDATION_FUNCTIONS_REPLY, zmq4.SNDMORE)
		if err == nil {
			encoded, _ := EncodeCommandFunctionReply(cfr)
			_, err = socket.SendBytes(encoded, 0)
			if err == nil {
				isSend = true
				return
			}
		}
		time.Sleep(2 * time.Second)
	}
}

type CommandMessageWait struct {
	WorkerSource string
	Value        string
	CommandType  string
}

func NewCommandMessageWait(workerSource, value, commandType string) (commandMessageWait *CommandMessageWait) {
	commandMessageWait = new(CommandMessageWait)
	commandMessageWait.WorkerSource = workerSource
	commandMessageWait.CommandType = commandType
	commandMessageWait.Value = value

	return
}

func (cmw CommandMessageWait) FromGrpc(commandType string, commandMessageWait pb.CommandMessageWait) {
	cmw.WorkerSource = commandMessageWait.GetWorkerSource()
	cmw.CommandType = commandType
	cmw.Value = commandMessageWait.GetValue()
}

func (cmw CommandMessageWait) SendWith(socket *zmq4.Socket) (isSend bool) {
	for {
		_, err := socket.Send(constant.COMMAND_WAIT, zmq4.SNDMORE)
		if err == nil {
			encoded, _ := EncodeCommandMessageWait(cmw)
			_, err = socket.SendBytes(encoded, 0)
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
		_, err := socket.Send(constant.COMMAND_READY, zmq4.SNDMORE)
		if err == nil {
			encoded, _ := EncodeCommandMessageReady(cry)
			_, err = socket.SendBytes(encoded, 0)
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

func EncodeCommandMessageWait(commandMessageWait CommandMessageWait) (bytesContent []byte, commandError error) {
	bytesContent, err := msgpack.Encode(commandMessageWait)
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
	err := msgpack.Decode(bytesContent, &commandMessage)
	if err != nil {
		commandError = fmt.Errorf("command %s", err)
		return
	}
	return
}

func DecodeCommandMessageReply(bytesContent []byte) (commandMessageReply CommandMessageReply, commandError error) {
	err := msgpack.Decode(bytesContent, &commandMessageReply)
	if err != nil {
		commandError = fmt.Errorf("CommandResponse %s", err)
		return
	}
	return
}

func DecodeCommandMessageReady(bytesContent []byte) (commandMessageReady CommandMessageReady, commandError error) {
	err := msgpack.Decode(bytesContent, &commandMessageReady)
	if err != nil {
		commandError = fmt.Errorf("CommandResponse %s", err)
		return
	}
	return
}

func DecodeCommandMessageWait(bytesContent []byte) (commandMessageWait CommandMessageWait, commandError error) {
	err := msgpack.Decode(bytesContent, &commandMessageWait)
	if err != nil {
		commandError = fmt.Errorf("CommandResponse %s", err)
		return
	}
	return
}

func DecodeCommandFunction(bytesContent []byte) (commandFunction CommandFunction, commandError error) {
	err := msgpack.Decode(bytesContent, &commandFunction)
	if err != nil {
		commandError = fmt.Errorf("CommandResponse %s", err)
		return
	}
	return
}

func DecodeCommandFunctionReply(bytesContent []byte) (commandFunctionReply CommandFunctionReply, commandError error) {
	err := msgpack.Decode(bytesContent, &commandFunctionReply)
	if err != nil {
		commandError = fmt.Errorf("CommandResponse %s", err)
		return
	}
	return
}
