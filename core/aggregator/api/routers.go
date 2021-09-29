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
		"/ditrit/Gandalf/1.0.0/authorization/{authorizationId}",
		IsAuthorized(DeleteAuthorization),
	},

	Route{
		"GetAuthorizationById",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/authorization/{authorizationId}",
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
		"/ditrit/Gandalf/1.0.0/domain/{domainId}",
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
		"CreateEventType",
		strings.ToUpper("Post"),
		"/ditrit/Gandalf/1.0.0/eventType",
		IsAuthorized(CreateEventType),
	},

	Route{
		"DeleteEventType",
		strings.ToUpper("Delete"),
		"/ditrit/Gandalf/1.0.0/eventType/{eventTypeId}",
		IsAuthorized(DeleteEventType),
	},

	Route{
		"GetEventTypeById",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/eventType/{eventTypeId}",
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
		"/ditrit/Gandalf/1.0.0/eventType/{eventTypeId}",
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
		"/ditrit/Gandalf/1.0.0/eventTypeToPoll/{eventTypeToPollId}",
		IsAuthorized(DeleteEventTypeToPoll),
	},

	Route{
		"GetEventTypeToPollById",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/eventTypeToPoll/{eventTypeToPollId}",
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
		"/ditrit/Gandalf/1.0.0/eventTypeToPoll/{eventTypeToPollId}",
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
		"CreateProduct",
		strings.ToUpper("Post"),
		"/ditrit/Gandalf/1.0.0/product",
		IsAuthorized(CreateProduct),
	},

	Route{
		"DeleteProduct",
		strings.ToUpper("Delete"),
		"/ditrit/Gandalf/1.0.0/product/{productId}",
		IsAuthorized(DeleteProduct),
	},

	Route{
		"GetProductById",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/product/{productId}",
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
		"/ditrit/Gandalf/1.0.0/product/{productId}",
		IsAuthorized(UpdateProduct),
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
		"/ditrit/Gandalf/1.0.0/resource/{resourceId}",
		IsAuthorized(DeleteResource),
	},

	Route{
		"GetResourceById",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/resource/{resourceId}",
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
		"/ditrit/Gandalf/1.0.0/resource/{resourceId}",
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
		"/ditrit/Gandalf/1.0.0/resourceType/{resourceTypeId}",
		IsAuthorized(DeleteResourceType),
	},

	Route{
		"GetResourceTypeById",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/resourceType/{resourceTypeId}",
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
		"/ditrit/Gandalf/1.0.0/resourceType/{resourceTypeId}",
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
		"/ditrit/Gandalf/1.0.0/role/{roleId}",
		IsAuthorized(DeleteRole),
	},

	Route{
		"GetRoleById",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/role/{roleId}",
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
		"/ditrit/Gandalf/1.0.0/role/{roleId}",
		IsAuthorized(UpdateRole),
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
		"/ditrit/Gandalf/1.0.0/tenant/{tenantId}",
		IsAuthorized(DeleteTenant),
	},

	Route{
		"GetTenantById",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/tenant/{tenantId}",
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
		"/ditrit/Gandalf/1.0.0/tenant/{tenantId}",
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
		"/ditrit/Gandalf/1.0.0/user/{userId}",
		IsAuthorized(DeleteUser),
	},

	Route{
		"GetUserById",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/user/{userId}",
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
		"/ditrit/Gandalf/1.0.0/user/{userId}",
		IsAuthorized(UpdateUser),
	},
}
