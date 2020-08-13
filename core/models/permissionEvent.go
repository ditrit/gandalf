package models

import "github.com/jinzhu/gorm"

// PermissionEvent : PermissionEvent struct.
type PermissionEvent struct {
	gorm.Model
	Role           Role
	ConnectorEvent ConnectorEvent
	Allow          bool
}
