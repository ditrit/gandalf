package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Tag struct {
	gorm.Model
	Name string `gorm:"not null"`
}

func GetTagDescendants(database *gorm.DB, id uint) (tags []Tag, err error) {
	err = database.Order("depth asc").Joins("JOIN tag_closures ON tags.id = tag_closures.descendant_id").Where("tag_closures.ancestor_id = ?", id).Find(&tags).Error

	return

}

func GetTagAncestors(database *gorm.DB, id uint) (tags []Tag, err error) {
	err = database.Order("depth desc").Joins("JOIN tag_closures ON tags.id = tag_closures.ancestor_id").Where("tag_closures.descendant_id = ?", id).Find(&tags).Error

	return
}

func InsertTagRoot(database *gorm.DB, tag Tag) (err error) {
	err = database.Transaction(func(tx *gorm.DB) error {

		if err := tx.Save(&tag).Error; err != nil {
			// return any error will rollback
			return err
		}
		if err := tx.Where("name = ?", tag.Name).First(&tag).Error; err != nil {
			// return any error will rollback
			return err
		}
		if err := tx.Exec("INSERT INTO tag_closures (ancestor_id, descendant_id, depth) SELECT ?,?,0;", tag.ID, tag.ID).Error; err != nil {
			// return any error will rollback
			return err
		}
		return nil
	})

	return
}

func InsertTagNewChild(database *gorm.DB, tag Tag, id uint) (err error) {

	err = database.Transaction(func(tx *gorm.DB) error {

		if err := tx.Save(&tag).Error; err != nil {
			// return any error will rollback
			return err
		}
		if err := tx.Where("name = ?", tag.Name).First(&tag).Error; err != nil {
			// return any error will rollback
			return err
		}
		if err := tx.Exec("INSERT INTO tag_closures (ancestor_id, descendant_id, depth) SELECT ancestor_id, ?, depth+1 FROM tag_closures WHERE descendant_id = ? UNION ALL SELECT ?,?,0;", tag.ID, id, tag.ID, tag.ID).Error; err != nil {
			// return any error will rollback
			return err
		}
		return nil
	})

	return
}

func DeleteTagChild(database *gorm.DB, tag Tag) (err error) {

	err = database.Transaction(func(tx *gorm.DB) error {

		if err := tx.Delete(&Tag{}, int(tag.ID)).Error; err != nil {
			fmt.Println(err)
			// return any error will rollback
			return err
		}
		var tagClosure TagClosure
		if err := tx.Where("descendant_id = ?", tag.ID).Delete(&tagClosure).Error; err != nil {
			// return any error will rollback
			fmt.Println(err)
			return err
		}
		return nil
	})
	return
}

func DeleteTagSubtree(database *gorm.DB, id uint) (err error) {

	err = database.Transaction(func(tx *gorm.DB) error {

		if err := tx.Exec("DELETE FROM tags WHERE tags.id IN (SELECT descendant_id FROM tag_closures WHERE ancestor_id = ?);", id).Error; err != nil {
			// return any error will rollback
			return err
		}
		if err := tx.Exec("DELETE FROM tag_closures WHERE descendant_id IN (SELECT descendant_id FROM tag_closures WHERE ancestor_id = ?);", id).Error; err != nil {
			// return any error will rollback
			return err
		}
		return nil
	})

	return
}

func UpdateTagChild(database *gorm.DB, tag Tag, newAncestor uint) (err error) {

	err = database.Transaction(func(tx *gorm.DB) error {

		if err := tx.Exec("DELETE FROM tag_closures WHERE descendant_id IN (SELECT descendant_id FROM tag_closures WHERE ancestor_id = ?) AND ancestor_id IN (SELECT ancestor_id FROM tag_closures WHERE descendant_id = ? AND ancestor_id != descendant_id);", tag.ID, tag.ID).Error; err != nil {
			// return any error will rollback
			return err
		}
		if err := tx.Exec("INSERT INTO tag_closures (ancestor_id, descendant_id, depth) SELECT supertree.ancestor_id, subtree.descendant_id,  supertree.depth + subtree.depth + 1 AS depth FROM tag_closures AS supertree CROSS JOIN tag_closures AS subtree WHERE supertree.descendant_id = ? AND subtree.ancestor_id = ?;", newAncestor, tag.ID).Error; err != nil {
			// return any error will rollback
			return err
		}
		return nil
	})

	return
}
