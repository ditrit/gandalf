package models

import (
	"github.com/jinzhu/gorm"
)

// Resource : Resource struct.
type Resource struct {
	gorm.Model
	Name     string `gorm:"unique"`
	DomainID uint
	Domain   Domain
}
