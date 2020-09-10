package dao

import (
	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"
)

func ListRole(database *gorm.DB) (roles []models.Role, err error) {
	err = database.Find(&roles).Error

	return
}

func CreateRole(database *gorm.DB, role models.Role) (err error) {
	err = database.Create(&role).Error

	return
}

func ReadRole(database *gorm.DB, id int) (role models.Role, err error) {
	err = database.First(&role, id).Error

	return
}

func UpdateRole(database *gorm.DB, role models.Role) (err error) {
	err = database.Save(&role).Error

	return
}

func DeleteRole(database *gorm.DB, id int) (err error) {
	var role models.Role
	err = database.Delete(&role, id).Error

	return
}
