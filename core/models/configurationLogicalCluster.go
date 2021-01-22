package models

import "github.com/jinzhu/gorm"

type ConfigurationLogicalCluster struct {
	gorm.Model
	LogicalName string
}

func NewConfigurationLogicalCluster(logicalName string) *ConfigurationLogicalCluster {
	configurationLogicalCluster := new(ConfigurationLogicalCluster)
	configurationLogicalCluster.LogicalName = logicalName

	return configurationLogicalCluster
}
