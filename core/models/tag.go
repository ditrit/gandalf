package models

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Tag struct {
	Model
	Name             string    `gorm:"not null"`
	ParentID         uuid.UUID `gorm:"type:uuid"`
	Parent           *Tag      `gorm:"constraint:OnDelete:CASCADE;"`
	ShortDescription string
	Description      string
	Logo             string
	Childs           []*Tag
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
