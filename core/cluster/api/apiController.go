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

	tenantsAuthenticationController *tenants.AuthenticationController
	tenantsAggregatorController     *tenants.AggregatorController
	tenantsConnectorController      *tenants.ConnectorController
	tenantsRoleController           *tenants.RoleController
	tenantsUserController           *tenants.UserController
}

// ReturnControllers :
func ReturnControllers(gandalfDatabase *gorm.DB, mapDatabase map[string]*gorm.DB, databasePath string) *Controllers {

	controllers := new(Controllers)
	controllers.gandalfAuthenticationController = gandalf.NewAuthenticationController(gandalfDatabase)
	controllers.gandalfClusterController = gandalf.NewClusterController(gandalfDatabase)
	controllers.gandalfTenantController = gandalf.NewTenantController(gandalfDatabase, databasePath)
	controllers.gandalfUserController = gandalf.NewUserController(gandalfDatabase)
	controllers.gandalfRoleController = gandalf.NewRoleController(gandalfDatabase)

	controllers.tenantsAuthenticationController = tenants.NewAuthenticationController(mapDatabase, databasePath)
	controllers.tenantsConnectorController = tenants.NewConnectorController(mapDatabase, databasePath)
	controllers.tenantsAggregatorController = tenants.NewAggregatorController(mapDatabase, databasePath)
	controllers.tenantsUserController = tenants.NewUserController(mapDatabase, databasePath)
	controllers.tenantsRoleController = tenants.NewRoleController(mapDatabase, databasePath)

	return controllers
}
