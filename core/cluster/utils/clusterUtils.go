//Package utils :
package utils

import (
	"log"

	"github.com/ditrit/gandalf-core/database"
	"github.com/ditrit/gandalf-core/models"

	"github.com/ditrit/shoset/msg"

	"github.com/jinzhu/gorm"
)

var gandalfDatabaseClient *gorm.DB = nil

// GetDatabaseClientByTenant : Cluster database client getter by tenant.
func GetDatabaseClientByTenant(tenant, databasePath string, mapDatabaseClient map[string]*gorm.DB) *gorm.DB {
	if _, ok := mapDatabaseClient[tenant]; !ok {
		mapDatabaseClient[tenant] = database.NewTenantDatabaseClient(tenant, databasePath)
	}

	return mapDatabaseClient[tenant]
}

/* // GetGandalfDatabaseClient : Database client constructor.
func GetGandalfDatabaseClient(databasePath string) *gorm.DB {

	if gandalfDatabaseClient == nil {
		gandalfDatabaseClient = database.NewGandalfDatabaseClient(databasePath)
	}
	return gandalfDatabaseClient
} */

// GetApplicationContext : Cluster application context getter.
func GetApplicationContext(cmd msg.Command, client *gorm.DB) (applicationContext models.Application) {
	var connectorType models.ConnectorType
	client.Where("name = ?", cmd.GetContext()["connectorType"].(string)).First(&connectorType)

	//var tenant models.Tenant
	//client.Where("name = ?", cmd.GetTenant()).First(&tenant)

	client.Where("connector_type_id = ?", connectorType.ID).Preload("Aggregator").Preload("Connector").Preload("ConnectorType").First(&applicationContext)

	return
}

// GetConnectorConfiguration : Cluster application context getter.
func GetConnectorsConfiguration(conf msg.Config, client *gorm.DB) (connectorsConfiguration []models.ConnectorConfig) {
	//client.Where("connector_type = ?", cmd.GetContext()["ConnectorType"].(string)).First(&connectorConfiguration)
	//var connectorsType []models.ConnectorType

	//client.Where("name = ?", conf.GetContext()["connectorType"].(string)).First(&connectorType)

	client.Order("connector_type_id, connector_product_id, version desc").Preload("ConnectorType").Preload("ConnectorProduct").Preload("ConnectorTypeCommands").Preload("ConnectorTypeEvents").Find(&connectorsConfiguration)

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
