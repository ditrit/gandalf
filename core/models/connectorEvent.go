//Package models :
package models

import (
	"github.com/jinzhu/gorm"
)

// ConnectorEvent : ConnectorEvent struct.
type ConnectorEvent struct {
	gorm.Model
	Name   string
	Schema string
}
