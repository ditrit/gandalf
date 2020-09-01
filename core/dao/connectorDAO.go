package dao

import (
	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"
)

type ConnectorDAO struct {
	GandalfDatabase *gorm.DB
}

func NewConnectorDAO(gandalfDatabase *gorm.DB) (connectorDAO *ConnectorDAO) {
	connectorDAO = new(ConnectorDAO)
	connectorDAO.GandalfDatabase = gandalfDatabase

	return
}

func (cd ConnectorDAO) list() (connectors []models.Connector) {
	cd.GandalfDatabase.Find(&connectors)

	return
}

func (cd ConnectorDAO) create(connector models.Connector) {
	cd.GandalfDatabase.Create(&connector)

}

func (cd ConnectorDAO) read(id int) (connector models.Connector) {
	cd.GandalfDatabase.First(&connector, id)

	return
}

func (cd ConnectorDAO) update(connector models.Connector) {
	cd.GandalfDatabase.Save(&connector)
}

func (cd ConnectorDAO) delete(id int) {
	var connector models.Connector
	cd.GandalfDatabase.Delete(&connector, id)

}
