package models

type KeyValue struct {
	gorm.Model
	Value interface
	KeyID uint
	Key Key
	LogicalConnectorID uint
	LogicalConnector LogicalConnector
}