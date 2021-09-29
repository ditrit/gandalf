package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Domain struct {
	gorm.Model
	ParentID         uint
	Parent           *Domain `gorm:"constraint:OnDelete:CASCADE;"`
	Name             string  `gorm:"not null"`
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
