package dao

import (
	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"
)

func ListAggregator(database *gorm.DB) (aggregators []models.Aggregator, err error) {
	err = database.Find(&aggregators).Error

	return
}

func CreateAggregator(database *gorm.DB, aggregator models.Aggregator) (err error) {
	err = database.Create(&aggregator).Error

	return
}

func ReadAggregator(database *gorm.DB, id int) (aggregator models.Aggregator, err error) {
	err = database.First(&aggregator, id).Error

	return
}

func UpdateAggregator(database *gorm.DB, aggregator models.Aggregator) (err error) {
	err = database.Save(&aggregator).Error

	return
}

func DeleteAggregator(database *gorm.DB, id int) (err error) {
	var aggregator models.Aggregator
	err = database.Delete(&aggregator, id).Error

	return
}
