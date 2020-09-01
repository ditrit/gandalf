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

func (ad AggregatorDAO) list() (aggregators []models.Aggregator) {
	ad.GandalfDatabase.Find(&aggregators)

	return
}

func (ad AggregatorDAO) create(aggregator models.Aggregator) {
	ad.GandalfDatabase.Create(&aggregator)

}

func (ad AggregatorDAO) read(id int) (aggregator models.Aggregator) {
	ad.GandalfDatabase.First(&aggregator, id)

	return
}

func (ad AggregatorDAO) update(aggregator models.Aggregator) {
	ad.GandalfDatabase.Save(&aggregator)
}

func (ad AggregatorDAO) delete(id int) {
	var aggregator models.Aggregator
	ad.GandalfDatabase.Delete(&aggregator, id)

}
