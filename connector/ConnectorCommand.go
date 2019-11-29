package connector

import (
	"github.com/zeromq/goczmq"
)

type ConnectorCommand struct {
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

func (c ConnectorCommand) init(identity, backEndSendRoutingAggregatorConnection, frontEndSendRoutingAggregatorConnection, backEndReceiveRoutingAggregatorConnection, frontEndReceiveRoutingAggregatorConnection string) {

}

func (c ConnectorCommand) close() {
}

func (c ConnectorCommand) reconnectToProxy() {

}
