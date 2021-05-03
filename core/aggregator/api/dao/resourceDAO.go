package dao

import (
	"errors"

	"github.com/ditrit/gandalf/core/aggregator/api/utils"

	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"
)

func ListResource(database *gorm.DB) (resources []models.Resource, err error) {
	err = database.Find(&resources).Error

	return
}

func CreateResource(database *gorm.DB, resource models.Resource) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			err = database.Create(&resource).Error
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}

func ReadResource(database *gorm.DB, id int) (resource models.Resource, err error) {
	err = database.First(&resource, id).Error

	return
}

func UpdateResource(database *gorm.DB, resource models.Resource) (err error) {
	err = database.Save(&resource).Error

	return
}

func DeleteResource(database *gorm.DB, id int) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			var resource models.Resource
			err = database.Delete(&resource, id).Error
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}