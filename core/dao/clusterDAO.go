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

func (cd ClusterDAO) list() (clusters []models.Cluster, err error) {
	err = cd.GandalfDatabase.Find(&clusters).Error

	return
}

func (cd ClusterDAO) create(cluster models.Cluster) (err error) {
	err = cd.GandalfDatabase.Create(&cluster).Error

	return

}

func (cd ClusterDAO) read(id int) (cluster models.Cluster, err error) {
	err = cd.GandalfDatabase.First(&cluster, id).Error

	return
}

func (cd ClusterDAO) update(cluster models.Cluster) (err error) {
	err = cd.GandalfDatabase.Save(&cluster).Error

	return
}

func (cd ClusterDAO) delete(id int) (err error) {
	var cluster models.Cluster
	err = cd.GandalfDatabase.Delete(&cluster, id).Error

	return
}
