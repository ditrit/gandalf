//Package models :
package models

import (
	"github.com/jinzhu/gorm"
)

// ConnectorConfig : ConnectorConfig struct.
type ConnectorConfig struct {
	gorm.Model
	Name string
}
