package models

type Product struct {
	gorm.Model
	Name string
}