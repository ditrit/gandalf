package api

import (
	"github.com/ditrit/gandalf/core/cluster/database"

	"github.com/go-chi/chi"
)

func GetRouter(databasePath string) *chi.Mux {

	gandalfDatabase := database.NewGandalfDatabaseClient(databasePath)

	//CONTROLLERS
	controllers := ReturnControllers(gandalfDatabase)

	//URLS
	urls := ReturnURLS()

	mux := chi.NewRouter()
	//mux.Handle("/images/*", http.StripPrefix("/images/", http.FileServer(http.Dir("./app/tmpl/images/"))))
	//mux.PathPrefix("/api/v1/").Subrouter()

	mux.Group(func(mux chi.Router) {

		//CLUSTER
		mux.MethodFunc("GET", urls.GANDALF_CLUSTER_PATH_LIST, controllers.clusterController.List)
		mux.MethodFunc("POST", urls.GANDALF_CLUSTER_PATH_CREATE, controllers.clusterController.Create)
		mux.MethodFunc("GET", urls.GANDALF_CLUSTER_PATH_READ, controllers.clusterController.Read)
		mux.MethodFunc("PUT", urls.GANDALF_CLUSTER_PATH_UPDATE, controllers.clusterController.Update)
		mux.MethodFunc("DELETE", urls.GANDALF_CLUSTER_PATH_DELETE, controllers.clusterController.Delete)

		//ROLE
		mux.MethodFunc("GET", urls.GANDALF_ROLE_PATH_LIST, controllers.roleController.List)
		mux.MethodFunc("POST", urls.GANDALF_ROLE_PATH_CREATE, controllers.roleController.Create)
		mux.MethodFunc("GET", urls.GANDALF_ROLE_PATH_READ, controllers.roleController.Read)
		mux.MethodFunc("PUT", urls.GANDALF_ROLE_PATH_UPDATE, controllers.roleController.Update)
		mux.MethodFunc("DELETE", urls.GANDALF_ROLE_PATH_DELETE, controllers.roleController.Delete)

		//TENANT
		mux.MethodFunc("GET", urls.GANDALF_TENANT_PATH_LIST, controllers.tenantController.List)
		mux.MethodFunc("POST", urls.GANDALF_TENANT_PATH_CREATE, controllers.tenantController.Create)
		mux.MethodFunc("GET", urls.GANDALF_TENANT_PATH_READ, controllers.tenantController.Read)
		mux.MethodFunc("PUT", urls.GANDALF_TENANT_PATH_UPDATE, controllers.tenantController.Update)
		mux.MethodFunc("DELETE", urls.GANDALF_TENANT_PATH_DELETE, controllers.tenantController.Delete)

		//USER
		mux.MethodFunc("GET", urls.GANDALF_USER_PATH_LIST, controllers.userController.List)
		mux.MethodFunc("POST", urls.GANDALF_USER_PATH_CREATE, controllers.userController.Create)
		mux.MethodFunc("GET", urls.GANDALF_USER_PATH_READ, controllers.userController.Read)
		mux.MethodFunc("PUT", urls.GANDALF_USER_PATH_UPDATE, controllers.userController.Update)
		mux.MethodFunc("DELETE", urls.GANDALF_USER_PATH_DELETE, controllers.userController.Delete)

	})

	mux.Group(func(mux chi.Router) {
		//AGGREGATOR
		mux.MethodFunc("GET", urls.TENANTS_AGGREGATOR_PATH_LIST, controllers.aggregatorController.List)
		mux.MethodFunc("POST", urls.TENANTS_AGGREGATOR_PATH_CREATE, controllers.aggregatorController.Create)
		mux.MethodFunc("GET", urls.TENANTS_AGGREGATOR_PATH_READ, controllers.aggregatorController.Read)
		mux.MethodFunc("PUT", urls.TENANTS_AGGREGATOR_PATH_UPDATE, controllers.aggregatorController.Update)
		mux.MethodFunc("DELETE", urls.TENANTS_AGGREGATOR_PATH_DELETE, controllers.aggregatorController.Delete)

		//CONNECTOR
		mux.MethodFunc("GET", urls.TENANTS_CONNECTOR_PATH_LIST, controllers.connectorController.List)
		mux.MethodFunc("POST", urls.TENANTS_CONNECTOR_PATH_CREATE, controllers.connectorController.Create)
		mux.MethodFunc("GET", urls.TENANTS_CONNECTOR_PATH_READ, controllers.connectorController.Read)
		mux.MethodFunc("PUT", urls.TENANTS_CONNECTOR_PATH_UPDATE, controllers.connectorController.Update)
		mux.MethodFunc("DELETE", urls.TENANTS_CONNECTOR_PATH_DELETE, controllers.connectorController.Delete)

		//ROLE
		mux.MethodFunc("GET", urls.TENANTS_ROLE_PATH_LIST, controllers.roleController.List)
		mux.MethodFunc("POST", urls.TENANTS_ROLE_PATH_CREATE, controllers.roleController.Create)
		mux.MethodFunc("GET", urls.TENANTS_ROLE_PATH_READ, controllers.roleController.Read)
		mux.MethodFunc("PUT", urls.TENANTS_ROLE_PATH_UPDATE, controllers.roleController.Update)
		mux.MethodFunc("DELETE", urls.TENANTS_ROLE_PATH_DELETE, controllers.roleController.Delete)

		//USER
		mux.MethodFunc("GET", urls.TENANTS_USER_PATH_LIST, controllers.userController.List)
		mux.MethodFunc("POST", urls.TENANTS_USER_PATH_CREATE, controllers.userController.Create)
		mux.MethodFunc("GET", urls.TENANTS_USER_PATH_READ, controllers.userController.Read)
		mux.MethodFunc("PUT", urls.TENANTS_USER_PATH_UPDATE, controllers.userController.Update)
		mux.MethodFunc("DELETE", urls.TENANTS_USER_PATH_DELETE, controllers.userController.Delete)

	})

	return mux
}
