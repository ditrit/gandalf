//Package utils :
package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ditrit/gandalf/core/models"
	"gopkg.in/yaml.v2"

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
func GetLogicalComponents(client *gorm.DB, logicalName string) (logicalComponenets *models.LogicalComponent) {
	client.Where("logical_name = ?", logicalName).Preload("KeyValues").First(&logicalComponenets)

	return
}

func GetPivots(client *gorm.DB, componentType string, version models.Version) (pivot *models.Pivot) {
	client.Where("name = ? and major = ? and minor = ?", componentType, version.Major, version.Minor).Preload("ResourceTypes").Preload("CommandTypes").Preload("EventTypes").Preload("Keys").First(&pivot)

	return
}

func GetProductConnectors(client *gorm.DB, product string, version models.Version) (productConnector models.ProductConnector) {
	client.Where("product.name = ? and major = ? and minor = ?", product, version.Major, version.Minor).Preload("Product").Preload("ResourceTypes").Preload("CommandTypes").Preload("EventTypes").Preload("Keys").Preload("Ressources").Preload("EventTypeToPolls").First(&productConnector)

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

// DownloadPivot : Download pivot from url
func DownloadPivot(url, ressource string) (pivot *models.Pivot, err error) {

	resp, err := http.Get(url + ressource)
	if err != nil {
		log.Printf("err: %s", err)
		return
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(bodyBytes, &pivot)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

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

func ValidateSecret(databaseClient *gorm.DB, secret, bindAddress string) (result bool, err error) {

	result = false

	var secretAssignement models.SecretAssignement
	err = databaseClient.Where("secret = ?", secret).First(&secretAssignement).Error
	if err == nil {
		if secretAssignement != (models.SecretAssignement{}) {
			if secretAssignement.AddressIP == "" {
				secretAssignement.AddressIP = bindAddress
				databaseClient.Save(secretAssignement)
				result = true
			} else if secretAssignement.AddressIP == bindAddress {
				result = true
			}
		}
	}

	return
}
