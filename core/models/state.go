package models

import "github.com/jinzhu/gorm"

type State struct {
	gorm.Model
	Admin bool
}
