package dao

import (
	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"
)

type ConnectorDAO struct {
}

func NewConnectorDAO() (connectorDAO *ConnectorDAO) {
	connectorDAO = new(ConnectorDAO)

	return
}

func (cd ConnectorDAO) List(database *gorm.DB) (connectors []models.Connector, err error) {
	err = database.Find(&connectors).Error

	return
}

func (cd ConnectorDAO) Create(database *gorm.DB, connector models.Connector) (err error) {
	err = database.Create(&connector).Error

	return
}

func (cd ConnectorDAO) Read(database *gorm.DB, id int) (connector models.Connector, err error) {
	err = database.First(&connector, id).Error

	return
}

func (cd ConnectorDAO) Update(database *gorm.DB, connector models.Connector) (err error) {
	err = database.Save(&connector).Error

	return
}

func (cd ConnectorDAO) Delete(database *gorm.DB, id int) (err error) {
	var connector models.Connector
	err = database.Delete(&connector, id).Error

	return
}
