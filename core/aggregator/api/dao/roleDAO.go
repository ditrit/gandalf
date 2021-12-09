package dao

import (
	"errors"

	"github.com/ditrit/gandalf/core/aggregator/api/utils"
	"github.com/google/uuid"

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

func ReadRole(database *gorm.DB, id uuid.UUID) (role models.Role, err error) {
	err = database.Where("id = ?", id).First(&role).Error

	return
}

func UpdateRole(database *gorm.DB, role models.Role) (err error) {
	err = database.Save(&role).Error

	return
}

func DeleteRole(database *gorm.DB, id uuid.UUID) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			var role models.Role
			err = database.Where("id = ?", id).First(&role).Error
			if err == nil {
				err = database.Unscoped().Delete(&role).Error
			}
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}
