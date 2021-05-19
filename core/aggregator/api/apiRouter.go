package api

import (
	"github.com/ditrit/gandalf/core/aggregator/database"
	net "github.com/ditrit/shoset"

	"github.com/gorilla/mux"
)

// GetRouter :
func GetRouter(databaseConnection *database.DatabaseConnection, shoset *net.Shoset) *mux.Router {

	//CONTROLLERS
	controllers := ReturnControllers(databaseConnection, shoset)

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
	subt.HandleFunc(urls.LOGICAL_COMPONENT_PATH_UPLOAD, controllers.LogicalComponentController.Upload).Methods("POST")
	subt.HandleFunc(urls.LOGICAL_COMPONENT_PATH_READ_BY_NAME, controllers.LogicalComponentController.ReadByName).Methods("GET")

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
	subt.HandleFunc(urls.USER_PATH_READ_BY_NAME, controllers.UserController.ReadByName).Methods("GET")
	subt.HandleFunc(urls.USER_PATH_UPDATE, controllers.UserController.Update).Methods("PUT")
	subt.HandleFunc(urls.USER_PATH_DELETE, controllers.UserController.Delete).Methods("DELETE")

	//SECRET
	subt.HandleFunc(urls.SECRET_PATH_LIST, controllers.SecretAssignementController.List).Methods("GET")
	subt.HandleFunc(urls.SECRET_PATH_CREATE, controllers.SecretAssignementController.Create).Methods("POST")

	//TENANT
	subt.HandleFunc(urls.TENANT_PATH_LIST, controllers.TenantController.List).Methods("GET")
	subt.HandleFunc(urls.TENANT_PATH_CREATE, controllers.TenantController.Create).Methods("POST")
	subt.HandleFunc(urls.TENANT_PATH_READ, controllers.TenantController.Read).Methods("GET")
	subt.HandleFunc(urls.TENANT_PATH_UPDATE, controllers.TenantController.Update).Methods("PUT")
	subt.HandleFunc(urls.TENANT_PATH_DELETE, controllers.TenantController.Delete).Methods("DELETE")

	//RESOURCE
	subt.HandleFunc(urls.RESOURCE_PATH_LIST, controllers.ResourceController.List).Methods("GET")
	subt.HandleFunc(urls.RESOURCE_PATH_CREATE, controllers.ResourceController.Create).Methods("POST")
	subt.HandleFunc(urls.RESOURCE_PATH_READ, controllers.ResourceController.Read).Methods("GET")
	subt.HandleFunc(urls.RESOURCE_PATH_READ_BY_NAME, controllers.ResourceController.ReadByName).Methods("GET")
	subt.HandleFunc(urls.RESOURCE_PATH_UPDATE, controllers.ResourceController.Update).Methods("PUT")
	subt.HandleFunc(urls.RESOURCE_PATH_DELETE, controllers.ResourceController.Delete).Methods("DELETE")

	//DOMAIN
	subt.HandleFunc(urls.DOMAIN_PATH_LIST, controllers.DomainController.List).Methods("GET")
	subt.HandleFunc(urls.DOMAIN_PATH_CREATE, controllers.DomainController.Create).Methods("POST")
	subt.HandleFunc(urls.DOMAIN_PATH_READ, controllers.DomainController.Read).Methods("GET")
	subt.HandleFunc(urls.DOMAIN_PATH_READ_BY_NAME, controllers.DomainController.ReadByName).Methods("GET")
	subt.HandleFunc(urls.DOMAIN_PATH_UPDATE, controllers.DomainController.Update).Methods("PUT")
	subt.HandleFunc(urls.DOMAIN_PATH_DELETE, controllers.DomainController.Delete).Methods("DELETE")

	//EVENTTYPETOPOLL
	subt.HandleFunc(urls.EVENT_TYPE_TO_POLL_PATH_LIST, controllers.EventTypeToPollController.List).Methods("GET")
	subt.HandleFunc(urls.EVENT_TYPE_TO_POLL_PATH_CREATE, controllers.EventTypeToPollController.Create).Methods("POST")
	subt.HandleFunc(urls.EVENT_TYPE_TO_POLL_PATH_READ, controllers.EventTypeToPollController.Read).Methods("GET")
	subt.HandleFunc(urls.EVENT_TYPE_TO_POLL_PATH_UPDATE, controllers.EventTypeToPollController.Update).Methods("PUT")
	subt.HandleFunc(urls.EVENT_TYPE_TO_POLL_PATH_DELETE, controllers.EventTypeToPollController.Delete).Methods("DELETE")

	//RESOURCE TYPE
	subt.HandleFunc(urls.RESOURCE_TYPE_PATH_LIST, controllers.ResourceTypeController.List).Methods("GET")
	subt.HandleFunc(urls.RESOURCE_TYPE_PATH_CREATE, controllers.ResourceTypeController.Create).Methods("POST")
	subt.HandleFunc(urls.RESOURCE_TYPE_PATH_READ, controllers.ResourceTypeController.Read).Methods("GET")
	subt.HandleFunc(urls.RESOURCE_TYPE_PATH_READ_BY_NAME, controllers.ResourceTypeController.ReadByName).Methods("GET")
	subt.HandleFunc(urls.RESOURCE_TYPE_PATH_UPDATE, controllers.ResourceTypeController.Update).Methods("PUT")
	subt.HandleFunc(urls.RESOURCE_TYPE_PATH_DELETE, controllers.ResourceTypeController.Delete).Methods("DELETE")

	//EVENT TYPE
	subt.HandleFunc(urls.EVENT_TYPE_PATH_LIST, controllers.EventTypeController.List).Methods("GET")
	subt.HandleFunc(urls.EVENT_TYPE_PATH_CREATE, controllers.EventTypeController.Create).Methods("POST")
	subt.HandleFunc(urls.EVENT_TYPE_PATH_READ, controllers.EventTypeController.Read).Methods("GET")
	subt.HandleFunc(urls.EVENT_TYPE_PATH_READ_BY_NAME, controllers.EventTypeController.ReadByName).Methods("GET")
	subt.HandleFunc(urls.EVENT_TYPE_PATH_UPDATE, controllers.EventTypeController.Update).Methods("PUT")
	subt.HandleFunc(urls.EVENT_TYPE_PATH_DELETE, controllers.EventTypeController.Delete).Methods("DELETE")

	//APPLICATION
	subt.HandleFunc(urls.APPLICATION_TYPE_PATH_LIST, controllers.ApplicationController.List).Methods("GET")
	subt.HandleFunc(urls.APPLICATION_TYPE_PATH_CREATE, controllers.ApplicationController.Create).Methods("POST")
	subt.HandleFunc(urls.APPLICATION_TYPE_PATH_READ, controllers.ApplicationController.Read).Methods("GET")
	subt.HandleFunc(urls.APPLICATION_TYPE_PATH_READ_BY_NAME, controllers.ApplicationController.ReadByName).Methods("GET")
	subt.HandleFunc(urls.APPLICATION_TYPE_PATH_UPDATE, controllers.ApplicationController.Update).Methods("PUT")
	subt.HandleFunc(urls.APPLICATION_TYPE_PATH_DELETE, controllers.ApplicationController.Delete).Methods("DELETE")

	return mux
}
