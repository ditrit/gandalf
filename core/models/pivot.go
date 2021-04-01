package models

import "github.com/jinzhu/gorm"

type Pivot struct {
	gorm.Model
	Name          string `gorm:"UNIQUE_INDEX:compositeindex;not null"`
	Major         int8   `gorm:"UNIQUE_INDEX:compositeindex;not null"`
	Minor         int8   `gorm:"UNIQUE_INDEX:compositeindex;not null"`
	ResourceTypes []ResourceType
	CommandTypes  []CommandType
	EventTypes    []EventType
	Keys          []Key
}
