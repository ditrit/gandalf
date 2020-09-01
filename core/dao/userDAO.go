package dao

import (
	"gandalf/core/models"

	"github.com/jinzhu/gorm"
)

type UserDAO struct {
	GandalfDatabase *gorm.DB
}

func NewUserDAO(gandalfDatabase *gorm.DB) (userDAO *UserDAO) {
	userDAO = new(UserDAO)
	userDAO.GandalfDatabase = gandalfDatabase

	return
}

func (ud UserDAO) List() (users []models.User, err error) {
	err = ud.GandalfDatabase.Find(&users).Error

	return
}

func (ud UserDAO) Create(user models.User) (err error) {
	err = ud.GandalfDatabase.Create(&user).Error

	return
}

func (ud UserDAO) Read(id int) (user models.User, err error) {
	err = ud.GandalfDatabase.First(&user, id).Error

	return
}

func (ud UserDAO) Update(user models.User) (err error) {
	err = ud.GandalfDatabase.Save(&user).Error

	return
}

func (ud UserDAO) Delete(id int) (err error) {
	var user models.User
	err = ud.GandalfDatabase.Delete(&user, id).Error

	return
}
