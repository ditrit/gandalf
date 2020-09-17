//Package models :
package models

import (
	"github.com/jinzhu/gorm"
)

// Cluster : Cluster struct.
type Cluster struct {
	gorm.Model
	LogicalName  string
	InstanceName string
	Secret       string
}
