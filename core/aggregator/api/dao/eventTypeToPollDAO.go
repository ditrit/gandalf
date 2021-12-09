package dao

import (
	"errors"

	"github.com/ditrit/gandalf/core/aggregator/api/utils"
	"github.com/google/uuid"

	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"
)

func ListEventTypeToPoll(database *gorm.DB) (eventTypeToPolls []models.EventTypeToPoll, err error) {
	err = database.Preload("Resource").Preload("EventType").Find(&eventTypeToPolls).Error

	return
}

func CreateEventTypeToPoll(database *gorm.DB, eventTypeToPoll *models.EventTypeToPoll) (err error) {
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

func ReadEventTypeToPoll(database *gorm.DB, id uuid.UUID) (eventTypeToPoll models.EventTypeToPoll, err error) {
	err = database.Where("id = ?", id).First(&eventTypeToPoll).Error

	return
}

func UpdateEventTypeToPoll(database *gorm.DB, eventTypeToPoll models.EventTypeToPoll) (err error) {
	err = database.Save(&eventTypeToPoll).Error

	return
}

func DeleteEventTypeToPoll(database *gorm.DB, id uuid.UUID) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			var eventTypeToPoll models.EventTypeToPoll
			err = database.Where("id = ?", id).First(&eventTypeToPoll).Error
			if err == nil {
				err = database.Unscoped().Delete(&eventTypeToPoll).Error
			}
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}
