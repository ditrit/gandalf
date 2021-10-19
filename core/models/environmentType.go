package models

import "github.com/jinzhu/gorm"

type EnvironmentType struct {
	gorm.Model
	Name             string `gorm:"not null"`
	ShortDescription string
	Description      string
	Logo             string
}
