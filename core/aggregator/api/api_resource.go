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
	"fmt"
	"net/http"
	"strconv"

	"github.com/ditrit/gandalf/core/aggregator/api/dao"
	apimodels "github.com/ditrit/gandalf/core/aggregator/api/models"
	"github.com/ditrit/gandalf/core/models"
	"github.com/golang-jwt/jwt/v4"

	"github.com/ditrit/gandalf/core/aggregator/api/utils"
	"github.com/gorilla/mux"
)

func CreateResource(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("gandalf")
	if err != nil {
		if err == http.ErrNoCookie {
			utils.RespondWithError(w, http.StatusUnauthorized, err.Error())
			return
		}
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	tokenStr := cookie.Value

	claims := &apimodels.Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return utils.GetJwtKey(), nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			utils.RespondWithError(w, http.StatusUnauthorized, err.Error())
			return
		}
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if !tkn.Valid {
		utils.RespondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}

	database := utils.DatabaseConnection.GetTenantDatabaseClient()
	if database != nil {
		var resource models.Resource
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&resource); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		defer r.Body.Close()

		if err := dao.CreateResource(database, resource); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusCreated, resource)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

func DeleteResource(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("gandalf")
	if err != nil {
		if err == http.ErrNoCookie {
			utils.RespondWithError(w, http.StatusUnauthorized, err.Error())
			return
		}
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	tokenStr := cookie.Value

	claims := &apimodels.Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return utils.GetJwtKey(), nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			utils.RespondWithError(w, http.StatusUnauthorized, err.Error())
			return
		}
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if !tkn.Valid {
		utils.RespondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}

	vars := mux.Vars(r)
	database := utils.DatabaseConnection.GetTenantDatabaseClient()
	if database != nil {
		id, err := strconv.Atoi(vars["resourceId"])
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid Product ID")
			return
		}

		if err := dao.DeleteResource(database, id); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

func GetResourceById(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("gandalf")
	if err != nil {
		if err == http.ErrNoCookie {
			utils.RespondWithError(w, http.StatusUnauthorized, err.Error())
			return
		}
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	tokenStr := cookie.Value

	claims := &apimodels.Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return utils.GetJwtKey(), nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			utils.RespondWithError(w, http.StatusUnauthorized, err.Error())
			return
		}
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if !tkn.Valid {
		utils.RespondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}

	vars := mux.Vars(r)
	database := utils.DatabaseConnection.GetTenantDatabaseClient()
	if database != nil {
		id, err := strconv.Atoi(vars["resourceId"])
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
			return
		}

		var resource models.Resource
		if resource, err = dao.ReadResource(database, id); err != nil {
			switch err {
			case sql.ErrNoRows:
				utils.RespondWithError(w, http.StatusNotFound, "Product not found")
			default:
				utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			}
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, resource)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

func GetResourceByName(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("gandalf")
	if err != nil {
		if err == http.ErrNoCookie {
			utils.RespondWithError(w, http.StatusUnauthorized, err.Error())
			return
		}
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	tokenStr := cookie.Value

	claims := &apimodels.Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return utils.GetJwtKey(), nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			utils.RespondWithError(w, http.StatusUnauthorized, err.Error())
			return
		}
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if !tkn.Valid {
		utils.RespondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}

	fmt.Println("READ_BY_NAME")
	vars := mux.Vars(r)
	name := vars["resourceName"]

	var resource models.Resource
	if resource, err = dao.ReadResourceByName(utils.DatabaseConnection.GetTenantDatabaseClient(), name); err != nil {
		switch err {
		case sql.ErrNoRows:
			utils.RespondWithError(w, http.StatusNotFound, "Product not found")
		default:
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, resource)
}

func ListResource(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("gandalf")
	if err != nil {
		if err == http.ErrNoCookie {
			utils.RespondWithError(w, http.StatusUnauthorized, err.Error())
			return
		}
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	tokenStr := cookie.Value

	claims := &apimodels.Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return utils.GetJwtKey(), nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			utils.RespondWithError(w, http.StatusUnauthorized, err.Error())
			return
		}
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if !tkn.Valid {
		utils.RespondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}

	database := utils.DatabaseConnection.GetTenantDatabaseClient()
	if database != nil {
		resources, err := dao.ListResource(database)
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

func UpdateResource(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("gandalf")
	if err != nil {
		if err == http.ErrNoCookie {
			utils.RespondWithError(w, http.StatusUnauthorized, err.Error())
			return
		}
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	tokenStr := cookie.Value

	claims := &apimodels.Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return utils.GetJwtKey(), nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			utils.RespondWithError(w, http.StatusUnauthorized, err.Error())
			return
		}
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if !tkn.Valid {
		utils.RespondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}

	vars := mux.Vars(r)
	database := utils.DatabaseConnection.GetTenantDatabaseClient()
	if database != nil {
		id, err := strconv.Atoi(vars["resourceId"])
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
			return
		}

		var resource models.Resource
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&resource); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
			return
		}
		defer r.Body.Close()
		resource.ID = uint(id)

		if err := dao.UpdateResource(database, resource); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, resource)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}
