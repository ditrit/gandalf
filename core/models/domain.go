package models

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Domain struct {
	Model
	Name             string    `gorm:"not null"`
	ParentID         uuid.UUID `gorm:"type:uuid"`
	Parent           *Domain   `gorm:"constraint:OnDelete:CASCADE;"`
	Authorizations   []Authorization
	Products         []Product
	Libraries        []Library     `gorm:"many2many:domain_libraries;"`
	Tags             []Tag         `gorm:"many2many:domain_tags;"`
	Environments     []Environment `gorm:"many2many:domain_environments;"`
	ShortDescription string
	Description      string
	Logo             string
	Childs           []*Domain
}

func (d *Domain) BeforeDelete(tx *gorm.DB) (err error) {
	var childs []Domain
	tx.Where("parent_id = ?", d.ID).Find(&childs)
	fmt.Println(childs)
	for _, child := range childs {
		tx.Unscoped().Delete(&child)
	}

	return
}
