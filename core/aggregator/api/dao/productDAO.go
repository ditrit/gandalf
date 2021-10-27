package dao

import (
	"errors"

	"github.com/ditrit/gandalf/core/aggregator/api/utils"
	"github.com/google/uuid"

	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"
)

func ListProduct(database *gorm.DB) (products []models.Product, err error) {
	err = database.Preload("Domain").Find(&products).Error

	return
}

func CreateProduct(database *gorm.DB, product *models.Product) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			err = database.Create(&product).Error
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}

func ReadProduct(database *gorm.DB, id uuid.UUID) (product models.Product, err error) {
	err = database.Where("id = ?", id).Preload("Domain").First(&product).Error

	return
}

func UpdateProduct(database *gorm.DB, product models.Product) (err error) {
	err = database.Save(&product).Error

	return
}

func DeleteProduct(database *gorm.DB, id uuid.UUID) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			var product models.Role
			err = database.Where("id = ?", id).First(&product).Error
			if err == nil {
				err = database.Unscoped().Delete(&product).Error
			}
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}
