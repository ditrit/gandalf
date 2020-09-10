package dao

import (
	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"
)

func ListCluster(database *gorm.DB) (clusters []models.Cluster, err error) {
	err = database.Find(&clusters).Error

	return
}

func CreateCluster(database *gorm.DB, cluster models.Cluster) (err error) {
	err = database.Create(&cluster).Error

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
	var cluster models.Cluster
	err = database.Delete(&cluster, id).Error

	return
}
