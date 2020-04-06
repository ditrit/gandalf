package models

import (
	"github.com/jinzhu/gorm"
)

type Application struct {
	gorm.Model
	Name string
	/* Aggregator      Aggregator
	AggregatorID    uint
	Connector       Connector
	ConnectorID     uint
	ConnectorType   ConnectorType
	ConnectorTypeID uint */
	Aggregator    string
	Connector     string
	ConnectorType string
}
