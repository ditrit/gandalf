package dao

import (
	"errors"

	"github.com/ditrit/gandalf/core/aggregator/api/utils"
	"github.com/google/uuid"

	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"
)

func ListLibrary(database *gorm.DB) (librarys []models.Library, err error) {
	err = database.Preload("Domain").Find(&librarys).Error

	return
}

func CreateLibrary(database *gorm.DB, library *models.Library) (err error) {
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

func ReadLibrary(database *gorm.DB, id uuid.UUID) (library models.Library, err error) {
	err = database.Where("id = ?", id).Preload("Domain").First(&library).Error

	return
}

func UpdateLibrary(database *gorm.DB, library models.Library) (err error) {
	err = database.Save(&library).Error

	return
}

func DeleteLibrary(database *gorm.DB, id uuid.UUID) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			var library models.Role
			err = database.Where("id = ?", id).First(&library).Error
			if err == nil {
				err = database.Unscoped().Delete(&library).Error
			}
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}
