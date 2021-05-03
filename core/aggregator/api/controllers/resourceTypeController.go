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

// ResourceTypeController :
type ResourceTypeController struct {
	databaseConnection *database.DatabaseConnection
}

// NewResourceTypeController :
func NewResourceTypeController(databaseConnection *database.DatabaseConnection) (resourceTypeController *ResourceTypeController) {
	resourceTypeController = new(ResourceTypeController)
	resourceTypeController.databaseConnection = databaseConnection

	return
}

// List :
func (rc ResourceTypeController) List(w http.ResponseWriter, r *http.Request) {
	database := rc.databaseConnection.GetTenantDatabaseClient()
	if database != nil {
		resourceTypes, err := dao.ListResourceType(database)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, resourceTypes)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

// Create :
func (rc ResourceTypeController) Create(w http.ResponseWriter, r *http.Request) {
	database := rc.databaseConnection.GetTenantDatabaseClient()
	if database != nil {
		var resourceType models.ResourceType
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&resourceType); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		defer r.Body.Close()

		if err := dao.CreateResourceType(database, resourceType); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusCreated, resourceType)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

// Read :
func (rc ResourceTypeController) Read(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	database := rc.databaseConnection.GetTenantDatabaseClient()
	if database != nil {
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
			return
		}

		var resourceType models.ResourceType
		if resourceType, err = dao.ReadResourceType(database, id); err != nil {
			switch err {
			case sql.ErrNoRows:
				utils.RespondWithError(w, http.StatusNotFound, "Product not found")
			default:
				utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			}
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, resourceType)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

// ReadByName :
func (rc ResourceTypeController) ReadByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	var resourceType models.ResourceType
	var err error
	if resourceType, err = dao.ReadResourceTypeByName(rc.databaseConnection.GetTenantDatabaseClient(), name); err != nil {
		switch err {
		case sql.ErrNoRows:
			utils.RespondWithError(w, http.StatusNotFound, "Product not found")
		default:
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, resourceType)
}

// Update :
func (rc ResourceTypeController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	database := rc.databaseConnection.GetTenantDatabaseClient()
	if database != nil {
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
			return
		}

		var resourceType models.ResourceType
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&resourceType); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
			return
		}
		defer r.Body.Close()
		resourceType.ID = uint(id)

		if err := dao.UpdateResourceType(database, resourceType); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, resourceType)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

// Delete :
func (rc ResourceTypeController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	database := rc.databaseConnection.GetTenantDatabaseClient()
	if database != nil {
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
			return
		}

		if err := dao.DeleteResourceType(database, id); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}
