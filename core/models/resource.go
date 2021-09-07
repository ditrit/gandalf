package models

import "github.com/jinzhu/gorm"

type Resource struct {
	gorm.Model
	Name               string `gorm:"unique;not null"`
	LogicalComponentID uint
	LogicalComponent   LogicalComponent
	DomainID           uint
	Domain             Domain
	ResourceTypeID     uint
	ResourceType       ResourceType
	EventTypeToPolls   []EventTypeToPoll `gorm:"ForeignKey:ResourceID"`
}
