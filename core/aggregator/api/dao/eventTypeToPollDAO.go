package dao

import (
	"errors"

	"github.com/ditrit/gandalf/core/aggregator/api/utils"

	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"
)

func ListEventTypeToPoll(database *gorm.DB) (eventTypeToPolls []models.EventTypeToPoll, err error) {
	err = database.Preload("Resource").Preload("EventType").Find(&eventTypeToPolls).Error

	return
}

func CreateEventTypeToPoll(database *gorm.DB, eventTypeToPoll models.EventTypeToPoll) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			err = database.Create(&eventTypeToPoll).Error
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}

func ReadEventTypeToPoll(database *gorm.DB, id int) (eventTypeToPoll models.EventTypeToPoll, err error) {
	err = database.First(&eventTypeToPoll, id).Error

	return
}

func UpdateEventTypeToPoll(database *gorm.DB, eventTypeToPoll models.EventTypeToPoll) (err error) {
	err = database.Save(&eventTypeToPoll).Error

	return
}

func DeleteEventTypeToPoll(database *gorm.DB, id int) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			var eventTypeToPoll models.EventTypeToPoll
			err = database.Delete(&eventTypeToPoll, id).Error
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}
