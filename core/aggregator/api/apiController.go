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
	ResourceController          *controllers.ResourceController
	DomainController            *controllers.DomainController
	EventTypeToPollController   *controllers.EventTypeToPollController
	ResourceTypeController      *controllers.ResourceTypeController
	EventTypeController         *controllers.EventTypeController
	ApplicationController       *controllers.ApplicationController
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
	aggregatorControllers.ResourceController = controllers.NewResourceController(databaseConnection)
	aggregatorControllers.DomainController = controllers.NewDomainController(databaseConnection)
	aggregatorControllers.EventTypeToPollController = controllers.NewEventTypeToPollController(databaseConnection)
	aggregatorControllers.ResourceTypeController = controllers.NewResourceTypeController(databaseConnection)
	aggregatorControllers.EventTypeController = controllers.NewEventTypeController(databaseConnection)
	aggregatorControllers.ApplicationController = controllers.NewApplicationController(databaseConnection)

	return aggregatorControllers
}
