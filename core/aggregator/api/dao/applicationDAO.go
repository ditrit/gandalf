package dao

import (
	"errors"
	"fmt"

	"github.com/ditrit/gandalf/core/aggregator/api/utils"

	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"
)

func ListApplication(database *gorm.DB) (applications []models.Application, err error) {
	err = database.Find(&applications).Error

	return
}

func CreateApplication(database *gorm.DB, application models.Application) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			err = database.Create(&application).Error
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}

func ReadApplication(database *gorm.DB, id int) (application models.Application, err error) {
	err = database.First(&application, id).Error

	return
}

func ReadApplicationByName(database *gorm.DB, name string) (application models.Application, err error) {
	fmt.Println("DAO")
	err = database.Where("name = ?", name).First(&application).Error
	fmt.Println(err)
	fmt.Println(application)
	return
}

func UpdateApplication(database *gorm.DB, application models.Application) (err error) {
	err = database.Save(&application).Error

	return
}

func DeleteApplication(database *gorm.DB, id int) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			var application models.Application
			err = database.Delete(&application, id).Error
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}
