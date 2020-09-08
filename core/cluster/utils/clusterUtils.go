//Package utils :
package utils

import (
	"log"

	"github.com/ditrit/gandalf/core/cluster/database"
	"github.com/ditrit/gandalf/core/models"

	"github.com/ditrit/shoset/msg"

	"github.com/jinzhu/gorm"
)

// GetDatabaseClientByTenant : Cluster database client getter by tenant.
func GetDatabaseClientByTenant(tenant, databasePath string, mapDatabaseClient map[string]*gorm.DB) *gorm.DB {
	if _, ok := mapDatabaseClient[tenant]; !ok {
		mapDatabaseClient[tenant] = database.NewTenantDatabaseClient(tenant, databasePath)
	}

	return mapDatabaseClient[tenant]
}

// GetApplicationContext : Cluster application context getter.
func GetApplicationContext(cmd msg.Command, client *gorm.DB) (applicationContext models.Application) {
	var connectorType models.ConnectorType
	client.Where("name = ?", cmd.GetContext()["connectorType"].(string)).First(&connectorType)

	client.Where("connector_type_id = ?", connectorType.ID).Preload("Aggregator").Preload("Connector").Preload("ConnectorType").First(&applicationContext)

	return
}

// GetConnectorConfiguration : Cluster application context getter.
func GetConnectorsConfiguration(client *gorm.DB) (connectorsConfiguration []models.ConnectorConfig) {
	client.Order("connector_type_id, connector_product_id, version desc").Preload("ConnectorType").Preload("ConnectorProduct").Preload("ConnectorCommands").Preload("ConnectorEvents").Find(&connectorsConfiguration)

	return
}

// GetConnectorConfiguration : Cluster application context getter.
func SaveConnectorsConfiguration(connectorConfig *models.ConnectorConfig, client *gorm.DB) {
	client.Save(connectorConfig)

	return
}

// CaptureMessage : Cluster capture message function.
func CaptureMessage(message msg.Message, msgType string, client *gorm.DB) bool {
	ok := true

	switch msgType {
	case "cmd":
		currentMsg := models.FromShosetCommand(message.(msg.Command))
		client.Create(&currentMsg)
	case "evt":
		currentMsg := models.FromShosetEvent(message.(msg.Event))
		client.Create(&currentMsg)
	case "config":
		currentMsg := models.FromShosetConfig(message.(msg.Config))
		client.Create(&currentMsg)
	default:
		ok = false

		log.Println("Can't capture this message")
	}

	return ok
}

func ValidateSecret(gandalfDatabaseClient *gorm.DB, componentType, logicalName, tenant, secret string) (result bool, err error) {

	result = false

	switch componentType {
	case "cluster":
		var cluster models.Cluster
		err = gandalfDatabaseClient.Where("name = ? and secret = ?", logicalName, secret).First(&cluster).Error
		if err == nil {
			if cluster != (models.Cluster{}) {
				result = true
			}
		}
		break
	case "aggregator":
		var aggregator models.Aggregator
		err = gandalfDatabaseClient.Where("name = ? and tenant.name = ? and secret = ?", logicalName, tenant, secret).Preload("Tenant").First(&aggregator).Error
		if err == nil {
			if aggregator != (models.Aggregator{}) {
				result = true
			}
		}
		break
	case "connector":
		var connector models.Connector
		err = gandalfDatabaseClient.Where("name = ? and tenant.name = ? and secret = ?", logicalName, tenant, secret).Preload("Tenant").First(&connector).Error
		if err == nil {
			if connector != (models.Connector{}) {
				result = true
			}
		}
		break
	}

	return
}
