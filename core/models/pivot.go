package models

import "github.com/jinzhu/gorm"

type Pivot struct {
	gorm.Model
	Name          string `gorm:"unique;not null"`
	Major         int8
	Minor         int8
	TypeID        uint
	Type          Type
	ResourceTypes []ResourceType
	CommandTypes  []CommandType
	EventTypes    []EventType
	Keys          []Key
}
