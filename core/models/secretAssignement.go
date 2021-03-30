package models

type SecretAssignement struct {
	Secret    string `gorm:"primaryKey"`
	AddressIP string
}
