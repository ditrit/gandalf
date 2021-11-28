package models

import "github.com/google/uuid"

type CommandType struct {
	Model
	Name               string    `gorm:"unique;not null"`
	Schema             string    //`gorm:"unique;not null"`
	PivotID            uuid.UUID `gorm:"type:uuid;check:(pivot_id IS NOT NULL AND product_connector_id IS NULL) OR (pivot_id IS NULL AND product_connector_id IS NOT NULL)"`
	Pivot              Pivot
	ProductConnectorID uuid.UUID `gorm:"type:uuid;check:(pivot_id IS NOT NULL AND product_connector_id IS NULL) OR (pivot_id IS NULL AND product_connector_id IS NOT NULL)"`
	ProductConnector   ProductConnector
}
