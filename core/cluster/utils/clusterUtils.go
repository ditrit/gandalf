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

func GetLogicalComponents(client *gorm.DB, logicalName string) (logicalComponent models.LogicalComponent, err error) {
	fmt.Println("logicalName")
	fmt.Println(logicalName)
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
	var productdb models.ConnectorProduct
	err = client.Where("name = ?", product).First(&productdb).Error
	fmt.Println("productdb")
	fmt.Println(productdb)
	if err == nil {
		err = client.Where("product_id = ? and major = ? and minor = ?", productdb.ID, version.Major, version.Minor).Preload("Product").Preload("ResourceTypes").Preload("CommandTypes").Preload("EventTypes").Preload("Keys").First(&productConnector).Error
		fmt.Println("productConnector")
		fmt.Println(productConnector)
	}

	return
}

func SaveOrUpdateHeartbeat(heartbeat models.Heartbeat, client *gorm.DB) {
	//IF ALREADY EXIST : UPDATE
	var heartbeatdb models.Heartbeat
	err := client.Where("logical_name = ? AND type = ? AND address = ?", heartbeat.LogicalName, heartbeat.Type, heartbeat.Address).First(&heartbeatdb).Error
	if err == nil {
		heartbeatdb.UpdatedAt = time.Now()
		fmt.Println("UPDATE HEARTBEAT")
		client.Save(&heartbeatdb)
	} else {
		heartbeat.CreatedAt = time.Now()
		heartbeat.UpdatedAt = time.Now()
		//IF NOT : SAVE
		fmt.Println("SAVE HEARTBEAT")
		client.Save(&heartbeat)
	}
}

func SavePivot(pivot models.Pivot, client *gorm.DB) {
	client.Save(&pivot)
}

func SaveProductConnector(productConnector *models.ProductConnector, client *gorm.DB) {
	var product models.ConnectorProduct
	client.Where("name = ?", productConnector.Product.Name).First(&product)
	if (product != models.ConnectorProduct{}) {
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
