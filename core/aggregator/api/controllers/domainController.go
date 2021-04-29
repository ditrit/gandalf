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

// DomainController :
type DomainController struct {
	databaseConnection *database.DatabaseConnection
}

// NewDomainController :
func NewDomainController(databaseConnection *database.DatabaseConnection) (domainController *DomainController) {
	domainController = new(DomainController)
	domainController.databaseConnection = databaseConnection

	return
}

// List :
func (rc DomainController) List(w http.ResponseWriter, r *http.Request) {
	database := rc.databaseConnection.GetTenantDatabaseClient()
	if database != nil {
		domains, err := dao.ListDomain(database)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, domains)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

// Create :
func (rc DomainController) Create(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	database := rc.databaseConnection.GetTenantDatabaseClient()
	if database != nil {
		parentDomainName := vars["name"]
		var domain models.Domain
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&domain); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		defer r.Body.Close()

		if err := dao.CreateDomain(database, domain, parentDomainName); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusCreated, domain)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

// Read :
func (rc DomainController) Read(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	database := rc.databaseConnection.GetTenantDatabaseClient()
	if database != nil {
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
			return
		}

		var domain models.Domain
		if domain, err = dao.ReadDomain(database, id); err != nil {
			switch err {
			case sql.ErrNoRows:
				utils.RespondWithError(w, http.StatusNotFound, "Product not found")
			default:
				utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			}
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, domain)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

// Update :
func (rc DomainController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	database := rc.databaseConnection.GetTenantDatabaseClient()
	if database != nil {
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
			return
		}

		var domain models.Domain
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&domain); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
			return
		}
		defer r.Body.Close()
		domain.ID = uint(id)

		if err := dao.UpdateDomain(database, domain); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, domain)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

// Delete :
func (rc DomainController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	database := rc.databaseConnection.GetTenantDatabaseClient()
	if database != nil {
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
			return
		}

		if err := dao.DeleteDomain(database, id); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}
