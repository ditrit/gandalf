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

func (cd ClusterDAO) list() (clusters []models.Cluster) {
	cd.GandalfDatabase.Find(&clusters)

	return
}

func (cd ClusterDAO) create(cluster models.Cluster) {
	cd.GandalfDatabase.Create(&cluster)

}

func (cd ClusterDAO) read(id int) (cluster models.Cluster) {
	cd.GandalfDatabase.First(&cluster, id)

	return
}

func (cd ClusterDAO) update(cluster models.Cluster) {
	cd.GandalfDatabase.Save(&cluster)
}

func (cd ClusterDAO) delete(id int) {
	var cluster models.Cluster
	cd.GandalfDatabase.Delete(&cluster, id)

}
