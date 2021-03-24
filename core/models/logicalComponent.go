package models

import "github.com/jinzhu/gorm"

type LogicalComponent struct {
	gorm.Model
	LogicalName        string
	Type               string //connector/cluster/aggregator
	PivotID            uint
	Pivot              Pivot
	ConnectorProductID uint
	ConnectorProduct   ConnectorProduct
	KeyValues          []KeyValue
}
