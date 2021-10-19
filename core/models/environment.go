package models

import "github.com/jinzhu/gorm"

type Environment struct {
	gorm.Model
	Name              string `gorm:"not null"`
	EnvironmentTypeID uint
	EnvironmentType   EnvironmentType
	ShortDescription  string
	Description       string
	Logo              string
}
