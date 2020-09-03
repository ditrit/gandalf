package dao

import (
	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"
)

type ClusterDAO struct {
	GandalfDatabase *gorm.DB
}

func NewClusterDAO(gandalfDatabase *gorm.DB) (clusterDAO *ClusterDAO) {
	clusterDAO = new(ClusterDAO)
	clusterDAO.GandalfDatabase = gandalfDatabase

	return
}

func (cd ClusterDAO) List() (clusters []models.Cluster, err error) {
	err = cd.GandalfDatabase.Find(&clusters).Error

	return
}

func (cd ClusterDAO) Create(cluster models.Cluster) (err error) {
	err = cd.GandalfDatabase.Create(&cluster).Error

	return

}

func (cd ClusterDAO) Read(id int) (cluster models.Cluster, err error) {
	err = cd.GandalfDatabase.First(&cluster, id).Error

	return
}

func (cd ClusterDAO) Update(cluster models.Cluster) (err error) {
	err = cd.GandalfDatabase.Save(&cluster).Error

	return
}

func (cd ClusterDAO) Delete(id int) (err error) {
	var cluster models.Cluster
	err = cd.GandalfDatabase.Delete(&cluster, id).Error

	return
}
