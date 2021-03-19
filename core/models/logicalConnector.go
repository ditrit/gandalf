package models

type LogicalConnector struct {
	gorm.Model
	LogicalName string
	ConnectorProductID uint
	ConnectorProduct ConnectorProduct
	KeyValues []KeyValue
}