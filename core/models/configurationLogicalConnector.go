package models

import (
	"github.com/jinzhu/gorm"
)

type ConfigurationLogicalConnector struct {
	gorm.Model
	LogicalName    string
	Tenant         string
	Secret         string
	ConnectorType  string
	Product        string
	WorkersUrl     string
	AutoUpdateTime string
	MaxTimeout     int64
	Versions       string
}

func NewConfigurationLogicalConnector(logicalName, tenant, connectorType, product, workersUrl, autoUpdateTime string, maxTimeout int64, versions string) *ConfigurationLogicalConnector {
	configurationLogicalConnector := new(ConfigurationLogicalConnector)
	configurationLogicalConnector.LogicalName = logicalName
	configurationLogicalConnector.Tenant = tenant
	configurationLogicalConnector.ConnectorType = connectorType
	configurationLogicalConnector.Product = product
	configurationLogicalConnector.WorkersUrl = workersUrl
	configurationLogicalConnector.AutoUpdateTime = autoUpdateTime
	configurationLogicalConnector.MaxTimeout = maxTimeout
	configurationLogicalConnector.Versions = versions

	return configurationLogicalConnector
}
