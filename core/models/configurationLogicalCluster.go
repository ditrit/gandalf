package models

import "github.com/jinzhu/gorm"

type ConfigurationLogicalCluster struct {
	gorm.Model
	LogicalName string
	Secret      string
	MaxTimeout  int64
}
