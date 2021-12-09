package dao

import (
	"errors"
	"fmt"

	"github.com/ditrit/gandalf/core/aggregator/api/utils"
	"github.com/google/uuid"

	"github.com/ditrit/gandalf/core/models"

	"github.com/jinzhu/gorm"
)

func ListUser(database *gorm.DB) (users []models.User, err error) {
	err = database.Find(&users).Error

	return
}

func CreateUser(database *gorm.DB, user *models.User) (err error) {
	err = database.Create(&user).Error
	if err == nil {
		err = utils.ChangeStateTenant(database)
	}
	return
}

func ReadUser(database *gorm.DB, id uuid.UUID) (user models.User, err error) {
	err = database.Where("id = ?", id).First(&user).Error

	return
}

func UpdateUser(database *gorm.DB, user models.User) (err error) {
	err = database.Save(&user).Error

	return
}

func DeleteUser(database *gorm.DB, id uuid.UUID) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			var user models.User
			err = database.Where("id = ?", id).First(&user).Error
			if err == nil {
				err = database.Unscoped().Delete(&user).Error
			}
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}

func ReadAdminByName(database *gorm.DB, name string) (user models.User, err error) {
	fmt.Println("DAO")
	if err = database.Where("Name = ?", name).First(&user).Error; err != nil {
		var admin models.Role
		if err = database.Where("name = ?", "Administrator").First(&admin).Error; err != nil {
			var root models.Domain
			if err = database.Where("name = ?", "Root").First(&root).Error; err != nil {
				var authorizations models.Authorization
				err = database.Where("role_id = ? and domain_id = ? and user_id = ?", admin.ID, root.ID, user.ID).Preload("User").Preload("Role").Preload("Domain").Find(&authorizations).Error
			}
		}
	}
	fmt.Println(err)
	fmt.Println(user)
	return
}

func ReadUserByName(database *gorm.DB, name string) (user models.User, err error) {
	fmt.Println("DAO")
	err = database.Where("name = ?", name).First(&user).Error
	fmt.Println(err)
	fmt.Println(user)
	return
}

func ReadUserByEmail(database *gorm.DB, email string) (user models.User, err error) {
	fmt.Println("DAO")
	err = database.Where("email = ?", email).First(&user).Error
	fmt.Println(err)
	fmt.Println(user)
	return
}
