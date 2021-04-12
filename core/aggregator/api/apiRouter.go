package api

import (
	"github.com/ditrit/gandalf/core/aggregator/database"

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
	mux.HandleFunc(urls.LOGIN_PATH, controllers.AuthenticationController.Login).Methods("POST")
	mux.HandleFunc(urls.CLI_PATH, controllers.CliController.Cli).Methods("GET")

	//mux.Handle("/images/*", http.StripPrefix("/images/", http.FileServer(http.Dir("./app/tmpl/images/"))))
	//mux.PathPrefix("/api/v1/").Subrouter()

	subt := mux.PathPrefix("/auth").Subrouter()
	subt.Use(TenantsJwtVerify)

	//LOGICAL COMPONENT
	subt.HandleFunc(urls.LOGICAL_COMPONENT_PAHT_UPLOAD, controllers.LogicalComponentController.Upload).Methods("POST")

	//ROLE
	subt.HandleFunc(urls.ROLE_PATH_LIST, controllers.RoleController.List).Methods("GET")
	subt.HandleFunc(urls.ROLE_PATH_CREATE, controllers.RoleController.Create).Methods("POST")
	subt.HandleFunc(urls.ROLE_PATH_READ, controllers.RoleController.Read).Methods("GET")
	subt.HandleFunc(urls.ROLE_PATH_UPDATE, controllers.RoleController.Update).Methods("PUT")
	subt.HandleFunc(urls.ROLE_PATH_DELETE, controllers.RoleController.Delete).Methods("DELETE")

	//USER
	subt.HandleFunc(urls.USER_PATH_LIST, controllers.UserController.List).Methods("GET")
	subt.HandleFunc(urls.USER_PATH_CREATE, controllers.UserController.Create).Methods("POST")
	subt.HandleFunc(urls.USER_PATH_READ, controllers.UserController.Read).Methods("GET")
	subt.HandleFunc(urls.USER_PATH_UPDATE, controllers.UserController.Update).Methods("PUT")
	subt.HandleFunc(urls.USER_PATH_DELETE, controllers.UserController.Delete).Methods("DELETE")

	//TENANTS
	/* //AGGREGATOR
	subt.HandleFunc(urls.AGGREGATOR_PATH_LIST, controllers.AggregatorController.List).Methods("GET")
	subt.HandleFunc(urls.AGGREGATOR_PATH_CREATE, controllers.AggregatorController.Create).Methods("POST")
	subt.HandleFunc(urls.AGGREGATOR_PATH_DECLARE_MEMBER, controllers.AggregatorController.DeclareMember).Methods("GET")
	subt.HandleFunc(urls.AGGREGATOR_PATH_READ, controllers.AggregatorController.Read).Methods("GET")
	subt.HandleFunc(urls.AGGREGATOR_PATH_UPDATE, controllers.AggregatorController.Update).Methods("PUT")
	subt.HandleFunc(urls.AGGREGATOR_PATH_DELETE, controllers.AggregatorController.Delete).Methods("DELETE")

	//CONNECTOR
	subt.HandleFunc(urls.CONNECTOR_PATH_LIST, controllers.ConnectorController.List).Methods("GET")
	subt.HandleFunc(urls.CONNECTOR_PATH_CREATE, controllers.ConnectorController.Create).Methods("POST")
	subt.HandleFunc(urls.CONNECTOR_PATH_DECLARE_MEMBER, controllers.ConnectorController.DeclareMember).Methods("GET")
	subt.HandleFunc(urls.CONNECTOR_PATH_READ, controllers.ConnectorController.Read).Methods("GET")
	subt.HandleFunc(urls.CONNECTOR_PATH_UPDATE, controllers.ConnectorController.Update).Methods("PUT")
	subt.HandleFunc(urls.CONNECTOR_PATH_DELETE, controllers.ConnectorController.Delete).Methods("DELETE")



	//CONFIGURATION AGGREGATOR
	subt.HandleFunc(urls.CONFIGURATION_AGGREGATOR_PATH_LIST, controllers.ConfigurationAggregatorController.List).Methods("GET")
	subt.HandleFunc(urls.CONFIGURATION_AGGREGATOR_PATH_CREATE, controllers.ConfigurationAggregatorController.Create).Methods("POST")
	subt.HandleFunc(urls.CONFIGURATION_AGGREGATOR_PATH_READ, controllers.ConfigurationAggregatorController.Read).Methods("GET")
	subt.HandleFunc(urls.CONFIGURATION_AGGREGATOR_PATH_UPDATE, controllers.ConfigurationAggregatorController.Update).Methods("PUT")
	subt.HandleFunc(urls.CONFIGURATION_AGGREGATOR_PATH_DELETE, controllers.ConfigurationAggregatorController.Delete).Methods("DELETE")
	subt.HandleFunc(urls.CONFIGURATION_AGGREGATOR_PATH_UPLOAD, controllers.ConfigurationAggregatorController.Upload).Methods("POST")

	//CONFIGURATION CONNECTOR
	subt.HandleFunc(urls.CONFIGURATION_CONNECTOR_PATH_LIST, controllers.ConfigurationConnectorController.List).Methods("GET")
	subt.HandleFunc(urls.CONFIGURATION_CONNECTOR_PATH_CREATE, controllers.ConfigurationConnectorController.Create).Methods("POST")
	subt.HandleFunc(urls.CONFIGURATION_CONNECTOR_PATH_READ, controllers.ConfigurationConnectorController.Read).Methods("GET")
	subt.HandleFunc(urls.CONFIGURATION_CONNECTOR_PATH_UPDATE, controllers.ConfigurationConnectorController.Update).Methods("PUT")
	subt.HandleFunc(urls.CONFIGURATION_CONNECTOR_PATH_DELETE, controllers.ConfigurationConnectorController.Delete).Methods("DELETE")
	subt.HandleFunc(urls.CONFIGURATION_CONNECTOR_PATH_UPLOAD, controllers.ConfigurationConnectorController.Upload).Methods("POST") */

	return mux
}
