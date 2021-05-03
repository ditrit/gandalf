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

// EventTypeController :
type EventTypeController struct {
	databaseConnection *database.DatabaseConnection
}

// NewEventTypeController :
func NewEventTypeController(databaseConnection *database.DatabaseConnection) (eventTypeController *EventTypeController) {
	eventTypeController = new(EventTypeController)
	eventTypeController.databaseConnection = databaseConnection

	return
}

// List :
func (ec EventTypeController) List(w http.ResponseWriter, r *http.Request) {
	database := ec.databaseConnection.GetTenantDatabaseClient()
	if database != nil {
		eventTypes, err := dao.ListEventType(database)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, eventTypes)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

// Create :
func (ec EventTypeController) Create(w http.ResponseWriter, r *http.Request) {
	database := ec.databaseConnection.GetTenantDatabaseClient()
	if database != nil {
		var eventType models.EventType
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&eventType); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		defer r.Body.Close()

		if err := dao.CreateEventType(database, eventType); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusCreated, eventType)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

// Read :
func (ec EventTypeController) Read(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	database := ec.databaseConnection.GetTenantDatabaseClient()
	if database != nil {
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
			return
		}

		var eventType models.EventType
		if eventType, err = dao.ReadEventType(database, id); err != nil {
			switch err {
			case sql.ErrNoRows:
				utils.RespondWithError(w, http.StatusNotFound, "Product not found")
			default:
				utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			}
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, eventType)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

// ReadByName :
func (ec EventTypeController) ReadByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	var eventType models.EventType
	var err error
	if eventType, err = dao.ReadEventTypeByName(ec.databaseConnection.GetTenantDatabaseClient(), name); err != nil {
		switch err {
		case sql.ErrNoRows:
			utils.RespondWithError(w, http.StatusNotFound, "Product not found")
		default:
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, eventType)
}

// Update :
func (ec EventTypeController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	database := ec.databaseConnection.GetTenantDatabaseClient()
	if database != nil {
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
			return
		}

		var eventType models.EventType
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&eventType); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
			return
		}
		defer r.Body.Close()
		eventType.ID = uint(id)

		if err := dao.UpdateEventType(database, eventType); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, eventType)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

// Delete :
func (ec EventTypeController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	database := ec.databaseConnection.GetTenantDatabaseClient()
	if database != nil {
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
			return
		}

		if err := dao.DeleteEventType(database, id); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}
