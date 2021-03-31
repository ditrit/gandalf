//Package models :
package models

import (
	"github.com/jinzhu/gorm"
)

// Tenant : Tenant struct.
type Type struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
}
