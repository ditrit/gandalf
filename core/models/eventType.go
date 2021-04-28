package models

import "github.com/jinzhu/gorm"

type EventType struct {
	gorm.Model
	Name               string `gorm:"unique;not null"`
	Schema             string //`gorm:"unique;not null"`
	PivotID            uint   `gorm:"check:(pivot_id IS NOT NULL AND connector_product_id IS NULL) OR (pivot_id IS NULL AND connector_product_id IS NOT NULL)"`
	Pivot              Pivot
	ConnectorProductID uint `gorm:"check:(pivot_id IS NOT NULL AND connector_product_id IS NULL) OR (pivot_id IS NULL AND connector_product_id IS NOT NULL)"`
	ConnectorProduct   ConnectorProduct
}
