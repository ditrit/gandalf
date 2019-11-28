package aggregator

import (
	"nanomsg.org/go/mangos/v2"
)

type ThreadAggregatorCommandMangos struct {
	AggregatorCommandMangos                    AggregatorCommandMangos
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
