package models

import "github.com/google/uuid"

type Key struct {
	Model
	Name               string
	DefaultValue       string
	Type               string
	Shortname          string
	Mandatory          bool
	PivotID            uuid.UUID `gorm:"type:uuid;check:(pivot_id IS NOT NULL AND product_connector_id IS NULL) OR (pivot_id IS NULL AND product_connector_id IS NOT NULL)"`
	Pivot              Pivot
	ProductConnectorID uuid.UUID `gorm:"type:uuid;check:(pivot_id IS NOT NULL AND product_connector_id IS NULL) OR (pivot_id IS NULL AND product_connector_id IS NOT NULL)"`
	ProductConnector   ProductConnector
}
