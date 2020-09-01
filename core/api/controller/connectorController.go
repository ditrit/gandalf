package controller

import (
	"database/sql"
	"encoding/json"
	"gandalf/core/api/utils"
	"gandalf/core/dao"
	"net/http"
	"strconv"

	"github.com/ditrit/gandalf/core/models"
	"github.com/gorilla/mux"
)

type ConnectorController struct {
	connectorDAO dao.ConnectorDAO
}

func (cc ConnectorController) List(w http.ResponseWriter, r *http.Request) {

	connectors, err := cc.connectorDAO.List()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, connectors)
}

func (cc ConnectorController) Create(w http.ResponseWriter, r *http.Request) {
	var connector models.Connector
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&connector); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := cc.connectorDAO.Create(connector); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, connector)
}

func (cc ConnectorController) Read(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var connector models.Connector
	if connector, err = cc.connectorDAO.Read(id); err != nil {
		switch err {
		case sql.ErrNoRows:
			utils.RespondWithError(w, http.StatusNotFound, "Product not found")
		default:
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, connector)
}

func (cc ConnectorController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var connector models.Connector
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&connector); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	connector.ID = uint(id)

	if err := cc.connectorDAO.Update(connector); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, connector)
}

func (cc ConnectorController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
		return
	}

	if err := cc.connectorDAO.Delete(id); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
