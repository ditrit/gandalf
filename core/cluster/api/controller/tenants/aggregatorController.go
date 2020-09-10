package controller

import (
	"database/sql"
	"encoding/json"
	"gandalf/core/api/dao"
	"gandalf/core/api/utils"
	"net/http"
	"strconv"

	"github.com/ditrit/gandalf/core/models"
	"github.com/jinzhu/gorm"

	"github.com/gorilla/mux"
)

type AggregatorController struct {
	aggregatorDAO *dao.AggregatorDAO
}

func NewAggregatorController(gandalfDatabase *gorm.DB) (aggregatorController *AggregatorController) {
	aggregatorController = new(AggregatorController)
	aggregatorController.aggregatorDAO = dao.NewAggregatorDAO(gandalfDatabase)

	return
}

func (ac AggregatorController) List(w http.ResponseWriter, r *http.Request) {
	aggregators, err := ac.aggregatorDAO.List()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, aggregators)
}

func (ac AggregatorController) Create(w http.ResponseWriter, r *http.Request) {
	var aggregator models.Aggregator
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&aggregator); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := ac.aggregatorDAO.Create(aggregator); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, aggregator)
}

func (ac AggregatorController) Read(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var aggregator models.Aggregator
	if aggregator, err = ac.aggregatorDAO.Read(id); err != nil {
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

func (ac AggregatorController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
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

	if err := ac.aggregatorDAO.Update(aggregator); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, aggregator)
}

func (ac AggregatorController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
		return
	}

	if err = ac.aggregatorDAO.Delete(id); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
