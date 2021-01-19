package api

import (
	"github.com/ditrit/gandalf/core/cluster/api/controllers/gandalf"
	"github.com/ditrit/gandalf/core/cluster/api/controllers/tenants"

	"github.com/jinzhu/gorm"
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
func ReturnControllers(gandalfDatabase *gorm.DB, mapDatabase map[string]*gorm.DB, databasePath, databaseBindAddr string) *Controllers {

	controllers := new(Controllers)
	controllers.gandalfAuthenticationController = gandalf.NewAuthenticationController(gandalfDatabase)
	controllers.gandalfClusterController = gandalf.NewClusterController(gandalfDatabase)
	controllers.gandalfTenantController = gandalf.NewTenantController(gandalfDatabase, mapDatabase, databasePath, databaseBindAddr)
	controllers.gandalfUserController = gandalf.NewUserController(gandalfDatabase)
	controllers.gandalfRoleController = gandalf.NewRoleController(gandalfDatabase)
	controllers.gandalfConfigurationController = gandalf.NewConfigurationController(gandalfDatabase)

	controllers.tenantsAuthenticationController = tenants.NewAuthenticationController(mapDatabase)
	controllers.tenantsConnectorController = tenants.NewConnectorController(mapDatabase)
	controllers.tenantsAggregatorController = tenants.NewAggregatorController(mapDatabase)
	controllers.tenantsUserController = tenants.NewUserController(mapDatabase)
	controllers.tenantsRoleController = tenants.NewRoleController(mapDatabase)
	controllers.tenantsConfigurationAggregatorController = tenants.NewConfigurationAggregatorController(mapDatabase)
	controllers.tenantsConfigurationConnectorController = tenants.NewConfigurationConnectorController(mapDatabase)

	return controllers
}
