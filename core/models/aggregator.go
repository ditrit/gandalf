//Package models :
package models

import (
	"github.com/jinzhu/gorm"
)

// Aggregator : Aggregator struct.
type Aggregator struct {
	gorm.Model
	LogicalName string
	Secret      string
	BindAddress string
}
