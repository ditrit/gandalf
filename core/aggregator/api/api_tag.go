/*
 * Swagger Gandalf
 *
 * This is a sample Petstore server.  You can find  out more about Swagger at  [http://swagger.io](http://swagger.io) or on  [irc.freenode.net, #swagger](http://swagger.io/irc/).
 *
 * API version: 1.0.0-oas3
 * Contact: romain.fairant@orness.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ditrit/gandalf/core/aggregator/api/dao"
	"github.com/ditrit/gandalf/core/models"

	"github.com/ditrit/gandalf/core/aggregator/api/utils"
	"github.com/gorilla/mux"
)

func CreateTag(w http.ResponseWriter, r *http.Request) {

	database := utils.DatabaseConnection.GetTenantDatabaseClient()
	if database != nil {
		var tag models.Tag
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&tag); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		defer r.Body.Close()

		if err := dao.CreateTag(database, tag); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusCreated, tag)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

func DeleteTag(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	database := utils.DatabaseConnection.GetTenantDatabaseClient()
	if database != nil {
		id, err := strconv.Atoi(vars["tagId"])
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
			return
		}

		if err := dao.DeleteTag(database, id); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

func GetTagById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	database := utils.DatabaseConnection.GetTenantDatabaseClient()
	if database != nil {
		id, err := strconv.Atoi(vars["tagId"])
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid ID supplied")
			return
		}

		var tag models.Tag
		if tag, err = dao.ReadTag(database, id); err != nil {
			switch err {
			case sql.ErrNoRows:
				utils.RespondWithError(w, http.StatusNotFound, "User not found")
			default:
				utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			}
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, tag)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

func GetTagByName(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	name := vars["tagName"]

	var tag models.Tag
	var err error
	if tag, err = dao.ReadTagByName(utils.DatabaseConnection.GetTenantDatabaseClient(), name); err != nil {
		switch err {
		case sql.ErrNoRows:
			utils.RespondWithError(w, http.StatusNotFound, "Product not found")
		default:
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, tag)
}

func ListTag(w http.ResponseWriter, r *http.Request) {

	database := utils.DatabaseConnection.GetTenantDatabaseClient()
	if database != nil {
		tags, err := dao.ListTag(database)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, tags)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

func UpdateTag(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	database := utils.DatabaseConnection.GetTenantDatabaseClient()
	if database != nil {
		id, err := strconv.Atoi(vars["tagId"])
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid ID supplied")
			return
		}

		var tag models.Tag
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&tag); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
			return
		}
		defer r.Body.Close()
		tag.ID = uint(id)

		if err := dao.UpdateTag(database, tag); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, tag)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}
