package aggregator

import (
	"nanomsg.org/go/mangos/v2"
)

type AggregatorCommandMangos struct {
	Context                                    mangos.Context
	BackEndSendRoutingAggregator               mangos.Socket
	BackEndSendRoutingAggregatorConnection     *string
	FrontEndReceiveRoutingAggregator           mangos.Socket
	FrontEndReceiveRoutingAggregatorConnection *string
	FrontEndSendRoutingAggregator              mangos.Socket
	FrontEndSendRoutingAggregatorConnection    string
	BackEndReceiveRoutingAggregator            mangos.Socket
	BackEndReceiveRoutingAggregatorConnection  string
	Identity                                   string
}

func (a AggregatorCommandMangos) init(identity, frontEndSendRoutingAggregatorConnection, backEndReceiveRoutingAggregatorConnection string, backEndSendRoutingAggregatorConnection, frontEndReceiveRoutingAggregatorConnection *string) {

}

func (a AggregatorCommandMangos) close() {
}

func (a AggregatorCommandMangos) reconnectToProxy() {

}
