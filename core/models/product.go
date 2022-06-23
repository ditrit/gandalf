package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Product struct {
	Model
	Name             string `gorm:"unique;not null"`
	ShortDescription string
	Description      string
	Logo             string
	RepositoryURL    string
	DomainID         uuid.UUID `gorm:"type:uuid"`
	Domain           Domain
}

	}
	return
}
