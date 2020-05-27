//Package models :
package models

import (
	"github.com/jinzhu/gorm"
)

// ConnectorType : ConnectorType struct.
type ConnectorType struct {
	gorm.Model
	Name string
	//Commands []ConnectorTypeCommand
}
