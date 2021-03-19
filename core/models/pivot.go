package models

type Pivot struct {
	gorm.Model
	Name string
	Major int8
	Minor int8
	ResourceTypes []ResourceType
	CommandTypes []CommandType
	EventTypes []EventType
	Keys []Key
}