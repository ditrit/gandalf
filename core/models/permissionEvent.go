package models

import "github.com/jinzhu/gorm"

// PermissionEvent : PermissionEvent struct.
type PermissionEvent struct {
	gorm.Model
	RoleID           uint
	Role             Role
	ConnectorEventID uint
	ConnectorEvent   ConnectorEvent
	Allow            bool
}
