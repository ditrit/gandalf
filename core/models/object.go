package models

import "github.com/jinzhu/gorm"

type Object struct {
	gorm.Model
	Name         string
	Schema       string
	Action       []Action `gorm:"many2many:object_actions;"`
	Domain       []Domain `gorm:"many2many:object_domains;"`
	ObjectTypeID uint
	ObjectType   *Object
}
