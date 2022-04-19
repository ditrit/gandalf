package api

import (
	"database/sql"
	"encoding/json"
	// "fmt"
	"net/http"

	"github.com/ditrit/gandalf/core/aggregator/api/dao"
	"github.com/ditrit/gandalf/core/models"
	"github.com/google/uuid"

	"github.com/ditrit/gandalf/core/aggregator/api/utils"
	"github.com/gorilla/mux"
)

func GetRequirementGroups(w http.ResponseWriter, r *http.Request) {

	database := utils.DatabaseConnection.GetTenantDatabaseClient()
	if database != nil {
		resources, err := dao.ListRequirementGroups(database)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, resources)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

func GetRequirementGroupById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	database := utils.DatabaseConnection.GetTenantDatabaseClient()
	if database != nil {
		id, err := uuid.Parse(vars["requirementGroupId"])
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid ID supplied")
			return
		}

		var requirementGroup models.RequirementGroup
		if requirementGroup, err = dao.ReadRequirementGroup(database, id); err != nil {
			switch err {
			case sql.ErrNoRows:
				utils.RespondWithError(w, http.StatusNotFound, "User not found")
			default:
				utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			}
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, requirementGroup)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

func CreateRequirementGroup(w http.ResponseWriter, r *http.Request) {

	database := utils.DatabaseConnection.GetTenantDatabaseClient()
	if database != nil {
		var requirementgroup *models.RequirementGroup
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&requirementgroup); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		defer r.Body.Close()

		if err := dao.CreateRequirementGroup(database, requirementgroup); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusCreated, requirementgroup)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

func UpdateRequirementGroup(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	database := utils.DatabaseConnection.GetTenantDatabaseClient()
	if database != nil {
		id, err := uuid.Parse(vars["requirementGroupId"])
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid ID supplied")
			return
		}

		var requirementGroup models.RequirementGroup
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&requirementGroup); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
			return
		}
		defer r.Body.Close()
		requirementGroup.ID = id

		if err := dao.UpdateRequirementGroup(database, requirementGroup); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, requirementGroup)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

func DeleteRequirementGroup(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	database := utils.DatabaseConnection.GetTenantDatabaseClient()
	if database != nil {
		id, err := uuid.Parse(vars["requirementGroupId"])
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
			return
		}

		if err := dao.DeleteRequirementGroup(database, id); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}