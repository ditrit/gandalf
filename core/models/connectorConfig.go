//Package models :
package models

import (
	"github.com/jinzhu/gorm"
)

// ConnectorConfig : ConnectorConfig struct.
type ConnectorConfig struct {
	gorm.Model
	Name                  string
	ConnectorTypeID       uint
	ConnectorType         ConnectorType
	Version               int
	ConnectorProductID    uint
	ConnectorProduct      ConnectorProduct
	ConnectorTypeCommands []ConnectorTypeCommand `gorm:"many2many:config_commands;"`
	ConnectorTypeEvents   []ConnectorTypeEvent   `gorm:"many2many:config_events;"`
	ConnectorTypeKeys     string
	ProductKeys           string
	VersionKeys           string
}
