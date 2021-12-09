package models

// Role : Role struct.
type Role struct {
	Model
	Name             string `gorm:"unique;not null"`
	ShortDescription string
	Description      string
	Logo             string
}
