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
	"fmt"
	"net/http"

	"github.com/ditrit/gandalf/core/aggregator/api/dao"
	apimodels "github.com/ditrit/gandalf/core/aggregator/api/models"
	"github.com/ditrit/gandalf/core/models"
	"github.com/golang-jwt/jwt/v4"

	"github.com/google/uuid"

	"github.com/ditrit/gandalf/core/aggregator/api/utils"
)

func CreateSecretAssignement(w http.ResponseWriter, r *http.Request) {
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
		var secretAssignement models.SecretAssignement
		secretAssignement.Secret = uuid.NewString()
		fmt.Println("SECRET")
		fmt.Println(secretAssignement)
		if err := dao.CreateSecretAssignement(database, secretAssignement); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		fmt.Println("SECRET2")
		fmt.Println(secretAssignement)
		utils.RespondWithJSON(w, http.StatusCreated, secretAssignement)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "database not found")
		return
	}
}

func ListSecretAssignement(w http.ResponseWriter, r *http.Request) {
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
		secrets, err := dao.ListSecretAssignement(database)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, secrets)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "database not found")
		return
	}
}
