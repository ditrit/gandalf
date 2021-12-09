package models

import (
	"github.com/google/uuid"
)

type Resource struct {
	Model
	Name               string    `gorm:"unique;not null"`
	LogicalComponentID uuid.UUID `gorm:"type:uuid"`
	LogicalComponent   LogicalComponent
	DomainID           uuid.UUID `gorm:"type:uuid"`
	Domain             Domain
	ResourceTypeID     uuid.UUID `gorm:"type:uuid"`
	ResourceType       ResourceType
	EventTypeToPolls   []EventTypeToPoll `gorm:"ForeignKey:ResourceID"`
}
