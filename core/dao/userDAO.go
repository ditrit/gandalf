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

func (ud UserDAO) list() (users []models.User, err error) {
	err = ud.GandalfDatabase.Find(&users).Error

	return
}

func (ud UserDAO) create(user models.User) (err error) {
	err = ud.GandalfDatabase.Create(&user).Error

	return
}

func (ud UserDAO) read(id int) (user models.User, err error) {
	err = ud.GandalfDatabase.First(&user, id).Error

	return
}

func (ud UserDAO) update(user models.User) (err error) {
	err = ud.GandalfDatabase.Save(&user).Error

	return
}

func (ud UserDAO) delete(id int) (err error) {
	var user models.User
	err = ud.GandalfDatabase.Delete(&user, id).Error

	return
}
