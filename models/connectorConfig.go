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
	ConnectorTypeCommands []ConnectorTypeCommand `gorm:"many2many:config_commands;"`
}
