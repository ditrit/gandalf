package models

type ProductConnector struct {
	gorm.Model
	Name string
	Major int8
	Minor int8
	PivotID uint
	Pivot Pivot
	ProductID uint
	Product Product
	ResourceTypes []ResourceType
	CommandTypes []CommandType
	EventTypes []EventType
	Keys []Key
	Resources []Resource
}