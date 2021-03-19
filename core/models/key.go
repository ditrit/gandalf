package models

type Key struct {
	gorm.Model
	Name         string
	DefaultValue string
	Type         string
	Mandatory    bool
	PivotID uint
	Pivot Pivot
	ConnectorProductID uint
	ConnectorProduct ConnectorProduct
}