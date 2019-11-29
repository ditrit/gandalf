package connector

import (
	"nanomsg.org/go/mangos/v2"
)

type ConnectorEventMangos struct {
	Context                                   mangos.Context
	BackEndSendRoutingConnector               mangos.Socket
	BackEndSendRoutingConnectorConnection     string
	FrontEndReceiveRoutingConnector           mangos.Socket
	FrontEndReceiveRoutingConnectorConnection string
	FrontEndSendRoutingConnector              mangos.Socket
	FrontEndSendRoutingConnectorConnection    string
	BackEndReceiveRoutingConnector            mangos.Socket
	BackEndReceiveRoutingConnectorConnection  string
	Identity                                  string
}

func (c ConnectorEventMangos) init(identity, backEndSendRoutingAggregatorConnection, frontEndSendRoutingAggregatorConnection, backEndReceiveRoutingAggregatorConnection, frontEndReceiveRoutingAggregatorConnection string) {

}

func (c ConnectorEventMangos) close() {
}

func (c ConnectorEventMangos) reconnectToProxy() {

}
