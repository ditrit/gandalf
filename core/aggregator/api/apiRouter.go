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
	mux.HandleFunc(urls.LOGIN_PATH, controllers.tenantsAuthenticationController.Login).Methods("POST")

	//mux.Handle("/images/*", http.StripPrefix("/images/", http.FileServer(http.Dir("./app/tmpl/images/"))))
	//mux.PathPrefix("/api/v1/").Subrouter()

	subt := mux.PathPrefix("/auth").Subrouter()
	subt.Use(TenantsJwtVerify)
	//TENANTS
	//AGGREGATOR
	subt.HandleFunc(urls.AGGREGATOR_PATH_LIST, controllers.tenantsAggregatorController.List).Methods("GET")
	subt.HandleFunc(urls.AGGREGATOR_PATH_CREATE, controllers.tenantsAggregatorController.Create).Methods("POST")
	subt.HandleFunc(urls.AGGREGATOR_PATH_DECLARE_MEMBER, controllers.tenantsAggregatorController.DeclareMember).Methods("GET")
	subt.HandleFunc(urls.AGGREGATOR_PATH_READ, controllers.tenantsAggregatorController.Read).Methods("GET")
	subt.HandleFunc(urls.AGGREGATOR_PATH_UPDATE, controllers.tenantsAggregatorController.Update).Methods("PUT")
	subt.HandleFunc(urls.AGGREGATOR_PATH_DELETE, controllers.tenantsAggregatorController.Delete).Methods("DELETE")

	//CONNECTOR
	subt.HandleFunc(urls.CONNECTOR_PATH_LIST, controllers.tenantsConnectorController.List).Methods("GET")
	subt.HandleFunc(urls.CONNECTOR_PATH_CREATE, controllers.tenantsConnectorController.Create).Methods("POST")
	subt.HandleFunc(urls.CONNECTOR_PATH_DECLARE_MEMBER, controllers.tenantsConnectorController.DeclareMember).Methods("GET")
	subt.HandleFunc(urls.CONNECTOR_PATH_READ, controllers.tenantsConnectorController.Read).Methods("GET")
	subt.HandleFunc(urls.CONNECTOR_PATH_UPDATE, controllers.tenantsConnectorController.Update).Methods("PUT")
	subt.HandleFunc(urls.CONNECTOR_PATH_DELETE, controllers.tenantsConnectorController.Delete).Methods("DELETE")

	//ROLE
	subt.HandleFunc(urls.ROLE_PATH_LIST, controllers.tenantsRoleController.List).Methods("GET")
	subt.HandleFunc(urls.ROLE_PATH_CREATE, controllers.tenantsRoleController.Create).Methods("POST")
	subt.HandleFunc(urls.ROLE_PATH_READ, controllers.tenantsRoleController.Read).Methods("GET")
	subt.HandleFunc(urls.ROLE_PATH_UPDATE, controllers.tenantsRoleController.Update).Methods("PUT")
	subt.HandleFunc(urls.ROLE_PATH_DELETE, controllers.tenantsRoleController.Delete).Methods("DELETE")

	//USER
	subt.HandleFunc(urls.USER_PATH_LIST, controllers.tenantsUserController.List).Methods("GET")
	subt.HandleFunc(urls.USER_PATH_CREATE, controllers.tenantsUserController.Create).Methods("POST")
	subt.HandleFunc(urls.USER_PATH_READ, controllers.tenantsUserController.Read).Methods("GET")
	subt.HandleFunc(urls.USER_PATH_UPDATE, controllers.tenantsUserController.Update).Methods("PUT")
	subt.HandleFunc(urls.USER_PATH_DELETE, controllers.tenantsUserController.Delete).Methods("DELETE")

	//CONFIGURATION AGGREGATOR
	subt.HandleFunc(urls.CONFIGURATION_AGGREGATOR_PATH_LIST, controllers.tenantsConfigurationAggregatorController.List).Methods("GET")
	subt.HandleFunc(urls.CONFIGURATION_AGGREGATOR_PATH_CREATE, controllers.tenantsConfigurationAggregatorController.Create).Methods("POST")
	subt.HandleFunc(urls.CONFIGURATION_AGGREGATOR_PATH_READ, controllers.tenantsConfigurationAggregatorController.Read).Methods("GET")
	subt.HandleFunc(urls.CONFIGURATION_AGGREGATOR_PATH_UPDATE, controllers.tenantsConfigurationAggregatorController.Update).Methods("PUT")
	subt.HandleFunc(urls.CONFIGURATION_AGGREGATOR_PATH_DELETE, controllers.tenantsConfigurationAggregatorController.Delete).Methods("DELETE")
	subt.HandleFunc(urls.CONFIGURATION_AGGREGATOR_PATH_UPLOAD, controllers.tenantsConfigurationAggregatorController.Upload).Methods("POST")

	//CONFIGURATION CONNECTOR
	subt.HandleFunc(urls.CONFIGURATION_CONNECTOR_PATH_LIST, controllers.tenantsConfigurationConnectorController.List).Methods("GET")
	subt.HandleFunc(urls.CONFIGURATION_CONNECTOR_PATH_CREATE, controllers.tenantsConfigurationConnectorController.Create).Methods("POST")
	subt.HandleFunc(urls.CONFIGURATION_CONNECTOR_PATH_READ, controllers.tenantsConfigurationConnectorController.Read).Methods("GET")
	subt.HandleFunc(urls.CONFIGURATION_CONNECTOR_PATH_UPDATE, controllers.tenantsConfigurationConnectorController.Update).Methods("PUT")
	subt.HandleFunc(urls.CONFIGURATION_CONNECTOR_PATH_DELETE, controllers.tenantsConfigurationConnectorController.Delete).Methods("DELETE")
	subt.HandleFunc(urls.CONFIGURATION_CONNECTOR_PATH_UPLOAD, controllers.tenantsConfigurationConnectorController.Upload).Methods("POST")

	return mux
}
