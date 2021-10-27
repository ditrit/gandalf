package dao

import (
	"errors"

	"github.com/ditrit/gandalf/core/aggregator/api/utils"
	"github.com/google/uuid"

	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"
)

func ListResourceAuthorization(database *gorm.DB) (resourceAuthorizations []models.ResourceAuthorization, err error) {
	err = database.Find(&resourceAuthorizations).Error

	return
}

func CreateResourceAuthorization(database *gorm.DB, resourceAuthorization *models.ResourceAuthorization) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			err = database.Create(&resourceAuthorization).Error
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}

func ReadResourceAuthorization(database *gorm.DB, id uuid.UUID) (resourceAuthorization models.ResourceAuthorization, err error) {
	err = database.Where("id = ?", id).First(&resourceAuthorization).Error

	return
}

func UpdateResourceAuthorization(database *gorm.DB, resourceAuthorization models.ResourceAuthorization) (err error) {
	err = database.Save(&resourceAuthorization).Error

	return
}

func DeleteResourceAuthorization(database *gorm.DB, id uuid.UUID) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			var resourceAuthorization models.ResourceAuthorization
			err = database.Where("id = ?", id).First(&resourceAuthorization).Error
			if err == nil {
				err = database.Unscoped().Delete(&resourceAuthorization).Error
			}
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}
