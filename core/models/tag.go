package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Tag struct {
	gorm.Model
	Name     string `gorm:"not null"`
	ParentID uint
	Parent   *Tag `gorm:"constraint:OnDelete:CASCADE;"`
}

func (t *Tag) BeforeDelete(tx *gorm.DB) (err error) {
	var childs []Tag
	tx.Where("parent_id = ?", t.ID).Find(&childs)
	fmt.Println(childs)
	for _, child := range childs {
		tx.Unscoped().Delete(&child)
	}

	return
}
