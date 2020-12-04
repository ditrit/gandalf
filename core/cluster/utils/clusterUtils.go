//Package utils :
package utils

import (
	"fmt"
	"log"

	"github.com/ditrit/gandalf/core/cluster/database"

	"github.com/ditrit/gandalf/core/models"

	"github.com/ditrit/shoset/msg"

	"github.com/jinzhu/gorm"
)

// GetDatabaseClientByTenant : Cluster database client getter by tenant.
func GetDatabaseClientByTenant(tenant, databasePath string, mapDatabaseClient map[string]*gorm.DB) *gorm.DB {
	fmt.Println("databasePath")
	fmt.Println(databasePath)
	fmt.Println("tenant")
	fmt.Println(tenant)
	if _, ok := mapDatabaseClient[tenant]; !ok {
		var databaseCreated, err = database.IsDatabaseCreated(databasePath, tenant)
		if err == nil {
			fmt.Println("databaseCreated")
			fmt.Println(databaseCreated)
			if databaseCreated {
				var tenantDatabaseClient *gorm.DB
				tenantDatabaseClient, err = database.NewTenantDatabaseClient(tenant, databasePath)
				if err == nil {
					mapDatabaseClient[tenant] = tenantDatabaseClient
				} else {
					log.Println("Can't create database client")
				}
			} else {
				return nil
			}
		} else {
			log.Println("Can't detect if the database is created or not")
		}
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
	client.Order("connector_type_id, connector_product_id, major desc").Preload("ConnectorType").Preload("ConnectorProduct").Preload("ConnectorCommands").Preload("ConnectorEvents").Find(&connectorsConfiguration)

	return
}

// GetConnectorConfiguration : Cluster application context getter.
func SaveConnectorsConfiguration(connectorConfig *models.ConnectorConfig, client *gorm.DB) {
	//fmt.Println(connectorConfig.ConnectorEvents)
	//fmt.Println(connectorConfig.Resources)

	var connectorType models.ConnectorType
	client.Where("name = ?", connectorConfig.ConnectorType.Name).First(&connectorType)
	connectorConfig.ConnectorType = connectorType

	//ConnectorCommands
	var connectorCommands []models.Object
	for _, connectorCommand := range connectorConfig.ConnectorCommands {
		var listAction []models.Action
		for _, action := range connectorCommand.Actions {
			var currentAction models.Action
			client.Where("name = ?", action.Name).First(&currentAction)
			listAction = append(listAction, currentAction)
		}
		connectorCommand.Actions = listAction
		connectorCommands = append(connectorCommands, connectorCommand)
	}
	connectorConfig.ConnectorCommands = connectorCommands

	//ConnectorEvents
	var connectorEvents []models.Object
	for _, connectorEvent := range connectorConfig.ConnectorEvents {
		var listAction []models.Action
		for _, action := range connectorEvent.Actions {
			var currentAction models.Action
			client.Where("name = ?", action.Name).First(&currentAction)
			listAction = append(listAction, currentAction)
		}
		connectorEvent.Actions = listAction
		connectorEvents = append(connectorEvents, connectorEvent)
	}
	connectorConfig.ConnectorEvents = connectorEvents

	//Resources
	var resources []models.Object
	for _, resource := range connectorConfig.Resources {
		var listAction []models.Action
		for _, action := range resource.Actions {
			var currentAction models.Action
			client.Where("name = ?", action.Name).First(&currentAction)
			listAction = append(listAction, currentAction)
		}
		resource.Actions = listAction
		resources = append(resources, resource)
	}
	connectorConfig.Resources = resources

	client.Save(connectorConfig)

	var connectorTypes []models.ConnectorType
	client.Find(&connectorTypes)
	fmt.Println("connectorTypes")
	fmt.Println(connectorTypes)

	var connectorConfig2 models.ConnectorConfig
	client.Where("name = ?", "ConnectorConfig7").Preload("ConnectorType").Preload("Resources.Actions").First(&connectorConfig2)
	fmt.Println(connectorConfig2)

	var connectorConfig3 models.ConnectorConfig
	client.Where("name = ?", "ConnectorConfig6").Preload("ConnectorType").Preload("ConnectorCommands").Preload("ConnectorEvents").Preload("Resources").First(&connectorConfig3)
	fmt.Println(connectorConfig3)
	/* 	var actions []models.Action
	   	client.Find(&actions)
	   	fmt.Println("actions")
	   	fmt.Println(actions) */

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

func ValidateSecret(databaseClient *gorm.DB, componentType, logicalName, instanceName, secret string) (result bool, err error) {

	result = false

	switch componentType {
	case "cluster":
		var cluster models.Cluster
		fmt.Println(logicalName)
		fmt.Println(instanceName)
		err = databaseClient.Where("logical_name = ? and instance_name = ?", logicalName, instanceName).First(&cluster).Error
		fmt.Println("err")
		fmt.Println(err)
		if err == nil {
			if cluster != (models.Cluster{}) {
				if cluster.Secret == secret {
					result = true
				}
			}
		}
		break
	case "aggregator":
		var aggregator models.Aggregator
		err = databaseClient.Where("logical_name = ? and instance_name = ?", logicalName, instanceName).First(&aggregator).Error
		if err == nil {
			if aggregator != (models.Aggregator{}) {
				if aggregator.Secret == secret {
					result = true
				}
			}
		}

		break
	case "connector":
		var connector models.Connector
		err = databaseClient.Where("logical_name = ? and instance_name = ?", logicalName, instanceName).First(&connector).Error
		if err == nil {
			if connector != (models.Connector{}) {
				if connector.Secret == secret {
					result = true
				}
			}
		}

		break
	}

	return
}
