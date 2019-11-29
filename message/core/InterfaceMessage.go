package message-core.core

import (
	"nanomsg.org/go/mangos/v2"
)

type InterfaceMessage struct {
    sendWith(socket mangos.Socket, routingInfo string)
}
