//Package utils :
package utils

import (
	"log"

	"github.com/ditrit/gandalf/core/models"

	"github.com/ditrit/shoset/msg"

	"github.com/jinzhu/gorm"
)

/* // GetDatabaseClientByTenant : Cluster database client getter by tenant.
func GetDatabaseClientByTenant(tenant, addr string, mapDatabaseClient map[string]*gorm.DB) *gorm.DB {
	if _, ok := mapDatabaseClient[tenant]; !ok {

		//var tenantDatabaseClient *gorm.DB
		tenantDatabaseClient, err := database.NewTenantDatabaseClient(addr, tenant)
		if err == nil {
			mapDatabaseClient[tenant] = tenantDatabaseClient
		} else {
			log.Println("Can't create database client")
			return nil
		}

	}

	return mapDatabaseClient[tenant]
} */

// GetApplicationContext : Cluster application context getter.
func GetApplicationContext(cmd msg.Command, client *gorm.DB) (applicationContext models.Application) {
	var connectorType models.ConnectorType
	client.Where("name = ?", cmd.GetContext()["connectorType"].(string)).First(&connectorType)

	client.Where("connector_type_id = ?", connectorType.ID).Preload("Aggregator").Preload("Connector").Preload("ConnectorType").First(&applicationContext)

	return
}

// GetConfigurationCluster :
func GetTenant(tenantName string, client *gorm.DB) (tenant models.Tenant, err error) {
	err = client.Where("name = ?", tenantName).First(&tenant).Error

	return
}

// GetConfigurationCluster :
func GetConfigurationCluster(logicalName string, client *gorm.DB) (configurationCluster models.ConfigurationLogicalCluster, err error) {
	err = client.Where("logical_name = ?", logicalName).First(&configurationCluster).Error

	return
}

// SaveConfigurationCluster :
func SaveConfigurationCluster(configurationCluster models.ConfigurationLogicalCluster, client *gorm.DB) (err error) {
	err = client.Create(&configurationCluster).Error

	return
}

// GetConfigurationAggregator :
func GetConfigurationAggregator(logicalName string, client *gorm.DB) (configurationAggregator models.ConfigurationLogicalAggregator, err error) {
	err = client.Where("logical_name = ?", logicalName).First(&configurationAggregator).Error

	return
}

// SaveConfigurationAggregator :
func SaveConfigurationAggregator(configurationAggregator models.ConfigurationLogicalAggregator, client *gorm.DB) (err error) {
	err = client.Create(&configurationAggregator).Error

	return
}

// GetConfigurationConnector :
func GetConfigurationConnector(logicalName string, client *gorm.DB) (configurationConnector models.ConfigurationLogicalConnector, err error) {
	err = client.Where("logical_name = ?", logicalName).First(&configurationConnector).Error

	return
}

// SaveConfigurationConnector :
func SaveConfigurationConnector(configurationConnector models.ConfigurationLogicalConnector, client *gorm.DB) (err error) {
	err = client.Create(&configurationConnector).Error

	return
}

/* // GetConnectorConfiguration : Cluster application context getter.
func GetConnectorsConfiguration(client *gorm.DB) (connectorsConfiguration []models.ConnectorConfig) {
	client.Order("connector_type_id, connector_product_id, major desc").Preload("ConnectorType").Preload("ConnectorProduct").Preload("ConnectorCommands").Preload("ConnectorEvents").Find(&connectorsConfiguration)

	return
} */
func GetLogicalComponents(client *gorm.DB, logicalName string) (logicalComponenets models.LogicalComponent) {
	client.Where("logical_name = ?", logicalName).Preload("KeyValues").First(&logicalComponenets)

	return
}

func GetPivots(client *gorm.DB, componentType string, version models.Version) (pivot models.Pivot) {
	client.Where("name = ? and major = ? and minor = ?", componentType, version.Major, version.Minor).Preload("ResourceTypes").Preload("CommandTypes").Preload("EventTypes").Preload("Keys").First(&pivot)

	return
}

func GetProductConnectors(client *gorm.DB, product string, version models.Version) (productConnector models.ProductConnector) {
	client.Where("product.name = ? and major = ? and minor = ?", product, version.Major, version.Minor).Preload("Product").Preload("ResourceTypes").Preload("CommandTypes").Preload("EventTypes").Preload("Keys").First(&productConnector)

	return
}

func SavePivot(pivot *models.Pivot, client *gorm.DB) {

	client.Save(pivot)
}

func SaveProductConnector(productConnector *models.ProductConnector, client *gorm.DB) {
	var product models.Product
	client.Where("name = ?", productConnector.Product.Name).First(&product)
	if (product != models.Product{}) {
		productConnector.Product = product
	}

	client.Save(productConnector)
}

/* // GetConnectorConfiguration : Cluster application context getter.
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
			fmt.Println("currentAction")
			fmt.Println(currentAction)
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
			fmt.Println("currentAction")
			fmt.Println(currentAction)
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
			fmt.Println("currentAction")
			fmt.Println(currentAction)
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

	return
} */

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

func ValidateSecret(databaseClient *gorm.DB, componentType, logicalName, secret, bindAddress string) (result bool, err error) {

	result = false

	switch componentType {
	case "cluster":
		var cluster models.Cluster
		err = databaseClient.Where("logical_name = ? and secret = ?", logicalName, secret).First(&cluster).Error
		if err == nil {
			if cluster != (models.Cluster{}) {
				if cluster.BindAddress == "" {
					cluster.BindAddress = bindAddress
					databaseClient.Save(cluster)
					result = true
				} else if cluster.BindAddress == bindAddress {
					result = true
				}
			}
		}
		break
	case "aggregator":
		var aggregator models.Aggregator
		err = databaseClient.Where("logical_name = ? and secret = ?", logicalName, secret).First(&aggregator).Error
		if err == nil {
			if aggregator != (models.Aggregator{}) {
				if aggregator.BindAddress == "" {
					aggregator.BindAddress = bindAddress
					databaseClient.Save(aggregator)
					result = true
				} else if aggregator.BindAddress == bindAddress {
					result = true
				}
			}
		}

		break
	case "connector":
		var connector models.Connector
		err = databaseClient.Where("logical_name = ? and secret = ?", logicalName, secret).First(&connector).Error
		if err == nil {
			if connector != (models.Connector{}) {
				if connector.BindAddress == "" {
					connector.BindAddress = bindAddress
					databaseClient.Save(connector)
					result = true
				} else if connector.BindAddress == bindAddress {
					result = true
				}
			}
		}

		break
	}

	return
}
