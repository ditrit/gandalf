package models

import "github.com/google/uuid"

type EventTypeToPoll struct {
	Model
	ResourceID  uuid.UUID `gorm:"type:uuid;UNIQUE_INDEX:compositeindex;not null"`
	Resource    Resource
	EventTypeID uuid.UUID `gorm:"type:uuid;UNIQUE_INDEX:compositeindex;not null"`
	EventType   EventType
}
