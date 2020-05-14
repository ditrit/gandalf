//Package models :
package models

import (
	"github.com/jinzhu/gorm"
)

// ConnectorTypeCommand : ConnectorTypeCommand struct.
type ConnectorTypeCommand struct {
	gorm.Model
	Name            string
	ConnectorTypeID int
	ConnectorType   ConnectorType
}
