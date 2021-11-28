package dao

import (
	"errors"

	"github.com/ditrit/gandalf/core/aggregator/api/utils"
	"github.com/google/uuid"

	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"
)

func ListEnvironment(database *gorm.DB) (environments []models.Environment, err error) {
	err = database.Preload("EnvironmentType").Find(&environments).Error

	return
}

func CreateEnvironment(database *gorm.DB, environment *models.Environment) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			err = database.Create(&environment).Error
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}

func ReadEnvironment(database *gorm.DB, id uuid.UUID) (environment models.Environment, err error) {
	err = database.Where("id = ?", id).Preload("EnvironmentType").First(&environment).Error

	return
}

func UpdateEnvironment(database *gorm.DB, environment models.Environment) (err error) {
	err = database.Save(&environment).Error

	return
}

func DeleteEnvironment(database *gorm.DB, id uuid.UUID) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			var environment models.Environment
			err = database.Where("id = ?", id).First(&environment).Error
			if err == nil {
				err = database.Unscoped().Delete(&environment).Error
			}
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}
