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

func (p *Product) AfterCreate(tx *gorm.DB) (err error) {
	if p.RepositoryURL == "" {
		var domain Domain
		tx.Where("id = ?", p.DomainID).First(&domain)
		if err == nil {
			p.RepositoryURL = domain.GitURL + "/" + p.Name + ".git"
		}
	}
	return
}
