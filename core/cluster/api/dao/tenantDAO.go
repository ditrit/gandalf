package dao

import (
	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"
)

type TenantDAO struct {
}

func NewTenantDAO() (tenantDAO *TenantDAO) {
	tenantDAO = new(TenantDAO)

	return
}

func (td TenantDAO) List(database *gorm.DB) (tenants []models.Tenant, err error) {
	err = database.Find(&tenants).Error

	return
}

func (td TenantDAO) Create(database *gorm.DB, tenant models.Tenant) (err error) {
	err = database.Create(&tenant).Error

	return
}

func (td TenantDAO) Read(database *gorm.DB, id int) (tenant models.Tenant, err error) {
	err = database.First(&tenant, id).Error

	return
}

func (td TenantDAO) Update(database *gorm.DB, tenant models.Tenant) (err error) {
	err = database.Save(&tenant).Error

	return
}

func (td TenantDAO) Delete(database *gorm.DB, id int) (err error) {
	var tenant models.Tenant
	err = database.Delete(&tenant, id).Error

	return
}
