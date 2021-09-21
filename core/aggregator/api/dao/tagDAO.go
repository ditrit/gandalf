package dao

import (
	"errors"

	"github.com/ditrit/gandalf/core/aggregator/api/utils"

	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"
)

func ListTag(database *gorm.DB) (tags []models.Tag, err error) {
	err = database.Find(&tags).Error

	return
}

func CreateTag(database *gorm.DB, tag models.Tag) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			err = database.Create(&tag).Error
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}

func ReadTag(database *gorm.DB, id int) (tag models.Tag, err error) {
	err = database.First(&tag, id).Error

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
			err = database.Unscoped().Delete(&tag, id).Error
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}
