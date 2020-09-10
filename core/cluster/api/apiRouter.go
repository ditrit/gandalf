package api

import (
	"github.com/jinzhu/gorm"

	"github.com/go-chi/chi"
)

func GetRouter(gandalfDatabase *gorm.DB, mapDatabase map[string]*gorm.DB, databasePath string) *chi.Mux {

	//CONTROLLERS
	controllers := ReturnControllers(gandalfDatabase, mapDatabase, databasePath)

	//URLS
	urls := ReturnURLS()

	mux := chi.NewRouter()
	//mux.Handle("/images/*", http.StripPrefix("/images/", http.FileServer(http.Dir("./app/tmpl/images/"))))
	//mux.PathPrefix("/api/v1/").Subrouter()

	mux.Group(func(mux chi.Router) {

		//CLUSTER
		mux.MethodFunc("GET", urls.GANDALF_CLUSTER_PATH_LIST, controllers.gandalfClusterController.List)
		mux.MethodFunc("POST", urls.GANDALF_CLUSTER_PATH_CREATE, controllers.gandalfClusterController.Create)
		mux.MethodFunc("GET", urls.GANDALF_CLUSTER_PATH_READ, controllers.gandalfClusterController.Read)
		mux.MethodFunc("PUT", urls.GANDALF_CLUSTER_PATH_UPDATE, controllers.gandalfClusterController.Update)
		mux.MethodFunc("DELETE", urls.GANDALF_CLUSTER_PATH_DELETE, controllers.gandalfClusterController.Delete)

		//ROLE
		mux.MethodFunc("GET", urls.GANDALF_ROLE_PATH_LIST, controllers.gandalfRoleController.List)
		mux.MethodFunc("POST", urls.GANDALF_ROLE_PATH_CREATE, controllers.gandalfRoleController.Create)
		mux.MethodFunc("GET", urls.GANDALF_ROLE_PATH_READ, controllers.gandalfRoleController.Read)
		mux.MethodFunc("PUT", urls.GANDALF_ROLE_PATH_UPDATE, controllers.gandalfRoleController.Update)
		mux.MethodFunc("DELETE", urls.GANDALF_ROLE_PATH_DELETE, controllers.gandalfRoleController.Delete)

		//TENANT
		mux.MethodFunc("GET", urls.GANDALF_TENANT_PATH_LIST, controllers.gandalfTenantController.List)
		mux.MethodFunc("POST", urls.GANDALF_TENANT_PATH_CREATE, controllers.gandalfTenantController.Create)
		mux.MethodFunc("GET", urls.GANDALF_TENANT_PATH_READ, controllers.gandalfTenantController.Read)
		mux.MethodFunc("PUT", urls.GANDALF_TENANT_PATH_UPDATE, controllers.gandalfTenantController.Update)
		mux.MethodFunc("DELETE", urls.GANDALF_TENANT_PATH_DELETE, controllers.gandalfTenantController.Delete)

		//USER
		mux.MethodFunc("GET", urls.GANDALF_USER_PATH_LIST, controllers.gandalfUserController.List)
		mux.MethodFunc("POST", urls.GANDALF_USER_PATH_CREATE, controllers.gandalfUserController.Create)
		mux.MethodFunc("GET", urls.GANDALF_USER_PATH_READ, controllers.gandalfUserController.Read)
		mux.MethodFunc("PUT", urls.GANDALF_USER_PATH_UPDATE, controllers.gandalfUserController.Update)
		mux.MethodFunc("DELETE", urls.GANDALF_USER_PATH_DELETE, controllers.gandalfUserController.Delete)

	})

	mux.Group(func(mux chi.Router) {
		//AGGREGATOR
		mux.MethodFunc("GET", urls.TENANTS_AGGREGATOR_PATH_LIST, controllers.tenantsAggregatorController.List)
		mux.MethodFunc("POST", urls.TENANTS_AGGREGATOR_PATH_CREATE, controllers.tenantsAggregatorController.Create)
		mux.MethodFunc("GET", urls.TENANTS_AGGREGATOR_PATH_READ, controllers.tenantsAggregatorController.Read)
		mux.MethodFunc("PUT", urls.TENANTS_AGGREGATOR_PATH_UPDATE, controllers.tenantsAggregatorController.Update)
		mux.MethodFunc("DELETE", urls.TENANTS_AGGREGATOR_PATH_DELETE, controllers.tenantsAggregatorController.Delete)

		//CONNECTOR
		mux.MethodFunc("GET", urls.TENANTS_CONNECTOR_PATH_LIST, controllers.tenantsConnectorController.List)
		mux.MethodFunc("POST", urls.TENANTS_CONNECTOR_PATH_CREATE, controllers.tenantsConnectorController.Create)
		mux.MethodFunc("GET", urls.TENANTS_CONNECTOR_PATH_READ, controllers.tenantsConnectorController.Read)
		mux.MethodFunc("PUT", urls.TENANTS_CONNECTOR_PATH_UPDATE, controllers.tenantsConnectorController.Update)
		mux.MethodFunc("DELETE", urls.TENANTS_CONNECTOR_PATH_DELETE, controllers.tenantsConnectorController.Delete)

		//ROLE
		mux.MethodFunc("GET", urls.TENANTS_ROLE_PATH_LIST, controllers.tenantsRoleController.List)
		mux.MethodFunc("POST", urls.TENANTS_ROLE_PATH_CREATE, controllers.tenantsRoleController.Create)
		mux.MethodFunc("GET", urls.TENANTS_ROLE_PATH_READ, controllers.tenantsRoleController.Read)
		mux.MethodFunc("PUT", urls.TENANTS_ROLE_PATH_UPDATE, controllers.tenantsRoleController.Update)
		mux.MethodFunc("DELETE", urls.TENANTS_ROLE_PATH_DELETE, controllers.tenantsRoleController.Delete)

		//USER
		mux.MethodFunc("GET", urls.TENANTS_USER_PATH_LIST, controllers.tenantsUserController.List)
		mux.MethodFunc("POST", urls.TENANTS_USER_PATH_CREATE, controllers.tenantsUserController.Create)
		mux.MethodFunc("GET", urls.TENANTS_USER_PATH_READ, controllers.tenantsUserController.Read)
		mux.MethodFunc("PUT", urls.TENANTS_USER_PATH_UPDATE, controllers.tenantsUserController.Update)
		mux.MethodFunc("DELETE", urls.TENANTS_USER_PATH_DELETE, controllers.tenantsUserController.Delete)

	})

	return mux
}
