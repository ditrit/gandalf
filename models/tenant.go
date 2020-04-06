package models

import (
	"github.com/jinzhu/gorm"
)

type Tenant struct {
	gorm.Model
	Name string `form:"name" json:"name" binding:"required" gorm:"type:varchar(255);not null"`
}
