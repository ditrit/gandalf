package models

import "github.com/jinzhu/gorm"

// State : State struct.
type State struct {
	gorm.Model
	Admin bool
}
