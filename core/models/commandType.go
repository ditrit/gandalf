package models

type CommandType struct {
	gorm.Model
	Name string
	Schema string
	PivotID uint
	Pivot Pivot
	ConnectorProductID uint
	ConnectorProduct ConnectorProduct
}