//Package models :
package models

import (
	"github.com/jinzhu/gorm"
)

// ConnectorTypeEvent : ConnectorTypeEvent struct.
type ConnectorTypeEvent struct {
	gorm.Model
	Name   string
	Schema string
}
