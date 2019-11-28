package connector

import (
	"nanomsg.org/go/mangos/v2"
)

type ConnectorCommandMangos struct {
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

func (c ConnectorCommandMangos) init(identity, backEndSendRoutingAggregatorConnection, frontEndSendRoutingAggregatorConnection, backEndReceiveRoutingAggregatorConnection, frontEndReceiveRoutingAggregatorConnection string) {

}

func (c ConnectorCommandMangos) close() {
}

func (c ConnectorCommandMangos) reconnectToProxy() {

}
