package dao

import (
	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"
)

type TenantDAO struct {
	GandalfDatabase *gorm.DB
}

func NewTenantDAO(gandalfDatabase *gorm.DB) (tenantDAO *TenantDAO) {
	tenantDAO = new(TenantDAO)
	tenantDAO.GandalfDatabase = gandalfDatabase

	return
}

func (td TenantDAO) List() (tenants []models.Tenant, err error) {
	err = td.GandalfDatabase.Find(&tenants).Error

	return
}

func (td TenantDAO) Create(tenant models.Tenant) (err error) {
	err = td.GandalfDatabase.Create(&tenant).Error

	return
}

func (td TenantDAO) Read(id int) (tenant models.Tenant, err error) {
	err = td.GandalfDatabase.First(&tenant, id).Error

	return
}

func (td TenantDAO) Update(tenant models.Tenant) (err error) {
	err = td.GandalfDatabase.Save(&tenant).Error

	return
}

func (td TenantDAO) Delete(id int) (err error) {
	var tenant models.Tenant
	err = td.GandalfDatabase.Delete(&tenant, id).Error

	return
}
