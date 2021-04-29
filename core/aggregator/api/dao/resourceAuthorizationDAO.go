package dao

import (
	"errors"

	"github.com/ditrit/gandalf/core/aggregator/api/utils"

	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"
)

func ListResourceAuthorization(database *gorm.DB) (resourceAuthorizations []models.ResourceAuthorization, err error) {
	err = database.Find(&resourceAuthorizations).Error

	return
}

func CreateResourceAuthorization(database *gorm.DB, resourceAuthorization models.ResourceAuthorization) (err error) {
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

func ReadResourceAuthorization(database *gorm.DB, id int) (resourceAuthorization models.ResourceAuthorization, err error) {
	err = database.First(&resourceAuthorization, id).Error

	return
}

func UpdateResourceAuthorization(database *gorm.DB, resourceAuthorization models.ResourceAuthorization) (err error) {
	err = database.Save(&resourceAuthorization).Error

	return
}

func DeleteResourceAuthorization(database *gorm.DB, id int) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			var resourceAuthorization models.ResourceAuthorization
			err = database.Delete(&resourceAuthorization, id).Error
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}
