package models

import "github.com/jinzhu/gorm"

type Perimeter struct {
	gorm.Model
	UserID   uint
	User     User
	RoleID   uint
	Role     Role
	DomainID uint
	Domain   Domain
}
