package dao

import (
	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"
)

type AggregatorDAO struct {
	GandalfDatabase *gorm.DB
}

func NewAggregatorDAO(gandalfDatabase *gorm.DB) (aggregatorDAO *AggregatorDAO) {
	aggregatorDAO = new(AggregatorDAO)
	aggregatorDAO.GandalfDatabase = gandalfDatabase

	return
}

func (ad AggregatorDAO) List() (aggregators []models.Aggregator, err error) {
	err = ad.GandalfDatabase.Find(&aggregators).Error

	return
}

func (ad AggregatorDAO) Create(aggregator models.Aggregator) (err error) {
	err = ad.GandalfDatabase.Create(&aggregator).Error

	return
}

func (ad AggregatorDAO) Read(id int) (aggregator models.Aggregator, err error) {
	err = ad.GandalfDatabase.First(&aggregator, id).Error

	return
}

func (ad AggregatorDAO) Update(aggregator models.Aggregator) (err error) {
	err = ad.GandalfDatabase.Save(&aggregator).Error

	return
}

func (ad AggregatorDAO) Delete(id int) (err error) {
	var aggregator models.Aggregator
	err = ad.GandalfDatabase.Delete(&aggregator, id).Error

	return
}
