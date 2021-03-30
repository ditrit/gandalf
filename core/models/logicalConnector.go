package models

import "github.com/jinzhu/gorm"

type LogicalConnector struct {
	gorm.Model
	LogicalName        string
	ConnectorProductID uint
	ConnectorProduct   ConnectorProduct
	KeyValues          []KeyValue
}
