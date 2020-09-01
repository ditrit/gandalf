package api

import (
	"gandalf/core/api/controller"
	"net/http"

	"github.com/ditrit/gandalf/core/database"

	"github.com/go-chi/chi"
)

func GetRouterWithUrl(databasePath string) *chi.Mux {

	gandalfDatabase := database.NewGandalfDatabaseClient(databasePath)

	//CONTROLLERS
	aggregatorController := controller.NewAggregatorController(gandalfDatabase)
	clusterController := controller.NewClusterController(gandalfDatabase)
	connectorController := controller.NewConnectorController(gandalfDatabase)
	roleController := controller.NewRoleController(gandalfDatabase)
	tenantController := controller.NewTenantController(gandalfDatabase)
	userController := controller.NewuserController(gandalfDatabase)

	//URLS
	urls := ReturnUrls()

	mux := chi.NewRouter()
	mux.Handle("/images/*", http.StripPrefix("/images/", http.FileServer(http.Dir("./app/tmpl/images/"))))
	mux.PathPrefix("/api/v1/").Subrouter()

	mux.Group(func(mux chi.Router) {
		//AGGREGATOR
		mux.MethodFunc("GET", urls.AGGREGATOR_PATH_LIST, aggregatorController.List)
		mux.MethodFunc("POST", urls.AGGREGATOR_PATH_CREATE, aggregatorController.Create)
		mux.MethodFunc("GET", urls.AGGREGATOR_PATH_READ, aggregatorController.Read)
		mux.MethodFunc("PUT", urls.AGGREGATOR_PATH_UPDATE, aggregatorController.Update)
		mux.MethodFunc("DELETE", urls.AGGREGATOR_PATH_DELETE, aggregatorController.Delete)

		//CLUSTER
		mux.MethodFunc("GET", urls.CLUSTER_PATH_LIST, clusterController.List)
		mux.MethodFunc("POST", urls.CLUSTER_PATH_CREATE, clusterController.Create)
		mux.MethodFunc("GET", urls.CLUSTER_PATH_READ, clusterController.Read)
		mux.MethodFunc("PUT", urls.CLUSTER_PATH_UPDATE, clusterController.Update)
		mux.MethodFunc("DELETE", urls.CLUSTER_PATH_DELETE, clusterController.Delete)

		//CONNECTOR
		mux.MethodFunc("GET", urls.CONNECTOR_PATH_LIST, connectorController.List)
		mux.MethodFunc("POST", urls.CONNECTOR_PATH_CREATE, connectorController.Create)
		mux.MethodFunc("GET", urls.CONNECTOR_PATH_READ, connectorController.Read)
		mux.MethodFunc("PUT", urls.CONNECTOR_PATH_UPDATE, connectorController.Update)
		mux.MethodFunc("DELETE", urls.CONNECTOR_PATH_DELETE, connectorController.Delete)

		//ROLE
		mux.MethodFunc("GET", urls.ROLE_PATH_LIST, roleController.List)
		mux.MethodFunc("POST", urls.ROLE_PATH_CREATE, roleController.Create)
		mux.MethodFunc("GET", urls.ROLE_PATH_READ, roleController.Read)
		mux.MethodFunc("PUT", urls.ROLE_PATH_UPDATE, roleController.Update)
		mux.MethodFunc("DELETE", urls.ROLE_PATH_DELETE, roleController.Delete)

		//TENANT
		mux.MethodFunc("GET", urls.TENANT_PATH_LIST, tenantController.List)
		mux.MethodFunc("POST", urls.TENANT_PATH_CREATE, tenantController.Create)
		mux.MethodFunc("GET", urls.TENANT_PATH_READ, tenantController.Read)
		mux.MethodFunc("PUT", urls.TENANT_PATH_UPDATE, tenantController.Update)
		mux.MethodFunc("DELETE", urls.TENANT_PATH_DELETE, tenantController.Delete)

		//USER
		mux.MethodFunc("GET", urls.USER_PATH_LIST, userController.List)
		mux.MethodFunc("POST", urls.USER_PATH_CREATE, userController.Create)
		mux.MethodFunc("GET", urls.USER_PATH_READ, userController.Read)
		mux.MethodFunc("PUT", urls.USER_PATH_UPDATE, userController.Update)
		mux.MethodFunc("DELETE", urls.USER_PATH_DELETE, userController.Delete)

	})
	return mux
}
