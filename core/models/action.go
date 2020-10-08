package models

import (
	"github.com/jinzhu/gorm"
)

// Action : Action struct.
type Action struct {
	gorm.Model
	Name string
}
