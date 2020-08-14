package models

import "github.com/jinzhu/gorm"

// PermissionCommand : PermissionCommand struct.
type PermissionCommand struct {
	gorm.Model
	Role             Role
	ConnectorCommand ConnectorCommand
	Allow            bool
}
