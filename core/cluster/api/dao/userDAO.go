package dao

import (
	"fmt"

	"github.com/ditrit/gandalf/core/models"

	"github.com/jinzhu/gorm"
)

func ListUser(database *gorm.DB) (users []models.User, err error) {
	err = database.Find(&users).Error

	return
}

func CreateUser(database *gorm.DB, user models.User) (err error) {
	err = database.Create(&user).Error

	return
}

func ReadUser(database *gorm.DB, id int) (user models.User, err error) {
	err = database.First(&user, id).Error

	return
}

func UpdateUser(database *gorm.DB, user models.User) (err error) {
	err = database.Save(&user).Error

	return
}

func DeleteUser(database *gorm.DB, id int) (err error) {
	var user models.User
	err = database.Delete(&user, id).Error

	return
}

func ReadUserByEmail(database *gorm.DB, email string) (user models.User, err error) {
	fmt.Println("DAO")
	err = database.Where("Email = ?", email).Preload("Role").First(&user).Error
	fmt.Println(err)
	fmt.Println(user)
	return
}
