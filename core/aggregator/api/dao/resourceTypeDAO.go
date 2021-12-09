package dao

import (
	"errors"
	"fmt"

	"github.com/ditrit/gandalf/core/aggregator/api/utils"
	"github.com/google/uuid"

	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"
)

func ListResourceType(database *gorm.DB) (resourceTypes []models.ResourceType, err error) {
	err = database.Find(&resourceTypes).Error

	return
}

func CreateResourceType(database *gorm.DB, resourceType *models.ResourceType) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			err = database.Create(&resourceType).Error
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}

func ReadResourceType(database *gorm.DB, id uuid.UUID) (resourceType models.ResourceType, err error) {
	err = database.Where("id = ?", id).First(&resourceType).Error

	return
}

func ReadResourceTypeByName(database *gorm.DB, name string) (resourceType models.ResourceType, err error) {
	fmt.Println("DAO")
	err = database.Where("name = ?", name).First(&resourceType).Error
	fmt.Println(err)
	fmt.Println(resourceType)
	return
}

func UpdateResourceType(database *gorm.DB, resourceType models.ResourceType) (err error) {
	err = database.Save(&resourceType).Error

	return
}

func DeleteResourceType(database *gorm.DB, id uuid.UUID) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			var resourceType models.ResourceType
			err = database.Where("id = ?", id).First(&resourceType).Error
			if err == nil {
				err = database.Unscoped().Delete(&resourceType).Error
			}
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}
