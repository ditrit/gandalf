package dao

import (
	"errors"

	"github.com/ditrit/gandalf/core/aggregator/api/utils"

	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"
)

func ListDomainProduct(database *gorm.DB) (domainProducts []models.DomainProduct, err error) {
	err = database.Preload("Domain").Find(&domainProducts).Error

	return
}

func CreateDomainProduct(database *gorm.DB, domainProduct *models.DomainProduct) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			err = database.Create(&domainProduct).Error
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}

func ReadDomainProduct(database *gorm.DB, id int) (domainProduct models.DomainProduct, err error) {
	err = database.Preload("Domain").First(&domainProduct, id).Error

	return
}

func UpdateDomainProduct(database *gorm.DB, domainProduct models.DomainProduct) (err error) {
	err = database.Save(&domainProduct).Error

	return
}

func DeleteDomainProduct(database *gorm.DB, id int) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			var domainProduct models.Role
			err = database.Unscoped().Delete(&domainProduct, id).Error
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}
