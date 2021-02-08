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
