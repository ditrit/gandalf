package api

import (
	"github.com/ditrit/gandalf/core/aggregator/database"
	net "github.com/ditrit/shoset"

	"github.com/ditrit/gandalf/core/aggregator/api/controllers"
)

// Controllers :
type Controllers struct {
	AuthenticationController    *controllers.AuthenticationController
	CliController               *controllers.CliController
	RoleController              *controllers.RoleController
	UserController              *controllers.UserController
	TenantController            *controllers.TenantController
	SecretAssignementController *controllers.SecretAssignementController
	LogicalComponentController  *controllers.LogicalComponentController
}

// ReturnControllers :
func ReturnControllers(databaseConnection *database.DatabaseConnection, shoset *net.Shoset) *Controllers {

	aggregatorControllers := new(Controllers)

	aggregatorControllers.AuthenticationController = controllers.NewAuthenticationController(databaseConnection)
	aggregatorControllers.CliController = controllers.NewCliController()
	aggregatorControllers.LogicalComponentController = controllers.NewLogicalComponentController(databaseConnection)

	aggregatorControllers.UserController = controllers.NewUserController(databaseConnection)
	aggregatorControllers.RoleController = controllers.NewRoleController(databaseConnection)
	aggregatorControllers.TenantController = controllers.NewTenantController(databaseConnection, shoset)
	aggregatorControllers.SecretAssignementController = controllers.NewSecretAssignementController(databaseConnection)

	return aggregatorControllers
}
