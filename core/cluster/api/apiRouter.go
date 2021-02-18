package api

import (
	"github.com/ditrit/gandalf/core/cluster/database"

	"github.com/gorilla/mux"
)

// GetRouter :
func GetRouter(databaseConnection *database.DatabaseConnection) *mux.Router {

	//CONTROLLERS
	controllers := ReturnControllers(databaseConnection)

	//URLS
	urls := ReturnURLS()

	mux := mux.NewRouter()
	mux.Use(CommonMiddleware)
	//TODO REVOIR
	mux.HandleFunc(urls.GANDALF_LOGIN_PATH, controllers.gandalfAuthenticationController.Login).Methods("POST")
	mux.HandleFunc(urls.TENANTS_LOGIN_PATH, controllers.tenantsAuthenticationController.Login).Methods("POST")

	//mux.Handle("/images/*", http.StripPrefix("/images/", http.FileServer(http.Dir("./app/tmpl/images/"))))
	//mux.PathPrefix("/api/v1/").Subrouter()

	subg := mux.PathPrefix("/auth").Subrouter()
	subg.Use(GandalfJwtVerify)

	//GANDALF
	//CLUSTER
	subg.HandleFunc(urls.GANDALF_CLUSTER_PATH_LIST, controllers.gandalfClusterController.List).Methods("GET")
	subg.HandleFunc(urls.GANDALF_CLUSTER_PATH_CREATE, controllers.gandalfClusterController.Create).Methods("POST")
	subg.HandleFunc(urls.GANDALF_CLUSTER_PATH_DECLARE_MEMBER, controllers.gandalfClusterController.DeclareMember).Methods("GET")
	subg.HandleFunc(urls.GANDALF_CLUSTER_PATH_READ, controllers.gandalfClusterController.Read).Methods("GET")
	subg.HandleFunc(urls.GANDALF_CLUSTER_PATH_UPDATE, controllers.gandalfClusterController.Update).Methods("PUT")
	subg.HandleFunc(urls.GANDALF_CLUSTER_PATH_DELETE, controllers.gandalfClusterController.Delete).Methods("DELETE")

	//ROLE
	//subg.HandleFunc(urls.GANDALF_ROLE_PATH_LIST, controllers.gandalfRoleController.List).Methods("GET")
	//subg.HandleFunc(urls.GANDALF_ROLE_PATH_CREATE, controllers.gandalfRoleController.Create).Methods("POST")
	//subg.HandleFunc(urls.GANDALF_ROLE_PATH_READ, controllers.gandalfRoleController.Read).Methods("GET")
	//subg.HandleFunc(urls.GANDALF_ROLE_PATH_UPDATE, controllers.gandalfRoleController.Update).Methods("PUT")
	//subg.HandleFunc(urls.GANDALF_ROLE_PATH_DELETE, controllers.gandalfRoleController.Delete).Methods("DELETE")

	//TENANT
	subg.HandleFunc(urls.GANDALF_TENANT_PATH_LIST, controllers.gandalfTenantController.List).Methods("GET")
	subg.HandleFunc(urls.GANDALF_TENANT_PATH_CREATE, controllers.gandalfTenantController.Create).Methods("POST")
	subg.HandleFunc(urls.GANDALF_TENANT_PATH_READ, controllers.gandalfTenantController.Read).Methods("GET")
	subg.HandleFunc(urls.GANDALF_TENANT_PATH_UPDATE, controllers.gandalfTenantController.Update).Methods("PUT")
	subg.HandleFunc(urls.GANDALF_TENANT_PATH_DELETE, controllers.gandalfTenantController.Delete).Methods("DELETE")

	//USER
	subg.HandleFunc(urls.GANDALF_USER_PATH_LIST, controllers.gandalfUserController.List).Methods("GET")
	subg.HandleFunc(urls.GANDALF_USER_PATH_CREATE, controllers.gandalfUserController.Create).Methods("POST")
	subg.HandleFunc(urls.GANDALF_USER_PATH_READ, controllers.gandalfUserController.Read).Methods("GET")
	subg.HandleFunc(urls.GANDALF_USER_PATH_READ_BY_NAME, controllers.gandalfUserController.ReadByName).Methods("GET")
	subg.HandleFunc(urls.GANDALF_USER_PATH_UPDATE, controllers.gandalfUserController.Update).Methods("PUT")
	subg.HandleFunc(urls.GANDALF_USER_PATH_DELETE, controllers.gandalfUserController.Delete).Methods("DELETE")

	//CONFIGURATION
	subg.HandleFunc(urls.GANDALF_CONFIGURATION_PATH_LIST, controllers.gandalfConfigurationController.List).Methods("GET")
	subg.HandleFunc(urls.GANDALF_CONFIGURATION_PATH_CREATE, controllers.gandalfConfigurationController.Create).Methods("POST")
	subg.HandleFunc(urls.GANDALF_CONFIGURATION_PATH_READ, controllers.gandalfConfigurationController.Read).Methods("GET")
	subg.HandleFunc(urls.GANDALF_CONFIGURATION_PATH_UPDATE, controllers.gandalfConfigurationController.Update).Methods("PUT")
	subg.HandleFunc(urls.GANDALF_CONFIGURATION_PATH_DELETE, controllers.gandalfConfigurationController.Delete).Methods("DELETE")
	subg.HandleFunc(urls.GANDALF_CONFIGURATION_PATH_UPLOAD, controllers.gandalfConfigurationController.Upload).Methods("POST")

	//TENANTS
	//AGGREGATOR
	subg.HandleFunc(urls.TENANTS_AGGREGATOR_PATH_LIST, controllers.tenantsAggregatorController.List).Methods("GET")
	subg.HandleFunc(urls.TENANTS_AGGREGATOR_PATH_CREATE, controllers.tenantsAggregatorController.Create).Methods("POST")
	subg.HandleFunc(urls.TENANTS_AGGREGATOR_PATH_DECLARE_MEMBER, controllers.tenantsAggregatorController.DeclareMember).Methods("GET")
	subg.HandleFunc(urls.TENANTS_AGGREGATOR_PATH_READ, controllers.tenantsAggregatorController.Read).Methods("GET")
	subg.HandleFunc(urls.TENANTS_AGGREGATOR_PATH_UPDATE, controllers.tenantsAggregatorController.Update).Methods("PUT")
	subg.HandleFunc(urls.TENANTS_AGGREGATOR_PATH_DELETE, controllers.tenantsAggregatorController.Delete).Methods("DELETE")

	//CONNECTOR
	subg.HandleFunc(urls.TENANTS_CONNECTOR_PATH_LIST, controllers.tenantsConnectorController.List).Methods("GET")
	subg.HandleFunc(urls.TENANTS_CONNECTOR_PATH_CREATE, controllers.tenantsConnectorController.Create).Methods("POST")
	subg.HandleFunc(urls.TENANTS_CONNECTOR_PATH_DECLARE_MEMBER, controllers.tenantsConnectorController.DeclareMember).Methods("GET")
	subg.HandleFunc(urls.TENANTS_CONNECTOR_PATH_READ, controllers.tenantsConnectorController.Read).Methods("GET")
	subg.HandleFunc(urls.TENANTS_CONNECTOR_PATH_UPDATE, controllers.tenantsConnectorController.Update).Methods("PUT")
	subg.HandleFunc(urls.TENANTS_CONNECTOR_PATH_DELETE, controllers.tenantsConnectorController.Delete).Methods("DELETE")

	//ROLE
	subg.HandleFunc(urls.TENANTS_ROLE_PATH_LIST, controllers.tenantsRoleController.List).Methods("GET")
	subg.HandleFunc(urls.TENANTS_ROLE_PATH_CREATE, controllers.tenantsRoleController.Create).Methods("POST")
	subg.HandleFunc(urls.TENANTS_ROLE_PATH_READ, controllers.tenantsRoleController.Read).Methods("GET")
	subg.HandleFunc(urls.TENANTS_ROLE_PATH_UPDATE, controllers.tenantsRoleController.Update).Methods("PUT")
	subg.HandleFunc(urls.TENANTS_ROLE_PATH_DELETE, controllers.tenantsRoleController.Delete).Methods("DELETE")

	//USER
	subg.HandleFunc(urls.TENANTS_USER_PATH_LIST, controllers.tenantsUserController.List).Methods("GET")
	subg.HandleFunc(urls.TENANTS_USER_PATH_CREATE, controllers.tenantsUserController.Create).Methods("POST")
	subg.HandleFunc(urls.TENANTS_USER_PATH_READ, controllers.tenantsUserController.Read).Methods("GET")
	subg.HandleFunc(urls.TENANTS_USER_PATH_UPDATE, controllers.tenantsUserController.Update).Methods("PUT")
	subg.HandleFunc(urls.TENANTS_USER_PATH_DELETE, controllers.tenantsUserController.Delete).Methods("DELETE")

	//CONFIGURATION AGGREGATOR
	subg.HandleFunc(urls.TENANTS_CONFIGURATION_AGGREGATOR_PATH_LIST, controllers.tenantsConfigurationAggregatorController.List).Methods("GET")
	subg.HandleFunc(urls.TENANTS_CONFIGURATION_AGGREGATOR_PATH_CREATE, controllers.tenantsConfigurationAggregatorController.Create).Methods("POST")
	subg.HandleFunc(urls.TENANTS_CONFIGURATION_AGGREGATOR_PATH_READ, controllers.tenantsConfigurationAggregatorController.Read).Methods("GET")
	subg.HandleFunc(urls.TENANTS_CONFIGURATION_AGGREGATOR_PATH_UPDATE, controllers.tenantsConfigurationAggregatorController.Update).Methods("PUT")
	subg.HandleFunc(urls.TENANTS_CONFIGURATION_AGGREGATOR_PATH_DELETE, controllers.tenantsConfigurationAggregatorController.Delete).Methods("DELETE")
	subg.HandleFunc(urls.TENANTS_CONFIGURATION_AGGREGATOR_PATH_UPLOAD, controllers.tenantsConfigurationAggregatorController.Upload).Methods("POST")

	//CONFIGURATION CONNECTOR
	subg.HandleFunc(urls.TENANTS_CONFIGURATION_CONNECTOR_PATH_LIST, controllers.tenantsConfigurationConnectorController.List).Methods("GET")
	subg.HandleFunc(urls.TENANTS_CONFIGURATION_CONNECTOR_PATH_CREATE, controllers.tenantsConfigurationConnectorController.Create).Methods("POST")
	subg.HandleFunc(urls.TENANTS_CONFIGURATION_CONNECTOR_PATH_READ, controllers.tenantsConfigurationConnectorController.Read).Methods("GET")
	subg.HandleFunc(urls.TENANTS_CONFIGURATION_CONNECTOR_PATH_UPDATE, controllers.tenantsConfigurationConnectorController.Update).Methods("PUT")
	subg.HandleFunc(urls.TENANTS_CONFIGURATION_CONNECTOR_PATH_DELETE, controllers.tenantsConfigurationConnectorController.Delete).Methods("DELETE")
	subg.HandleFunc(urls.TENANTS_CONFIGURATION_CONNECTOR_PATH_UPLOAD, controllers.tenantsConfigurationConnectorController.Upload).Methods("POST")

	return mux
}
