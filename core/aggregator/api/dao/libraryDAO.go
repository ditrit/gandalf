package dao

import (
	"errors"

	"github.com/ditrit/gandalf/core/aggregator/api/utils"

	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"
)

func ListLibrary(database *gorm.DB) (librarys []models.Library, err error) {
	err = database.Find(&librarys).Error

	return
}

func CreateLibrary(database *gorm.DB, library models.Library) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			err = database.Create(&library).Error
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}

func ReadLibrary(database *gorm.DB, id int) (library models.Library, err error) {
	err = database.First(&library, id).Error

	return
}

func UpdateLibrary(database *gorm.DB, library models.Library) (err error) {
	err = database.Save(&library).Error

	return
}

func DeleteLibrary(database *gorm.DB, id int) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			var library models.Role
			err = database.Unscoped().Delete(&library, id).Error
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}
