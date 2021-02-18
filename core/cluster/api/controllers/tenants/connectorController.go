package tenants

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ditrit/gandalf/core/cluster/database"

	"github.com/ditrit/gandalf/core/cluster/api/dao"
	"github.com/ditrit/gandalf/core/cluster/api/utils"
	"github.com/ditrit/gandalf/core/models"

	"github.com/gorilla/mux"
)

// ConnectorController :
type ConnectorController struct {
	databaseConnection *database.DatabaseConnection
}

// NewConnectorController :
func NewConnectorController(databaseConnection *database.DatabaseConnection) (connectorController *ConnectorController) {
	connectorController = new(ConnectorController)
	connectorController.databaseConnection = databaseConnection

	return
}

// List :
func (cc ConnectorController) List(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tenant := vars["tenant"]
	database := cc.databaseConnection.GetDatabaseClientByTenant(tenant)
	if database != nil {
		connectors, err := dao.ListConnector(database)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, connectors)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

// Create :
func (cc ConnectorController) Create(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tenant := vars["tenant"]
	database := cc.databaseConnection.GetDatabaseClientByTenant(tenant)
	if database != nil {
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
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

// DeclareMember :
func (cc ConnectorController) DeclareMember(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tenant := vars["tenant"]
	name := vars["name"]
	database := cc.databaseConnection.GetDatabaseClientByTenant(tenant)
	if database != nil {
		connector, err := dao.ReadConnectorByName(database, name)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		//var newConnector models.Connector
		//newConnector.LogicalName = connector.LogicalName
		connector.Secret = utils.GenerateHash()

		if err := dao.UpdateConnector(database, connector); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusCreated, connector)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

// Read :
func (cc ConnectorController) Read(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tenant := vars["tenant"]
	database := cc.databaseConnection.GetDatabaseClientByTenant(tenant)
	if database != nil {
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
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

// Update :
func (cc ConnectorController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tenant := vars["tenant"]
	database := cc.databaseConnection.GetDatabaseClientByTenant(tenant)
	if database != nil {
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
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

// Delete :
func (cc ConnectorController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tenant := vars["tenant"]
	database := cc.databaseConnection.GetDatabaseClientByTenant(tenant)
	if database != nil {
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
			return
		}

		if err := dao.DeleteConnector(database, id); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}
