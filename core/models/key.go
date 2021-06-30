package models

import "github.com/jinzhu/gorm"

type Key struct {
	gorm.Model
	Name               string
	DefaultValue       string
	Type               string
	Shortname          string
	Mandatory          bool
	PivotID            uint `gorm:"check:(pivot_id IS NOT NULL AND product_connector_id IS NULL) OR (pivot_id IS NULL AND product_connector_id IS NOT NULL)"`
	Pivot              Pivot
	ProductConnectorID uint `gorm:"check:(pivot_id IS NOT NULL AND product_connector_id IS NULL) OR (pivot_id IS NULL AND product_connector_id IS NOT NULL)"`
	ProductConnector   ProductConnector
}
