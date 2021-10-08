package models

import "github.com/jinzhu/gorm"

type Library struct {
	gorm.Model
	Name             string `gorm:"not null"`
	ShortDescription string
	Description      string
	Logo             string
	DomainID         uint
	Domain           Domain
}
