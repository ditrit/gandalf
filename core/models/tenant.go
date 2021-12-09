//Package models :
package models

// Tenant : Tenant struct.
type Tenant struct {
	Model
	Name             string `gorm:"unique"`
	Password         string
	ShortDescription string
	Description      string
	Logo             string
}
