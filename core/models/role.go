package models

import (
	"github.com/jinzhu/gorm"
)

// Role : Role struct.
type Role struct {
	gorm.Model
	Name string `gorm:"unique"`
}
