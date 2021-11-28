package models

type ConnectorProduct struct {
	Model
	Name string `gorm:"unique;not null"`
}
