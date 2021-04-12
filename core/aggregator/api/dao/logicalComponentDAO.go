package dao

import (
	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"
)

func ListLogicalComponent(database *gorm.DB) (logicalComponent []models.LogicalComponent, err error) {
	err = database.Find(&logicalComponent).Error

	return
}

func ListLogicalComponentConnector(database *gorm.DB) (logicalComponent []models.LogicalComponent, err error) {
	err = database.Where("type = ?", "connector").Find(&logicalComponent).Error

	return
}

func ListLogicalComponentAggregator(database *gorm.DB) (logicalComponent []models.LogicalComponent, err error) {
	err = database.Where("type = ?", "aggregator").Find(&logicalComponent).Error

	return
}
