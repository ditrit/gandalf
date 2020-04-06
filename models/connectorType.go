package models

import (
	"github.com/jinzhu/gorm"
)

type ConnectorType struct {
	gorm.Model
	Name string
}
