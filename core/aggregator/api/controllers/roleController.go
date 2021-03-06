package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ditrit/gandalf/core/aggregator/database"

	"github.com/ditrit/gandalf/core/aggregator/api/dao"
	"github.com/ditrit/gandalf/core/aggregator/api/utils"
	"github.com/ditrit/gandalf/core/models"

	"github.com/gorilla/mux"
)

// RoleController :
type RoleController struct {
	databaseConnection *database.DatabaseConnection
}

// NewRoleController :
func NewRoleController(databaseConnection *database.DatabaseConnection) (roleController *RoleController) {
	roleController = new(RoleController)
	roleController.databaseConnection = databaseConnection

	return
}

// List :
func (rc RoleController) List(w http.ResponseWriter, r *http.Request) {
	database := rc.databaseConnection.GetTenantDatabaseClient()
	if database != nil {
		roles, err := dao.ListRole(database)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, roles)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

// Create :
func (rc RoleController) Create(w http.ResponseWriter, r *http.Request) {
	database := rc.databaseConnection.GetTenantDatabaseClient()
	if database != nil {
		var role models.Role
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&role); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		defer r.Body.Close()

		if err := dao.CreateRole(database, role); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusCreated, role)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

// Read :
func (rc RoleController) Read(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	database := rc.databaseConnection.GetTenantDatabaseClient()
	if database != nil {
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
			return
		}

		var role models.Role
		if role, err = dao.ReadRole(database, id); err != nil {
			switch err {
			case sql.ErrNoRows:
				utils.RespondWithError(w, http.StatusNotFound, "Product not found")
			default:
				utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			}
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, role)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

// Update :
func (rc RoleController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	database := rc.databaseConnection.GetTenantDatabaseClient()
	if database != nil {
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
			return
		}

		var role models.Role
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&role); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
			return
		}
		defer r.Body.Close()
		role.ID = uint(id)

		if err := dao.UpdateRole(database, role); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, role)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

// Delete :
func (rc RoleController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	database := rc.databaseConnection.GetTenantDatabaseClient()
	if database != nil {
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
			return
		}

		if err := dao.DeleteRole(database, id); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}
