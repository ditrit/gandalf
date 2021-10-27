package models

type EnvironmentType struct {
	Model
	Name             string `gorm:"not null"`
	ShortDescription string
	Description      string
	Logo             string
}
