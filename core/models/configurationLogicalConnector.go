package models

import "github.com/jinzhu/gorm"

type ConfigurationLogicalConnector struct {
	gorm.Model
	LogicalName    string
	Tenant         string
	ConnectorType  string
	Product        string
	WorkersUrl     string
	AutoUpdateTime string
	AutoUpdate     bool
	MaxTimeout     int64
	VersionsMajor  int8
	VersionsMinor  int8
}

func NewConfigurationLogicalConnector(logicalName, connectorType, product, workersUrl, autoUpdateTime string, autoUpdate bool, maxTimeout int64, versionsMajor, versionsMinor int8) *ConfigurationLogicalConnector {
	configurationLogicalConnector := new(ConfigurationLogicalConnector)
	configurationLogicalConnector.LogicalName = logicalName
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
