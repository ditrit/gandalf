package dao

import (
	"errors"

	"github.com/ditrit/gandalf/core/aggregator/api/utils"

	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"
)

func ListConnector(database *gorm.DB) (connectors []models.Connector, err error) {
	err = database.Find(&connectors).Error

	return
}

func CreateConnector(database *gorm.DB, connector models.Connector) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			err = database.Create(&connector).Error
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}

func ReadConnector(database *gorm.DB, id int) (connector models.Connector, err error) {
	err = database.First(&connector, id).Error

	return
}

func ReadConnectorByName(database *gorm.DB, name string) (connector models.Connector, err error) {
	err = database.Where("logical_name = ?", name).First(&connector).Error

	return
}

func UpdateConnector(database *gorm.DB, connector models.Connector) (err error) {
	err = database.Save(&connector).Error

	return
}

func DeleteConnector(database *gorm.DB, id int) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			var connector models.Connector
			err = database.Delete(&connector, id).Error
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}
