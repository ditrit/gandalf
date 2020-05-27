//Package models :
package models

import (
	"github.com/jinzhu/gorm"
)

// Application : Application struct.
type Application struct {
	gorm.Model
	Name          string
	Tenant        Tenant
	Aggregator    Aggregator
	Connector     Connector
	ConnectorType ConnectorType
}
