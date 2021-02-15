package api

import (
	"github.com/ditrit/gandalf/core/cluster/database"

	"github.com/ditrit/gandalf/core/cluster/api/controllers/gandalf"
	"github.com/ditrit/gandalf/core/cluster/api/controllers/tenants"
)

// Controllers :
type Controllers struct {
	gandalfAuthenticationController *gandalf.AuthenticationController
	gandalfClusterController        *gandalf.ClusterController
	gandalfTenantController         *gandalf.TenantController
	gandalfRoleController           *gandalf.RoleController
	gandalfUserController           *gandalf.UserController
	gandalfConfigurationController  *gandalf.ConfigurationController

	tenantsAuthenticationController          *tenants.AuthenticationController
	tenantsAggregatorController              *tenants.AggregatorController
	tenantsConnectorController               *tenants.ConnectorController
	tenantsRoleController                    *tenants.RoleController
	tenantsUserController                    *tenants.UserController
	tenantsConfigurationAggregatorController *tenants.ConfigurationAggregatorController
	tenantsConfigurationConnectorController  *tenants.ConfigurationConnectorController
}

// ReturnControllers :
func ReturnControllers(databaseConnection *database.DatabaseConnection) *Controllers {

	controllers := new(Controllers)
	controllers.gandalfAuthenticationController = gandalf.NewAuthenticationController(databaseConnection)
	controllers.gandalfClusterController = gandalf.NewClusterController(databaseConnection)
	controllers.gandalfTenantController = gandalf.NewTenantController(databaseConnection)
	controllers.gandalfUserController = gandalf.NewUserController(databaseConnection)
	controllers.gandalfRoleController = gandalf.NewRoleController(databaseConnection)
	controllers.gandalfConfigurationController = gandalf.NewConfigurationController(databaseConnection)

	controllers.tenantsAuthenticationController = tenants.NewAuthenticationController(databaseConnection)
	controllers.tenantsConnectorController = tenants.NewConnectorController(databaseConnection)
	controllers.tenantsAggregatorController = tenants.NewAggregatorController(databaseConnection)
	controllers.tenantsUserController = tenants.NewUserController(databaseConnection)
	controllers.tenantsRoleController = tenants.NewRoleController(databaseConnection)
	controllers.tenantsConfigurationAggregatorController = tenants.NewConfigurationAggregatorController(databaseConnection)
	controllers.tenantsConfigurationConnectorController = tenants.NewConfigurationConnectorController(databaseConnection)

	return controllers
}
