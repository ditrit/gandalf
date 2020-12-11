//Package models :
package models

import (
	"github.com/jinzhu/gorm"
)

// Connector : Connector struct.
type Connector struct {
	gorm.Model
	LogicalName string
	Secret      string
	BindAddress string
}
