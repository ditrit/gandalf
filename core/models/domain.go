package models

import (

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Domain struct {
	Model
	Name             string    `gorm:"not null"`
	ShortDescription string
	Description      string
	Logo             string

	ParentID         uuid.UUID `gorm:"type:uuid"`
	Parent           *Domain   `gorm:"constraint:OnDelete:CASCADE;"`

	Authorizations   []Authorization
	Libraries        []Library `gorm:"many2many:domain_libraries;"`
	Tags             []Tag     `gorm:"many2many:domain_tags;"`
	Environments     []Environment
	
	GitServerURL string
	GitPersonalAccessToken string
	GitOrganization string
	
	Products         []Product
	Childs           []*Domain


}

func (d *Domain) BeforeDelete(tx *gorm.DB) (err error) {
	// TODO : delete only if empty, if products in descendants domain we do not delete domain (cascade delete)
	var childs []Domain
	tx.Where("parent_id = ?", d.ID).Find(&childs)
	for _, child := range childs {
		tx.Unscoped().Delete(&child)
	}

	return
}
