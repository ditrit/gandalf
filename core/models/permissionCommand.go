package models

import "github.com/jinzhu/gorm"

// PermissionMessage : PermissionMessage struct.
type PermissionMessage struct {
	gorm.Model
	Role             Role
	ConnectorCommand ConnectorCommand
	Allow            bool
}
