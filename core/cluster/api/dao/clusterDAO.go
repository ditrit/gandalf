package dao

import (
	"errors"

	"github.com/ditrit/gandalf/core/cluster/api/utils"

	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"
)

func ListCluster(database *gorm.DB) (clusters []models.Cluster, err error) {
	err = database.Find(&clusters).Error

	return
}

func CreateCluster(database *gorm.DB, cluster models.Cluster) (err error) {
	admin, err := utils.GetStateGandalf(database)
	if err == nil {
		if admin {
			err = database.Create(&cluster).Error
		} else {
			err = errors.New("Invalid state")
		}
	}

	return

}

func ReadCluster(database *gorm.DB, id int) (cluster models.Cluster, err error) {
	err = database.First(&cluster, id).Error

	return
}

func UpdateCluster(database *gorm.DB, cluster models.Cluster) (err error) {
	err = database.Save(&cluster).Error

	return
}

func DeleteCluster(database *gorm.DB, id int) (err error) {
	admin, err := utils.GetStateGandalf(database)
	if err == nil {
		if admin {
			var cluster models.Cluster
			err = database.Delete(&cluster, id).Error
		} else {
			err = errors.New("Invalid state")
		}
	}

	return
}
