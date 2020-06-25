//Package models :
package models

import (
	"github.com/jinzhu/gorm"
)

// ConnectorConfigProduct : ConnectorConfigProduct struct.
type ConnectorProduct struct {
	gorm.Model
	Name    string
	Version string
}
