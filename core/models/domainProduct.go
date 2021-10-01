package models

import "github.com/jinzhu/gorm"

type DomainProduct struct {
	gorm.Model
	Name             string `gorm:"unique;not null"`
	ShortDescription string
	Description      string
	Logo             string
	DomainID         uint
	Domain           Domain
}
