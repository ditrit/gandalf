package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/ditrit/gandalf/core/aggregator/shoset"

	"github.com/ditrit/gandalf/core/aggregator/database"

	"github.com/ditrit/gandalf/core/aggregator/api/dao"
	"github.com/ditrit/gandalf/core/aggregator/api/utils"
	"github.com/ditrit/gandalf/core/models"
	net "github.com/ditrit/shoset"

	"github.com/gorilla/mux"
)

// TenantController :
type TenantController struct {
	databaseConnection *database.DatabaseConnection
	shoset             *net.Shoset
}

// NewTenantController :
func NewTenantController(databaseConnection *database.DatabaseConnection, shoset *net.Shoset) (tenantController *TenantController) {
	tenantController = new(TenantController)
	tenantController.databaseConnection = databaseConnection
	tenantController.shoset = shoset

	return
}

// List :
func (tc TenantController) List(w http.ResponseWriter, r *http.Request) {

	tenants, err := dao.ListTenant(tc.databaseConnection.GetTenantDatabaseClient())
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, tenants)
}

// Create :
func (tc TenantController) Create(w http.ResponseWriter, r *http.Request) {
	var result map[string]interface{}
	result = make(map[string]interface{})

	var tenant models.Tenant
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&tenant); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	tenant.Password = utils.GenerateHash()

	err := dao.CreateTenant(tc.databaseConnection.GetTenantDatabaseClient(), tenant)
	fmt.Println("tenant.Password")
	fmt.Println(tenant.Password)
	fmt.Println(err)
	var createDatabase *models.CreateDatabase
	createDatabase, err = tc.CreateDatabase(tenant.Name)
	fmt.Println("createDatabase0")
	fmt.Println(createDatabase)
	fmt.Println(err)
	if err == nil {
		fmt.Println("CREATE 0")

		tenantDatabaseClient := tc.databaseConnection.GetDatabaseClientByTenant(tenant.Name)
		//tc.mapTenantDatabase[tenant.Name] = tenantDatabaseClient
		fmt.Println("CREATE 1")
		if tenantDatabaseClient != nil {

			//utils.ChangeStateTenant(tenantDatabaseClient)

			if err == nil {

				//CREATE SECRET
				var secretAssignement models.SecretAssignement
				secretAssignement.Secret = utils.GenerateHash()
				err := dao.CreateSecretAssignement(tc.databaseConnection.GetTenantDatabaseClient(), secretAssignement)
				fmt.Println("CREATE 2")
				if err == nil {
					//GET PIVOT AGGREGATOR
					version := models.Version{Major: tc.databaseConnection.GetPivot().Major, Minor: tc.databaseConnection.GetPivot().Minor}
					fmt.Println("test")
					fmt.Println(tc.databaseConnection.GetLogicalComponent())
					fmt.Println(tc.databaseConnection.GetLogicalComponent().GetKeyValueByKey("repository_url").Value)
					fmt.Println(version)
					pivot, _ := utils.GetPivot(tenantDatabaseClient, tc.databaseConnection.GetLogicalComponent().GetKeyValueByKey("repository_url").Value, "aggregator", version)
					fmt.Println("CREATE 3")
					//CREATE AGGREGATOR LOGICAL COMPONENT
					logicalComponent, _ := utils.SaveLogicalComponent(tenantDatabaseClient, tenant.Name, tc.databaseConnection.GetLogicalComponent().GetKeyValueByKey("repository_url").Value, pivot)
					fmt.Println(logicalComponent)
					fmt.Println("CREATE 4")
				} else {
					dao.DeleteTenant(tc.databaseConnection.GetTenantDatabaseClient(), int(tenant.ID))
					utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
					return
				}
				fmt.Println("CREATE 5")

				result["login"] = createDatabase.Login
				result["password"] = createDatabase.Password
				result["tenant"] = tenant

			} else {
				dao.DeleteTenant(tc.databaseConnection.GetTenantDatabaseClient(), int(tenant.ID))
				utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
				return
			}

		} else {

		}
	} else {

	}

	utils.RespondWithJSON(w, http.StatusCreated, result)
}

func (tc TenantController) CreateDatabase(tenant string) (*models.CreateDatabase, error) {
	fmt.Println("SEND DATABASE")
	shoset.SendCreateDatabase(tc.shoset, tenant)
	time.Sleep(time.Second * time.Duration(5))

	createDatabase, ok := tc.shoset.Context["databaseCreation"].(*models.CreateDatabase)
	if ok {
		return createDatabase, nil
	}
	return nil, fmt.Errorf("Creation database nil")
}

// Read :
func (tc TenantController) Read(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var tenant models.Tenant
	if tenant, err = dao.ReadTenant(tc.databaseConnection.GetTenantDatabaseClient(), id); err != nil {
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

// Update :
func (tc TenantController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
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

	if err := dao.UpdateTenant(tc.databaseConnection.GetTenantDatabaseClient(), tenant); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, tenant)
}

// Delete :
func (tc TenantController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
		return
	}

	if err := dao.DeleteTenant(tc.databaseConnection.GetTenantDatabaseClient(), id); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
