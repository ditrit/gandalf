/*
 * Swagger Gandalf
 *
 * This is a sample Petstore server.  You can find  out more about Swagger at  [http://swagger.io](http://swagger.io) or on  [irc.freenode.net, #swagger](http://swagger.io/irc/).
 *
 * API version: 1.0.0-oas3
 * Contact: romain.fairant@orness.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ditrit/gandalf/core/aggregator/shoset"

	"github.com/ditrit/gandalf/core/aggregator/api/dao"
	"github.com/ditrit/gandalf/core/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"gopkg.in/yaml.v2"

	"github.com/ditrit/gandalf/core/aggregator/api/utils"
)

func CreateTenant(w http.ResponseWriter, r *http.Request) {

	var result map[string]interface{}
	result = make(map[string]interface{})
	var err error
	var tenant models.Tenant
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&tenant); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	tenant.Password = utils.GenerateHash()

	err = dao.CreateTenant(utils.DatabaseConnection.GetTenantDatabaseClient(), tenant)
	fmt.Println("tenant.Password")
	fmt.Println(tenant.Password)
	fmt.Println(err)
	var createDatabase *models.CreateDatabase
	createDatabase, err = TenantCreateDatabase(tenant.Name)
	fmt.Println("createDatabase0")
	fmt.Println(createDatabase)
	fmt.Println(err)
	if err == nil {
		fmt.Println("CREATE 0")

		tenantDatabaseClient := utils.DatabaseConnection.GetDatabaseClientByTenant(tenant.Name)
		//tc.mapTenantDatabase[tenant.Name] = tenantDatabaseClient
		fmt.Println("CREATE 1")
		if tenantDatabaseClient != nil {

			//utils.ChangeStateTenant(tenantDatabaseClient)

			if err == nil {

				//CREATE SECRET
				var secretAssignement models.SecretAssignement
				secretAssignement.Secret = utils.GenerateHash()
				err := dao.CreateSecretAssignement(utils.DatabaseConnection.GetTenantDatabaseClient(), secretAssignement)
				fmt.Println("CREATE 2")
				if err == nil {
					//GET PIVOT AGGREGATOR
					version := models.Version{Major: utils.DatabaseConnection.GetPivot().Major, Minor: utils.DatabaseConnection.GetPivot().Minor}
					fmt.Println("test")
					fmt.Println(version)
					pivot, _ := TenantGetPivot(tenantDatabaseClient, utils.DatabaseConnection.GetLogicalComponent().GetKeyValueByKey("repository_url").Value, "aggregator", version)
					fmt.Println("CREATE 3")
					//CREATE AGGREGATOR LOGICAL COMPONENT
					logicalComponent, _ := TenantSaveLogicalComponent(tenantDatabaseClient, tenant.Name, utils.DatabaseConnection.GetLogicalComponent().GetKeyValueByKey("repository_url").Value, pivot)
					fmt.Println(logicalComponent)
					fmt.Println("CREATE 4")
				} else {
					dao.DeleteTenant(utils.DatabaseConnection.GetTenantDatabaseClient(), int(tenant.ID))
					utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
					return
				}
				fmt.Println("CREATE 5")

				result["login"] = createDatabase.Login
				result["password"] = createDatabase.Password
				result["tenant"] = tenant

			} else {
				dao.DeleteTenant(utils.DatabaseConnection.GetTenantDatabaseClient(), int(tenant.ID))
				utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
				return
			}

		} else {

		}
	} else {

	}

	utils.RespondWithJSON(w, http.StatusCreated, result)
}

func DeleteTenant(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["tenantId"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
		return
	}

	if err := dao.DeleteTenant(utils.DatabaseConnection.GetTenantDatabaseClient(), id); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func GetTenantById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["tenantId"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var tenant models.Tenant
	if tenant, err = dao.ReadTenant(utils.DatabaseConnection.GetTenantDatabaseClient(), id); err != nil {
		switch err {
		case sql.ErrNoRows:
			utils.RespondWithError(w, http.StatusNotFound, "Product not found")
		default:
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, tenant)
}

func ListTenant(w http.ResponseWriter, r *http.Request) {

	tenants, err := dao.ListTenant(utils.DatabaseConnection.GetTenantDatabaseClient())
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, tenants)
}

func UpdateTenant(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["tenantId"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var tenant models.Tenant
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&tenant); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	tenant.ID = uint(id)

	if err := dao.UpdateTenant(utils.DatabaseConnection.GetTenantDatabaseClient(), tenant); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, tenant)
}

func TenantGetPivot(client *gorm.DB, baseurl, componentType string, version models.Version) (models.Pivot, error) {
	var pivot models.Pivot
	err := client.Where("name = ? and major = ? and minor = ?", componentType, version.Major, version.Minor).Preload("ResourceTypes").Preload("CommandTypes").Preload("EventTypes").Preload("Keys").First(&pivot).Error
	fmt.Println(err)
	if err != nil {
		pivot, _ = TenantDownloadPivot(baseurl, "/configurations/"+strings.ToLower(componentType)+"/"+strconv.Itoa(int(version.Major))+"_"+strconv.Itoa(int(version.Minor))+"_pivot.yaml")
		client.Create(&pivot)
	}

	return pivot, nil
}

// DownloadPivot : Download pivot from url
func TenantDownloadPivot(url, ressource string) (pivot models.Pivot, err error) {

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

func TenantSaveLogicalComponent(client *gorm.DB, logicalName, repositoryURL string, pivot models.Pivot) (*models.LogicalComponent, error) {
	logicalComponent := new(models.LogicalComponent)
	logicalComponent.LogicalName = logicalName
	logicalComponent.Type = "aggregator"
	logicalComponent.Pivot = pivot
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

	client.Create(&logicalComponent)

	return logicalComponent, nil
}

func TenantCreateDatabase(tenant string) (*models.CreateDatabase, error) {
	fmt.Println("SEND DATABASE")
	shoset.SendCreateDatabase(utils.Shoset, tenant)
	time.Sleep(time.Second * time.Duration(5))

	createDatabase, ok := utils.Shoset.Context["databaseCreation"].(*models.CreateDatabase)
	if ok {
		return createDatabase, nil
	}
	return nil, fmt.Errorf("Creation database nil")
}
