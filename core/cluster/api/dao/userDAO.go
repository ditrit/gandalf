package dao

import (
	"gandalf/core/models"

	"github.com/jinzhu/gorm"
)

type UserDAO struct {
}

func NewUserDAO(gandalfDatabase *gorm.DB) (userDAO *UserDAO) {
	userDAO = new(UserDAO)

	return
}

func (ud UserDAO) List(database *gorm.DB) (users []models.User, err error) {
	err = database.Find(&users).Error

	return
}

func (ud UserDAO) Create(database *gorm.DB, user models.User) (err error) {
	err = database.Create(&user).Error

	return
}

func (ud UserDAO) Read(database *gorm.DB, id int) (user models.User, err error) {
	err = database.First(&user, id).Error

	return
}

func (ud UserDAO) Update(database *gorm.DB, user models.User) (err error) {
	err = database.Save(&user).Error

	return
}

func (ud UserDAO) Delete(database *gorm.DB, id int) (err error) {
	var user models.User
	err = database.Delete(&user, id).Error

	return
}
