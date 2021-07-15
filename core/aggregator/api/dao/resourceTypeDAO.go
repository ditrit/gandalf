package dao

import (
	"errors"
	"fmt"

	"github.com/ditrit/gandalf/core/aggregator/api/utils"

	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"
)

func ListResourceType(database *gorm.DB) (resourceTypes []models.ResourceType, err error) {
	err = database.Find(&resourceTypes).Error

	return
}

func CreateResourceType(database *gorm.DB, resourceType models.ResourceType) (err error) {
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

func ReadResourceType(database *gorm.DB, id int) (resourceType models.ResourceType, err error) {
	err = database.First(&resourceType, id).Error

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

func DeleteResourceType(database *gorm.DB, id int) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			var resourceType models.ResourceType
			err = database.Unscoped().Delete(&resourceType, id).Error
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}
