package models

import "github.com/jinzhu/gorm"

type EventTypeToPoll struct {
	gorm.Model
	ResourceID  uint `gorm:"UNIQUE_INDEX:compositeindex;not null"`
	Resource    Resource
	EventTypeID uint `gorm:"UNIQUE_INDEX:compositeindex;not null"`
	EventType   EventType
}
