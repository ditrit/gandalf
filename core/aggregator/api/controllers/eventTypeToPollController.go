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

// EventTypeToPollController :
type EventTypeToPollController struct {
	databaseConnection *database.DatabaseConnection
}

// NewEventTypeToPollController :
func NewEventTypeToPollController(databaseConnection *database.DatabaseConnection) (eventTypeToPollController *EventTypeToPollController) {
	eventTypeToPollController = new(EventTypeToPollController)
	eventTypeToPollController.databaseConnection = databaseConnection

	return
}

// List :
func (rc EventTypeToPollController) List(w http.ResponseWriter, r *http.Request) {
	database := rc.databaseConnection.GetTenantDatabaseClient()
	if database != nil {
		eventTypeToPolls, err := dao.ListEventTypeToPoll(database)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, eventTypeToPolls)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

// Create :
func (rc EventTypeToPollController) Create(w http.ResponseWriter, r *http.Request) {
	database := rc.databaseConnection.GetTenantDatabaseClient()
	if database != nil {
		var eventTypeToPoll models.EventTypeToPoll
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&eventTypeToPoll); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		defer r.Body.Close()

		if err := dao.CreateEventTypeToPoll(database, eventTypeToPoll); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusCreated, eventTypeToPoll)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

// Read :
func (rc EventTypeToPollController) Read(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	database := rc.databaseConnection.GetTenantDatabaseClient()
	if database != nil {
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
			return
		}

		var eventTypeToPoll models.EventTypeToPoll
		if eventTypeToPoll, err = dao.ReadEventTypeToPoll(database, id); err != nil {
			switch err {
			case sql.ErrNoRows:
				utils.RespondWithError(w, http.StatusNotFound, "Product not found")
			default:
				utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			}
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, eventTypeToPoll)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

// Update :
func (rc EventTypeToPollController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	database := rc.databaseConnection.GetTenantDatabaseClient()
	if database != nil {
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
			return
		}

		var eventTypeToPoll models.EventTypeToPoll
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&eventTypeToPoll); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
			return
		}
		defer r.Body.Close()
		eventTypeToPoll.ID = uint(id)

		if err := dao.UpdateEventTypeToPoll(database, eventTypeToPoll); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, eventTypeToPoll)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

// Delete :
func (rc EventTypeToPollController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	database := rc.databaseConnection.GetTenantDatabaseClient()
	if database != nil {
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
			return
		}

		if err := dao.DeleteEventTypeToPoll(database, id); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}
