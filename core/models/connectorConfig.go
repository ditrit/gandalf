//Package models :
package models

import (
	"github.com/jinzhu/gorm"
)

// ConnectorConfig : ConnectorConfig struct.
type ConnectorConfig struct {
	gorm.Model
	Name               string
	ConnectorTypeID    uint
	ConnectorType      ConnectorType
	Version            int
	ConnectorProductID uint
	ConnectorProduct   ConnectorProduct
	ConnectorCommands  []ConnectorCommand `gorm:"many2many:config_commands;"`
	ConnectorEvents    []ConnectorEvent   `gorm:"many2many:config_events;"`
	Actions            []Action           `gorm:"many2many:config_actions;"`
	//Resources          []Resource         `gorm:"many2many:config_resources;"`
	ConnectorTypeKeys string
	ProductKeys       string
	VersionKeys       string
}
