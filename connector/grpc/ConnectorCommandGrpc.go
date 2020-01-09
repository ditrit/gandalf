package connector

import (
	"golang.org/x/net/context"
)

type ConnectorCommandGrpc struct {
}

func (ccg *ConnectorCommandGrpc) SendCommandMessage(ctx context.Context, in *CommandMessage) (*CommandMessageUUID, error) {
	
}

func (ccg *ConnectorCommandGrpc) SendCommandMessage(ctx context.Context, in *CommandMessage) (*CommandMessageUUID, error) {
	
}

func (ccg *ConnectorCommandGrpc) SendCommandMessageReply(ctx context.Context, in *CommandMessageReply) (*Empty, error) {
	
}

func (ccg *ConnectorCommandGrpc) WaitCommandMessage(ctx context.Context, in *CommandMessageRequest) (*CommandMessage, error) {
	
}

func (ccg *ConnectorCommandGrpc) WaitCommandMessageReply(ctx context.Context, in *CommandMessageUUID) (*CommandMessageReply, error) {

}