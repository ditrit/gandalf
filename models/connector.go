//Package models :
package models

import (
	"github.com/jinzhu/gorm"
)

type Connector struct {
	gorm.Model
	Name string
}
