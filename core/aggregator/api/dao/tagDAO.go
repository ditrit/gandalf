package dao

import (
	"errors"
	"fmt"

	"github.com/ditrit/gandalf/core/aggregator/api/utils"

	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"
)

func ListTag(database *gorm.DB) (tags []models.Tag, err error) {
	var root models.Tag
	err = database.Where("name = ?", "root").First(&root).Error
	if err == nil {
		//tags, err = models.GetTagAncestors(database, root.ID)
		//tags, err = models.GetTagDescendants(database, root.ID)
		//tags, err = models.GetTagTree(database, root.ID)
	}
	err = database.Preload("Parent").Find(&tags).Error
	fmt.Println(err)
	return
}

func CreateTag(database *gorm.DB, tag models.Tag, parentTagID uint) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			// if parentTagName == "root" {
			// 	err = models.InsertTagRoot(database, tag)
			// } else {
			// var parentTag models.Tag
			// err = database.Where("name = ?", parentTagName).First(&parentTag).Error
			// if err == nil {
			tag.ParentID = parentTagID
			err = database.Save(&tag).Error
			//}
			//}
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}

func TreeTag(database *gorm.DB) (result *models.TagTree, err error) {
	var results []models.Tag
	database.Raw("select * from tags order by parent_id").Scan(&results)

	tagTree := new(models.TagTree)
	tagTree.Tag = results[0]
	TreeRecursiveTag(tagTree, results)

	result = tagTree
	return
}

func TreeRecursiveTag(tagtree *models.TagTree, results []models.Tag) {
	for _, result := range results {
		if result.ParentID == tagtree.Tag.ID {
			currentTagTree := new(models.TagTree)
			currentTagTree.Tag = result
			tagtree.Childs = append(tagtree.Childs, currentTagTree)
		}
	}
	for _, child := range tagtree.Childs {
		TreeRecursiveTag(child, results)
	}
}

func ReadTag(database *gorm.DB, id int) (tag models.Tag, err error) {
	err = database.First(&tag, id).Error

	return
}

func ReadTagByName(database *gorm.DB, name string) (tag models.Tag, err error) {
	fmt.Println("DAO")
	err = database.Where("name = ?", name).First(&tag).Error
	fmt.Println(err)
	fmt.Println(tag)
	return
}

func UpdateTag(database *gorm.DB, tag models.Tag) (err error) {
	err = database.Save(&tag).Error

	return
}

func DeleteTag(database *gorm.DB, id int) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			var tag models.Tag
			err = database.First(&tag, id).Error
			if err == nil {
				err = database.Unscoped().Delete(&tag).Error
			}
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}
