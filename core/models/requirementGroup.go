package models

type RequirementGroup struct {
	Model
	Name        string `gorm:"unique;not null"`
	Description string
	Logo        string
}
