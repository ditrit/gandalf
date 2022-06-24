package dao

import (
	"errors"

	"github.com/ditrit/gandalf/core/aggregator/api/utils"
	"github.com/google/uuid"

	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"
)

const INVALID_STATE = "invalid state"

func ListProduct(database *gorm.DB) (products []models.Product, err error) {
	err = database.Preload("Domain").Preload("Libraries").Preload("Authorizations.User").Preload("Authorizations.Role").Preload("Tags").Preload("Environments").Find(&products).Error

	return
}

func CreateProduct(database *gorm.DB, product *models.Product, parentDomainID uuid.UUID) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			product.DomainID = parentDomainID
			err = database.Save(&product).Error
		} else {
			err = errors.New(INVALID_STATE)
		}
	}

	return
}

func ReadProduct(database *gorm.DB, id uuid.UUID) (product models.Product, err error) {
	err = database.Where("id = ?", id).Preload("Domain").Preload("Libraries").Preload("Authorizations.User").Preload("Authorizations.Role").Preload("Tags").Preload("Environments.EnvironmentType").Preload("Environments.Product").First(&product).Error

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
			err = errors.New(INVALID_STATE)
		}
	}

	return
}

func AddProductLibrary(database *gorm.DB, product models.Product, library models.Library) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			err = database.Model(&product).Association("Libraries").Append(&library).Error
		} else {
			err = errors.New(INVALID_STATE)
		}
	}

	return
}

func AddProductTag(database *gorm.DB, product models.Product, tag models.Tag) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			err = database.Model(&product).Association("Tags").Append(&tag).Error
		} else {
			err = errors.New(INVALID_STATE)
		}
	}

	return
}

func RemoveProductTag(database *gorm.DB, product models.Product, tag models.Tag) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			err = database.Model(&product).Association("Tags").Delete(&tag).Error

		} else {
			err = errors.New(INVALID_STATE)
		}
	}

	return
}
