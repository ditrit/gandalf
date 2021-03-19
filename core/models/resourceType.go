package models

type ResourceType struct {
	gorm.Model
	Name string
	PivotID uint
	Pivot Pivot
	ConnectorProductID uint
	ConnectorProduct ConnectorProduct
}