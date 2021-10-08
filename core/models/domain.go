package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Domain struct {
	gorm.Model
	Name             string `gorm:"not null"`
	ParentID         uint
	Parent           *Domain `gorm:"constraint:OnDelete:CASCADE;"`
	Products         []Product
	Libraries        []Library
	Authorizations   []Authorization
	Tags             []Tag `gorm:"many2many:domain_tags;"`
	Environments     string
	ShortDescription string
	Description      string
	Logo             string
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
