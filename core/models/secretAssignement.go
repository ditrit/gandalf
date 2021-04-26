package models

type SecretAssignement struct {
	Secret    string `gorm:"primaryKey;unique"`
	AddressIP string
}
