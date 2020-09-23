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
	"github.com/jinzhu/gorm"
)

type TenantController struct {
	gandalfDatabase *gorm.DB
	databasePath    string
}

func NewTenantController(gandalfDatabase *gorm.DB, databasePath string) (tenantController *TenantController) {
	tenantController = new(TenantController)
	tenantController.gandalfDatabase = gandalfDatabase
	tenantController.databasePath = databasePath

	return
}

func (tc TenantController) List(w http.ResponseWriter, r *http.Request) {

	tenants, err := dao.ListTenant(tc.gandalfDatabase)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, tenants)
}

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

	err := dao.CreateTenant(tc.gandalfDatabase, tenant)
	if err == nil {
		var tenantDatabaseClient *gorm.DB
		tenantDatabaseClient, err := database.NewTenantDatabaseClient(tenant.Name, tc.databasePath)
		if err == nil {

			login, password, err = database.InitTenantDatabase(tenantDatabaseClient)

			if err == nil {
				result["login"] = login
				result["password"] = password
				result["tenant"] = tenant

			} else {
				dao.DeleteTenant(tc.gandalfDatabase, int(tenant.ID))
				utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
				return
			}
		} else {
			dao.DeleteTenant(tc.gandalfDatabase, int(tenant.ID))
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, result)
}

func (tc TenantController) Read(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var tenant models.Tenant
	if tenant, err = dao.ReadTenant(tc.gandalfDatabase, id); err != nil {
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

	if err := dao.UpdateTenant(tc.gandalfDatabase, tenant); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, tenant)
}

func (tc TenantController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
		return
	}

	if err := dao.DeleteTenant(tc.gandalfDatabase, id); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
