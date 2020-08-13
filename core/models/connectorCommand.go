//Package models :
package models

import (
	"github.com/jinzhu/gorm"
)

// ConnectorCommand : ConnectorCommand struct.
type ConnectorCommand struct {
	gorm.Model
	Name   string
	Schema string
}
