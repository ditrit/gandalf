//Package models :
package models

import (
	"github.com/jinzhu/gorm"
)

// Aggregator : Aggregator struct.
type Aggregator struct {
	gorm.Model
	Name     string
	TenantID uint
	Tenant   Tenant
	Secret   string
}
