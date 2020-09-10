package api

import (
	"github.com/jinzhu/gorm"

	"github.com/gorilla/mux"
)

func GetRouter(gandalfDatabase *gorm.DB, mapDatabase map[string]*gorm.DB, databasePath string) *mux.Router {

	//CONTROLLERS
	controllers := ReturnControllers(gandalfDatabase, mapDatabase, databasePath)

	//URLS
	urls := ReturnURLS()

	mux := mux.NewRouter()
	//mux.Handle("/images/*", http.StripPrefix("/images/", http.FileServer(http.Dir("./app/tmpl/images/"))))
	//mux.PathPrefix("/api/v1/").Subrouter()

	//GANDALF
	//CLUSTER
	mux.HandleFunc(urls.GANDALF_CLUSTER_PATH_LIST, controllers.gandalfClusterController.List).Methods("GET")
	mux.HandleFunc(urls.GANDALF_CLUSTER_PATH_CREATE, controllers.gandalfClusterController.Create).Methods("POST")
	mux.HandleFunc(urls.GANDALF_CLUSTER_PATH_READ, controllers.gandalfClusterController.Read).Methods("GET")
	mux.HandleFunc(urls.GANDALF_CLUSTER_PATH_UPDATE, controllers.gandalfClusterController.Update).Methods("PUT")
	mux.HandleFunc(urls.GANDALF_CLUSTER_PATH_DELETE, controllers.gandalfClusterController.Delete).Methods("DELETE")

	//ROLE
	mux.HandleFunc(urls.GANDALF_ROLE_PATH_LIST, controllers.gandalfRoleController.List).Methods("GET")
	mux.HandleFunc(urls.GANDALF_ROLE_PATH_CREATE, controllers.gandalfRoleController.Create).Methods("POST")
	mux.HandleFunc(urls.GANDALF_ROLE_PATH_READ, controllers.gandalfRoleController.Read).Methods("GET")
	mux.HandleFunc(urls.GANDALF_ROLE_PATH_UPDATE, controllers.gandalfRoleController.Update).Methods("PUT")
	mux.HandleFunc(urls.GANDALF_ROLE_PATH_DELETE, controllers.gandalfRoleController.Delete).Methods("DELETE")

	//TENANT
	mux.HandleFunc(urls.GANDALF_TENANT_PATH_LIST, controllers.gandalfTenantController.List).Methods("GET")
	mux.HandleFunc(urls.GANDALF_TENANT_PATH_CREATE, controllers.gandalfTenantController.Create).Methods("POST")
	mux.HandleFunc(urls.GANDALF_TENANT_PATH_READ, controllers.gandalfTenantController.Read).Methods("GET")
	mux.HandleFunc(urls.GANDALF_TENANT_PATH_UPDATE, controllers.gandalfTenantController.Update).Methods("PUT")
	mux.HandleFunc(urls.GANDALF_TENANT_PATH_DELETE, controllers.gandalfTenantController.Delete).Methods("DELETE")

	//USER
	mux.HandleFunc(urls.GANDALF_USER_PATH_LIST, controllers.gandalfUserController.List).Methods("GET")
	mux.HandleFunc(urls.GANDALF_USER_PATH_CREATE, controllers.gandalfUserController.Create).Methods("POST")
	mux.HandleFunc(urls.GANDALF_USER_PATH_READ, controllers.gandalfUserController.Read).Methods("GET")
	mux.HandleFunc(urls.GANDALF_USER_PATH_UPDATE, controllers.gandalfUserController.Update).Methods("PUT")
	mux.HandleFunc(urls.GANDALF_USER_PATH_DELETE, controllers.gandalfUserController.Delete).Methods("DELETE")

	//TENANTS
	//AGGREGATOR
	mux.HandleFunc(urls.TENANTS_AGGREGATOR_PATH_LIST, controllers.tenantsAggregatorController.List).Methods("GET")
	mux.HandleFunc(urls.TENANTS_AGGREGATOR_PATH_CREATE, controllers.tenantsAggregatorController.Create).Methods("POST")
	mux.HandleFunc(urls.TENANTS_AGGREGATOR_PATH_READ, controllers.tenantsAggregatorController.Read).Methods("GET")
	mux.HandleFunc(urls.TENANTS_AGGREGATOR_PATH_UPDATE, controllers.tenantsAggregatorController.Update).Methods("PUT")
	mux.HandleFunc(urls.TENANTS_AGGREGATOR_PATH_DELETE, controllers.tenantsAggregatorController.Delete).Methods("DELETE")

	//CONNECTOR
	mux.HandleFunc(urls.TENANTS_CONNECTOR_PATH_LIST, controllers.tenantsConnectorController.List).Methods("GET")
	mux.HandleFunc(urls.TENANTS_CONNECTOR_PATH_CREATE, controllers.tenantsConnectorController.Create).Methods("POST")
	mux.HandleFunc(urls.TENANTS_CONNECTOR_PATH_READ, controllers.tenantsConnectorController.Read).Methods("GET")
	mux.HandleFunc(urls.TENANTS_CONNECTOR_PATH_UPDATE, controllers.tenantsConnectorController.Update).Methods("PUT")
	mux.HandleFunc(urls.TENANTS_CONNECTOR_PATH_DELETE, controllers.tenantsConnectorController.Delete).Methods("DELETE")

	//ROLE
	mux.HandleFunc(urls.TENANTS_ROLE_PATH_LIST, controllers.tenantsRoleController.List).Methods("GET")
	mux.HandleFunc(urls.TENANTS_ROLE_PATH_CREATE, controllers.tenantsRoleController.Create).Methods("POST")
	mux.HandleFunc(urls.TENANTS_ROLE_PATH_READ, controllers.tenantsRoleController.Read).Methods("GET")
	mux.HandleFunc(urls.TENANTS_ROLE_PATH_UPDATE, controllers.tenantsRoleController.Update).Methods("PUT")
	mux.HandleFunc(urls.TENANTS_ROLE_PATH_DELETE, controllers.tenantsRoleController.Delete).Methods("DELETE")

	//USER
	mux.HandleFunc(urls.TENANTS_USER_PATH_LIST, controllers.tenantsUserController.List).Methods("GET")
	mux.HandleFunc(urls.TENANTS_USER_PATH_CREATE, controllers.tenantsUserController.Create).Methods("POST")
	mux.HandleFunc(urls.TENANTS_USER_PATH_READ, controllers.tenantsUserController.Read).Methods("GET")
	mux.HandleFunc(urls.TENANTS_USER_PATH_UPDATE, controllers.tenantsUserController.Update).Methods("PUT")
	mux.HandleFunc(urls.TENANTS_USER_PATH_DELETE, controllers.tenantsUserController.Delete).Methods("DELETE")

	return mux
}
