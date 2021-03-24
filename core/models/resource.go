package models

import "github.com/jinzhu/gorm"

type Resource struct {
	gorm.Model
	Name               string
	ProductConnectorID uint
	ProductConnector   ProductConnector
	DomainID           uint
	Domain             Domain
	ResourceTypeID     uint
	ResourceType       ResourceType
}
