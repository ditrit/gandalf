package models

import "github.com/jinzhu/gorm"

type ConfigurationConnector struct {
	gorm.Model
	LogicalName     string
	GRPCBindAddress string
	ConnectorType   string
	Product         string
	WorkersUrl      string
	Workers         string
	AutoUpdateTime  string
	AutoUpdate      bool
	MaxTimeout      int64
	VersionsMajor   int64
	VersionsMinor   int64
}
