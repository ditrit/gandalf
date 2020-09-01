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

func (ud UserDAO) list() (users []models.User) {
	ud.GandalfDatabase.Find(&users)

	return
}

func (ud UserDAO) create(user models.User) {
	ud.GandalfDatabase.Create(&user)

}

func (ud UserDAO) read(id int) (user models.User) {
	ud.GandalfDatabase.First(&user, id)

	return
}

func (ud UserDAO) update(user models.User) {
	ud.GandalfDatabase.Save(&user)
}

func (ud UserDAO) delete(id int) {
	var user models.User
	ud.GandalfDatabase.Delete(&user, id)

}
