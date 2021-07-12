//Package utils :
package utils

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

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

/* // GetApplicationContext : Cluster application context getter.
func GetApplicationContext(cmd msg.Command, client *gorm.DB) (applicationContext models.Application) {
	var connectorType models.ConnectorType
	client.Where("name = ?", cmd.GetContext()["connectorType"].(string)).First(&connectorType)

	client.Where("connector_type_id = ?", connectorType.ID).Preload("Aggregator").Preload("Connector").Preload("ConnectorType").First(&applicationContext)

	return
} */

// GetApplicationContext : Cluster application context getter.
func GetApplicationContext(cmd msg.Command, client *gorm.DB) (connector models.LogicalComponent) {
	connectorType := cmd.GetContext()["connectorType"].(string)
	var connectors []models.LogicalComponent
	client.Where("type = ?", "connector").Preload("ProductConnector.Pivot").Find(&connectors)
	fmt.Println("connectors")
	fmt.Println(connectors)

	for _, currentConnector := range connectors {
		if currentConnector.ProductConnector.Pivot.Name == connectorType {
			connector = currentConnector
			fmt.Println("FIND")
			return
		}
	}

	return
}

// GetConfigurationCluster :
func GetTenant(tenantName string, client *gorm.DB) (tenant models.Tenant, err error) {
	err = client.Where("name = ?", tenantName).First(&tenant).Error

	return
}

/* // GetConfigurationCluster :
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
}*/

/* // GetConnectorConfiguration : Cluster application context getter.
func GetConnectorsConfiguration(client *gorm.DB) (connectorsConfiguration []models.ConnectorConfig) {
	client.Order("connector_type_id, connector_product_id, major desc").Preload("ConnectorType").Preload("ConnectorProduct").Preload("ConnectorCommands").Preload("ConnectorEvents").Find(&connectorsConfiguration)

	return
} */
func GetLogicalComponents(client *gorm.DB, logicalName string) (logicalComponent models.LogicalComponent, err error) {
	err = client.Where("logical_name = ?", logicalName).Preload("KeyValues.Key").Preload("Resources.EventTypeToPolls.Resource").Preload("Resources.EventTypeToPolls.EventType").First(&logicalComponent).Error
	fmt.Println("logicalComponent")
	fmt.Println(logicalComponent)
	return
}

func GetPivots(client *gorm.DB, componentType string, version models.Version) (pivot models.Pivot, err error) {
	fmt.Println("GET PIVOT")
	fmt.Println(componentType)
	fmt.Println(version.Major)
	fmt.Println(version.Minor)
	err = client.Where("name = ? and major = ? and minor = ?", componentType, version.Major, version.Minor).Preload("ResourceTypes").Preload("CommandTypes").Preload("EventTypes").Preload("Keys").First(&pivot).Error
	fmt.Println(err)
	fmt.Println(pivot)
	return
}

func GetProductConnectors(client *gorm.DB, product string, version models.Version) (productConnector models.ProductConnector, err error) {
	var productdb models.Product
	err = client.Where("name = ?", product).First(&productdb).Error
	fmt.Println("productdb")
	fmt.Println(productdb)
	if err == nil {
		err = client.Where("product_id = ? and major = ? and minor = ?", productdb.ID, version.Major, version.Minor).Preload("Product").Preload("ResourceTypes").Preload("CommandTypes").Preload("EventTypes").Preload("Keys").First(&productConnector).Error
		fmt.Println("productConnector")
		fmt.Println(productConnector)
		//client.Where("product.name = ? and major = ? and minor = ?", product, version.Major, version.Minor).Preload("Product").Preload("ResourceTypes").Preload("CommandTypes").Preload("EventTypes").Preload("Keys").Preload("Ressources").Preload("EventTypeToPolls").First(&productConnector)

	}

	return
}

func SaveOrUpdateHeartbeat(heartbeat models.Heartbeat, client *gorm.DB) {
	//IF ALREADY EXIST : UPDATE
	var heartbeatdb models.Heartbeat
	err := client.Where("logical_name = ?, type = ?, address = ?", heartbeat.LogicalName, heartbeat.Type, heartbeat.Address).First(&heartbeatdb).Error
	if err == nil {
		heartbeatdb.UpdatedAt = time.Now()
		client.Save(&heartbeatdb)
	} else {
		heartbeat.CreatedAt = time.Now()
		heartbeat.UpdatedAt = time.Now()
		//IF NOT : SAVE
		client.Save(&heartbeat)
	}

}

func SavePivot(pivot models.Pivot, client *gorm.DB) {
	client.Save(&pivot)
}

func SaveProductConnector(productConnector *models.ProductConnector, client *gorm.DB) {
	var product models.Product
	client.Where("name = ?", productConnector.Product.Name).First(&product)
	if (product != models.Product{}) {
		productConnector.Product = product
	}

	client.Save(productConnector)
}

func CreateAggregatorLogicalComponent(logicalName, repositoryURL string, pivot *models.Pivot) *models.LogicalComponent {
	logicalComponent := new(models.LogicalComponent)
	logicalComponent.LogicalName = logicalName
	logicalComponent.Type = "aggregator"
	logicalComponent.Pivot = *pivot
	var keyValues []models.KeyValue
	for _, key := range pivot.Keys {
		keyValue := new(models.KeyValue)
		switch key.Name {
		case "repository_url":
			keyValue.Value = repositoryURL
			keyValue.Key = key
			keyValues = append(keyValues, *keyValue)
		}

	}

	logicalComponent.KeyValues = keyValues

	return logicalComponent
}

func GetAggregatorPivot(baseurl, componentType string, version models.Version) (*models.Pivot, error) {

	pivot, _ := DownloadPivot(baseurl, "/configurations/"+strings.ToLower(componentType)+"/"+strconv.Itoa(int(version.Major))+"_"+strconv.Itoa(int(version.Minor))+"_pivot.yaml")

	return pivot, nil
}

// DownloadPivot : Download pivot from url
func DownloadPivot(url, ressource string) (pivot *models.Pivot, err error) {

	resp, err := http.Get(url + ressource)
	if err != nil {
		log.Printf("Error : %s", err)
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

		log.Println("Error : Can't capture message")
	}

	return ok
}

func ValidateSecret(databaseClient *gorm.DB, secret, bindAddress string) (result bool, err error) {

	result = false

	var secretAssignements []models.SecretAssignement
	err = databaseClient.Find(&secretAssignements).Error
	fmt.Println(err)
	fmt.Println("secretAssignments")
	fmt.Println(secretAssignements)
	var secretAssignement models.SecretAssignement
	err = databaseClient.Where("secret = ?", secret).First(&secretAssignement).Error
	fmt.Println("err")
	fmt.Println(err)
	fmt.Println(secretAssignement)
	if err == nil {
		if secretAssignement != (models.SecretAssignement{}) {
			if secretAssignement.AddressIP == "" {
				//secretAssignement.AddressIP = bindAddress
				//databaseClient.Save(secretAssignement)
				databaseClient.Model(&models.SecretAssignement{}).Where("secret = ?", secretAssignement.Secret).Update("address_ip", bindAddress)
				err = databaseClient.Where("secret = ?", secret).First(&secretAssignement).Error
				fmt.Println("err2")
				fmt.Println(err)
				fmt.Println(secretAssignement)
				result = true
			} else if secretAssignement.AddressIP == bindAddress {
				result = true
			}
		}
	}

	return
}

func GenerateHash() string {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	concatenated := fmt.Sprint(random.Intn(100))
	sha512 := sha512.New()
	sha512.Write([]byte(concatenated))
	hash := base64.URLEncoding.EncodeToString(sha512.Sum(nil))
	return hash
}
