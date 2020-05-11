//Package models :
package models

import (
	"github.com/jinzhu/gorm"
)

type Aggregator struct {
	gorm.Model
	Name string
}
