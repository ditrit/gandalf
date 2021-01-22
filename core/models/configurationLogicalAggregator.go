package models

import "github.com/jinzhu/gorm"

type ConfigurationLogicalAggregator struct {
	gorm.Model
	LogicalName string
	Tenant      string
}

func NewConfigurationLogicalAggregator(logicalName, tenant string) *ConfigurationLogicalAggregator {
	configurationLogicalAggregator := new(ConfigurationLogicalAggregator)
	configurationLogicalAggregator.LogicalName = logicalName
	configurationLogicalAggregator.Tenant = tenant

	return configurationLogicalAggregator
}
