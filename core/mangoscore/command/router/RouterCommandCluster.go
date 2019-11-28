package connector

import (
	"nanomsg.org/go/mangos/v2"
)

type RouterCommandCluster interface {
	getCommandTarget(command mangos.Message) mangos.Message
}
