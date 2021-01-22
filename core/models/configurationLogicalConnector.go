package models

import "github.com/jinzhu/gorm"

type ConfigurationLogicalConnector struct {
	gorm.Model
	LogicalName    string
	ConnectorType  string
	Product        string
	WorkersUrl     string
	AutoUpdateTime string
	AutoUpdate     bool
	MaxTimeout     int64
	VersionsMajor  int8
	VersionsMinor  int8
}
