package models

import "github.com/jinzhu/gorm"

type ConfigurationAggregator struct {
	gorm.Model
	LogicalName string
}
