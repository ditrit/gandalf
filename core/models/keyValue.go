package models

import (
	"github.com/google/uuid"
)

type KeyValue struct {
	Model
	Value              string
	KeyID              uuid.UUID `gorm:"type:uuid"`
	Key                Key
	LogicalComponentID uuid.UUID `gorm:"type:uuid"`
	LogicalComponent   LogicalComponent
}
