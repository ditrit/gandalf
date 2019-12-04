package aggregator

import (
	zmq "github.com/zeromq/goczmq"
)

type AggregatorCommand struct {
    aggregatorCommandSendW2C           zmq.Sock
    aggregatorCommandReceiveC2WConnection string
    aggregatorEventReceiveC2W             zmq.Sock
    aggregatorEventReceiveC2WConnection   string
    aggregatorCommandReceiveC2W           zmq.Sock
    aggregatorCommandReceiveC2WConnection string
    aggregatorEventReceiveC2W             zmq.Sock
    aggregatorEventReceiveC2WConnection   string
	Identity                                   string
}

func (a AggregatorCommandMangos) init(identity, frontEndSendRoutingAggregatorConnection, backEndReceiveRoutingAggregatorConnection string, backEndSendRoutingAggregatorConnection, frontEndReceiveRoutingAggregatorConnection *string) {

}

func (a AggregatorCommandMangos) close() {
}

func (a AggregatorCommandMangos) reconnectToProxy() {

}
