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

// ConnectorController :
type ConnectorController struct {
	mapDatabase  map[string]*gorm.DB
	databasePath string
}

// NewConnectorController :
func NewConnectorController(mapDatabase map[string]*gorm.DB, databasePath string) (connectorController *ConnectorController) {
	connectorController = new(ConnectorController)
	connectorController.mapDatabase = mapDatabase
	connectorController.databasePath = databasePath

	return
}

// List :
func (cc ConnectorController) List(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tenant := vars["tenant"]
	database := utils.GetDatabase(cc.mapDatabase, cc.databasePath, tenant)

	connectors, err := dao.ListConnector(database)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, connectors)
}

// Create :
func (cc ConnectorController) Create(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tenant := vars["tenant"]
	database := utils.GetDatabase(cc.mapDatabase, cc.databasePath, tenant)

	var connector models.Connector
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&connector); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := dao.CreateConnector(database, connector); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, connector)
}

// Read :
func (cc ConnectorController) Read(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tenant := vars["tenant"]
	database := utils.GetDatabase(cc.mapDatabase, cc.databasePath, tenant)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var connector models.Connector
	if connector, err = dao.ReadConnector(database, id); err != nil {
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

// Update :
func (cc ConnectorController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tenant := vars["tenant"]
	database := utils.GetDatabase(cc.mapDatabase, cc.databasePath, tenant)

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

	if err := dao.UpdateConnector(database, connector); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, connector)
}

// Delete :
func (cc ConnectorController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tenant := vars["tenant"]
	database := utils.GetDatabase(cc.mapDatabase, cc.databasePath, tenant)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
		return
	}

	if err := dao.DeleteAggregator(database, id); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
