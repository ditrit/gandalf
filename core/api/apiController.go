package api

import (
	"gandalf/core/api/controller"

	"github.com/jinzhu/gorm"
)

type Controllers struct {
	aggregatorController *controller.AggregatorController
	clusterController    *controller.ClusterController
	connectorController  *controller.ConnectorController
	roleController       *controller.RoleController
	tenantController     *controller.TenantController
	userController       *controller.UserController
}

func ReturnControllers(gandalfDatabase *gorm.DB) *Controllers {

	controllers := new(Controllers)
	controllers.aggregatorController = controller.NewAggregatorController(gandalfDatabase)
	controllers.clusterController = controller.NewClusterController(gandalfDatabase)
	controllers.connectorController = controller.NewConnectorController(gandalfDatabase)
	controllers.roleController = controller.NewRoleController(gandalfDatabase)
	controllers.tenantController = controller.NewTenantController(gandalfDatabase)
	controllers.userController = controller.NewUserController(gandalfDatabase)

	return controllers
}
