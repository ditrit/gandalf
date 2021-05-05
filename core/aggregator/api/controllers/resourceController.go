package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ditrit/gandalf/core/aggregator/database"

	"github.com/ditrit/gandalf/core/aggregator/api/dao"
	"github.com/ditrit/gandalf/core/aggregator/api/utils"
	"github.com/ditrit/gandalf/core/models"

	"github.com/gorilla/mux"
)

// ResourceController :
type ResourceController struct {
	databaseConnection *database.DatabaseConnection
}

// NewResourceController :
func NewResourceController(databaseConnection *database.DatabaseConnection) (resourceController *ResourceController) {
	resourceController = new(ResourceController)
	resourceController.databaseConnection = databaseConnection

	return
}

// List :
func (rc ResourceController) List(w http.ResponseWriter, r *http.Request) {
	database := rc.databaseConnection.GetTenantDatabaseClient()
	if database != nil {
		resources, err := dao.ListResource(database)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, resources)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

// Create :
func (rc ResourceController) Create(w http.ResponseWriter, r *http.Request) {
	database := rc.databaseConnection.GetTenantDatabaseClient()
	if database != nil {
		var resource models.Resource
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&resource); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		defer r.Body.Close()

		if err := dao.CreateResource(database, resource); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusCreated, resource)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

// Read :
func (rc ResourceController) Read(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	database := rc.databaseConnection.GetTenantDatabaseClient()
	if database != nil {
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
			return
		}

		var resource models.Resource
		if resource, err = dao.ReadResource(database, id); err != nil {
			switch err {
			case sql.ErrNoRows:
				utils.RespondWithError(w, http.StatusNotFound, "Product not found")
			default:
				utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			}
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, resource)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

// ReadByName :
func (rc ResourceController) ReadByName(w http.ResponseWriter, r *http.Request) {
	fmt.Println("READ_BY_NAME")
	vars := mux.Vars(r)
	name := vars["name"]

	var resource models.Resource
	var err error
	if resource, err = dao.ReadResourceByName(rc.databaseConnection.GetTenantDatabaseClient(), name); err != nil {
		switch err {
		case sql.ErrNoRows:
			utils.RespondWithError(w, http.StatusNotFound, "Product not found")
		default:
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, resource)
}

// Update :
func (rc ResourceController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	database := rc.databaseConnection.GetTenantDatabaseClient()
	if database != nil {
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
			return
		}

		var resource models.Resource
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&resource); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
			return
		}
		defer r.Body.Close()
		resource.ID = uint(id)

		if err := dao.UpdateResource(database, resource); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, resource)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

// Delete :
func (rc ResourceController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	database := rc.databaseConnection.GetTenantDatabaseClient()
	if database != nil {
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
			return
		}

		if err := dao.DeleteResource(database, id); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}
