package connector

import (
	"github.com/zeromq/goczmq"
)

type ConnectorEvent struct {
	context                                   goczmq.Context
	backEndSendRoutingConnector               goczmq.Socket
	backEndSendRoutingConnectorConnection     string
	frontEndReceiveRoutingConnector           goczmq.Socket
	frontEndReceiveRoutingConnectorConnection string
	frontEndSendRoutingConnector              goczmq.Socket
	frontEndSendRoutingConnectorConnection    string
	backEndReceiveRoutingConnector            goczmq.Socket
	backEndReceiveRoutingConnectorConnection  string
	identity                                  string
}

func (c ConnectorEvent) init(identity, backEndSendRoutingAggregatorConnection, frontEndSendRoutingAggregatorConnection, backEndReceiveRoutingAggregatorConnection, frontEndReceiveRoutingAggregatorConnection string) {

}

func (c ConnectorEvent) close() {
}

func (c ConnectorEvent) reconnectToProxy() {

}
