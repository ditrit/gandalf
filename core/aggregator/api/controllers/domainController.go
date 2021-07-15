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
func (dc DomainController) List(w http.ResponseWriter, r *http.Request) {
	database := dc.databaseConnection.GetTenantDatabaseClient()
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
func (dc DomainController) Create(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	database := dc.databaseConnection.GetTenantDatabaseClient()
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
func (dc DomainController) Read(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	database := dc.databaseConnection.GetTenantDatabaseClient()
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

// ReadByName :
func (dc DomainController) ReadByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	var domain models.Domain
	var err error
	if domain, err = dao.ReadDomainByName(dc.databaseConnection.GetTenantDatabaseClient(), name); err != nil {
		switch err {
		case sql.ErrNoRows:
			utils.RespondWithError(w, http.StatusNotFound, "Product not found")
		default:
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, domain)
}

// Update :
func (dc DomainController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	database := dc.databaseConnection.GetTenantDatabaseClient()
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
func (dc DomainController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	database := dc.databaseConnection.GetTenantDatabaseClient()
	if database != nil {
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
			return
		}

		//TEST
		var domains []models.Domain
		database.Unscoped().Find(&domains)
		fmt.Println("domains")
		for _, domain := range domains {
			fmt.Println(domain)

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
