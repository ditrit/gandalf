package controllers

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/ditrit/gandalf/core/aggregator/database"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	"github.com/ditrit/gandalf/core/aggregator/api/dao"
	"github.com/ditrit/gandalf/core/aggregator/api/utils"
	"github.com/ditrit/gandalf/core/models"

	"gopkg.in/yaml.v2"
)

// ConfigurationAggregatorController :
type LogicalComponentController struct {
	databaseConnection *database.DatabaseConnection
}

// NewConfigurationController :
func NewLogicalComponentController(databaseConnection *database.DatabaseConnection) (logicalComponentController *LogicalComponentController) {
	logicalComponentController = new(LogicalComponentController)
	logicalComponentController.databaseConnection = databaseConnection

	return
}

// ReadByName :
func (lc LogicalComponentController) ReadByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	var logicalComponent models.LogicalComponent
	var err error
	if logicalComponent, err = dao.ReadLogicalComponentByName(lc.databaseConnection.GetTenantDatabaseClient(), name); err != nil {
		switch err {
		case sql.ErrNoRows:
			utils.RespondWithError(w, http.StatusNotFound, "Product not found")
		default:
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, logicalComponent)
}

func (lc LogicalComponentController) Upload(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	typeComponent := vars["type"]
	tenant := vars["tenant"]
	database := lc.databaseConnection.GetDatabaseClientByTenant(tenant)
	if database != nil {
		fmt.Println("File Upload Endpoint Hit")

		r.ParseMultipartForm(10 << 20)

		file, handler, err := r.FormFile("myFile")
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		defer file.Close()

		fmt.Printf("Uploaded File: %+v\n", handler.Filename)
		fmt.Printf("File Size: %+v\n", handler.Size)
		fmt.Printf("MIME Header: %+v\n", handler.Header)

		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		var logicalComponent *models.LogicalComponent
		err = yaml.Unmarshal(fileBytes, &logicalComponent)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		if typeComponent == "aggregator" {
			version := models.Version{Major: int8(logicalComponent.Pivot.Major), Minor: int8(logicalComponent.Pivot.Minor)}

			pivot, err := lc.GetPivot(database, lc.databaseConnection.GetLogicalComponent().GetKeyValueByKey("repository_url").Value, typeComponent, version)
			if err == nil {
				//TODO ADD VALIDATION
				lc.SaveAggregatorLogicalComponent(database, logicalComponent, pivot)
			} else {
				utils.RespondWithError(w, http.StatusInternalServerError, "pivot not found")
				return
			}

		} else if typeComponent == "connector" {
			version := models.Version{Major: int8(logicalComponent.ProductConnector.Major), Minor: int8(logicalComponent.ProductConnector.Minor)}
			fmt.Println("GET PIVOT")
			pivot, err := lc.GetPivot(database, lc.databaseConnection.GetLogicalComponent().GetKeyValueByKey("repository_url").Value, logicalComponent.ProductConnector.Name, version)
			if err == nil {
				var productConnector models.ProductConnector
				fmt.Println("GET PRODUCT CONNECTOR")
				productConnector, err = lc.GetProductConnector(database, lc.databaseConnection.GetLogicalComponent().GetKeyValueByKey("repository_url").Value, logicalComponent.ProductConnector.Name, logicalComponent.ProductConnector.Product.Name, version, pivot)
				if err == nil {
					//TODO ADD VALIDATION
					fmt.Println("SAVE")
					lc.SaveConnectorLogicalComponent(database, logicalComponent, productConnector)

				} else {
					utils.RespondWithError(w, http.StatusInternalServerError, "product connector not found")
					return
				}
			} else {
				utils.RespondWithError(w, http.StatusInternalServerError, "pivot not found")
				return
			}
		}

	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}

	fmt.Fprintf(w, "Successfully Uploaded File\n")

}

func (lc LogicalComponentController) GetPivot(client *gorm.DB, baseurl, componentType string, version models.Version) (models.Pivot, error) {
	var pivot models.Pivot
	err := client.Where("name = ? and major = ? and minor = ?", componentType, version.Major, version.Minor).Preload("ResourceTypes").Preload("CommandTypes").Preload("EventTypes").Preload("Keys").First(&pivot).Error
	fmt.Println("err pivot")
	fmt.Println(err)
	if err != nil {
		pivot, _ = lc.DownloadPivot(baseurl, "/configurations/"+strings.ToLower(componentType)+"/"+strconv.Itoa(int(version.Major))+"_"+strconv.Itoa(int(version.Minor))+"_pivot.yaml")
		client.Create(&pivot)
	}

	return pivot, nil
}

// DownloadPivot : Download pivot from url
func (lc LogicalComponentController) DownloadPivot(url, ressource string) (pivot models.Pivot, err error) {

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

func (lc LogicalComponentController) GetProductConnector(client *gorm.DB, baseurl, productType, product string, version models.Version, pivot models.Pivot) (models.ProductConnector, error) {
	var productConnector models.ProductConnector
	var productConnectors []models.ProductConnector
	var productDB models.Product
	err := client.Where("name = ?", product).First(&productDB).Error
	fmt.Println(productDB)
	fmt.Println(productDB.ID)
	fmt.Println("errproduct")
	fmt.Println(err)
	if err == nil {
		err := client.Find(&productConnectors).Error
		fmt.Println("productConnectors")
		fmt.Println(productConnectors)
		for _, toto := range productConnectors {
			fmt.Println(toto.ProductID)
		}
		//client.Joins("Product").Where("product.name = ? and major = ? and minor = ?", product, version.Major, version.Minor).Preload("Product").Preload("ResourceTypes").Preload("CommandTypes").Preload("EventTypes").Preload("Keys").First(&productConnector).Error

		err = client.Where("product_id = ? and major = ? and minor = ?", productDB.ID, version.Major, version.Minor).Preload("Product").Preload("ResourceTypes").Preload("CommandTypes").Preload("EventTypes").Preload("Keys").First(&productConnector).Error
		fmt.Println("err product connector")
		fmt.Println(err)
		fmt.Println("productConnectorDB")
		fmt.Println(productConnector)
		if err != nil {
			fmt.Println("DB PRODUCT")
			productConnector, _ = lc.DownloadProductConnector(baseurl, "/configurations/"+strings.ToLower(productType)+"/"+strings.ToLower(product)+"/"+strconv.Itoa(int(version.Major))+"_"+strconv.Itoa(int(version.Minor))+"_product_connector.yaml")
			fmt.Println("productConnector")
			fmt.Println(productConnector)
			productConnector.Pivot = pivot
			productConnector.Product = productDB
			fmt.Println("productConnector.Name")
			fmt.Println(productConnector.Name)
			client.Create(&productConnector)
		}
	} else {
		fmt.Println("LOGICAL PRODUCT")
		productConnector, _ = lc.DownloadProductConnector(baseurl, "/configurations/"+strings.ToLower(productType)+"/"+strings.ToLower(product)+"/"+strconv.Itoa(int(version.Major))+"_"+strconv.Itoa(int(version.Minor))+"_product_connector.yaml")
		fmt.Println("productConnector")
		fmt.Println(productConnector)
		productConnector.Pivot = pivot
		productDB.Name = product
		productConnector.Product = productDB
		fmt.Println("productConnector.Name")
		fmt.Println(productConnector.Name)
		fmt.Println(productConnector.ResourceTypes)
		fmt.Println(productConnector.Keys)
		fmt.Println(productConnector.CommandTypes)
		fmt.Println(productConnector.EventTypes)
		client.Create(&productConnector)

	}

	return productConnector, nil
}

// DownloadPivot : Download pivot from url
func (lc LogicalComponentController) DownloadProductConnector(url, ressource string) (productConnector models.ProductConnector, err error) {
	fmt.Println("url + ressource")
	fmt.Println(url + ressource)
	resp, err := http.Get(url + ressource)
	if err != nil {
		log.Printf("Error : %s", err)
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		fmt.Println("err200")
		fmt.Println(resp.StatusCode)
		return
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string("bodyBytes"))
	fmt.Println(string(bodyBytes))
	err = yaml.Unmarshal(bodyBytes, &productConnector)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	return
}

func (lc LogicalComponentController) SaveAggregatorLogicalComponent(client *gorm.DB, logicalComponent *models.LogicalComponent, pivot models.Pivot) error {
	fmt.Println("test")
	fmt.Println(logicalComponent.KeyValues)
	logicalComponent.Pivot = pivot
	for _, keyValue := range logicalComponent.KeyValues {
		for _, key := range pivot.Keys {
			if keyValue.Key.Name == key.Name {
				keyValue.Key = key
				break
			}
		}
	}
	fmt.Println("logicalComponent")
	fmt.Println(logicalComponent)
	err := client.Create(&logicalComponent).Error
	fmt.Println(err)
	return nil
}

func (lc LogicalComponentController) SaveConnectorLogicalComponent(client *gorm.DB, logicalComponent *models.LogicalComponent, productConnector models.ProductConnector) error {
	fmt.Println("test")
	fmt.Println(logicalComponent.KeyValues)
	logicalComponent.ProductConnector = productConnector
	fmt.Println("productConnector.Product")
	fmt.Println(productConnector.Product)
	for _, keyValue := range logicalComponent.KeyValues {
		for _, key := range productConnector.Pivot.Keys {
			if keyValue.Key.Name == key.Name {
				keyValue.Key = key
				break
			}
		}
		for _, key := range productConnector.Keys {
			if keyValue.Key.Name == key.Name {
				keyValue.Key = key
				break
			}
		}
	}
	fmt.Println("logicalComponent")
	fmt.Println(logicalComponent)
	err := client.Create(&logicalComponent).Error
	fmt.Println(err)
	return nil
}
