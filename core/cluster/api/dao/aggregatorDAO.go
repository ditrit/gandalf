package dao

import (
	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"
)

type AggregatorDAO struct {
}

func NewAggregatorDAO() (aggregatorDAO *AggregatorDAO) {
	aggregatorDAO = new(AggregatorDAO)

	return
}

func (ad AggregatorDAO) List(database *gorm.DB) (aggregators []models.Aggregator, err error) {
	err = database.Find(&aggregators).Error

	return
}

func (ad AggregatorDAO) Create(database *gorm.DB, aggregator models.Aggregator) (err error) {
	err = database.Create(&aggregator).Error

	return
}

func (ad AggregatorDAO) Read(database *gorm.DB, id int) (aggregator models.Aggregator, err error) {
	err = database.First(&aggregator, id).Error

	return
}

func (ad AggregatorDAO) Update(database *gorm.DB, aggregator models.Aggregator) (err error) {
	err = database.Save(&aggregator).Error

	return
}

func (ad AggregatorDAO) Delete(database *gorm.DB, id int) (err error) {
	var aggregator models.Aggregator
	err = database.Delete(&aggregator, id).Error

	return
}
