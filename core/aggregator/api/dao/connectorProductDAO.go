package dao

import (
	"errors"

	"github.com/ditrit/gandalf/core/aggregator/api/utils"

	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"
)

func ListConnectorProduct(database *gorm.DB) (connectorProducts []models.ConnectorProduct, err error) {
	err = database.Find(&connectorProducts).Error

	return
}

func CreateConnectorProduct(database *gorm.DB, connectorProduct *models.ConnectorProduct) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			err = database.Create(&connectorProduct).Error
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}

func ReadConnectorProduct(database *gorm.DB, id int) (connectorProduct models.ConnectorProduct, err error) {
	err = database.First(&connectorProduct, id).Error

	return
}

func UpdateConnectorProduct(database *gorm.DB, connectorProduct models.ConnectorProduct) (err error) {
	err = database.Save(&connectorProduct).Error

	return
}

func DeleteConnectorProduct(database *gorm.DB, id int) (err error) {
	admin, err := utils.GetState(database)
	if err == nil {
		if admin {
			var connectorProduct models.Role
			err = database.Unscoped().Delete(&connectorProduct, id).Error
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}
