package library

import (
	"nanomsg.org/go/mangos/v2"
)

type ClientMangos struct {
	Context                  mangos.Context
	BackEndClient            mangos.Socket
	BackEndClientConnections *string
	BackEndClientConnection  string
	Identity                 string
	Responses                *mangos.Message
}

func (c ClientMangos) init(identity, backEndClientConnection string) {

}

func (c ClientMangos) initList(identity string, backEndClientConnections *string) {

}

func (c ClientMangos) sendCommandSync(context, timeout, uuid, commandtype, command, payload string) mangos.Message {
}

func (c ClientMangos) getCommandResultSync() mangos.Message {

}

func (c ClientMangos) getCommandResultAsync() mangos.Message {
}

func (c ClientMangos) close() {
}
