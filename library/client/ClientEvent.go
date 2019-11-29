package client

import (
	"nanomsg.org/go/mangos/v2"
)

type PublisherMangos struct {
	Context                     mangos.Context
	BackEndPublisher            mangos.Socket
	BackEndPublisherConnections string
	BackEndPublisherConnection  *string
	Identity                    string
	Responses                   *mangos.Message
}

func (p PublisherMangos) init(identity, backEndClientConnection string) {

}

func (p PublisherMangos) sendEvent(topic, timeout, event, payload string) {

}

func (p PublisherMangos) close() {
}
