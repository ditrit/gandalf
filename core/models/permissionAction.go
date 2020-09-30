package models

import "github.com/jinzhu/gorm"

// PermissionAction : PermissionAction struct.
type PermissionAction struct {
	gorm.Model
	RoleID   uint
	Role     Role
	ActioniD uint
	Action   Action
	Allow    bool
}
