package api

import (
	"gandalf/core/cluster/api/controller/gandalf"
	"gandalf/core/cluster/api/controller/tenants"

	"github.com/jinzhu/gorm"
)

type Controllers struct {
	gandalfClusterController *gandalf.ClusterController
	gandalfTenantController  *gandalf.TenantController
	gandalfRoleController    *gandalf.RoleController
	gandalfUserController    *gandalf.UserController

	tenantsAggregatorController *tenants.AggregatorController
	tenantsConnectorController  *tenants.ConnectorController
	tenantsRoleController       *tenants.RoleController
	tenantsUserController       *tenants.UserController
}

func ReturnControllers(gandalfDatabase *gorm.DB, mapDatabase map[string]*gorm.DB, databasePath string) *Controllers {

	controllers := new(Controllers)
	controllers.gandalfClusterController = gandalf.NewClusterController(gandalfDatabase)
	controllers.gandalfTenantController = gandalf.NewTenantController(gandalfDatabase)
	controllers.gandalfUserController = gandalf.NewUserController(gandalfDatabase)
	controllers.gandalfRoleController = gandalf.NewRoleController(gandalfDatabase)

	controllers.tenantsConnectorController = tenants.NewConnectorController(mapDatabase, databasePath)
	controllers.tenantsAggregatorController = tenants.NewAggregatorController(mapDatabase, databasePath)
	controllers.tenantsUserController = tenants.NewUserController(mapDatabase, databasePath)
	controllers.tenantsRoleController = tenants.NewRoleController(mapDatabase, databasePath)

	return controllers
}
