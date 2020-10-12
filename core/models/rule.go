package models

import "github.com/jinzhu/gorm"

type Rule struct {
	gorm.Model
	RoleID   uint
	Role     Role
	DomainID uint
	Domain   Domain
	ObjectID uint
	Object   Object
	ActionID uint
	Action   Action
	Allow    bool
}
