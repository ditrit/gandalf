package models

import "github.com/jinzhu/gorm"

type ResourceAuthorization struct {
	gorm.Model
	RoleID         uint
	Role           Role
	DomainID       uint
	Domain         Domain
	ResourceTypeID uint
	ResourceType   ResourceType
	Allow          bool
}
