package models

import "github.com/jinzhu/gorm"

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
	Versions       []models.Version
}

func NewConfigurationLogicalConnector(logicalName, tenant, connectorType, product, workersUrl, autoUpdateTime string, autoUpdate bool, maxTimeout int64, versionsMajor, versionsMinor int8) *ConfigurationLogicalConnector {
	configurationLogicalConnector := new(ConfigurationLogicalConnector)
	configurationLogicalConnector.LogicalName = logicalName
	configurationLogicalConnector.Tenant = tenant
	configurationLogicalConnector.ConnectorType = connectorType
	configurationLogicalConnector.Product = product
	configurationLogicalConnector.WorkersUrl = workersUrl
	configurationLogicalConnector.AutoUpdateTime = autoUpdateTime
	configurationLogicalConnector.AutoUpdate = autoUpdate
	configurationLogicalConnector.MaxTimeout = maxTimeout
	configurationLogicalConnector.VersionsMajor = versionsMajor
	configurationLogicalConnector.VersionsMinor = versionsMinor

	return configurationLogicalConnector
}
