package models

import "github.com/jinzhu/gorm"

type Authorization struct {
	gorm.Model
	UserID   uint
	User     User
	RoleID   uint
	Role     Role
	DomainID uint
	Domain   Domain
}
