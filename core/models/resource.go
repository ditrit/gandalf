package models

type Resource struct {
	gorm.Model
	ProductConnectorID uint
	ProductConnector ProductConnector
	DomainID uint
	Domain Domain
	ResourceTypeID uint
	ResourceType ResourceType
}