package api

import (
	"github.com/ditrit/gandalf/core/aggregator/database"

	"github.com/ditrit/gandalf/core/aggregator/api/controllers"
)

// Controllers :
type Controllers struct {
	tenantsAuthenticationController          *controllers.AuthenticationController
	tenantsAggregatorController              *controllers.AggregatorController
	tenantsConnectorController               *controllers.ConnectorController
	tenantsRoleController                    *controllers.RoleController
	tenantsUserController                    *controllers.UserController
	tenantsConfigurationAggregatorController *controllers.ConfigurationAggregatorController
	tenantsConfigurationConnectorController  *controllers.ConfigurationConnectorController
}

// ReturnControllers :
func ReturnControllers(databaseConnection *database.DatabaseConnection) *Controllers {

	controllers := new(Controllers)

	controllers.tenantsAuthenticationController = controllers.NewAuthenticationController(databaseConnection)
	controllers.tenantsConnectorController = controllers.NewConnectorController(databaseConnection)
	controllers.tenantsAggregatorController = controllers.NewAggregatorController(databaseConnection)
	controllers.tenantsUserController = controllers.NewUserController(databaseConnection)
	controllers.tenantsRoleController = controllers.NewRoleController(databaseConnection)
	controllers.tenantsConfigurationAggregatorController = controllers.NewConfigurationAggregatorController(databaseConnection)
	controllers.tenantsConfigurationConnectorController = controllers.NewConfigurationConnectorController(databaseConnection)

	return controllers
}
