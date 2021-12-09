package dao

import (
	"errors"

	"github.com/ditrit/gandalf/core/aggregator/api/utils"
	"github.com/google/uuid"

	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"
)

func ListEnvironmentType(database *gorm.DB) (environmentTypes []models.EnvironmentType, err error) {
	err = database.Find(&environmentTypes).Error

	return
}

func CreateEnvironmentType(database *gorm.DB, environmentType *models.EnvironmentType) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			err = database.Create(&environmentType).Error
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}

func ReadEnvironmentType(database *gorm.DB, id uuid.UUID) (environmentType models.EnvironmentType, err error) {
	err = database.Where("id = ?", id).First(&environmentType).Error

	return
}

func UpdateEnvironmentType(database *gorm.DB, environmentType models.EnvironmentType) (err error) {
	err = database.Save(&environmentType).Error

	return
}

func DeleteEnvironmentType(database *gorm.DB, id uuid.UUID) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			var environmentType models.EnvironmentType
			err = database.Where("id = ?", id).First(&environmentType).Error
			if err == nil {
				err = database.Unscoped().Delete(&environmentType).Error
			}
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}
