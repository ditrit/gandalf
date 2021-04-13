package models

import "github.com/jinzhu/gorm"

type Resource struct {
	gorm.Model
	Name               string `gorm:"unique;not null"`
	ProductConnectorID uint
	ProductConnector   ProductConnector
	DomainID           uint
	Domain             Domain
	ResourceTypeID     uint
	ResourceType       ResourceType
	EventTypeToPolls   []EventTypeToPoll
}
