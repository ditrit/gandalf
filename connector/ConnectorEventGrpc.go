package connector

import (
	"golang.org/x/net/context"
)

type ConnectorEventGrpc struct {
}

func (ceg *ConnectorEventGrpc) SendEventMessage(ctx context.Context, in *EventMessage) (*Empty, error) {

}

func (ceg *ConnectorEventGrpc) WaitEventMessage(ctx context.Context, in *EventMessageRequest) (*EventMessage, error) {

}
