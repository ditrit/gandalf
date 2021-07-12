package models

import (
	"time"

	cmsg "github.com/ditrit/gandalf/core/msg"

	"github.com/jinzhu/gorm"
)

type Heartbeat struct {
	gorm.Model
	LogicalName string
	Type        string
	Address     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// FromShosetCommand : Shoset command to core command.
func FromShosetHeartbeat(mheartbeat cmsg.Heartbeat) (heartbeat Heartbeat) {
	heartbeat.LogicalName = mheartbeat.GetContext()["logicalName"].(string)
	heartbeat.Type = mheartbeat.GetContext()["componentType"].(string)
	heartbeat.Address = mheartbeat.GetContext()["bindAddress"].(string)

	return
}
