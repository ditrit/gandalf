package models

import "github.com/jinzhu/gorm"

type ConfigurationConnector struct {
	gorm.Model
	LogicalName    string
	ConnectorType  string
	Product        string
	GRPCSocketDir  string
	WorkersUrl     string
	WorkersPath    string
	AutoUpdateTime string
	AutoUpdate     bool
	MaxTimeout     int64
	VersionsMajor  int8
	VersionsMinor  int8
}
