//Package models :
package models

import (
	"github.com/jinzhu/gorm"
)

// Connector : Connector struct.
type Connector struct {
	gorm.Model
	LogicalName  string
	InstanceName string
	Secret       string
}
