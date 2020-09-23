//Package models :
package models

import (
	"github.com/jinzhu/gorm"
)

// Tenant : Tenant struct.
type Tenant struct {
	gorm.Model
	Name string `gorm:"unique"`
}
