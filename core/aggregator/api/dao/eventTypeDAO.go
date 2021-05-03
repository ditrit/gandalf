package dao

import (
	"errors"
	"fmt"

	"github.com/ditrit/gandalf/core/aggregator/api/utils"

	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"
)

func ListEventType(database *gorm.DB) (eventTypes []models.EventType, err error) {
	err = database.Find(&eventTypes).Error

	return
}

func CreateEventType(database *gorm.DB, eventType models.EventType) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			err = database.Create(&eventType).Error
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}

func ReadEventType(database *gorm.DB, id int) (eventType models.EventType, err error) {
	err = database.First(&eventType, id).Error

	return
}

func ReadEventTypeByName(database *gorm.DB, name string) (eventType models.EventType, err error) {
	fmt.Println("DAO")
	err = database.Where("name = ?", name).First(&eventType).Error
	fmt.Println(err)
	fmt.Println(eventType)
	return
}

func UpdateEventType(database *gorm.DB, eventType models.EventType) (err error) {
	err = database.Save(&eventType).Error

	return
}

func DeleteEventType(database *gorm.DB, id int) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			var eventType models.EventType
			err = database.Delete(&eventType, id).Error
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}
