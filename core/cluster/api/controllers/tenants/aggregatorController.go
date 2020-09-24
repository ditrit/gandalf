package tenants

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ditrit/gandalf/core/cluster/api/dao"
	"github.com/ditrit/gandalf/core/cluster/api/utils"
	"github.com/ditrit/gandalf/core/models"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// AggregatorController :
type AggregatorController struct {
	mapDatabase  map[string]*gorm.DB
	databasePath string
}

// NewAggregatorController :
func NewAggregatorController(mapDatabase map[string]*gorm.DB, databasePath string) (aggregatorController *AggregatorController) {
	aggregatorController = new(AggregatorController)
	aggregatorController.mapDatabase = mapDatabase
	aggregatorController.databasePath = databasePath

	return
}

// List :
func (ac AggregatorController) List(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tenant := vars["tenant"]
	database := utils.GetDatabase(ac.mapDatabase, ac.databasePath, tenant)

	aggregators, err := dao.ListAggregator(database)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, aggregators)
}

// Create :
func (ac AggregatorController) Create(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tenant := vars["tenant"]
	database := utils.GetDatabase(ac.mapDatabase, ac.databasePath, tenant)

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
}

// Read :
func (ac AggregatorController) Read(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tenant := vars["tenant"]
	database := utils.GetDatabase(ac.mapDatabase, ac.databasePath, tenant)

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
}

// Update :
func (ac AggregatorController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tenant := vars["tenant"]
	database := utils.GetDatabase(ac.mapDatabase, ac.databasePath, tenant)

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
}

// Delete :
func (ac AggregatorController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tenant := vars["tenant"]
	database := utils.GetDatabase(ac.mapDatabase, ac.databasePath, tenant)
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
}
