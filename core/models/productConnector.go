package models

import (
	"github.com/google/uuid"
)

type ProductConnector struct {
	Model
	Name          string `gorm:"unique;not null"`
	Major         int8
	Minor         int8
	PivotID       uuid.UUID `gorm:"type:uuid"`
	Pivot         Pivot
	ProductID     uuid.UUID `gorm:"type:uuid"`
	Product       ConnectorProduct
	ResourceTypes []ResourceType `gorm:"ForeignKey:ProductConnectorID"`
	CommandTypes  []CommandType  `gorm:"ForeignKey:ProductConnectorID"`
	EventTypes    []EventType    `gorm:"ForeignKey:ProductConnectorID"`
	Keys          []Key          `gorm:"ForeignKey:ProductConnectorID"`
}
