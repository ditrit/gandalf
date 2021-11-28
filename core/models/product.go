package models

import (
	"github.com/google/uuid"
)

type Product struct {
	Model
	Name             string `gorm:"unique;not null"`
	ShortDescription string
	Description      string
	Logo             string
	DomainID         uuid.UUID `gorm:"type:uuid"`
	Domain           Domain
}
