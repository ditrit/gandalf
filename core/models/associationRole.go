package models

import "github.com/jinzhu/gorm"

type AssociationRole struct {
	gorm.Model
	UserID   uint
	User     User
	RoleID   uint
	Role     Role
	DomainID uint
	Domain   Domain
}
