//Package models :
package models

import (
	"github.com/jinzhu/gorm"
)

// ConnectorProduct : ConnectorProduct struct.
type ConnectorProduct struct {
	gorm.Model
	Name            string
	Version         string
	ConnectorTypeID uint
	ConnectorType   ConnectorType
}
