package dao

import (
	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"
)

type RoleDAO struct {
	GandalfDatabase *gorm.DB
}

func NewRoleDAO(gandalfDatabase *gorm.DB) (roleDAO *RoleDAO) {
	roleDAO = new(RoleDAO)
	roleDAO.GandalfDatabase = gandalfDatabase

	return
}

func (rd RoleDAO) list() (roles []models.Role) {
	rd.GandalfDatabase.Find(&roles)

	return
}

func (rd RoleDAO) create(role models.Role) {
	rd.GandalfDatabase.Create(&role)

}

func (rd RoleDAO) read(id int) (role models.Role) {
	rd.GandalfDatabase.First(&role, id)

	return
}

func (rd RoleDAO) update(role models.Role) {
	rd.GandalfDatabase.Save(&role)
}

func (rd RoleDAO) delete(id int) {
	var role models.Role
	rd.GandalfDatabase.Delete(&role, id)

}
