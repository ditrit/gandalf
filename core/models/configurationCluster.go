package models

import "github.com/jinzhu/gorm"

type ConfigurationCluster struct {
	gorm.Model
	LogicalName string
	DBPath      string
	DBName      string
}
