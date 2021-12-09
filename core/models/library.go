package models

type Library struct {
	Model
	Name             string `gorm:"unique;not null"`
	ShortDescription string
	Description      string
	Logo             string
}
