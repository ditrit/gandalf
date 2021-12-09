package models

import (
	"github.com/google/uuid"
)

type Authorization struct {
	Model
	UserID   uuid.UUID `gorm:"type:uuid"`
	User     User
	RoleID   uuid.UUID `gorm:"type:uuid"`
	Role     Role
	DomainID uuid.UUID `gorm:"type:uuid"`
	Domain   Domain
}
