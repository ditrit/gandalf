package controllers

import (
	"net/http"

	"github.com/ditrit/gandalf/core/cluster/database"

	"github.com/ditrit/gandalf/core/cluster/api/dao"
	"github.com/ditrit/gandalf/core/cluster/api/utils"
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
	database := sac.databaseConnection.GetGandalfDatabaseClient()
	if database != nil {
		roles, err := dao.ListSecretAssignement(database)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, roles)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}

// Create :
func (sac SecretAssignementController) Create(w http.ResponseWriter, r *http.Request) {
	database := sac.databaseConnection.GetGandalfDatabaseClient()
	if database != nil {
		var secretAssignement models.SecretAssignement
		secretAssignement.Secret = utils.GenerateHash()

		if err := dao.CreateSecretAssignement(database, secretAssignement); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.RespondWithJSON(w, http.StatusCreated, secretAssignement)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, "tenant not found")
		return
	}
}
