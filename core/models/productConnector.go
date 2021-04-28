package models

import "github.com/jinzhu/gorm"

type ProductConnector struct {
	gorm.Model
	Name          string `gorm:"unique;not null"`
	Major         int8
	Minor         int8
	PivotID       uint
	Pivot         Pivot
	ProductID     uint
	Product       Product
	ResourceTypes []ResourceType `gorm:"ForeignKey:ConnectorProductID"`
	CommandTypes  []CommandType  `gorm:"ForeignKey:ConnectorProductID"`
	EventTypes    []EventType    `gorm:"ForeignKey:ConnectorProductID"`
	Keys          []Key          `gorm:"ForeignKey:ConnectorProductID"`
	Resources     []Resource
}
