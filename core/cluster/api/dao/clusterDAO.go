package dao

import (
	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"
)

type ClusterDAO struct {
}

func NewClusterDAO() (clusterDAO *ClusterDAO) {
	clusterDAO = new(ClusterDAO)

	return
}

func (cd ClusterDAO) List(database *gorm.DB) (clusters []models.Cluster, err error) {
	err = database.Find(&clusters).Error

	return
}

func (cd ClusterDAO) Create(database *gorm.DB, cluster models.Cluster) (err error) {
	err = database.Create(&cluster).Error

	return

}

func (cd ClusterDAO) Read(database *gorm.DB, id int) (cluster models.Cluster, err error) {
	err = database.First(&cluster, id).Error

	return
}

func (cd ClusterDAO) Update(database *gorm.DB, cluster models.Cluster) (err error) {
	err = database.Save(&cluster).Error

	return
}

func (cd ClusterDAO) Delete(database *gorm.DB, id int) (err error) {
	var cluster models.Cluster
	err = database.Delete(&cluster, id).Error

	return
}
