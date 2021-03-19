package models

type EventType struct {
	gorm.Model
	Name string
	Schema string
	PivotID uint
	Pivot Pivot
	ConnectorProductID uint
	ConnectorProduct ConnectorProduct
}