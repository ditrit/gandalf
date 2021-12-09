package dao

import (
	"errors"

	"github.com/ditrit/gandalf/core/aggregator/api/utils"
	"github.com/google/uuid"

	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"
)

func ListTenant(database *gorm.DB) (tenants []models.Tenant, err error) {
	err = database.Find(&tenants).Error

	return
}

func CreateTenant(database *gorm.DB, tenant *models.Tenant) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			err = database.Create(&tenant).Error
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}

func ReadTenant(database *gorm.DB, id uuid.UUID) (tenant models.Tenant, err error) {
	err = database.Where("id = ?", id).First(&tenant).Error

	return
}

func UpdateTenant(database *gorm.DB, tenant models.Tenant) (err error) {
	err = database.Save(&tenant).Error

	return
}

func DeleteTenant(database *gorm.DB, id uuid.UUID) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			var tenant models.Tenant
			err = database.Where("id = ?", id).First(&tenant).Error
			if err == nil {
				err = database.Unscoped().Delete(&tenant).Error
			}
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}
