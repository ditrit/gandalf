package models

import "github.com/jinzhu/gorm"

type ConfigurationLogicalAggregator struct {
	gorm.Model
	LogicalName string
	Tenant      string
	Secret      string
	MaxTimeout  int64
}
