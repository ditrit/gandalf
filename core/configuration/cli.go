/*
Copyright © 2020 DitRit community <contact@ditrit.io>
This file is part of Gandalf
*/

// Package configuration manages commands and configuration
package configuration

import (
	"fmt"

	"github.com/ditrit/gandalf/verdeter"
	"github.com/google/uuid"

	"github.com/ditrit/gandalf/core/cli"
	cmodels "github.com/ditrit/gandalf/core/configuration/models"
	"github.com/ditrit/gandalf/core/models"

	"github.com/spf13/viper"
)

// cliCmd represents the cli command
var cliCfg = verdeter.NewConfigCmd("cli", "Launch gandalf in 'cli' mode.", `Gandalf is launched as CLI (Command Line Interface) to interact with a Gandalf system.`, nil)

var cliCreate = verdeter.NewConfigCmd("create", "create user|tenant|role|domain|resource|eventTypeToPoll|resourceType|eventType", "create command allows the creation of Gandalf objects (users, tenants, roles and domains).", nil)
var cliList = verdeter.NewConfigCmd("list", "list users|tenants|roles|domain|resource|eventTypeToPoll|resourceType|eventType", "list command allows to list Gandalf objects (users, tenants, roles and domains).", nil)
var cliUpdate = verdeter.NewConfigCmd("update", "update user|tenant|role|domain|resource|eventTypeToPoll|resourceType|eventType", "update command allows update of Gandalf objects (users, tenants, roles and domains).", nil)
var cliDelete = verdeter.NewConfigCmd("delete", "delete user|tenant|role|domain|resource|eventTypeToPoll|resourceType|eventType", "update command allows deleting of Gandalf objects (users, tenants, roles and domains).", nil)
var cliLogin = verdeter.NewConfigCmd("login", "log in as a user into Gandalf", "login command allows user to authenticate using its credentials.", runLogin)

var cliCreateAuthorization = verdeter.NewConfigCmd("authorization", "create authorization <UserID> <RoleID> <DomainID>", "create authorization command allows the creation of a new authorization", runCreateAuthorization)
var cliListAuthorizations = verdeter.NewConfigCmd("authorization", "list authorizations", "list authorizations command allows to list Gandalf authorizations.", runListAuthorizations)
var cliUpdateAuthorization = verdeter.NewConfigCmd("authorization", "update authorization <AuthorizationID> [options]", "update authorization command allows to update a Gandalf authorization.", runUpdateAuthorization)
var cliDeleteAuthorization = verdeter.NewConfigCmd("authorization", "delete authorization <AuthorizationID>", "delete authorization command allows to delete a Gandalf authorization.", runDeleteAuthorization)

var cliCreateDomain = verdeter.NewConfigCmd("domain", "create domain <DomainID>", "create domain command allows the creation of a new domain (in the form <[name.]*name>)", runCreateDomain)
var cliListDomains = verdeter.NewConfigCmd("domain", "list domains ", "list domains command allows to list Gandalf domains.", runListDomains)
var cliUpdateDomain = verdeter.NewConfigCmd("domain", "update domain <DomainID> [options]", "update domain command allows to update a Gandalf domain.", runUpdateDomain)
var cliDeleteDomain = verdeter.NewConfigCmd("domain", "delete domain <DomainID>", "delete domain command allows to delete a Gandalf domain.", runDeleteDomain)

var cliCreateEnvironment = verdeter.NewConfigCmd("environment", "create environment <username> <email> <password>", "create environment command allows the creation of a new environment", runCreateEnvironment)
var cliListEnvironments = verdeter.NewConfigCmd("environment", "list environments", "list environments command allows to list Gandalf environments.", runListEnvironments)
var cliUpdateEnvironment = verdeter.NewConfigCmd("environment", "update environment <username> [options]", "update environment command allows to update a Gandalf environment.", runUpdateEnvironment)
var cliDeleteEnvironment = verdeter.NewConfigCmd("environment", "delete environment <username>", "delete environment command allows to delete a Gandalf environment.", runDeleteEnvironment)

var cliCreateEnvironmentType = verdeter.NewConfigCmd("environmentType", "create environmentType <username> <email> <password>", "create environmentType command allows the creation of a new environmentType", runCreateEnvironmentType)
var cliListEnvironmentTypes = verdeter.NewConfigCmd("environmentType", "list environmentTypes", "list users command allows to list Gandalf environmentTypes.", runListEnvironmentTypes)
var cliUpdateEnvironmentType = verdeter.NewConfigCmd("environmentType", "update environmentType <username> [options]", "update environmentType command allows to update a Gandalf environmentType.", runUpdateEnvironmentType)
var cliDeleteEnvironmentType = verdeter.NewConfigCmd("environmentType", "delete environmentType <username>", "delete environmentType command allows to delete a Gandalf environmentType.", runDeleteEnvironmentType)

var cliCreateEventType = verdeter.NewConfigCmd("eventtype", "create eventtype <eventtypename>", "create eventtype command allows the creation of a new resource (in the form <[name.]*name>)", runCreateEventType)
var cliListEventTypes = verdeter.NewConfigCmd("eventtype", "list eventtypes ", "list eventtypes command allows to list Gandalf eventtype.", runListEventTypes)
var cliUpdateEventType = verdeter.NewConfigCmd("eventtype", "update eventtype <eventtypename> [options]", "update resource command allows to update a Gandalf eventtype.", runUpdateEventType)
var cliDeleteEventType = verdeter.NewConfigCmd("eventtype", "delete eventtype <eventtypename>", "delete eventtype command allows to delete a Gandalf eventtype.", runDeleteEventType)

var cliCreateEventTypeToPoll = verdeter.NewConfigCmd("eventtypetopoll", "create eventtypetopoll <eventtypetopollname>", "create eventtypetopoll command allows the creation of a new resource (in the form <[name.]*name>)", runCreateEventTypeToPoll)
var cliListEventTypeToPolls = verdeter.NewConfigCmd("eventtypetopoll", "list eventtypetopolls ", "list eventtypetopolls command allows to list Gandalf eventtypetopolls.", runListEventTypeToPolls)
var cliUpdateEventTypeToPoll = verdeter.NewConfigCmd("eventtypetopoll", "update eventtypetopoll <eventtypetopollname> [options]", "update resource command allows to update a Gandalf eventtypetopoll.", runUpdateEventTypeToPoll)
var cliDeleteEventTypeToPoll = verdeter.NewConfigCmd("eventtypetopoll", "delete eventtypetopoll <eventtypetopollname>", "delete eventtypetopoll command allows to delete a Gandalf eventtypetopoll.", runDeleteEventTypeToPoll)

var cliCreateLibrary = verdeter.NewConfigCmd("library", "create library <username> <email> <password>", "create library command allows the creation of a new library", runCreateLibrary)
var cliListLibraries = verdeter.NewConfigCmd("library", "list libraries", "list libraries command allows to list Gandalf libraries.", runListLibraries)
var cliUpdateLibrary = verdeter.NewConfigCmd("library", "update library <username> [options]", "update library command allows to update a Gandalf library.", runUpdateLibrary)
var cliDeleteLibrary = verdeter.NewConfigCmd("library", "delete library <username>", "delete library command allows to delete a Gandalf library.", runDeleteLibrary)

//Logical component ?

var cliCreateProduct = verdeter.NewConfigCmd("product", "create product <username> <email> <password>", "create product command allows the creation of a new product", runCreateProduct)
var cliListProducts = verdeter.NewConfigCmd("product", "list products", "list products command allows to list Gandalf products.", runListProducts)
var cliUpdateProduct = verdeter.NewConfigCmd("product", "update product <username> [options]", "update product command allows to update a Gandalf product.", runUpdateProduct)
var cliDeleteProduct = verdeter.NewConfigCmd("product", "delete product <username>", "delete product command allows to delete a Gandalf product.", runDeleteProduct)

var cliCreateResource = verdeter.NewConfigCmd("resource", "create resource <resourcename>", "create resource command allows the creation of a new resource (in the form <[name.]*name>)", runCreateResource)
var cliListResources = verdeter.NewConfigCmd("resource", "list resources ", "list resource command allows to list Gandalf resources.", runListResources)
var cliUpdateResource = verdeter.NewConfigCmd("resource", "update resource <resourcename> [options]", "update resource command allows to update a Gandalf resource.", runUpdateResource)
var cliDeleteResource = verdeter.NewConfigCmd("resource", "delete resource <resourcename>", "delete resource command allows to delete a Gandalf resource.", runDeleteResource)

var cliCreateResourceType = verdeter.NewConfigCmd("resourcetype", "create resourcetype <resourcetypename>", "create resourcetype command allows the creation of a new resource (in the form <[name.]*name>)", runCreateResourceType)
var cliListResourceTypes = verdeter.NewConfigCmd("resourcetype", "list resourcetype ", "list resourcetype command allows to list Gandalf resourcetypes.", runListResourceTypes)
var cliUpdateResourceType = verdeter.NewConfigCmd("resourcetype", "update resourcetype <resourcetypename> [options]", "update resource command allows to update a Gandalf resourcetype.", runUpdateResourceType)
var cliDeleteResourceType = verdeter.NewConfigCmd("resourcetype", "delete resourcetype <resourcetypename>", "delete resourcetype command allows to delete a Gandalf resourcetype.", runDeleteResourceType)

var cliCreateRole = verdeter.NewConfigCmd("role", "create role <rolename> ", "create role command allows the creation of a new role", runCreateRole)
var cliListRoles = verdeter.NewConfigCmd("role", "list roles <rolename> ", "list roles command allows to list Gandalf roles.", runListRoles)
var cliUpdateRole = verdeter.NewConfigCmd("role", "update role <rolename> [options]", "update role command allows to update a Gandalf role.", runUpdateRole)
var cliDeleteRole = verdeter.NewConfigCmd("role", "delete role <rolename>", "delete role command allows to delete a Gandalf role.", runDeleteRole)

var cliCreateSecret = verdeter.NewConfigCmd("secret", "create secret", "declare  name command allows to declare the name of a new connector.", runCreateSecret)
var cliListSecret = verdeter.NewConfigCmd("secret", "list secret", "declare  member command allows to declare a new member for an existing connector.", runListSecret)

var cliCreateTag = verdeter.NewConfigCmd("tag", "create tag <domainname>", "create tag command allows the creation of a new tag (in the form <[name.]*name>)", runCreateTag)
var cliListTags = verdeter.NewConfigCmd("tag", "list tags ", "list tags command allows to list Gandalf domains.", runListTags)
var cliUpdateTag = verdeter.NewConfigCmd("tag", "update tag <domainname> [options]", "update tag command allows to update a Gandalf tag.", runUpdateTag)
var cliDeleteTag = verdeter.NewConfigCmd("tag", "delete tag <domainname>", "delete tag command allows to delete a Gandalf tag.", runDeleteTag)

var cliCreateTenant = verdeter.NewConfigCmd("tenant", "create tenant <tenantname>", "create tenant command allows the creation of a new tenant", runCreateTenant)
var cliListTenants = verdeter.NewConfigCmd("tenant", "list tenants <tenantname>", "list tenants command allows to list Gandalf tenants.", runListTenants)
var cliUpdateTenant = verdeter.NewConfigCmd("tenant", "update tenant <tenantname> [options]", "update tenant command allows to update a Gandalf tenant.", runUpdateTenant)
var cliDeleteTenant = verdeter.NewConfigCmd("tenant", "delete tenant <tenantname>", "delete tenant command allows to delete a Gandalf tenant.", runDeleteTenant)

var cliCreateUser = verdeter.NewConfigCmd("user", "create user <username> <email> <password>", "create user command allows the creation of a new user", runCreateUser)
var cliListUsers = verdeter.NewConfigCmd("user", "list users", "list users command allows to list Gandalf users.", runListUsers)
var cliUpdateUser = verdeter.NewConfigCmd("user", "update user <username> [options]", "update user command allows to update a Gandalf user.", runUpdateUser)
var cliDeleteUser = verdeter.NewConfigCmd("user", "delete user <username>", "delete user command allows to delete a Gandalf user.", runDeleteUser)

//

func init() {

	rootCfg.AddConfig(cliCfg)

	cliCfg.GKey("endpoint", verdeter.IsStr, "e", "Gandalf endpoint")
	cliCfg.SetRequired("endpoint")
	cliCfg.GKey("token", verdeter.IsStr, "t", "Gandalf auth token")
	//cliCfg.SetRequired("token")

	cliCfg.AddConfig(cliCreate)
	cliCfg.AddConfig(cliList)
	cliCfg.AddConfig(cliUpdate)
	cliCfg.AddConfig(cliDelete)
	cliCfg.AddConfig(cliLogin)

	cliCreate.AddConfig(cliCreateAuthorization)
	cliList.AddConfig(cliListAuthorizations)
	cliUpdate.AddConfig(cliUpdateAuthorization)
	cliDelete.AddConfig(cliDeleteAuthorization)

	cliCreate.AddConfig(cliCreateDomain)
	cliList.AddConfig(cliListDomains)
	cliUpdate.AddConfig(cliUpdateDomain)
	cliDelete.AddConfig(cliDeleteDomain)

	cliCreate.AddConfig(cliCreateEnvironment)
	cliList.AddConfig(cliListEnvironments)
	cliUpdate.AddConfig(cliUpdateEnvironment)
	cliDelete.AddConfig(cliDeleteEnvironment)

	cliCreate.AddConfig(cliCreateEnvironmentType)
	cliList.AddConfig(cliListEnvironmentTypes)
	cliUpdate.AddConfig(cliUpdateEnvironmentType)
	cliDelete.AddConfig(cliDeleteEnvironmentType)

	cliCreate.AddConfig(cliCreateEventType)
	cliList.AddConfig(cliListEventTypes)
	cliUpdate.AddConfig(cliUpdateEventType)
	cliDelete.AddConfig(cliDeleteEventType)

	cliCreate.AddConfig(cliCreateEventTypeToPoll)
	cliList.AddConfig(cliListEventTypeToPolls)
	cliUpdate.AddConfig(cliUpdateEventTypeToPoll)
	cliDelete.AddConfig(cliDeleteEventTypeToPoll)

	cliCreate.AddConfig(cliCreateLibrary)
	cliList.AddConfig(cliListLibraries)
	cliUpdate.AddConfig(cliUpdateLibrary)
	cliDelete.AddConfig(cliDeleteLibrary)

	cliCreate.AddConfig(cliCreateProduct)
	cliList.AddConfig(cliListProducts)
	cliUpdate.AddConfig(cliUpdateProduct)
	cliDelete.AddConfig(cliDeleteProduct)

	cliCreate.AddConfig(cliCreateResource)
	cliList.AddConfig(cliListResources)
	cliUpdate.AddConfig(cliUpdateResource)
	cliDelete.AddConfig(cliDeleteResource)

	cliCreate.AddConfig(cliCreateResourceType)
	cliList.AddConfig(cliListResourceTypes)
	cliUpdate.AddConfig(cliUpdateResourceType)
	cliDelete.AddConfig(cliDeleteResourceType)

	cliCreate.AddConfig(cliCreateRole)
	cliList.AddConfig(cliListRoles)
	cliUpdate.AddConfig(cliUpdateRole)
	cliDelete.AddConfig(cliDeleteRole)

	cliCreate.AddConfig(cliCreateSecret)
	cliList.AddConfig(cliListSecret)

	cliCreate.AddConfig(cliCreateTag)
	cliList.AddConfig(cliListTags)
	cliUpdate.AddConfig(cliUpdateTag)
	cliDelete.AddConfig(cliDeleteTag)

	cliCreate.AddConfig(cliCreateTenant)
	cliList.AddConfig(cliListTenants)
	cliUpdate.AddConfig(cliUpdateTenant)
	cliDelete.AddConfig(cliDeleteTenant)

	cliCreate.AddConfig(cliCreateUser)
	cliList.AddConfig(cliListUsers)
	cliUpdate.AddConfig(cliUpdateUser)
	cliDelete.AddConfig(cliDeleteUser)

	cliLogin.SetNbArgs(2)

	cliCreateAuthorization.SetNbArgs(3)
	cliListAuthorizations.SetNbArgs(0)
	cliUpdateAuthorization.SetNbArgs(1)
	cliUpdateAuthorization.LKey("userID", verdeter.IsStr, "u", "name of the user")
	cliUpdateAuthorization.LKey("roleID", verdeter.IsStr, "r", "mail of the user")
	cliUpdateAuthorization.LKey("domainID", verdeter.IsStr, "d", "firstname of the user")
	cliDeleteAuthorization.SetNbArgs(1)

	cliCreateDomain.SetNbArgs(3)
	cliListDomains.SetNbArgs(0)
	cliUpdateDomain.SetNbArgs(1)
	cliDeleteDomain.SetNbArgs(1)

	cliCreateEnvironment.SetNbArgs(6)
	cliListEnvironments.SetNbArgs(0)
	cliUpdateEnvironment.SetNbArgs(1)
	cliUpdateEnvironment.LKey("name", verdeter.IsStr, "n", "name of the user")
	cliUpdateEnvironment.LKey("environmentTypeID", verdeter.IsStr, "e", "mail of the user")
	cliUpdateEnvironment.LKey("shortDescription", verdeter.IsStr, "s", "firstname of the user")
	cliUpdateEnvironment.LKey("description", verdeter.IsStr, "u", "firstname of the user")
	cliUpdateEnvironment.LKey("logo", verdeter.IsStr, "l", "firstname of the user")
	cliUpdateEnvironment.LKey("domainID", verdeter.IsStr, "d", "firstname of the user")
	cliDeleteEnvironment.SetNbArgs(1)

	cliCreateEnvironmentType.SetNbArgs(4)
	cliListEnvironmentTypes.SetNbArgs(0)
	cliUpdateEnvironmentType.SetNbArgs(1)
	cliUpdateEnvironmentType.LKey("name", verdeter.IsStr, "n", "name of the user")
	cliUpdateEnvironmentType.LKey("shortDescription", verdeter.IsStr, "s", "firstname of the user")
	cliUpdateEnvironmentType.LKey("description", verdeter.IsStr, "u", "firstname of the user")
	cliUpdateEnvironmentType.LKey("logo", verdeter.IsStr, "l", "firstname of the user")
	cliDeleteEnvironmentType.SetNbArgs(1)

	cliCreateEventType.SetNbArgs(4)
	cliListEventTypes.SetNbArgs(0)
	cliUpdateEventType.SetNbArgs(1)
	cliUpdateEventType.LKey("name", verdeter.IsStr, "n", "name of the user")
	cliUpdateEventType.LKey("schema", verdeter.IsStr, "s", "firstname of the user")
	cliUpdateEventType.LKey("pivotID", verdeter.IsStr, "p", "firstname of the user")
	cliUpdateEventType.LKey("productConnectorID", verdeter.IsStr, "c", "firstname of the user")
	cliDeleteEventType.SetNbArgs(1)

	cliCreateLibrary.SetNbArgs(4)
	cliListLibraries.SetNbArgs(0)
	cliUpdateLibrary.SetNbArgs(1)
	cliUpdateLibrary.LKey("name", verdeter.IsStr, "n", "name of the user")
	cliUpdateLibrary.LKey("shortDescription", verdeter.IsStr, "s", "firstname of the user")
	cliUpdateLibrary.LKey("description", verdeter.IsStr, "u", "firstname of the user")
	cliUpdateLibrary.LKey("logo", verdeter.IsStr, "l", "firstname of the user")
	cliDeleteLibrary.SetNbArgs(1)

	cliCreateProduct.SetNbArgs(6)
	cliListProducts.SetNbArgs(0)
	cliUpdateProduct.SetNbArgs(1)
	cliUpdateProduct.LKey("name", verdeter.IsStr, "n", "name of the user")
	cliUpdateProduct.LKey("shortDescription", verdeter.IsStr, "s", "firstname of the user")
	cliUpdateProduct.LKey("description", verdeter.IsStr, "u", "firstname of the user")
	cliUpdateProduct.LKey("logo", verdeter.IsStr, "l", "firstname of the user")
	cliUpdateProduct.LKey("repositoryURL", verdeter.IsStr, "r", "firstname of the user")
	cliUpdateProduct.LKey("domainID", verdeter.IsStr, "d", "firstname of the user")
	cliDeleteProduct.SetNbArgs(1)

	cliCreateResource.SetNbArgs(4)
	cliListResources.SetNbArgs(0)
	cliUpdateResource.SetNbArgs(1)
	cliUpdateResource.LKey("name", verdeter.IsStr, "n", "name of the user")
	cliUpdateResource.LKey("logicalComponentID", verdeter.IsStr, "l", "firstname of the user")
	cliUpdateResource.LKey("domainID", verdeter.IsStr, "d", "firstname of the user")
	cliUpdateResource.LKey("resourceTypeID", verdeter.IsStr, "r", "firstname of the user")
	cliDeleteResource.SetNbArgs(1)

	cliCreateResourceType.SetNbArgs(3)
	cliListResourceTypes.SetNbArgs(0)
	cliUpdateResourceType.SetNbArgs(1)
	cliUpdateResourceType.LKey("name", verdeter.IsStr, "n", "name of the user")
	cliUpdateResourceType.LKey("pivotID", verdeter.IsStr, "l", "firstname of the user")
	cliUpdateResourceType.LKey("productConnectorID", verdeter.IsStr, "d", "firstname of the user")
	cliDeleteResourceType.SetNbArgs(1)

	cliCreateRole.SetNbArgs(4)
	cliListRoles.SetNbArgs(0)
	cliUpdateRole.SetNbArgs(1)
	cliUpdateRole.LKey("name", verdeter.IsStr, "n", "name of the user")
	cliUpdateRole.LKey("shortDescription", verdeter.IsStr, "s", "firstname of the user")
	cliUpdateRole.LKey("description", verdeter.IsStr, "u", "firstname of the user")
	cliUpdateRole.LKey("logo", verdeter.IsStr, "l", "firstname of the user")
	cliDeleteRole.SetNbArgs(1)

	cliCreateSecret.SetNbArgs(0)
	cliListSecret.SetNbArgs(0)

	cliCreateTag.SetNbArgs(5)
	cliListTags.SetNbArgs(0)
	cliUpdateTag.SetNbArgs(1)
	cliUpdateTag.LKey("name", verdeter.IsStr, "n", "name of the user")
	cliUpdateTag.LKey("shortDescription", verdeter.IsStr, "s", "firstname of the user")
	cliUpdateTag.LKey("description", verdeter.IsStr, "u", "firstname of the user")
	cliUpdateTag.LKey("logo", verdeter.IsStr, "l", "firstname of the user")
	cliUpdateTag.LKey("parentID", verdeter.IsStr, "p", "firstname of the user")
	cliDeleteTag.SetNbArgs(1)

	cliCreateTenant.SetNbArgs(4)
	cliListTenants.SetNbArgs(0)
	cliUpdateTenant.SetNbArgs(1)
	cliUpdateTenant.LKey("name", verdeter.IsStr, "n", "name of the Tenant")
	cliUpdateTenant.LKey("shortDescription", verdeter.IsStr, "s", "firstname of the user")
	cliUpdateTenant.LKey("description", verdeter.IsStr, "u", "firstname of the user")
	cliUpdateTenant.LKey("logo", verdeter.IsStr, "l", "firstname of the user")
	cliDeleteTenant.SetNbArgs(1)

	cliCreateUser.SetNbArgs(5)
	cliListUsers.SetNbArgs(0)
	cliUpdateUser.SetNbArgs(1)
	cliUpdateUser.LKey("email", verdeter.IsStr, "m", "mail of the user")
	cliUpdateUser.LKey("firstName", verdeter.IsStr, "f", "firstname of the user")
	cliUpdateUser.LKey("lastName", verdeter.IsStr, "l", "secondname of the user")
	cliUpdateUser.LKey("companyID", verdeter.IsStr, "c", "companyid of the user")
	cliUpdateUser.LKey("password", verdeter.IsStr, "p", "password of the user")
	cliDeleteUser.SetNbArgs(1)

	// EVENT type to poll

}

func runLogin(cfg *verdeter.ConfigCmd, args []string) {
	mail := args[0]
	password := args[1]

	fmt.Printf("gandalf cli login called with mail=%s and password=%s\n", mail, password)
	configurationCli := cmodels.NewConfigurationCli()
	fmt.Println(configurationCli.GetEndpoint())
	cliClient := cli.NewClient(configurationCli.GetEndpoint())

	var user models.User
	user.Email = mail
	user.Password = password
	//user := models.NewUser(name, name, password)
	token, err := cliClient.UserService.Login(user)
	if err == nil {
		fmt.Println("Token: " + token)
	} else {
		fmt.Println(err)
	}
}

func runCreateAuthorization(cfg *verdeter.ConfigCmd, args []string) {
	userID, err := uuid.Parse(args[0])
	if err == nil {
		roleID, err := uuid.Parse(args[1])
		if err == nil {
			domainID, err := uuid.Parse(args[2])
			if err == nil {

				fmt.Printf("gandalf cli create authorization called with userID=%s, roleID=%s, domainID=%s\n", userID, roleID, domainID)
				configurationCli := cmodels.NewConfigurationCli()
				cliClient := cli.NewClient(configurationCli.GetEndpoint())

				authorization := models.Authorization{UserID: userID, RoleID: roleID, DomainID: domainID}
				err := cliClient.AuthorizationService.Create(configurationCli.GetToken(), authorization)
				if err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Println(err)
			}
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}
}

func runListAuthorizations(cfg *verdeter.ConfigCmd, args []string) {
	fmt.Printf("gandalf cli list authorizations\n")
	configurationCli := cmodels.NewConfigurationCli()
	cliClient := cli.NewClient(configurationCli.GetEndpoint())

	authorizations, err := cliClient.AuthorizationService.List(configurationCli.GetToken())
	if err == nil {
		for _, authorization := range authorizations {
			fmt.Println(authorization)
		}
	} else {
		fmt.Println(err)
	}

}

func runUpdateAuthorization(cfg *verdeter.ConfigCmd, args []string) {
	authorizationID, err := uuid.Parse(args[0])
	if err == nil {
		userID, err := uuid.Parse(viper.GetString("userID"))
		if err == nil {
			roleID, err := uuid.Parse(viper.GetViper().GetString("roleID"))
			if err == nil {
				domainID, err := uuid.Parse(viper.GetViper().GetString("domainID"))
				if err == nil {

					fmt.Printf("gandalf cli update authorization called with userID=%s, roleID=%s, domainID=%s,\n", userID, roleID, domainID)
					configurationCli := cmodels.NewConfigurationCli()
					cliClient := cli.NewClient(configurationCli.GetEndpoint())

					oldAuthorization, err := cliClient.AuthorizationService.Read(configurationCli.GetToken(), authorizationID)
					if err == nil {
						oldAuthorization.UserID = userID
						oldAuthorization.RoleID = roleID
						oldAuthorization.DomainID = domainID
						err = cliClient.AuthorizationService.Update(configurationCli.GetToken(), authorizationID, *oldAuthorization)
						if err != nil {
							fmt.Println(err)
						}
					} else {
						fmt.Println(err)
					}
				} else {
					fmt.Println(err)
				}
			} else {
				fmt.Println(err)
			}
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}
}

func runDeleteAuthorization(cfg *verdeter.ConfigCmd, args []string) {
	authorizationID, err := uuid.Parse(args[0])
	if err == nil {
		fmt.Printf("gandalf cli delete authorization called with authorizationID=%s\n", authorizationID)
		configurationCli := cmodels.NewConfigurationCli()
		cliClient := cli.NewClient(configurationCli.GetEndpoint())

		err = cliClient.AuthorizationService.Delete(configurationCli.GetToken(), authorizationID)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}

}

///

func runCreateEnvironment(cfg *verdeter.ConfigCmd, args []string) {
	name := args[0]
	environmentTypeID, err := uuid.Parse(args[1])
	if err == nil {
		shortDescription := args[2]
		description := args[3]
		logo := args[4]
		domainID, err := uuid.Parse(args[5])
		if err == nil {

			fmt.Printf("gandalf cli create environment called with name=%s, environmentTypeID=%s, shortDescription=%s, description=%s, logo=%s, domainID=%s\n", name, environmentTypeID, shortDescription, description, logo, domainID)
			configurationCli := cmodels.NewConfigurationCli()
			cliClient := cli.NewClient(configurationCli.GetEndpoint())

			environment := models.Environment{Name: name, EnvironmentTypeID: environmentTypeID, ShortDescription: shortDescription, Description: description, Logo: logo, DomainID: domainID}
			err := cliClient.EnvironmentService.Create(configurationCli.GetToken(), environment)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}

}

func runListEnvironments(cfg *verdeter.ConfigCmd, args []string) {
	fmt.Printf("gandalf cli list environments\n")
	configurationCli := cmodels.NewConfigurationCli()
	cliClient := cli.NewClient(configurationCli.GetEndpoint())

	environments, err := cliClient.EnvironmentService.List(configurationCli.GetToken())
	if err == nil {
		for _, environment := range environments {
			fmt.Println(environment)
		}
	} else {
		fmt.Println(err)
	}

}

func runUpdateEnvironment(cfg *verdeter.ConfigCmd, args []string) {
	environmentID, err := uuid.Parse(args[0])
	if err == nil {
		name := viper.GetViper().GetString("name")
		environmentTypeID, err := uuid.Parse(viper.GetViper().GetString("environmentTypeID"))
		if err == nil {
			shortDescription := viper.GetViper().GetString("shortDescription")
			description := viper.GetViper().GetString("description")
			logo := viper.GetViper().GetString("logo")

			domainID, err := uuid.Parse(viper.GetViper().GetString("domainID"))
			if err == nil {
				fmt.Printf("gandalf cli update environment called with name=%s, environmentTypeID=%s, shortDescription=%s, description=%s, logo=%s, domainID=%s\n", name, environmentTypeID, shortDescription, description, logo, domainID)
				configurationCli := cmodels.NewConfigurationCli()
				cliClient := cli.NewClient(configurationCli.GetEndpoint())

				oldEnvironment, err := cliClient.EnvironmentService.Read(configurationCli.GetToken(), environmentID)
				if err == nil {
					oldEnvironment.Name = name
					oldEnvironment.EnvironmentTypeID = environmentTypeID
					oldEnvironment.ShortDescription = shortDescription
					oldEnvironment.Description = description
					oldEnvironment.Logo = logo
					oldEnvironment.DomainID = domainID
					err = cliClient.EnvironmentService.Update(configurationCli.GetToken(), environmentID, *oldEnvironment)
					if err != nil {
						fmt.Println(err)
					}
				} else {
					fmt.Println(err)
				}
			} else {
				fmt.Println(err)
			}
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}

}

func runDeleteEnvironment(cfg *verdeter.ConfigCmd, args []string) {
	environmentID, err := uuid.Parse(args[0])
	fmt.Printf("gandalf cli delete environment called with environmentID=%s\n", environmentID)
	configurationCli := cmodels.NewConfigurationCli()
	cliClient := cli.NewClient(configurationCli.GetEndpoint())

	err = cliClient.EnvironmentService.Delete(configurationCli.GetToken(), environmentID)
	if err != nil {
		fmt.Println(err)
	}

}

///

func runCreateProduct(cfg *verdeter.ConfigCmd, args []string) {
	name := args[0]
	shortDescription := args[1]
	description := args[2]
	logo := args[3]
	repositoryURL := args[4]
	domainID, err := uuid.Parse(args[5])
	if err == nil {

		fmt.Printf("gandalf cli create product called with name=%s, shortDescription=%s, description=%s, logo=%s, repositoryURL=%s, domainID=%s\n", name, shortDescription, description, logo, repositoryURL, domainID)
		configurationCli := cmodels.NewConfigurationCli()
		cliClient := cli.NewClient(configurationCli.GetEndpoint())

		product := models.Product{Name: name, ShortDescription: shortDescription, Description: description, Logo: logo, RepositoryURL: repositoryURL, DomainID: domainID}
		err := cliClient.ProductService.Create(configurationCli.GetToken(), product)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}
}

func runListProducts(cfg *verdeter.ConfigCmd, args []string) {
	fmt.Printf("gandalf cli list products\n")
	configurationCli := cmodels.NewConfigurationCli()
	cliClient := cli.NewClient(configurationCli.GetEndpoint())

	products, err := cliClient.ProductService.List(configurationCli.GetToken())
	if err == nil {
		for _, product := range products {
			fmt.Println(product)
		}
	} else {
		fmt.Println(err)
	}

}

func runUpdateProduct(cfg *verdeter.ConfigCmd, args []string) {
	productID, err := uuid.Parse(args[0])
	if err == nil {
		name := viper.GetViper().GetString("name")
		shortDescription := viper.GetViper().GetString("shortDescription")
		description := viper.GetViper().GetString("description")
		logo := viper.GetViper().GetString("logo")
		repositoryURL := viper.GetViper().GetString("repositoryURL")
		domainID, err := uuid.Parse(viper.GetViper().GetString("domainID"))
		if err == nil {
			fmt.Printf("gandalf cli update user called with name=%s, shortDescription=%s, description=%s, logo=%s, repositoryURL=%s, domainID=%s\n", name, shortDescription, description, logo, repositoryURL, domainID)
			configurationCli := cmodels.NewConfigurationCli()
			cliClient := cli.NewClient(configurationCli.GetEndpoint())

			oldProduct, err := cliClient.ProductService.Read(configurationCli.GetToken(), productID)
			if err == nil {
				oldProduct.Name = name
				oldProduct.ShortDescription = shortDescription
				oldProduct.Description = description
				oldProduct.Logo = logo
				oldProduct.RepositoryURL = repositoryURL
				oldProduct.DomainID = domainID
				err = cliClient.ProductService.Update(configurationCli.GetToken(), productID, *oldProduct)
				if err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Println(err)
			}
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}
}

func runDeleteProduct(cfg *verdeter.ConfigCmd, args []string) {
	productID, err := uuid.Parse(args[0])
	if err == nil {
		fmt.Printf("gandalf cli delete product called with productID=%s\n", productID)
		configurationCli := cmodels.NewConfigurationCli()
		cliClient := cli.NewClient(configurationCli.GetEndpoint())

		err = cliClient.ProductService.Delete(configurationCli.GetToken(), productID)
		if err != nil {
			fmt.Println(err)
		}

	} else {
		fmt.Println(err)
	}
}

///

func runCreateLibrary(cfg *verdeter.ConfigCmd, args []string) {
	name := args[0]
	shortDescription := args[1]
	description := args[2]
	logo := args[3]

	fmt.Printf("gandalf cli create library called with name=%s, shortDescription=%s, description=%s, logo=%s\n", name, shortDescription, description, logo)
	configurationCli := cmodels.NewConfigurationCli()
	cliClient := cli.NewClient(configurationCli.GetEndpoint())

	library := models.Library{Name: name, ShortDescription: shortDescription, Description: description, Logo: logo}
	err := cliClient.LibraryService.Create(configurationCli.GetToken(), library)
	if err != nil {
		fmt.Println(err)
	}

}

func runListLibraries(cfg *verdeter.ConfigCmd, args []string) {
	fmt.Printf("gandalf cli list libraries\n")
	configurationCli := cmodels.NewConfigurationCli()
	cliClient := cli.NewClient(configurationCli.GetEndpoint())

	libraries, err := cliClient.LibraryService.List(configurationCli.GetToken())
	fmt.Println(libraries)
	if err == nil {
		for _, library := range libraries {
			fmt.Println(library)
		}
	} else {
		fmt.Println(err)
	}

}

func runUpdateLibrary(cfg *verdeter.ConfigCmd, args []string) {
	libraryID, err := uuid.Parse(args[0])
	if err == nil {
		name := viper.GetViper().GetString("name")
		shortDescription := viper.GetViper().GetString("shortDescription")
		description := viper.GetViper().GetString("description")
		logo := viper.GetViper().GetString("logo")

		fmt.Printf("gandalf cli update library called with name=%s, shortDescription=%s, description=%s, logo=%s\n", name, shortDescription, description, logo)
		configurationCli := cmodels.NewConfigurationCli()
		cliClient := cli.NewClient(configurationCli.GetEndpoint())

		oldLibrary, err := cliClient.LibraryService.Read(configurationCli.GetToken(), libraryID)
		if err == nil {
			oldLibrary.Name = name
			oldLibrary.ShortDescription = shortDescription
			oldLibrary.Description = description
			oldLibrary.Logo = logo
			err = cliClient.LibraryService.Update(configurationCli.GetToken(), libraryID, *oldLibrary)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}
}

func runDeleteLibrary(cfg *verdeter.ConfigCmd, args []string) {
	libraryID, err := uuid.Parse(args[0])
	if err == nil {
		fmt.Printf("gandalf cli delete library called with libraryID=%s\n", libraryID)
		configurationCli := cmodels.NewConfigurationCli()
		cliClient := cli.NewClient(configurationCli.GetEndpoint())

		err = cliClient.LibraryService.Delete(configurationCli.GetToken(), libraryID)
		if err != nil {
			fmt.Println(err)
		}

	} else {
		fmt.Println(err)
	}

}

///

func runCreateEnvironmentType(cfg *verdeter.ConfigCmd, args []string) {
	name := args[0]
	shortDescription := args[1]
	description := args[2]
	logo := args[3]

	fmt.Printf("gandalf cli create environmentType called with name=%s, shortDescription=%s, description=%s, logo=%s\n", name, shortDescription, description, logo)
	configurationCli := cmodels.NewConfigurationCli()
	cliClient := cli.NewClient(configurationCli.GetEndpoint())

	environmentType := models.EnvironmentType{Name: name, ShortDescription: shortDescription, Description: description, Logo: logo}
	err := cliClient.EnvironmentTypeService.Create(configurationCli.GetToken(), environmentType)
	if err != nil {
		fmt.Println(err)
	}

}

func runListEnvironmentTypes(cfg *verdeter.ConfigCmd, args []string) {
	fmt.Printf("gandalf cli list environmentTypes\n")
	configurationCli := cmodels.NewConfigurationCli()
	cliClient := cli.NewClient(configurationCli.GetEndpoint())

	environments, err := cliClient.EnvironmentTypeService.List(configurationCli.GetToken())
	if err == nil {
		for _, environment := range environments {
			fmt.Println(environment)
		}
	} else {
		fmt.Println(err)
	}

}

func runUpdateEnvironmentType(cfg *verdeter.ConfigCmd, args []string) {
	environmentTypeID, err := uuid.Parse(args[0])
	if err == nil {
		name := viper.GetViper().GetString("name")
		shortDescription := viper.GetViper().GetString("shortDescription")
		description := viper.GetViper().GetString("description")
		logo := viper.GetViper().GetString("logo")

		fmt.Printf("gandalf cli update environmentType called with name=%s, shortDescription=%s, description=%s, logo=%s\n", name, shortDescription, description, logo)
		configurationCli := cmodels.NewConfigurationCli()
		cliClient := cli.NewClient(configurationCli.GetEndpoint())

		oldEnvironmentType, err := cliClient.EnvironmentTypeService.Read(configurationCli.GetToken(), environmentTypeID)
		if err == nil {
			oldEnvironmentType.Name = name
			oldEnvironmentType.ShortDescription = shortDescription
			oldEnvironmentType.Description = description
			oldEnvironmentType.Logo = logo
			err = cliClient.EnvironmentTypeService.Update(configurationCli.GetToken(), environmentTypeID, *oldEnvironmentType)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}

}

func runDeleteEnvironmentType(cfg *verdeter.ConfigCmd, args []string) {
	environmentTypeID, err := uuid.Parse(args[0])
	if err == nil {
		fmt.Printf("gandalf cli delete environmentType called with environmentTypeID=%s\n", environmentTypeID)
		configurationCli := cmodels.NewConfigurationCli()
		cliClient := cli.NewClient(configurationCli.GetEndpoint())

		err = cliClient.EnvironmentTypeService.Delete(configurationCli.GetToken(), environmentTypeID)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}

}

///

func runCreateTag(cfg *verdeter.ConfigCmd, args []string) {
	name := args[0]
	shortDescription := args[1]
	description := args[2]
	logo := args[3]
	parentID, err := uuid.Parse(args[4])
	if err == nil {

		fmt.Printf("gandalf cli create authorization called with name=%s, shortDescription=%s, description=%s, logo=%s, parentID=%s\n", name, shortDescription, description, logo, parentID)
		configurationCli := cmodels.NewConfigurationCli()
		cliClient := cli.NewClient(configurationCli.GetEndpoint())

		tag := models.Tag{Name: name, ShortDescription: shortDescription, Description: description, Logo: logo, ParentID: parentID}
		err := cliClient.TagService.Create(configurationCli.GetToken(), tag, parentID)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}
}

func runListTags(cfg *verdeter.ConfigCmd, args []string) {
	fmt.Printf("gandalf cli list users\n")
	configurationCli := cmodels.NewConfigurationCli()
	cliClient := cli.NewClient(configurationCli.GetEndpoint())

	tags, err := cliClient.TagService.List(configurationCli.GetToken())
	if err == nil {
		for _, tag := range tags {
			fmt.Println(tag)
		}
	} else {
		fmt.Println(err)
	}

}

func runUpdateTag(cfg *verdeter.ConfigCmd, args []string) {
	tagID, err := uuid.Parse(args[0])
	if err == nil {
		name := viper.GetViper().GetString("name")
		shortDescription := viper.GetViper().GetString("shortDescription")
		description := viper.GetViper().GetString("description")
		logo := viper.GetViper().GetString("logo")
		parentID, err := uuid.Parse(viper.GetViper().GetString("parentID"))
		if err == nil {
			fmt.Printf("gandalf cli update user called with name=%s, shortDescription=%s, description=%s, logo=%s, parentID=%s\n", name, shortDescription, description, logo, parentID)
			configurationCli := cmodels.NewConfigurationCli()
			cliClient := cli.NewClient(configurationCli.GetEndpoint())

			oldTag, err := cliClient.TagService.Read(configurationCli.GetToken(), tagID)
			if err == nil {
				oldTag.Name = name
				oldTag.ShortDescription = shortDescription
				oldTag.Description = description
				oldTag.Logo = logo
				oldTag.ParentID = parentID
				err = cliClient.TagService.Update(configurationCli.GetToken(), tagID, *oldTag)
				if err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Println(err)
			}
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}

}

func runDeleteTag(cfg *verdeter.ConfigCmd, args []string) {
	tagID, err := uuid.Parse(args[0])
	if err == nil {
		fmt.Printf("gandalf cli delete user called with tagID=%s\n", tagID)
		configurationCli := cmodels.NewConfigurationCli()
		cliClient := cli.NewClient(configurationCli.GetEndpoint())

		err = cliClient.TagService.Delete(configurationCli.GetToken(), tagID)
		if err != nil {
			fmt.Println(err)
		}

	} else {
		fmt.Println(err)
	}
}

///
func runCreateSecret(cfg *verdeter.ConfigCmd, args []string) {

	fmt.Printf("gandalf cli create secret called \n")
	configurationCli := cmodels.NewConfigurationCli()
	fmt.Println(configurationCli.GetEndpoint())
	cliClient := cli.NewClient(configurationCli.GetEndpoint())

	secret, err := cliClient.SecretAssignementService.Create(configurationCli.GetToken())
	if err == nil {
		fmt.Println(secret)
	} else {
		fmt.Println(err)
	}
}

func runListSecret(cfg *verdeter.ConfigCmd, args []string) {
	fmt.Printf("gandalf cli list secret called \n")
	configurationCli := cmodels.NewConfigurationCli()
	fmt.Println(configurationCli.GetEndpoint())
	cliClient := cli.NewClient(configurationCli.GetEndpoint())

	secrets, err := cliClient.SecretAssignementService.List(configurationCli.GetToken())
	if err == nil {
		for _, secret := range secrets {
			fmt.Println(secret)
		}
	} else {
		fmt.Println(err)
	}
}

func runCreateUser(cfg *verdeter.ConfigCmd, args []string) {
	email := args[0]
	firstname := args[1]
	secondname := args[2]
	companyid := args[3]
	password := args[4]

	fmt.Printf("gandalf cli create user called with email=%s, firstname=%s, secondname=%s, companyid=%s, password=%s\n", email, firstname, secondname, companyid, password)
	configurationCli := cmodels.NewConfigurationCli()
	cliClient := cli.NewClient(configurationCli.GetEndpoint())

	user := models.NewUser(email, firstname, secondname, companyid, password)
	err := cliClient.UserService.Create(configurationCli.GetToken(), user)
	if err != nil {
		fmt.Println(err)
	}

}

func runListUsers(cfg *verdeter.ConfigCmd, args []string) {
	fmt.Printf("gandalf cli list users\n")
	configurationCli := cmodels.NewConfigurationCli()
	cliClient := cli.NewClient(configurationCli.GetEndpoint())

	users, err := cliClient.UserService.List(configurationCli.GetToken())
	if err == nil {
		for _, user := range users {
			fmt.Println(user)
		}
	} else {
		fmt.Println(err)
	}

}

func runUpdateUser(cfg *verdeter.ConfigCmd, args []string) {
	userID, err := uuid.Parse(args[0])
	if err == nil {
		email := viper.GetViper().GetString("email")
		firstname := viper.GetViper().GetString("firstname")
		secondname := viper.GetViper().GetString("secondname")
		companyid := viper.GetViper().GetString("companyid")
		password := viper.GetViper().GetString("password")
		fmt.Printf("gandalf cli update user called with email=%s, firstname=%s, secondname=%s, companyid=%s, password=%s\n", email, firstname, secondname, companyid, password)
		configurationCli := cmodels.NewConfigurationCli()
		cliClient := cli.NewClient(configurationCli.GetEndpoint())

		_, err := cliClient.UserService.Read(configurationCli.GetToken(), userID)
		if err == nil {
			user := models.NewUser(email, firstname, secondname, companyid, password)
			err = cliClient.UserService.Update(configurationCli.GetToken(), userID, user)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}

}

func runDeleteUser(cfg *verdeter.ConfigCmd, args []string) {
	userID, err := uuid.Parse(args[0])
	if err == nil {
		fmt.Printf("gandalf cli delete user called with userID=%s\n", userID)
		configurationCli := cmodels.NewConfigurationCli()
		cliClient := cli.NewClient(configurationCli.GetEndpoint())

		err = cliClient.UserService.Delete(configurationCli.GetToken(), userID)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}

}

func runCreateTenant(cfg *verdeter.ConfigCmd, args []string) {
	name := args[0]
	shortDescription := args[1]
	description := args[2]
	logo := args[3]
	fmt.Printf("gandalf cli create tenant called with name=%s, shortDescription=%s, description=%s, logo=%s\n", name, shortDescription, description, logo)
	configurationCli := cmodels.NewConfigurationCli()
	cliClient := cli.NewClient(configurationCli.GetEndpoint())

	tenant := models.Tenant{Name: name}
	login, password, err := cliClient.TenantService.Create(configurationCli.GetToken(), tenant)
	if err == nil {
		fmt.Println("login : " + login)
		fmt.Println("password : " + password)
	} else {
		fmt.Println(err)
	}

}

func runListTenants(cfg *verdeter.ConfigCmd, args []string) {
	fmt.Printf("gandalf cli list tenants\n")
	configurationCli := cmodels.NewConfigurationCli()
	cliClient := cli.NewClient(configurationCli.GetEndpoint())

	tenants, err := cliClient.TenantService.List(configurationCli.GetToken())
	if err == nil {
		for _, tenant := range tenants {
			fmt.Println(tenant)
		}
	} else {
		fmt.Println(err)
	}

}

func runUpdateTenant(cfg *verdeter.ConfigCmd, args []string) {
	name := args[0]
	newName := viper.GetString("name")
	fmt.Printf("gandalf cli update tenant called with tenant=%s, newName=%s\n", name, newName)
}

func runDeleteTenant(cfg *verdeter.ConfigCmd, args []string) {
	name := args[0]
	fmt.Printf("gandalf cli delete tenant called with tenant=%s\n", name)
}

func runCreateRole(cfg *verdeter.ConfigCmd, args []string) {
	name := args[0]
	shortDescription := args[1]
	description := args[2]
	logo := args[3]
	fmt.Printf("gandalf cli create role called with name=%s, shortDescription=%s, description=%s, logo=%s\n", name, shortDescription, description, logo)

	configurationCli := cmodels.NewConfigurationCli()
	cliClient := cli.NewClient(configurationCli.GetEndpoint())

	role := models.Role{Name: name, ShortDescription: shortDescription, Description: description, Logo: logo}
	err := cliClient.RoleService.Create(configurationCli.GetToken(), role)
	if err != nil {
		fmt.Println(err)
	}
}

func runListRoles(cfg *verdeter.ConfigCmd, args []string) {
	fmt.Printf("gandalf cli list roles\n")
	configurationCli := cmodels.NewConfigurationCli()
	cliClient := cli.NewClient(configurationCli.GetEndpoint())

	roles, err := cliClient.RoleService.List(configurationCli.GetToken())
	if err == nil {
		for _, role := range roles {
			fmt.Println(role)
		}
	} else {
		fmt.Println(err)
	}
}

func runUpdateRole(cfg *verdeter.ConfigCmd, args []string) {
	roleID, err := uuid.Parse(args[0])
	if err == nil {
		name := viper.GetViper().GetString("name")
		shortDescription := viper.GetViper().GetString("shortDescription")
		description := viper.GetViper().GetString("description")
		logo := viper.GetViper().GetString("logo")
		fmt.Printf("gandalf cli update role called with name=%s, shortDescription=%s, description=%s, logo=%s\n", name, shortDescription, description, logo)
		configurationCli := cmodels.NewConfigurationCli()
		cliClient := cli.NewClient(configurationCli.GetEndpoint())

		oldRole, err := cliClient.RoleService.Read(configurationCli.GetToken(), roleID)
		if err == nil {
			oldRole.Name = name
			oldRole.ShortDescription = shortDescription
			oldRole.Description = description
			oldRole.Logo = logo
			err = cliClient.RoleService.Update(configurationCli.GetToken(), roleID, *oldRole)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}
}

func runDeleteRole(cfg *verdeter.ConfigCmd, args []string) {
	roleID, err := uuid.Parse(args[0])
	if err == nil {
		fmt.Printf("gandalf cli delete role called with roleID=%s\n", roleID)
		configurationCli := cmodels.NewConfigurationCli()
		cliClient := cli.NewClient(configurationCli.GetEndpoint())

		err = cliClient.RoleService.Delete(configurationCli.GetToken(), roleID)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}
}

func runCreateDomain(cfg *verdeter.ConfigCmd, args []string) {
	name := args[0]
	shortDescription := args[1]
	description := args[2]
	logo := args[3]
	parentID, err := uuid.Parse(args[4])
	if err == nil {
		fmt.Printf("gandalf cli create domain called with name=%s, shortDescription=%s, description=%s, logo=%s, parentID=%s\n", name, shortDescription, description, logo, parentID)

		configurationCli := cmodels.NewConfigurationCli()
		cliClient := cli.NewClient(configurationCli.GetEndpoint())
		domain := models.Domain{Name: name, ShortDescription: shortDescription, Description: description, Logo: logo, ParentID: parentID}
		err := cliClient.DomainService.Create(configurationCli.GetToken(), domain, parentID)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}
}

func runListDomains(cfg *verdeter.ConfigCmd, args []string) {
	fmt.Printf("gandalf cli list domains\n")
	configurationCli := cmodels.NewConfigurationCli()
	cliClient := cli.NewClient(configurationCli.GetEndpoint())

	domains, err := cliClient.DomainService.List(configurationCli.GetToken())
	if err == nil {
		for _, domain := range domains {
			fmt.Println(domain)
		}
	} else {
		fmt.Println(err)
	}

}

func runUpdateDomain(cfg *verdeter.ConfigCmd, args []string) {
	domainID, err := uuid.Parse(args[0])
	if err == nil {
		name := viper.GetViper().GetString("name")
		shortDescription := viper.GetViper().GetString("shortDescription")
		description := viper.GetViper().GetString("description")
		logo := viper.GetViper().GetString("logo")
		parentID, err := uuid.Parse(viper.GetViper().GetString("parentID"))
		if err == nil {
			fmt.Printf("gandalf cli update domain called with name=%s, shortDescription=%s, description=%s, logo=%s, parentID=%s\n", name, shortDescription, description, logo, parentID)
			configurationCli := cmodels.NewConfigurationCli()
			cliClient := cli.NewClient(configurationCli.GetEndpoint())

			oldDomain, err := cliClient.DomainService.Read(configurationCli.GetToken(), domainID)
			if err == nil {
				oldDomain.Name = name
				oldDomain.ShortDescription = shortDescription
				oldDomain.Description = description
				oldDomain.Logo = logo
				oldDomain.ParentID = parentID
				err = cliClient.DomainService.Update(configurationCli.GetToken(), domainID, *oldDomain)
				if err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Println(err)
			}
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}
}

func runDeleteDomain(cfg *verdeter.ConfigCmd, args []string) {
	domainID, err := uuid.Parse(args[0])
	if err == nil {
		fmt.Printf("gandalf cli delete domain called with domainID=%s\n", domainID)
		configurationCli := cmodels.NewConfigurationCli()
		cliClient := cli.NewClient(configurationCli.GetEndpoint())

		err = cliClient.DomainService.Delete(configurationCli.GetToken(), domainID)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}
}

func runCreateResource(cfg *verdeter.ConfigCmd, args []string) {
	name := args[0]
	logicalComponentID, err := uuid.Parse(args[1])
	if err == nil {
		domainID, err := uuid.Parse(args[2])
		if err == nil {
			resourceTypeID, err := uuid.Parse(args[3])
			if err == nil {
				fmt.Printf("gandalf cli create resource called with name=%s, logicalComponentID=%s, domainID=%s, resourceTypeID=%s\n", name, logicalComponentID, domainID, resourceTypeID)

				configurationCli := cmodels.NewConfigurationCli()
				cliClient := cli.NewClient(configurationCli.GetEndpoint())
				resource := models.Resource{Name: name, LogicalComponentID: logicalComponentID, DomainID: domainID, ResourceTypeID: resourceTypeID}
				err := cliClient.ResourceService.Create(configurationCli.GetToken(), resource)
				if err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Println(err)
			}
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}
}

func runListResources(cfg *verdeter.ConfigCmd, args []string) {
	fmt.Printf("gandalf cli list resources\n")
	configurationCli := cmodels.NewConfigurationCli()
	cliClient := cli.NewClient(configurationCli.GetEndpoint())

	resources, err := cliClient.ResourceService.List(configurationCli.GetToken())
	if err == nil {
		for _, resource := range resources {
			fmt.Println(resource)
		}
	} else {
		fmt.Println(err)
	}

}

func runUpdateResource(cfg *verdeter.ConfigCmd, args []string) {
	resourceID, err := uuid.Parse(args[0])
	if err == nil {
		name := viper.GetViper().GetString("name")
		logicalComponentID, err := uuid.Parse(viper.GetViper().GetString("logicalComponentID"))
		if err == nil {
			domainID, err := uuid.Parse(viper.GetViper().GetString("domainID"))
			if err == nil {
				resourceTypeID, err := uuid.Parse(viper.GetViper().GetString("resourceTypeID"))
				if err == nil {
					fmt.Printf("gandalf cli update resource called with name=%s, logicalComponentID=%s, domainID=%s, resourceTypeID=%s\n", name, logicalComponentID, domainID, resourceTypeID)
					configurationCli := cmodels.NewConfigurationCli()
					cliClient := cli.NewClient(configurationCli.GetEndpoint())

					oldResource, err := cliClient.ResourceService.Read(configurationCli.GetToken(), resourceID)
					if err == nil {
						oldResource.Name = name
						oldResource.LogicalComponentID = logicalComponentID
						oldResource.DomainID = domainID
						oldResource.ResourceTypeID = resourceTypeID
						err = cliClient.ResourceService.Update(configurationCli.GetToken(), resourceID, *oldResource)
						if err != nil {
							fmt.Println(err)
						}
					} else {
						fmt.Println(err)
					}
				} else {
					fmt.Println(err)
				}
			} else {
				fmt.Println(err)
			}
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}
}

func runDeleteResource(cfg *verdeter.ConfigCmd, args []string) {
	resourceID, err := uuid.Parse(args[0])
	if err == nil {
		fmt.Printf("gandalf cli delete resource called with resourceID=%s\n", resourceID)
		configurationCli := cmodels.NewConfigurationCli()
		cliClient := cli.NewClient(configurationCli.GetEndpoint())

		err = cliClient.ResourceService.Delete(configurationCli.GetToken(), resourceID)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}
}

func runCreateEventTypeToPoll(cfg *verdeter.ConfigCmd, args []string) {
	resourceID, err := uuid.Parse(args[0])
	if err == nil {
		eventTypeID, err := uuid.Parse(args[1])
		if err == nil {
			fmt.Printf("gandalf cli create eventtypetopoll called with resourceID=%s and eventTypeID=%s \n", resourceID, eventTypeID)

			configurationCli := cmodels.NewConfigurationCli()
			cliClient := cli.NewClient(configurationCli.GetEndpoint())

			eventTypeToPoll := models.EventTypeToPoll{ResourceID: resourceID, EventTypeID: eventTypeID}
			err = cliClient.EventTypeToPollService.Create(configurationCli.GetToken(), eventTypeToPoll)
			if err != nil {
				fmt.Println(err)
			}

		} else {
			fmt.Println(err)
		}

	} else {
		fmt.Println(err)
	}
}

func runListEventTypeToPolls(cfg *verdeter.ConfigCmd, args []string) {
	fmt.Printf("gandalf cli list eventtypetopoll\n")
	configurationCli := cmodels.NewConfigurationCli()
	cliClient := cli.NewClient(configurationCli.GetEndpoint())

	eventTypeToPolls, err := cliClient.EventTypeToPollService.List(configurationCli.GetToken())
	if err == nil {
		for _, eventTypeToPoll := range eventTypeToPolls {
			fmt.Println(eventTypeToPoll)
		}
	} else {
		fmt.Println(err)
	}

}

func runUpdateEventTypeToPoll(cfg *verdeter.ConfigCmd, args []string) {
	name := args[0]
	newName := viper.GetString("name")
	parent := viper.GetString("parent")
	fmt.Printf("gandalf cli update eventtypetopoll called with eventtypetopoll=%s, newName=%s, parent=%s\n", name, newName, parent)
}

func runDeleteEventTypeToPoll(cfg *verdeter.ConfigCmd, args []string) {
	name := args[0]
	fmt.Printf("gandalf cli delete eventtypetopoll called with eventtypetopoll=%s\n", name)
}

func runCreateResourceType(cfg *verdeter.ConfigCmd, args []string) {
	name := args[0]
	pivotID, err := uuid.Parse(args[1])
	if err == nil {
		typePivot := args[2]
		fmt.Printf("gandalf cli create resourceType called with name=%s, pivotID=%s\n", name, pivotID)

		configurationCli := cmodels.NewConfigurationCli()
		cliClient := cli.NewClient(configurationCli.GetEndpoint())
		if typePivot == "pivot" {
			resourceType := models.ResourceType{Name: name, PivotID: pivotID}
			err := cliClient.ResourceTypeService.Create(configurationCli.GetToken(), resourceType)
			if err != nil {
				fmt.Println(err)
			}
		} else if typePivot == "productconnector" {
			resourceType := models.ResourceType{Name: name, ProductConnectorID: pivotID}
			err := cliClient.ResourceTypeService.Create(configurationCli.GetToken(), resourceType)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Wrong type")
		}

	} else {
		fmt.Println(err)
	}
}

func runListResourceTypes(cfg *verdeter.ConfigCmd, args []string) {
	fmt.Printf("gandalf cli list resourceTypes\n")
	configurationCli := cmodels.NewConfigurationCli()
	cliClient := cli.NewClient(configurationCli.GetEndpoint())

	resourceTypes, err := cliClient.ResourceTypeService.List(configurationCli.GetToken())
	if err == nil {
		for _, resourceType := range resourceTypes {
			fmt.Println(resourceType)
		}
	} else {
		fmt.Println(err)
	}

}

func runUpdateResourceType(cfg *verdeter.ConfigCmd, args []string) {
	resourceTypeID, err := uuid.Parse(args[0])
	if err == nil {
		name := viper.GetViper().GetString("name")
		pivotID, err := uuid.Parse(viper.GetViper().GetString("pivotID"))
		if err == nil {
			typePivot := viper.GetViper().GetString("typePivot")

			fmt.Printf("gandalf cli update resourceType called with name=%s, pivotID=%s, typePivot=%s\n", name, pivotID, typePivot)
			configurationCli := cmodels.NewConfigurationCli()
			cliClient := cli.NewClient(configurationCli.GetEndpoint())

			oldResourceType, err := cliClient.ResourceTypeService.Read(configurationCli.GetToken(), resourceTypeID)
			if err == nil {
				if typePivot == "pivot" {
					oldResourceType.Name = name
					oldResourceType.PivotID = pivotID
					err = cliClient.ResourceTypeService.Update(configurationCli.GetToken(), resourceTypeID, *oldResourceType)
					if err != nil {
						fmt.Println(err)
					}
				} else if typePivot == "productconnector" {
					oldResourceType.Name = name
					oldResourceType.ProductConnectorID = pivotID
					err = cliClient.ResourceTypeService.Update(configurationCli.GetToken(), resourceTypeID, *oldResourceType)
					if err != nil {
						fmt.Println(err)
					}
				} else {
					fmt.Println("Wrong type")
				}

			} else {
				fmt.Println(err)
			}
		} else {
			fmt.Println(err)
		}

	} else {
		fmt.Println(err)
	}
}

func runDeleteResourceType(cfg *verdeter.ConfigCmd, args []string) {
	resourceTypeID, err := uuid.Parse(args[0])
	if err == nil {
		fmt.Printf("gandalf cli delete resourceType called with resourceID=%s\n", resourceTypeID)
		configurationCli := cmodels.NewConfigurationCli()
		cliClient := cli.NewClient(configurationCli.GetEndpoint())

		err = cliClient.ResourceTypeService.Delete(configurationCli.GetToken(), resourceTypeID)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}
}

func runCreateEventType(cfg *verdeter.ConfigCmd, args []string) {
	name := args[0]
	schema := args[1]
	pivotID, err := uuid.Parse(args[2])
	if err == nil {
		typePivot := args[3]
		fmt.Printf("gandalf cli create eventType called with name=%s, schema=%s, pivotID=%s, typePivot=%s\n", name, schema, pivotID, typePivot)

		configurationCli := cmodels.NewConfigurationCli()
		cliClient := cli.NewClient(configurationCli.GetEndpoint())
		if typePivot == "pivot" {
			eventType := models.EventType{Name: name, Schema: schema, PivotID: pivotID}
			err := cliClient.EventTypeService.Create(configurationCli.GetToken(), eventType)
			if err != nil {
				fmt.Println(err)
			}
		} else if typePivot == "productconnector" {
			eventType := models.EventType{Name: name, Schema: schema, ProductConnectorID: pivotID}
			err := cliClient.EventTypeService.Create(configurationCli.GetToken(), eventType)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Wrong type")
		}

	} else {
		fmt.Println(err)
	}
}

func runListEventTypes(cfg *verdeter.ConfigCmd, args []string) {
	fmt.Printf("gandalf cli list eventtypes\n")
	configurationCli := cmodels.NewConfigurationCli()
	cliClient := cli.NewClient(configurationCli.GetEndpoint())

	eventTypeToPolls, err := cliClient.EventTypeService.List(configurationCli.GetToken())
	if err == nil {
		for _, eventTypeToPoll := range eventTypeToPolls {
			fmt.Println(eventTypeToPoll)
		}
	} else {
		fmt.Println(err)
	}

}

func runUpdateEventType(cfg *verdeter.ConfigCmd, args []string) {
	eventTypeID, err := uuid.Parse(args[0])
	if err == nil {
		name := viper.GetViper().GetString("name")
		schema := viper.GetViper().GetString("schema")
		pivotID, err := uuid.Parse(viper.GetViper().GetString("pivotID"))
		if err == nil {
			typePivot := viper.GetViper().GetString("typePivot")

			fmt.Printf("gandalf cli update eventType called with name=%s, schema=%s, pivotID=%s, typePivot=%s\n", name, schema, pivotID, typePivot)
			configurationCli := cmodels.NewConfigurationCli()
			cliClient := cli.NewClient(configurationCli.GetEndpoint())

			oldEventType, err := cliClient.EventTypeService.Read(configurationCli.GetToken(), eventTypeID)
			if err == nil {
				if typePivot == "pivot" {
					oldEventType.Name = name
					oldEventType.Schema = schema
					oldEventType.PivotID = pivotID
					err = cliClient.EventTypeService.Update(configurationCli.GetToken(), eventTypeID, *oldEventType)
					if err != nil {
						fmt.Println(err)
					}
				} else if typePivot == "productconnector" {
					oldEventType.Name = name
					oldEventType.Schema = schema
					oldEventType.ProductConnectorID = pivotID
					err = cliClient.EventTypeService.Update(configurationCli.GetToken(), eventTypeID, *oldEventType)
					if err != nil {
						fmt.Println(err)
					}
				} else {
					fmt.Println("Wrong type")
				}

			} else {
				fmt.Println(err)
			}
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}
}

func runDeleteEventType(cfg *verdeter.ConfigCmd, args []string) {
	eventTypeID, err := uuid.Parse(args[0])
	if err == nil {
		fmt.Printf("gandalf cli delete eventType called with eventTypeID=%s\n", eventTypeID)
		configurationCli := cmodels.NewConfigurationCli()
		cliClient := cli.NewClient(configurationCli.GetEndpoint())

		err = cliClient.EventTypeService.Delete(configurationCli.GetToken(), eventTypeID)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}
}

/*
func init() {
	rootCfg.AddConfig(cliCfg)

	cliCfg.Key("api_port", verdeter.IsInt, "", "Address to bind (default is *:9199)")
	cliCfg.SetDefault("api_port", 9199+verdeter.GetOffset())

	cliCfg.Key("database_mode", verdeter.IsStr, "", "database mode (gandalf|tenant)")
	cliCfg.SetDefault("database_mode", "gandalf")
	cliCfg.SetCheck("database_mode", func(val interface{}) bool {
		strVal := strings.ToLower(strings.TrimSpace(val.(string)))
		return map[string]bool{"gandalf": true, "tenant": true}[strVal]
	})

	cliCfg.Key("tenant", verdeter.IsStr, "", "database mode (gandalf|tenant)")
	cliCfg.SetConstraint("a tenant should be provided if database_mode == tenant",
		func() bool {
			return viper.IsSet("database_mode") && viper.GetString("database_mode") == "tenant" && viper.IsSet("tenant")
		})

	cliCfg.Key("model", verdeter.IsStr, "", "models  gandalf(authentication|cluster|tenant|role|user) || tenant(authentication|aggregator|connector|role|user)")
	cliCfg.SetCheck("model", func(val interface{}) bool {
		strVal := strings.ToLower(strings.TrimSpace(val.(string)))
		return map[string]bool{"authentication": true, "cluster": true, "tenant": true, "role": true, "user": true, "aggregator": true, "connector": true}[strVal]
	})
*/

//TODO REVOIR
/* 	cliCfg.SetConstraint("a id should be provided if command == (delete|update|read)",
func() bool {
	return viper.IsSet("database_mode") || viper.GetString("database_mode") == "gandalf" || viper.GetString("command") == "update" || viper.GetString("command") == "read" || viper.IsSet("id")
}) */

/*
	cliCfg.Key("command", verdeter.IsStr, "", "command  (list|read|create|update|delete|upload)")
	cliCfg.SetCheck("command", func(val interface{}) bool {
		strVal := strings.ToLower(strings.TrimSpace(val.(string)))
		return map[string]bool{"list": true, "read": true, "create": true, "update": true, "delete": true, "upload": true}[strVal]
	})

	cliCfg.Key("token", verdeter.IsStr, "", "")
	cliCfg.SetConstraint("a token should be provided if model != authenticaion",
		func() bool {
			return viper.IsSet("model") && viper.GetString("model") != "authentication" && viper.IsSet("token")
		})

	cliCfg.Key("id", verdeter.IsStr, "", "id")
	cliCfg.SetConstraint("a id should be provided if command == (delete|update|read)",
		func() bool {
			return viper.IsSet("command") && (viper.GetString("command") == "delete" || viper.GetString("command") == "update" || viper.GetString("command") == "read") && viper.IsSet("id")
		})

	cliCfg.Key("value", verdeter.IsStr, "", "json")
	cliCfg.SetConstraint("a value should be provided if command == (create|update)",
		func() bool {
			return viper.IsSet("command") && (viper.GetString("command") == "create" || viper.GetString("command") == "update") && viper.IsSet("value")
		})
}
*/
