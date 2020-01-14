package sender

import (
	"gandalf-go/message"
)

type SenderGandalf struct {
	Identity                string
	SenderCommandConnection string
	SenderEventConnection   string
	SenderCommandRoutine    *SenderCommandRoutine
	SenderEventRoutine      *SenderEventRoutine
}

func NewSenderGandalf(identity, senderCommandConnection, senderEventConnection string) (senderGandalf *SenderGandalf) {
	senderGandalf = new(SenderGandalf)
	senderGandalf.Identity = identity
	senderGandalf.SenderCommandConnection = senderCommandConnection
	senderGandalf.SenderEventConnection = senderEventConnection
	senderGandalf.SenderCommandRoutine = NewSenderCommandRoutine(identity, senderCommandConnection)
	senderGandalf.SenderEventRoutine = NewSenderEventRoutine(identity, senderEventConnection)

	return
}

func (sg SenderGandalf) SendEvent(topic, timeout, uuid, event, payload string) {
	sg.SenderEventRoutine.SendEvent(topic, timeout, uuid, event, payload)
}

func (sg SenderGandalf) SendCommand(context, timeout, uuid, connectorType, commandType, command, payload string) {
	sg.SenderCommandRoutine.SendCommand(context, timeout, uuid, connectorType, commandType, command, payload)
}

func (sg SenderGandalf) SendCommandReply(commandMessage message.CommandMessage, reply, payload string) {
	sg.SenderCommandRoutine.SendCommandReply(commandMessage, reply, payload)
}
