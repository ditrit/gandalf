//Package models :
package models

import (
	"github.com/jinzhu/gorm"
)

// Application : Application struct.
type Application struct {
	gorm.Model
	Name            string
	AggregatorID    uint
	Aggregator      Aggregator
	ConnectorID     uint
	Connector       Connector
	ConnectorTypeID uint
	ConnectorType   ConnectorType
}
