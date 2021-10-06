package dao

import (
	"errors"

	"github.com/ditrit/gandalf/core/aggregator/api/utils"

	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"
)

func ListRole(database *gorm.DB) (roles []models.Role, err error) {
	err = database.Find(&roles).Error

	return
}

func CreateRole(database *gorm.DB, role *models.Role) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			err = database.Create(&role).Error
		} else {
			err = errors.New("Invalid state")
		}
	}

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
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			var role models.Role
			err = database.Unscoped().Delete(&role, id).Error
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}
