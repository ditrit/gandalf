package models

import "github.com/jinzhu/gorm"

// PermissionCommand : PermissionCommand struct.
type PermissionCommand struct {
	gorm.Model
	RoleID             uint
	Role               Role
	ConnectorCommandID uint
	ConnectorCommand   ConnectorCommand
	Allow              bool
}
