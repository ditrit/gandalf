/*
 * Swagger Gandalf
 *
 * This is a sample Petstore server.  You can find  out more about Swagger at  [http://swagger.io](http://swagger.io) or on  [irc.freenode.net, #swagger](http://swagger.io/irc/).
 *
 * API version: 1.0.0-oas3
 * Contact: romain.fairant@orness.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/ditrit/Gandalf/1.0.0/",
		Index,
	},

	Route{
		"CreateAuthorization",
		strings.ToUpper("Post"),
		"/ditrit/Gandalf/1.0.0/authorization",
		IsAuthorized(CreateAuthorization),
	},

	Route{
		"DeleteAuthorization",
		strings.ToUpper("Delete"),
		"/ditrit/Gandalf/1.0.0/authorization/{authorizationId:[0-9]+}",
		IsAuthorized(DeleteAuthorization),
	},

	Route{
		"GetAuthorizationById",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/authorization/{authorizationId:[0-9]+}",
		IsAuthorized(GetAuthorizationById),
	},

	Route{
		"ListAuthorization",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/authorization",
		IsAuthorized(ListAuthorization),
	},

	Route{
		"UpdateAuthorization",
		strings.ToUpper("Put"),
		"/ditrit/Gandalf/1.0.0/authorization/{authorizationId}",
		IsAuthorized(UpdateAuthorization),
	},

	Route{
		"CreateDomain",
		strings.ToUpper("Post"),
		"/ditrit/Gandalf/1.0.0/domain/{domainId:[0-9]+}",
		IsAuthorized(CreateDomain),
	},

	Route{
		"DeleteDomain",
		strings.ToUpper("Delete"),
		"/ditrit/Gandalf/1.0.0/domain/{domainId:[0-9]+}",
		IsAuthorized(DeleteDomain),
	},

	Route{
		"GetDomainById",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/domain/{domainId:[0-9]+}",
		IsAuthorized(GetDomainById),
	},

	Route{
		"ListDomain",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/domain",
		IsAuthorized(ListDomain),
	},

	Route{
		"TreeDomain",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/domain/tree",
		IsAuthorized(GetDomainTree),
	},

	Route{
		"UpdateDomain",
		strings.ToUpper("Put"),
		"/ditrit/Gandalf/1.0.0/domain/{domainId:[0-9]+}",
		IsAuthorized(UpdateDomain),
	},

	Route{
		"ListDomainTag",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/domain/{domainId:[0-9]+}/tag",
		IsAuthorized(ListDomainTag),
	},

	Route{
		"CreateDomainTag",
		strings.ToUpper("Post"),
		"/ditrit/Gandalf/1.0.0/domain/{domainId:[0-9]+}/tag/{tagId:[0-9]+}",
		IsAuthorized(CreateDomainTag),
	},

	Route{
		"DeleteDomainTag",
		strings.ToUpper("Delete"),
		"/ditrit/Gandalf/1.0.0/domain/{domainId:[0-9]+}/tag/{tagId:[0-9]+}",
		IsAuthorized(DeleteDomainTag),
	},

	Route{
		"ListDomainEnvironmentg",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/domain/{domainId:[0-9]+}/environment",
		IsAuthorized(ListDomainEnvironment),
	},

	Route{
		"CreateDomainEnvironment",
		strings.ToUpper("Post"),
		"/ditrit/Gandalf/1.0.0/domain/{domainId:[0-9]+}/environment/{environmentId:[0-9]+}",
		IsAuthorized(CreateDomainEnvironment),
	},

	Route{
		"DeleteDomainEnvironment",
		strings.ToUpper("Delete"),
		"/ditrit/Gandalf/1.0.0/domain/{domainId:[0-9]+}/environment/{environmentId:[0-9]+}",
		IsAuthorized(DeleteDomainEnvironment),
	},
	///

	Route{
		"CreateEventType",
		strings.ToUpper("Post"),
		"/ditrit/Gandalf/1.0.0/eventType",
		IsAuthorized(CreateEventType),
	},

	Route{
		"DeleteEventType",
		strings.ToUpper("Delete"),
		"/ditrit/Gandalf/1.0.0/eventType/{eventTypeId:[0-9]+}",
		IsAuthorized(DeleteEventType),
	},

	Route{
		"GetEventTypeById",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/eventType/{eventTypeId:[0-9]+}",
		IsAuthorized(GetEventTypeById),
	},

	Route{
		"GetEventTypeByName",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/eventType/{eventTypeName}",
		IsAuthorized(GetEventTypeByName),
	},

	Route{
		"ListEventType",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/eventType",
		IsAuthorized(ListEventType),
	},

	Route{
		"UpdateEventType",
		strings.ToUpper("Put"),
		"/ditrit/Gandalf/1.0.0/eventType/{eventTypeId:[0-9]+}",
		IsAuthorized(UpdateEventType),
	},

	Route{
		"CreateEventTypeToPoll",
		strings.ToUpper("Post"),
		"/ditrit/Gandalf/1.0.0/eventTypeToPoll",
		IsAuthorized(CreateEventTypeToPoll),
	},

	Route{
		"DeleteEventTypeToPoll",
		strings.ToUpper("Delete"),
		"/ditrit/Gandalf/1.0.0/eventTypeToPoll/{eventTypeToPollId:[0-9]+}",
		IsAuthorized(DeleteEventTypeToPoll),
	},

	Route{
		"GetEventTypeToPollById",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/eventTypeToPoll/{eventTypeToPollId:[0-9]+}",
		IsAuthorized(GetEventTypeToPollById),
	},

	Route{
		"ListEventTypeToPoll",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/eventTypeToPoll",
		IsAuthorized(ListEventTypeToPoll),
	},

	Route{
		"UpdateEventTypeToPoll",
		strings.ToUpper("Put"),
		"/ditrit/Gandalf/1.0.0/eventTypeToPoll/{eventTypeToPollId:[0-9]+}",
		IsAuthorized(UpdateEventTypeToPoll),
	},

	Route{
		"GetLogicalComponentByName",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/logicalcomponent/{logicalComponentName}",
		IsAuthorized(GetLogicalComponentByName),
	},

	Route{
		"UploadLogicalComponentByTenantAndType",
		strings.ToUpper("Post"),
		"/ditrit/Gandalf/1.0.0/logicalcomponent/upload/{tenantName}/{typeName}",
		IsAuthorized(UploadLogicalComponentByTenantAndType),
	},

	Route{
		"CreateLibrary",
		strings.ToUpper("Post"),
		"/ditrit/Gandalf/1.0.0/library",
		IsAuthorized(CreateLibrary),
	},

	Route{
		"DeleteLibrary",
		strings.ToUpper("Delete"),
		"/ditrit/Gandalf/1.0.0/library/{libraryId:[0-9]+}",
		IsAuthorized(DeleteLibrary),
	},

	Route{
		"GetLibraryById",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/library/{libraryId:[0-9]+}",
		IsAuthorized(GetLibraryById),
	},

	Route{
		"ListLibrary",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/library",
		IsAuthorized(ListLibrary),
	},

	Route{
		"UpdateLibrary",
		strings.ToUpper("Put"),
		"/ditrit/Gandalf/1.0.0/library/{libraryId:[0-9]+}",
		IsAuthorized(UpdateLibrary),
	},

	Route{
		"CreateProduct",
		strings.ToUpper("Post"),
		"/ditrit/Gandalf/1.0.0/product",
		IsAuthorized(CreateProduct),
	},

	Route{
		"DeleteProduct",
		strings.ToUpper("Delete"),
		"/ditrit/Gandalf/1.0.0/product/{productId:[0-9]+}",
		IsAuthorized(DeleteProduct),
	},

	Route{
		"GetProductById",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/product/{productId:[0-9]+}",
		IsAuthorized(GetProductById),
	},

	Route{
		"ListProduct",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/product",
		IsAuthorized(ListProduct),
	},

	Route{
		"UpdateProduct",
		strings.ToUpper("Put"),
		"/ditrit/Gandalf/1.0.0/product/{productId:[0-9]+}",
		IsAuthorized(UpdateProduct),
	},

	Route{
		"CreateConnectorProduct",
		strings.ToUpper("Post"),
		"/ditrit/Gandalf/1.0.0/connectorProduct",
		IsAuthorized(CreateConnectorProduct),
	},

	Route{
		"DeleteConnectorProduct",
		strings.ToUpper("Delete"),
		"/ditrit/Gandalf/1.0.0/connectorProduct/{connectorProductId:[0-9]+}",
		IsAuthorized(DeleteConnectorProduct),
	},

	Route{
		"GetConnectorProductById",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/connectorProduct/{connectorProductId:[0-9]+}",
		IsAuthorized(GetConnectorProductById),
	},

	Route{
		"ListConnectorProduct",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/connectorProduct",
		IsAuthorized(ListConnectorProduct),
	},

	Route{
		"UpdateConnectorProduct",
		strings.ToUpper("Put"),
		"/ditrit/Gandalf/1.0.0/connectorProduct/{connectorProductId:[0-9]+",
		IsAuthorized(UpdateConnectorProduct),
	},

	Route{
		"CreateResource",
		strings.ToUpper("Post"),
		"/ditrit/Gandalf/1.0.0/resource",
		IsAuthorized(CreateResource),
	},

	Route{
		"DeleteResource",
		strings.ToUpper("Delete"),
		"/ditrit/Gandalf/1.0.0/resource/{resourceId:[0-9]+}",
		IsAuthorized(DeleteResource),
	},

	Route{
		"GetResourceById",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/resource/{resourceId:[0-9]+}",
		IsAuthorized(GetResourceById),
	},

	Route{
		"GetResourceByName",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/resource/{resourceName}",
		IsAuthorized(GetResourceByName),
	},

	Route{
		"ListResource",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/resource",
		IsAuthorized(ListResource),
	},

	Route{
		"UpdateResource",
		strings.ToUpper("Put"),
		"/ditrit/Gandalf/1.0.0/resource/{resourceId:[0-9]+}",
		IsAuthorized(UpdateResource),
	},

	Route{
		"CreateResourceType",
		strings.ToUpper("Post"),
		"/ditrit/Gandalf/1.0.0/resourceType",
		IsAuthorized(CreateResourceType),
	},

	Route{
		"DeleteResourceType",
		strings.ToUpper("Delete"),
		"/ditrit/Gandalf/1.0.0/resourceType/{resourceTypeId:[0-9]+}",
		IsAuthorized(DeleteResourceType),
	},

	Route{
		"GetResourceTypeById",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/resourceType/{resourceTypeId:[0-9]+}",
		IsAuthorized(GetResourceTypeById),
	},

	Route{
		"GetResourceTypeByName",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/resourceType/{resourceTypeName}",
		IsAuthorized(GetResourceTypeByName),
	},

	Route{
		"ListResourceType",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/resourceType",
		IsAuthorized(ListResourceType),
	},

	Route{
		"UpdateResourceType",
		strings.ToUpper("Put"),
		"/ditrit/Gandalf/1.0.0/resourceType/{resourceTypeId:[0-9]+}",
		IsAuthorized(UpdateResourceType),
	},

	Route{
		"CreateRole",
		strings.ToUpper("Post"),
		"/ditrit/Gandalf/1.0.0/role",
		IsAuthorized(CreateRole),
	},

	Route{
		"DeleteRole",
		strings.ToUpper("Delete"),
		"/ditrit/Gandalf/1.0.0/role/{roleId:[0-9]+}",
		IsAuthorized(DeleteRole),
	},

	Route{
		"GetRoleById",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/role/{roleId:[0-9]+}",
		IsAuthorized(GetRoleById),
	},

	Route{
		"ListRole",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/role",
		IsAuthorized(ListRole),
	},

	Route{
		"UpdateRole",
		strings.ToUpper("Put"),
		"/ditrit/Gandalf/1.0.0/role/{roleId:[0-9]+}",
		IsAuthorized(UpdateRole),
	},

	Route{
		"CreateEnvironment",
		strings.ToUpper("Post"),
		"/ditrit/Gandalf/1.0.0/environment",
		IsAuthorized(CreateEnvironment),
	},

	Route{
		"DeleteEnvironment",
		strings.ToUpper("Delete"),
		"/ditrit/Gandalf/1.0.0/environment/{environmentId:[0-9]+}",
		IsAuthorized(DeleteEnvironment),
	},

	Route{
		"GetEnvironmentById",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/environment/{environmentId:[0-9]+}",
		IsAuthorized(GetEnvironmentById),
	},

	Route{
		"ListEnvironment",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/environment",
		IsAuthorized(ListEnvironment),
	},

	Route{
		"UpdateEnvironment",
		strings.ToUpper("Put"),
		"/ditrit/Gandalf/1.0.0/environment/{environmentId:[0-9]+}",
		IsAuthorized(UpdateEnvironment),
	},

	Route{
		"CreateEnvironmentType",
		strings.ToUpper("Post"),
		"/ditrit/Gandalf/1.0.0/environmentType",
		IsAuthorized(CreateEnvironmentType),
	},

	Route{
		"DeleteEnvironmentType",
		strings.ToUpper("Delete"),
		"/ditrit/Gandalf/1.0.0/environmentType/{environmentTypeId:[0-9]+}",
		IsAuthorized(DeleteEnvironmentType),
	},

	Route{
		"GetEnvironmentTypeById",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/environmentType/{environmentTypeId:[0-9]+}",
		IsAuthorized(GetEnvironmentTypeById),
	},

	Route{
		"ListEnvironmentType",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/environmentType",
		IsAuthorized(ListEnvironmentType),
	},

	Route{
		"UpdateEnvironmentType",
		strings.ToUpper("Put"),
		"/ditrit/Gandalf/1.0.0/environmentType/{environmentTypeId:[0-9]+}",
		IsAuthorized(UpdateEnvironmentType),
	},

	Route{
		"CreateSecretAssignement",
		strings.ToUpper("Post"),
		"/ditrit/Gandalf/1.0.0/secretAssignement",
		IsAuthorized(CreateSecretAssignement),
	},

	Route{
		"ListSecretAssignement",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/secretAssignement",
		IsAuthorized(ListSecretAssignement),
	},

	Route{
		"CreateTag",
		strings.ToUpper("Post"),
		"/ditrit/Gandalf/1.0.0/tag/{tagId:[0-9]+}",
		IsAuthorized(CreateTag),
	},

	Route{
		"DeleteTag",
		strings.ToUpper("Delete"),
		"/ditrit/Gandalf/1.0.0/tag/{tagId:[0-9]+}",
		IsAuthorized(DeleteTag),
	},

	Route{
		"GetTagById",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/tag/{tagId:[0-9]+}",
		IsAuthorized(GetTagById),
	},

	Route{
		"ListTag",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/tag",
		IsAuthorized(ListTag),
	},

	Route{
		"UpdateTag",
		strings.ToUpper("Put"),
		"/ditrit/Gandalf/1.0.0/tag/{tagId:[0-9]+}",
		IsAuthorized(UpdateTag),
	},

	Route{
		"TreeTag",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/tag/tree",
		IsAuthorized(GetTagTree),
	},

	Route{
		"CreateTenant",
		strings.ToUpper("Post"),
		"/ditrit/Gandalf/1.0.0/tenant",
		IsAuthorized(CreateTenant),
	},

	Route{
		"DeleteTenant",
		strings.ToUpper("Delete"),
		"/ditrit/Gandalf/1.0.0/tenant/{tenantId:[0-9]+}",
		IsAuthorized(DeleteTenant),
	},

	Route{
		"GetTenantById",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/tenant/{tenantId:[0-9]+}",
		IsAuthorized(GetTenantById),
	},

	Route{
		"ListTenant",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/tenant",
		IsAuthorized(ListTenant),
	},

	Route{
		"UpdateTenant",
		strings.ToUpper("Put"),
		"/ditrit/Gandalf/1.0.0/tenant/{tenantId:[0-9]+}",
		IsAuthorized(UpdateTenant),
	},

	Route{
		"CreateUser",
		strings.ToUpper("Post"),
		"/ditrit/Gandalf/1.0.0/user",
		IsAuthorized(CreateUser),
	},

	Route{
		"DeleteUser",
		strings.ToUpper("Delete"),
		"/ditrit/Gandalf/1.0.0/user/{userId:[0-9]+}",
		IsAuthorized(DeleteUser),
	},

	Route{
		"GetUserById",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/user/{userId:[0-9]+}",
		IsAuthorized(GetUserById),
	},

	Route{
		"GetUserByName",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/user/{userName}",
		IsAuthorized(GetUserByName),
	},

	Route{
		"ListUser",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/user",
		IsAuthorized(ListUser),
	},

	Route{
		"LoginUser",
		strings.ToUpper("Post"),
		"/ditrit/Gandalf/1.0.0/user/login",
		LoginUser,
	},

	Route{
		"LogoutUser",
		strings.ToUpper("Post"),
		"/ditrit/Gandalf/1.0.0/user/logout",
		LogoutUser,
	},

	Route{
		"RegisterUser",
		strings.ToUpper("Post"),
		"/ditrit/Gandalf/1.0.0/user/register",
		RegisterUser,
	},

	Route{
		"UpdateUser",
		strings.ToUpper("Put"),
		"/ditrit/Gandalf/1.0.0/user/{userId:[0-9]+}",
		IsAuthorized(UpdateUser),
	},
}
