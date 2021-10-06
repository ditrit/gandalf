package dao

import (
	"errors"

	"github.com/ditrit/gandalf/core/aggregator/api/utils"

	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"
)

func ListDomainLibrary(database *gorm.DB) (domainLibrarys []models.DomainLibrary, err error) {
	err = database.Preload("Domain").Find(&domainLibrarys).Error

	return
}

func CreateDomainLibrary(database *gorm.DB, domainLibrary *models.DomainLibrary) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			err = database.Create(&domainLibrary).Error
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}

func ReadDomainLibrary(database *gorm.DB, id int) (domainLibrary models.DomainLibrary, err error) {
	err = database.Preload("Domain").First(&domainLibrary, id).Error

	return
}

func UpdateDomainLibrary(database *gorm.DB, domainLibrary models.DomainLibrary) (err error) {
	err = database.Save(&domainLibrary).Error

	return
}

func DeleteDomainLibrary(database *gorm.DB, id int) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			var domainLibrary models.Role
			err = database.Unscoped().Delete(&domainLibrary, id).Error
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}
