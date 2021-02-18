package dao

import (
	"errors"

	"github.com/ditrit/gandalf/core/aggregator/api/utils"

	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"
)

func ListConfigurationConnector(database *gorm.DB) (configurationConnectors []models.ConfigurationLogicalConnector, err error) {
	err = database.Find(&configurationConnectors).Error

	return
}

func CreateConfigurationConnector(database *gorm.DB, configurationConnector models.ConfigurationLogicalConnector) (err error) {
	admin, err := utils.GetStateGandalf(database)
	if err == nil {
		if admin {
			err = database.Create(&configurationConnector).Error
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}

func ReadConfigurationConnector(database *gorm.DB, id int) (configurationConnector models.ConfigurationLogicalConnector, err error) {
	err = database.First(&configurationConnector, id).Error

	return
}

func UpdateConfigurationConnector(database *gorm.DB, configurationConnector models.ConfigurationLogicalConnector) (err error) {
	err = database.Save(&configurationConnector).Error

	return
}

func DeleteConfigurationConnector(database *gorm.DB, id int) (err error) {
	admin, err := utils.GetStateGandalf(database)
	if err == nil {
		if admin {
			var configurationConnector models.ConfigurationLogicalConnector
			err = database.Delete(&configurationConnector, id).Error
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}
