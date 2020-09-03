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

func (cd ConnectorDAO) List() (connectors []models.Connector, err error) {
	err = cd.GandalfDatabase.Find(&connectors).Error

	return
}

func (cd ConnectorDAO) Create(connector models.Connector) (err error) {
	err = cd.GandalfDatabase.Create(&connector).Error

	return
}

func (cd ConnectorDAO) Read(id int) (connector models.Connector, err error) {
	err = cd.GandalfDatabase.First(&connector, id).Error

	return
}

func (cd ConnectorDAO) Update(connector models.Connector) (err error) {
	err = cd.GandalfDatabase.Save(&connector).Error

	return
}

func (cd ConnectorDAO) Delete(id int) (err error) {
	var connector models.Connector
	err = cd.GandalfDatabase.Delete(&connector, id).Error

	return
}
