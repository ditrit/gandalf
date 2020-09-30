package models

import (
	"github.com/jinzhu/gorm"
)

// ResourceType : ResourceType struct.
type ResourceType struct {
	gorm.Model
	Name string `gorm:"unique"`
}
