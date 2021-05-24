//Package models :
package models

import (
	"github.com/jinzhu/gorm"
)

// Application : Application struct.
type Application struct {
	gorm.Model
	Name            string `gorm:"unique"`
	ConnectorID     uint
	Connector       LogicalComponent
	ConnectorTypeID uint
	ConnectorType   ConnectorType
}
