package gandalf

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ditrit/gandalf/core/cluster/database"

	"github.com/ditrit/gandalf/core/cluster/api/dao"
	"github.com/ditrit/gandalf/core/cluster/api/utils"
	"github.com/ditrit/gandalf/core/models"

	"github.com/gorilla/mux"
)

// TenantController :
type TenantController struct {
	databaseConnection *database.DatabaseConnection
}

// NewTenantController :
func NewTenantController(databaseConnection *database.DatabaseConnection) (tenantController *TenantController) {
	tenantController = new(TenantController)
	tenantController.databaseConnection = databaseConnection

	return
}

// List :
func (tc TenantController) List(w http.ResponseWriter, r *http.Request) {

	tenants, err := dao.ListTenant(tc.databaseConnection.GetGandalfDatabaseClient())
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
	var login, password string

	var tenant models.Tenant
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&tenant); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	err := dao.CreateTenant(tc.databaseConnection.GetGandalfDatabaseClient(), tenant)
	if err == nil {

		err = tc.databaseConnection.NewDatabase(tenant.Name)
		if err == nil {

			//var tenantDatabaseClient *gorm.DB
			tenantDatabaseClient := tc.databaseConnection.GetDatabaseClientByTenant(tenant.Name)
			//tc.mapTenantDatabase[tenant.Name] = tenantDatabaseClient

			if tenantDatabaseClient != nil {

				login, password, err = tc.databaseConnection.InitTenantDatabase(tenantDatabaseClient)

				if err == nil {
					result["login"] = login
					result["password"] = password
					result["tenant"] = tenant

				} else {
					dao.DeleteTenant(tc.databaseConnection.GetGandalfDatabaseClient(), int(tenant.ID))
					utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
					return
				}
			} else {
				dao.DeleteTenant(tc.databaseConnection.GetGandalfDatabaseClient(), int(tenant.ID))
				utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
				return
			}
		} else {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, result)
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
	if tenant, err = dao.ReadTenant(tc.databaseConnection.GetGandalfDatabaseClient(), id); err != nil {
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

	if err := dao.UpdateTenant(tc.databaseConnection.GetGandalfDatabaseClient(), tenant); err != nil {
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

	if err := dao.DeleteTenant(tc.databaseConnection.GetGandalfDatabaseClient(), id); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
