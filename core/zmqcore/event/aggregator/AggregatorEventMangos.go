package aggregator

import (
	"nanomsg.org/go/mangos/v2"
)

type AggregatorEventMangos struct {
	Context                                    mangos.Context
	BackEndSendRoutingAggregator               mangos.Socket
	BackEndSendRoutingAggregatorConnection     string
	FrontEndReceiveRoutingAggregator           mangos.Socket
	FrontEndReceiveRoutingAggregatorConnection *string
	FrontEndSendRoutingAggregator              mangos.Socket
	FrontEndSendRoutingAggregatorConnection    string
	BackEndReceiveRoutingAggregator            mangos.Socket
	BackEndReceiveRoutingAggregatorConnection  string
	Identity                                   string
}

func (a AggregatorEventMangos) init(identity, backEndSendRoutingAggregatorConnection, frontEndSendRoutingAggregatorConnection, backEndReceiveRoutingAggregatorConnection string, frontEndReceiveRoutingAggregatorConnection *string) {

}

func (a AggregatorEventMangos) close() {
}

func (a AggregatorEventMangos) reconnectToProxy() {

}
