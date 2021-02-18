package dao

import (
	"errors"

	"github.com/ditrit/gandalf/core/aggregator/api/utils"

	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"
)

func ListAggregator(database *gorm.DB) (aggregators []models.Aggregator, err error) {
	err = database.Find(&aggregators).Error

	return
}

func CreateAggregator(database *gorm.DB, aggregator models.Aggregator) (err error) {
	admin, err := utils.GetStateGandalf(database)
	if err == nil {
		if admin {
			err = database.Create(&aggregator).Error
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}

func ReadAggregator(database *gorm.DB, id int) (aggregator models.Aggregator, err error) {
	err = database.First(&aggregator, id).Error

	return
}

func ReadAggregatorByName(database *gorm.DB, name string) (aggregator models.Aggregator, err error) {
	err = database.Where("logical_name = ?", name).First(&aggregator).Error

	return
}

func UpdateAggregator(database *gorm.DB, aggregator models.Aggregator) (err error) {
	err = database.Save(&aggregator).Error

	return
}

func DeleteAggregator(database *gorm.DB, id int) (err error) {
	admin, err := utils.GetStateGandalf(database)
	if err == nil {
		if admin {
			var aggregator models.Aggregator
			err = database.Delete(&aggregator, id).Error
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}
