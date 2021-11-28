package dao

import (
	"errors"

	"github.com/ditrit/gandalf/core/aggregator/api/utils"
	"github.com/google/uuid"

	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"
)

func ListAuthorization(database *gorm.DB) (authorizations []models.Authorization, err error) {
	err = database.Preload("User").Preload("Role").Preload("Domain").Find(&authorizations).Error

	return
}

func CreateAuthorization(database *gorm.DB, authorization *models.Authorization) (err error) {
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

func ReadAuthorization(database *gorm.DB, id uuid.UUID) (authorization models.Authorization, err error) {
	err = database.Where("id = ?", id).Preload("User").Preload("Role").Preload("Domain").First(&authorization).Error

	return
}

func UpdateAuthorization(database *gorm.DB, authorization models.Authorization) (err error) {
	err = database.Save(&authorization).Error

	return
}

func DeleteAuthorization(database *gorm.DB, id uuid.UUID) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			var authorization models.Authorization
			err = database.Where("id = ?", id).First(&authorization).Error
			if err == nil {
				err = database.Unscoped().Delete(&authorization).Error
			}
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}
