package dao

import (
	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"
)

type RoleDAO struct {
}

func NewRoleDAO() (roleDAO *RoleDAO) {
	roleDAO = new(RoleDAO)

	return
}

func (rd RoleDAO) List(database *gorm.DB) (roles []models.Role, err error) {
	err = database.Find(&roles).Error

	return
}

func (rd RoleDAO) Create(database *gorm.DB, role models.Role) (err error) {
	err = database.Create(&role).Error

	return
}

func (rd RoleDAO) Read(database *gorm.DB, id int) (role models.Role, err error) {
	err = database.First(&role, id).Error

	return
}

func (rd RoleDAO) Update(database *gorm.DB, role models.Role) (err error) {
	err = database.Save(&role).Error

	return
}

func (rd RoleDAO) Delete(database *gorm.DB, id int) (err error) {
	var role models.Role
	err = database.Delete(&role, id).Error

	return
}
