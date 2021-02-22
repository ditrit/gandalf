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

	mux.HandleFunc(urls.LOGIN_PATH, controllers.AuthenticationController.Login).Methods("POST")
	mux.HandleFunc(urls.CLI_PATH, controllers.CliController.Cli).Methods("GET")

	//mux.Handle("/images/*", http.StripPrefix("/images/", http.FileServer(http.Dir("./app/tmpl/images/"))))
	//mux.PathPrefix("/api/v1/").Subrouter()

	subg := mux.PathPrefix("/auth").Subrouter()
	subg.Use(GandalfJwtVerify)

	//GANDALF
	//CLUSTER
	subg.HandleFunc(urls.CLUSTER_PATH_LIST, controllers.ClusterController.List).Methods("GET")
	subg.HandleFunc(urls.CLUSTER_PATH_CREATE, controllers.ClusterController.Create).Methods("POST")
	subg.HandleFunc(urls.CLUSTER_PATH_DECLARE_MEMBER, controllers.ClusterController.DeclareMember).Methods("GET")
	subg.HandleFunc(urls.CLUSTER_PATH_READ, controllers.ClusterController.Read).Methods("GET")
	subg.HandleFunc(urls.CLUSTER_PATH_UPDATE, controllers.ClusterController.Update).Methods("PUT")
	subg.HandleFunc(urls.CLUSTER_PATH_DELETE, controllers.ClusterController.Delete).Methods("DELETE")

	//TENANT
	subg.HandleFunc(urls.TENANT_PATH_LIST, controllers.TenantController.List).Methods("GET")
	subg.HandleFunc(urls.TENANT_PATH_CREATE, controllers.TenantController.Create).Methods("POST")
	subg.HandleFunc(urls.TENANT_PATH_READ, controllers.TenantController.Read).Methods("GET")
	subg.HandleFunc(urls.TENANT_PATH_UPDATE, controllers.TenantController.Update).Methods("PUT")
	subg.HandleFunc(urls.TENANT_PATH_DELETE, controllers.TenantController.Delete).Methods("DELETE")

	//USER
	subg.HandleFunc(urls.USER_PATH_LIST, controllers.UserController.List).Methods("GET")
	subg.HandleFunc(urls.USER_PATH_CREATE, controllers.UserController.Create).Methods("POST")
	subg.HandleFunc(urls.USER_PATH_READ, controllers.UserController.Read).Methods("GET")
	subg.HandleFunc(urls.USER_PATH_READ_BY_NAME, controllers.UserController.ReadByName).Methods("GET")
	subg.HandleFunc(urls.USER_PATH_UPDATE, controllers.UserController.Update).Methods("PUT")
	subg.HandleFunc(urls.USER_PATH_DELETE, controllers.UserController.Delete).Methods("DELETE")

	//CONFIGURATION
	subg.HandleFunc(urls.CONFIGURATION_PATH_LIST, controllers.ConfigurationController.List).Methods("GET")
	subg.HandleFunc(urls.CONFIGURATION_PATH_CREATE, controllers.ConfigurationController.Create).Methods("POST")
	subg.HandleFunc(urls.CONFIGURATION_PATH_READ, controllers.ConfigurationController.Read).Methods("GET")
	subg.HandleFunc(urls.CONFIGURATION_PATH_UPDATE, controllers.ConfigurationController.Update).Methods("PUT")
	subg.HandleFunc(urls.CONFIGURATION_PATH_DELETE, controllers.ConfigurationController.Delete).Methods("DELETE")
	subg.HandleFunc(urls.CONFIGURATION_PATH_UPLOAD, controllers.ConfigurationController.Upload).Methods("POST")

	//TENANTS
	//AGGREGATOR
	subg.HandleFunc(urls.AGGREGATOR_PATH_LIST, controllers.AggregatorController.List).Methods("GET")
	subg.HandleFunc(urls.AGGREGATOR_PATH_CREATE, controllers.AggregatorController.Create).Methods("POST")
	subg.HandleFunc(urls.AGGREGATOR_PATH_DECLARE_MEMBER, controllers.AggregatorController.DeclareMember).Methods("GET")
	subg.HandleFunc(urls.AGGREGATOR_PATH_READ, controllers.AggregatorController.Read).Methods("GET")
	subg.HandleFunc(urls.AGGREGATOR_PATH_UPDATE, controllers.AggregatorController.Update).Methods("PUT")
	subg.HandleFunc(urls.AGGREGATOR_PATH_DELETE, controllers.AggregatorController.Delete).Methods("DELETE")

	//CONNECTOR
	subg.HandleFunc(urls.CONNECTOR_PATH_LIST, controllers.ConnectorController.List).Methods("GET")
	subg.HandleFunc(urls.CONNECTOR_PATH_CREATE, controllers.ConnectorController.Create).Methods("POST")
	subg.HandleFunc(urls.CONNECTOR_PATH_DECLARE_MEMBER, controllers.ConnectorController.DeclareMember).Methods("GET")
	subg.HandleFunc(urls.CONNECTOR_PATH_READ, controllers.ConnectorController.Read).Methods("GET")
	subg.HandleFunc(urls.CONNECTOR_PATH_UPDATE, controllers.ConnectorController.Update).Methods("PUT")
	subg.HandleFunc(urls.CONNECTOR_PATH_DELETE, controllers.ConnectorController.Delete).Methods("DELETE")

	//USER
	subg.HandleFunc(urls.ADMIN_TENANT_PATH_LIST, controllers.AdminTenantController.List).Methods("GET")
	subg.HandleFunc(urls.ADMIN_TENANT_PATH_CREATE, controllers.AdminTenantController.Create).Methods("POST")
	//subg.HandleFunc(urls.ADMIN_TENANT_PATH_READ, controllers.AdminTenantController.Read).Methods("GET")
	//subg.HandleFunc(urls.ADMIN_TENANT_PATH_UPDATE, controllers.AdminTenantController.Update).Methods("PUT")
	//subg.HandleFunc(urls.ADMIN_TENANT_PATH_DELETE, controllers.AdminTenantController.Delete).Methods("DELETE")

	//CONFIGURATION AGGREGATOR
	subg.HandleFunc(urls.CONFIGURATION_AGGREGATOR_PATH_LIST, controllers.ConfigurationAggregatorController.List).Methods("GET")
	subg.HandleFunc(urls.CONFIGURATION_AGGREGATOR_PATH_CREATE, controllers.ConfigurationAggregatorController.Create).Methods("POST")
	subg.HandleFunc(urls.CONFIGURATION_AGGREGATOR_PATH_READ, controllers.ConfigurationAggregatorController.Read).Methods("GET")
	subg.HandleFunc(urls.CONFIGURATION_AGGREGATOR_PATH_UPDATE, controllers.ConfigurationAggregatorController.Update).Methods("PUT")
	subg.HandleFunc(urls.CONFIGURATION_AGGREGATOR_PATH_DELETE, controllers.ConfigurationAggregatorController.Delete).Methods("DELETE")
	subg.HandleFunc(urls.CONFIGURATION_AGGREGATOR_PATH_UPLOAD, controllers.ConfigurationAggregatorController.Upload).Methods("POST")

	//CONFIGURATION CONNECTOR
	subg.HandleFunc(urls.CONFIGURATION_CONNECTOR_PATH_LIST, controllers.ConfigurationConnectorController.List).Methods("GET")
	subg.HandleFunc(urls.CONFIGURATION_CONNECTOR_PATH_CREATE, controllers.ConfigurationConnectorController.Create).Methods("POST")
	subg.HandleFunc(urls.CONFIGURATION_CONNECTOR_PATH_READ, controllers.ConfigurationConnectorController.Read).Methods("GET")
	subg.HandleFunc(urls.CONFIGURATION_CONNECTOR_PATH_UPDATE, controllers.ConfigurationConnectorController.Update).Methods("PUT")
	subg.HandleFunc(urls.CONFIGURATION_CONNECTOR_PATH_DELETE, controllers.ConfigurationConnectorController.Delete).Methods("DELETE")
	subg.HandleFunc(urls.CONFIGURATION_CONNECTOR_PATH_UPLOAD, controllers.ConfigurationConnectorController.Upload).Methods("POST")

	return mux
}
