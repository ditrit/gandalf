package models

import "github.com/jinzhu/gorm"

type KeyValue struct {
	gorm.Model
	Value              string
	KeyID              uint
	Key                Key
	LogicalConnectorID uint
	//LogicalConnector   LogicalConnector
}
