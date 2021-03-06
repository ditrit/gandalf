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

// AggregatorController :
type AggregatorController struct {
	databaseConnection *database.DatabaseConnection
}

// NewAggregatorController :
func NewAggregatorController(databaseConnection *database.DatabaseConnection) (aggregatorController *AggregatorController) {
	aggregatorController = new(AggregatorController)
	aggregatorController.databaseConnection = databaseConnection

	return
}

// List :
func (ac AggregatorController) List(w http.ResponseWriter, r *http.Request) {
	database := ac.databaseConnection.GetTenantDatabaseClient()
	if database != nil {
		aggregators, err := dao.ListAggregator(database)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		utils.RespondWithJSON(w, http.StatusOK, aggregators)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}

}

// Create :
func (ac AggregatorController) Create(w http.ResponseWriter, r *http.Request) {
	database := ac.databaseConnection.GetTenantDatabaseClient()
	if database != nil {

		var aggregator models.Aggregator
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&aggregator); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		defer r.Body.Close()

		if err := dao.CreateAggregator(database, aggregator); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusCreated, aggregator)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

// DeclareMember :
func (ac AggregatorController) DeclareMember(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	database := ac.databaseConnection.GetTenantDatabaseClient()
	if database != nil {
		aggregator, err := dao.ReadAggregatorByName(database, name)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		var newAggregator models.Aggregator
		newAggregator.LogicalName = aggregator.LogicalName
		newAggregator.Secret = utils.GenerateHash()

		if err := dao.CreateAggregator(database, newAggregator); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusCreated, newAggregator)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

// Read :
func (ac AggregatorController) Read(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	database := ac.databaseConnection.GetTenantDatabaseClient()
	if database != nil {

		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
			return
		}

		var aggregator models.Aggregator
		if aggregator, err = dao.ReadAggregator(database, id); err != nil {
			switch err {
			case sql.ErrNoRows:
				utils.RespondWithError(w, http.StatusNotFound, "Product not found")
			default:
				utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			}
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, aggregator)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

// Update :
func (ac AggregatorController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	database := ac.databaseConnection.GetTenantDatabaseClient()
	if database != nil {

		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
			return
		}

		var aggregator models.Aggregator
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&aggregator); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
			return
		}
		defer r.Body.Close()
		aggregator.ID = uint(id)

		if err := dao.UpdateAggregator(database, aggregator); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, aggregator)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

// Delete :
func (ac AggregatorController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	database := ac.databaseConnection.GetTenantDatabaseClient()
	if database != nil {
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
			return
		}

		if err = dao.DeleteAggregator(database, id); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}
