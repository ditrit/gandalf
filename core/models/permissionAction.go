package models

import "github.com/jinzhu/gorm"

// PermissionAction : PermissionAction struct.
type PermissionAction struct {
	gorm.Model
	Role   Role
	Action Action
	Allow  bool
}
