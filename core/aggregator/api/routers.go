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
		"CreateDomain",
		strings.ToUpper("Post"),
		"/ditrit/Gandalf/1.0.0/domain",
		CreateDomain,
	},

	Route{
		"DeleteDomain",
		strings.ToUpper("Delete"),
		"/ditrit/Gandalf/1.0.0/domain/{domainId}",
		DeleteDomain,
	},

	Route{
		"GetDomainById",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/domain/{domainId}",
		GetDomainById,
	},

	Route{
		"GetDomainByName",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/domain/{domainName}",
		GetDomainByName,
	},

	Route{
		"ListDomain",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/domain",
		ListDomain,
	},

	Route{
		"UpdateDomain",
		strings.ToUpper("Put"),
		"/ditrit/Gandalf/1.0.0/domain/{domainId}",
		UpdateDomain,
	},

	Route{
		"CreateEventType",
		strings.ToUpper("Post"),
		"/ditrit/Gandalf/1.0.0/eventType",
		CreateEventType,
	},

	Route{
		"DeleteEventType",
		strings.ToUpper("Delete"),
		"/ditrit/Gandalf/1.0.0/eventType/{eventTypeId}",
		DeleteEventType,
	},

	Route{
		"GetEventTypeById",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/eventType/{eventTypeId}",
		GetEventTypeById,
	},

	Route{
		"GetEventTypeByName",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/eventType/{eventTypeName}",
		GetEventTypeByName,
	},

	Route{
		"ListEventType",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/eventType",
		ListEventType,
	},

	Route{
		"UpdateEventType",
		strings.ToUpper("Put"),
		"/ditrit/Gandalf/1.0.0/eventType/{eventTypeId}",
		UpdateEventType,
	},

	Route{
		"CreateEventTypeToPoll",
		strings.ToUpper("Post"),
		"/ditrit/Gandalf/1.0.0/eventTypeToPoll",
		CreateEventTypeToPoll,
	},

	Route{
		"DeleteEventTypeToPoll",
		strings.ToUpper("Delete"),
		"/ditrit/Gandalf/1.0.0/eventTypeToPoll/{eventTypeToPollId}",
		DeleteEventTypeToPoll,
	},

	Route{
		"GetEventTypeToPollById",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/eventTypeToPoll/{eventTypeToPollId}",
		GetEventTypeToPollById,
	},

	Route{
		"ListEventTypeToPoll",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/eventTypeToPoll",
		ListEventTypeToPoll,
	},

	Route{
		"UpdateEventTypeToPoll",
		strings.ToUpper("Put"),
		"/ditrit/Gandalf/1.0.0/eventTypeToPoll/{eventTypeToPollId}",
		UpdateEventTypeToPoll,
	},

	Route{
		"GetLogicalComponentByName",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/logicalcomponent/{logicalComponentName}",
		GetLogicalComponentByName,
	},

	Route{
		"UploadLogicalComponentByTenantAndType",
		strings.ToUpper("Post"),
		"/ditrit/Gandalf/1.0.0/logicalcomponent/upload/{tenantName}/{typeName}",
		UploadLogicalComponentByTenantAndType,
	},

	Route{
		"CreateResource",
		strings.ToUpper("Post"),
		"/ditrit/Gandalf/1.0.0/resource",
		CreateResource,
	},

	Route{
		"DeleteResource",
		strings.ToUpper("Delete"),
		"/ditrit/Gandalf/1.0.0/resource/{resourceId}",
		DeleteResource,
	},

	Route{
		"GetResourceById",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/resource/{resourceId}",
		GetResourceById,
	},

	Route{
		"GetResourceByName",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/resource/{resourceName}",
		GetResourceByName,
	},

	Route{
		"ListResource",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/resource",
		ListResource,
	},

	Route{
		"UpdateResource",
		strings.ToUpper("Put"),
		"/ditrit/Gandalf/1.0.0/resource/{resourceId}",
		UpdateResource,
	},

	Route{
		"CreateResourceType",
		strings.ToUpper("Post"),
		"/ditrit/Gandalf/1.0.0/resourceType",
		CreateResourceType,
	},

	Route{
		"DeleteResourceType",
		strings.ToUpper("Delete"),
		"/ditrit/Gandalf/1.0.0/resourceType/{resourceTypeId}",
		DeleteResourceType,
	},

	Route{
		"GetResourceTypeById",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/resourceType/{resourceTypeId}",
		GetResourceTypeById,
	},

	Route{
		"GetResourceTypeByName",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/resourceType/{resourceTypeName}",
		GetResourceTypeByName,
	},

	Route{
		"ListResourceType",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/resourceType",
		ListResourceType,
	},

	Route{
		"UpdateResourceType",
		strings.ToUpper("Put"),
		"/ditrit/Gandalf/1.0.0/resourceType/{resourceTypeId}",
		UpdateResourceType,
	},

	Route{
		"CreateRole",
		strings.ToUpper("Post"),
		"/ditrit/Gandalf/1.0.0/role",
		CreateRole,
	},

	Route{
		"DeleteRole",
		strings.ToUpper("Delete"),
		"/ditrit/Gandalf/1.0.0/role/{roleId}",
		DeleteRole,
	},

	Route{
		"GetRoleById",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/role/{roleId}",
		GetRoleById,
	},

	Route{
		"ListRole",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/role",
		ListRole,
	},

	Route{
		"UpdateRole",
		strings.ToUpper("Put"),
		"/ditrit/Gandalf/1.0.0/role/{roleId}",
		UpdateRole,
	},

	Route{
		"CreateSecretAssignement",
		strings.ToUpper("Post"),
		"/ditrit/Gandalf/1.0.0/secretAssignement",
		CreateSecretAssignement,
	},

	Route{
		"ListSecretAssignement",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/secretAssignement",
		ListSecretAssignement,
	},

	Route{
		"CreateTenant",
		strings.ToUpper("Post"),
		"/ditrit/Gandalf/1.0.0/tenant",
		CreateTenant,
	},

	Route{
		"DeleteTenant",
		strings.ToUpper("Delete"),
		"/ditrit/Gandalf/1.0.0/tenant/{tenantId}",
		DeleteTenant,
	},

	Route{
		"GetTenantById",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/tenant/{tenantId}",
		GetTenantById,
	},

	Route{
		"ListTenant",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/tenant",
		ListTenant,
	},

	Route{
		"UpdateTenant",
		strings.ToUpper("Put"),
		"/ditrit/Gandalf/1.0.0/tenant/{tenantId}",
		UpdateTenant,
	},

	Route{
		"CreateUser",
		strings.ToUpper("Post"),
		"/ditrit/Gandalf/1.0.0/user",
		CreateUser,
	},

	Route{
		"DeleteUser",
		strings.ToUpper("Delete"),
		"/ditrit/Gandalf/1.0.0/user/{userId}",
		DeleteUser,
	},

	Route{
		"GetUserById",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/user/{userId}",
		GetUserById,
	},

	Route{
		"GetUserByName",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/user/{userName}",
		GetUserByName,
	},

	Route{
		"ListUser",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/user",
		ListUser,
	},

	Route{
		"LoginUser",
		strings.ToUpper("Post"),
		"/ditrit/Gandalf/1.0.0/user/login",
		LoginUser,
	},

	Route{
		"LogoutUser",
		strings.ToUpper("Get"),
		"/ditrit/Gandalf/1.0.0/user/logout",
		LogoutUser,
	},

	Route{
		"UpdateUser",
		strings.ToUpper("Put"),
		"/ditrit/Gandalf/1.0.0/user/{userId}",
		UpdateUser,
	},
}
