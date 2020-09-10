package dao

import (
	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"
)

func ListConnector(database *gorm.DB) (connectors []models.Connector, err error) {
	err = database.Find(&connectors).Error

	return
}

func CreateConnector(database *gorm.DB, connector models.Connector) (err error) {
	err = database.Create(&connector).Error

	return
}

func ReadConnector(database *gorm.DB, id int) (connector models.Connector, err error) {
	err = database.First(&connector, id).Error

	return
}

func UpdateConnector(database *gorm.DB, connector models.Connector) (err error) {
	err = database.Save(&connector).Error

	return
}

func DeleteConnector(database *gorm.DB, id int) (err error) {
	var connector models.Connector
	err = database.Delete(&connector, id).Error

	return
}
