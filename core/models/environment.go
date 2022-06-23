package models

import (
	"github.com/google/uuid"
)

type Environment struct {
	Model
	Name              string    `gorm:"not null"`
	EnvironmentTypeID uuid.UUID `gorm:"type:uuid"`
	EnvironmentType   EnvironmentType
	ShortDescription  string
	Description       string
	Logo              string
	DomainID          uuid.UUID `gorm:"type:uuid"`
	Domain           Domain
}
