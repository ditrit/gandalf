package models

type EventTypeToPoll struct {
	ResourceID  uint `gorm:"UNIQUE_INDEX:compositeindex;not null"`
	Resource    Resource
	EventTypeID uint `gorm:"UNIQUE_INDEX:compositeindex;not null"`
	EventType   EventType
}
