//Package models :
package models

import (
	"github.com/jinzhu/gorm"
)

// Connector : Connector struct.
type Connector struct {
	gorm.Model
	Name     string
	TenantID uint
	Tenant   Tenant
	Secret   string
}
