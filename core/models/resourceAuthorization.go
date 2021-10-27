package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type ResourceAuthorization struct {
	gorm.Model
	RoleID         uuid.UUID `gorm:"type:uuid"`
	Role           Role
	DomainID       uuid.UUID `gorm:"type:uuid"`
	Domain         Domain
	ResourceTypeID uuid.UUID `gorm:"type:uuid"`
	ResourceType   ResourceType
	Allow          bool
}
