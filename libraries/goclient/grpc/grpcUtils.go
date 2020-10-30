package grpc

import (
	"strconv"

	"github.com/ditrit/shoset/msg"
)

//CommandMessageFromGrpc :
func CommandFromGrpc(commandMessage *CommandMessage) (c msg.Command) {
	//c.SourceAggregator = commandMessage.GetSourceAggregator()
	//c.SourceConnector = commandMessage.GetSourceConnector()
	//c.SourceWorker = commandMessage.GetSourceWorker()
	//c.DestinationAggregator = commandMessage.GetDestinationAggregator()
	//c.DestinationConnector = commandMessage.GetDestinationConnector()
	//c.DestinationWorker = commandMessage.GetDestinationWorker()
	c.Tenant = commandMessage.GetTenant()
	c.Token = commandMessage.GetToken()
	//c.Context = commandMessage.GetContext()
	c.Timeout, _ = strconv.ParseInt(commandMessage.GetTimeout(), 10, 64)
	c.Timestamp, _ = strconv.ParseInt(commandMessage.GetTimestamp(), 10, 64)
	c.Major = int8(commandMessage.GetMajor())
	//c.Minor = int8(commandMessage.GetMinor())
	c.UUID = commandMessage.GetUUID()
	c.Context = make(map[string]interface{})
	c.Context["ConnectorType"] = commandMessage.GetConnectorType()
	//c.CommandType = commandMessage.GetCommandType()
	c.Command = commandMessage.GetCommand()
	c.Payload = commandMessage.GetPayload()

	return
}

//CommandMessageToGrpc :
func CommandToGrpc(c msg.Command) (commandMessage *CommandMessage) {
	commandMessage = new(CommandMessage)
	//commandMessage.SourceAggregator = c.SourceAggregator
	//commandMessage.SourceConnector = c.SourceConnector
	//commandMessage.SourceWorker = c.SourceWorker
	//commandMessage.DestinationAggregator = c.DestinationAggregator
	//commandMessage.DestinationConnector = c.DestinationConnector
	//commandMessage.DestinationWorker = c.DestinationWorker
	commandMessage.Tenant = c.Tenant
	commandMessage.Token = c.Token
	//commandMessage.Context = c.Context
	commandMessage.Timeout = strconv.Itoa(int(c.Timeout))
	commandMessage.Timestamp = strconv.Itoa(int(c.Timestamp))
	commandMessage.Major = int64(c.Major)
	//commandMessage.Minor = int64(c.Minor)
	commandMessage.UUID = c.UUID
	//commandMessage.ConnectorType = c.ConnectorType
	//commandMessage.CommandType = c.CommandType
	commandMessage.Command = c.Command
	commandMessage.Payload = c.Payload

	return
}

//EventMessageFromGrpc :
func EventFromGrpc(eventMessage *EventMessage) (e msg.Event) {
	e.Tenant = eventMessage.GetTenant()
	e.Token = eventMessage.GetToken()
	e.Timeout, _ = strconv.ParseInt(eventMessage.GetTimeout(), 10, 64)
	e.Timestamp, _ = strconv.ParseInt(eventMessage.GetTimestamp(), 10, 64)
	e.UUID = eventMessage.GetUUID()
	e.Topic = eventMessage.GetTopic()
	e.Event = eventMessage.GetEvent()
	e.Payload = eventMessage.GetPayload()
	e.ReferenceUUID = eventMessage.GetReferenceUUID()

	return
}

//EventMessageToGrpc :
func EventToGrpc(e msg.Event) (eventMessage *EventMessage) {
	eventMessage = new(EventMessage)
	eventMessage.Tenant = e.Tenant
	eventMessage.Token = e.Token
	eventMessage.Timeout = strconv.Itoa(int(e.Timeout))
	eventMessage.Timestamp = strconv.Itoa(int(e.Timestamp))
	eventMessage.UUID = e.UUID
	eventMessage.Topic = e.Topic
	eventMessage.Event = e.Event
	eventMessage.Payload = e.Payload
	eventMessage.Payload = e.Payload
	eventMessage.ReferenceUUID = e.GetReferenceUUID()

	return
}
