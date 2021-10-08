package models

import "github.com/jinzhu/gorm"

type ConnectorProduct struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
}
