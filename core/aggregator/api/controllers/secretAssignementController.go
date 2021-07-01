package controllers

import (
	"fmt"
	"net/http"

	"github.com/ditrit/gandalf/core/aggregator/database"
	"github.com/google/uuid"

	"github.com/ditrit/gandalf/core/aggregator/api/dao"
	"github.com/ditrit/gandalf/core/aggregator/api/utils"
	"github.com/ditrit/gandalf/core/models"
)

// RoleController :
type SecretAssignementController struct {
	databaseConnection *database.DatabaseConnection
}

// NewRoleController :
func NewSecretAssignementController(databaseConnection *database.DatabaseConnection) (secretAssignementController *SecretAssignementController) {
	secretAssignementController = new(SecretAssignementController)
	secretAssignementController.databaseConnection = databaseConnection

	return
}

// List :
func (sac SecretAssignementController) List(w http.ResponseWriter, r *http.Request) {
	database := sac.databaseConnection.GetTenantDatabaseClient()
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

// Create :
func (sac SecretAssignementController) Create(w http.ResponseWriter, r *http.Request) {
	database := sac.databaseConnection.GetTenantDatabaseClient()
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
