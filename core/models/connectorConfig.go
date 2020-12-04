//Package models :
package models

import (
	"github.com/jinzhu/gorm"
)

// ConnectorConfig : ConnectorConfig struct.
type ConnectorConfig struct {
	gorm.Model
	Name               string `gorm:"unique"`
	ConnectorTypeID    uint
	ConnectorType      ConnectorType
	Major              int8
	ConnectorProductID uint
	ConnectorProduct   ConnectorProduct
	ConnectorCommands  []Object `gorm:"many2many:config_commands;"`
	ConnectorEvents    []Object `gorm:"many2many:config_events;"`
	//Actions            []Action           `gorm:"many2many:config_actions;"`
	Resources         []Object `gorm:"many2many:config_resources;"`
	ConnectorTypeKeys string
	ProductKeys       string
	VersionMajorKeys  string
	VersionMinorKeys  string
}
