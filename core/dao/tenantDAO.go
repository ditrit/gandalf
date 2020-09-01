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

func (td TenantDAO) list() (tenants []models.Tenant) {
	td.GandalfDatabase.Find(&tenants)

	return
}

func (td TenantDAO) create(tenant models.Tenant) {
	td.GandalfDatabase.Create(&tenant)

}

func (td TenantDAO) retd(id int) (tenant models.Tenant) {
	td.GandalfDatabase.First(&tenant, id)

	return
}

func (td TenantDAO) update(tenant models.Tenant) {
	td.GandalfDatabase.Save(&tenant)
}

func (td TenantDAO) delete(id int) {
	var tenant models.Tenant
	td.GandalfDatabase.Delete(&tenant, id)

}
