package dao

import (
	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"
)

func ListTenant(database *gorm.DB) (tenants []models.Tenant, err error) {
	err = database.Find(&tenants).Error

	return
}

func CreateTenant(database *gorm.DB, tenant models.Tenant) (err error) {
	err = database.Create(&tenant).Error

	return
}

func ReadTenant(database *gorm.DB, id int) (tenant models.Tenant, err error) {
	err = database.First(&tenant, id).Error

	return
}

func UpdateTenant(database *gorm.DB, tenant models.Tenant) (err error) {
	err = database.Save(&tenant).Error

	return
}

func DeleteTenant(database *gorm.DB, id int) (err error) {
	var tenant models.Tenant
	err = database.Delete(&tenant, id).Error

	return
}
