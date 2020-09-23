package models

import (
	"github.com/jinzhu/gorm"
)

// Action : Action struct.
type Role struct {
	gorm.Model
	Name string `gorm:"unique"`
}
