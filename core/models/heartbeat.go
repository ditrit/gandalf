package models

import (
	cmsg "github.com/ditrit/gandalf/core/msg"
)

type Heartbeat struct {
	Model
	LogicalName string
	Type        string
	Address     string
}

// FromShosetCommand : Shoset command to core command.
func FromShosetHeartbeat(mheartbeat cmsg.Heartbeat) (heartbeat Heartbeat) {
	heartbeat.LogicalName = mheartbeat.GetContext()["logicalName"].(string)
	heartbeat.Type = mheartbeat.GetContext()["componentType"].(string)
	heartbeat.Address = mheartbeat.GetContext()["bindAddress"].(string)

	return
}
