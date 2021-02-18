package api

import (
	"github.com/ditrit/gandalf/core/aggregator/database"

	"github.com/ditrit/gandalf/core/aggregator/api/controllers"
)

// Controllers :
type Controllers struct {
	AuthenticationController          *controllers.AuthenticationController
	AggregatorController              *controllers.AggregatorController
	ConnectorController               *controllers.ConnectorController
	RoleController                    *controllers.RoleController
	UserController                    *controllers.UserController
	ConfigurationAggregatorController *controllers.ConfigurationAggregatorController
	ConfigurationConnectorController  *controllers.ConfigurationConnectorController
}

// ReturnControllers :
func ReturnControllers(databaseConnection *database.DatabaseConnection) *Controllers {

	aggregatorControllers := new(Controllers)

	aggregatorControllers.AuthenticationController = controllers.NewAuthenticationController(databaseConnection)
	aggregatorControllers.ConnectorController = controllers.NewConnectorController(databaseConnection)
	aggregatorControllers.AggregatorController = controllers.NewAggregatorController(databaseConnection)
	aggregatorControllers.UserController = controllers.NewUserController(databaseConnection)
	aggregatorControllers.RoleController = controllers.NewRoleController(databaseConnection)
	aggregatorControllers.ConfigurationAggregatorController = controllers.NewConfigurationAggregatorController(databaseConnection)
	aggregatorControllers.ConfigurationConnectorController = controllers.NewConfigurationConnectorController(databaseConnection)

	return aggregatorControllers
}
