package models

import "github.com/jinzhu/gorm"

type ConfigurationLogicalCluster struct {
	gorm.Model
	LogicalName string
	Secret      string
	MaxTimeout  int64
}

func NewConfigurationLogicalCluster(logicalName string) *ConfigurationLogicalCluster {
	configurationLogicalCluster := new(ConfigurationLogicalCluster)
	configurationLogicalCluster.LogicalName = logicalName

	return configurationLogicalCluster
}
