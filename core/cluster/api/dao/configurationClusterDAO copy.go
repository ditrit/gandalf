package dao

import (
	"errors"

	"github.com/ditrit/gandalf/core/cluster/api/utils"

	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"
)

func ListConfigurationCluster(database *gorm.DB) (configurationClusters []models.ConfigurationCluster, err error) {
	err = database.Find(&configurationClusters).Error

	return
}

func CreateConfigurationCluster(database *gorm.DB, configurationCluster models.ConfigurationCluster) (err error) {
	admin, err := utils.GetStateGandalf(database)
	if err == nil {
		if admin {
			err = database.Create(&configurationCluster).Error
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}

func ReadConfigurationCluster(database *gorm.DB, id int) (configurationCluster models.ConfigurationCluster, err error) {
	err = database.First(&configurationCluster, id).Error

	return
}

func UpdateConfigurationCluster(database *gorm.DB, configurationCluster models.ConfigurationCluster) (err error) {
	err = database.Save(&configurationCluster).Error

	return
}

func DeleteConfigurationCluster(database *gorm.DB, id int) (err error) {
	admin, err := utils.GetStateGandalf(database)
	if err == nil {
		if admin {
			var configurationCluster models.ConfigurationCluster
			err = database.Delete(&configurationCluster, id).Error
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}
