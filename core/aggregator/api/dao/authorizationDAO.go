package dao

import (
	"errors"

	"github.com/ditrit/gandalf/core/aggregator/api/utils"

	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"
)

func ListAuthorization(database *gorm.DB) (authorizations []models.Authorization, err error) {
	err = database.Find(&authorizations).Error

	return
}

func CreateAuthorization(database *gorm.DB, authorization models.Authorization) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			err = database.Create(&authorization).Error
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}

func ReadAuthorization(database *gorm.DB, id int) (authorization models.Authorization, err error) {
	err = database.First(&authorization, id).Error

	return
}

func UpdateAuthorization(database *gorm.DB, authorization models.Authorization) (err error) {
	err = database.Save(&authorization).Error

	return
}

func DeleteAuthorization(database *gorm.DB, id int) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			var authorization models.Authorization
			err = database.Unscoped().Delete(&authorization, id).Error
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}
