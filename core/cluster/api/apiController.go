package api

import (
	"github.com/ditrit/gandalf/core/cluster/database"

	"github.com/ditrit/gandalf/core/cluster/api/controllers"
)

// Controllers :
type Controllers struct {
	AuthenticationController          *controllers.AuthenticationController
	CliController                     *controllers.CliController
	ClusterController                 *controllers.ClusterController
	TenantController                  *controllers.TenantController
	UserController                    *controllers.UserController
	ConfigurationController           *controllers.ConfigurationController
	AggregatorController              *controllers.AggregatorController
	ConnectorController               *controllers.ConnectorController
	AdminTenantController             *controllers.AdminTenantController
	ConfigurationAggregatorController *controllers.ConfigurationAggregatorController
	ConfigurationConnectorController  *controllers.ConfigurationConnectorController
}

// ReturnControllers :
func ReturnControllers(databaseConnection *database.DatabaseConnection) *Controllers {

	clusterControllers := new(Controllers)
	clusterControllers.AuthenticationController = controllers.NewAuthenticationController(databaseConnection)
	clusterControllers.CliController = controllers.NewCliController()
	clusterControllers.ClusterController = controllers.NewClusterController(databaseConnection)
	clusterControllers.TenantController = controllers.NewTenantController(databaseConnection)
	clusterControllers.UserController = controllers.NewUserController(databaseConnection)
	clusterControllers.ConfigurationController = controllers.NewConfigurationController(databaseConnection)

	clusterControllers.ConnectorController = controllers.NewConnectorController(databaseConnection)
	clusterControllers.AggregatorController = controllers.NewAggregatorController(databaseConnection)
	clusterControllers.UserController = controllers.NewUserController(databaseConnection)
	clusterControllers.ConfigurationAggregatorController = controllers.NewConfigurationAggregatorController(databaseConnection)
	clusterControllers.ConfigurationConnectorController = controllers.NewConfigurationConnectorController(databaseConnection)

	return clusterControllers
}
