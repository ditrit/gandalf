/*
Copyright © 2020 DitRit community <contact@ditrit.io>
This file is part of Gandalf
*/

// Package cmd manages commands and configuration
package configuration

import (
	"github.com/ditrit/gandalf/core/configuration/config"
)

// cliCmd represents the cli command
var cliCfg = config.NewConfigCmd("cli", "Launch gandalf in 'cli' mode.", `Gandalf is launched as CLI (Command Line Interface) to interact with a Gandalf system.`, nil)

var cliCreate = config.NewConfigCmd("create", "create user|tenant|role|domain", "create command allows the creation of Gandalf objects (users, tenants, roles and domains).", nil)
var cliList = config.NewConfigCmd("list", "list users|tenants|roles|domains", "list command allows to list Gandalf objects (users, tenants, roles and domains).", nil)
var cliUpdate = config.NewConfigCmd("update", "update user|tenant|role|domain", "update command allows update of Gandalf objects (users, tenants, roles and domains).", nil)
var cliDelete = config.NewConfigCmd("delete", "delete user|tenant|role|domain", "update command allows deleting of Gandalf objects (users, tenants, roles and domains).", nil)
var cliDeclare = config.NewConfigCmd("declare", "declare cluster|agregator|connector", "declare command allows to declare a new Gandalf component name or member (cluster, aggregator, connector).", nil)
var cliLogin = config.NewConfigCmd("login", "log in as a user into Gandalf", "login command allows user to authenticate using its credentials.", nil)

var cliCreateUser = config.NewConfigCmd("user", "create user <username> [options]", "create user command allows the creation of a new user", nil)
var cliListUsers = config.NewConfigCmd("user", "list users <username> [options]", "list users command allows to list all or filtered (using regexp) Gandalf users.", nil)
var cliUpdateUser = config.NewConfigCmd("user", "update user <username> [options]", "update user command allows to update a Gandalf user.", nil)
var cliDeleteUser = config.NewConfigCmd("user", "delete user <username>", "delete user command allows to delete a Gandalf user.", nil)

var cliCreateTenant = config.NewConfigCmd("tenant", "create tenant <tenantname> [options]", "create tenant command allows the creation of a new tenant", nil)
var cliListTenants = config.NewConfigCmd("tenant", "list tenants <tenantname> [options]", "list tenants command allows to list all or filtered (using regexp) Gandalf tenants.", nil)
var cliUpdateTenant = config.NewConfigCmd("tenant", "update tenant <tenantname> [options]", "update tenant command allows to update a Gandalf tenant.", nil)
var cliDeleteTenant = config.NewConfigCmd("tenant", "delete tenant <tenantname>", "delete tenant command allows to delete a Gandalf tenant.", nil)

var cliCreateRole = config.NewConfigCmd("role", "create role <rolename> [options]", "create role command allows the creation of a new role", nil)
var cliListRoles = config.NewConfigCmd("role", "list roles <rolename> [options]", "list roles command allows to list all or filtered (using regexp) Gandalf roles.", nil)
var cliUpdateRole = config.NewConfigCmd("role", "update role <rolename> [options]", "update role command allows to update a Gandalf role.", nil)
var cliDeleteRole = config.NewConfigCmd("role", "delete role <rolename>", "delete role command allows to delete a Gandalf role.", nil)

var cliCreateDomain = config.NewConfigCmd("domain", "create domain <domainname> [options]", "create domain command allows the creation of a new domain", nil)
var cliListDomains = config.NewConfigCmd("domain", "list domains <domainname> [options]", "list domains command allows to list all or filtered (using regexp) Gandalf domains.", nil)
var cliUpdateDomain = config.NewConfigCmd("domain", "update domain <domainname> [options]", "update domain command allows to update a Gandalf domain.", nil)
var cliDeleteDomain = config.NewConfigCmd("domain", "delete domain <domainname>", "delete domain command allows to delete a Gandalf domain.", nil)

var cliDeclareCluster = config.NewConfigCmd("custer", "declare cluster", "declare cluster command allows to declare a new cluster memeber", nil)
var cliDeclareAggregator = config.NewConfigCmd("aggregator", "declare aggregator name|member", "declare aggregator command allows to declare the name or a new member for an aggragator.", nil)
var cliDeclareConnector = config.NewConfigCmd("connector", "declare connector name|member", "declare connector command allows to declare the name or a new member for a connector.", nil)

var cliDeclareClusterMember = config.NewConfigCmd("member", "declare cluster member", "declare cluster command allows to declare a new cluster member", nil)

var cliDeclareAggregatorName = config.NewConfigCmd("name", "declare aggregator name <name>", "declare aggregator name command allows to declare the name of a new aggregator.", nil)
var cliDeclareAggregatorMember = config.NewConfigCmd("member", "declare aggregator member <name>", "declare aggregator member command allows to declare a new member for an existing aggregator.", nil)

var cliDeclareConnectorName = config.NewConfigCmd("name", "declare  name <name>", "declare  name command allows to declare the name of a new connector.", nil)
var cliDeclareConnectorMember = config.NewConfigCmd("member", "declare  member <name>", "declare  member command allows to declare a new member for an existing connector.", nil)

func init() {
	rootCfg.AddConfig(cliCfg)
	cliCfg.AddConfig(cliCreate)
	cliCfg.AddConfig(cliList)
	cliCfg.AddConfig(cliUpdate)
	cliCfg.AddConfig(cliDelete)
	cliCfg.AddConfig(cliDeclare)
	cliCfg.AddConfig(cliLogin)

	cliCreate.AddConfig(cliCreateUser)
	cliList.AddConfig(cliListUsers)
	cliUpdate.AddConfig(cliUpdateUser)
	cliDelete.AddConfig(cliDeleteUser)

	cliCreate.AddConfig(cliCreateTenant)
	cliList.AddConfig(cliListTenants)
	cliUpdate.AddConfig(cliUpdateTenant)
	cliDelete.AddConfig(cliDeleteTenant)

	cliCreate.AddConfig(cliCreateRole)
	cliList.AddConfig(cliListRoles)
	cliUpdate.AddConfig(cliUpdateRole)
	cliDelete.AddConfig(cliDeleteRole)

	cliCreate.AddConfig(cliCreateDomain)
	cliList.AddConfig(cliListDomains)
	cliUpdate.AddConfig(cliUpdateDomain)
	cliDelete.AddConfig(cliDeleteDomain)

	cliDeclare.AddConfig(cliDeclareCluster)
	cliDeclare.AddConfig(cliDeclareAggregator)
	cliDeclare.AddConfig(cliDeclareConnector)

	cliDeclareCluster.AddConfig(cliDeclareClusterMember)

	cliDeclareAggregator.AddConfig(cliDeclareAggregatorName)
	cliDeclareAggregator.AddConfig(cliDeclareAggregatorMember)

	cliDeclareConnector.AddConfig(cliDeclareConnectorName)
	cliDeclareConnector.AddConfig(cliDeclareConnectorMember)

	cliCfg.Key("endpoint", config.IsStr, "e", "Gandalf auth token")
	cliCfg.SetRequired("endpoint")
	cliCfg.Key("token", config.IsStr, "t", "Gandalf auth token")
	cliCfg.SetRequired("token")

	cliCreateUser.Key("email", config.IsStr, "m", "mail of the user")
	cliCreateUser.SetRequired("email")
	cliCreateUser.Key("password", config.IsStr, "p", "password of the user")
	cliCreateUser.SetRequired("password")
	cliUpdateUser.Key("name", config.IsStr, "n", "name of the user")
	cliUpdateUser.Key("email", config.IsStr, "m", "mail of the user")
	cliUpdateUser.Key("password", config.IsStr, "p", "password of the user")
	cliListUsers.Key("filter", config.IsStr, "f", "regexp to filter results")

	cliUpdateTenant.Key("name", config.IsStr, "n", "name of the Tenant")
	cliListTenants.Key("filter", config.IsStr, "f", "regexp to filter results")

	cliUpdateRole.Key("name", config.IsStr, "n", "name of the Role")
	cliListRoles.Key("filter", config.IsStr, "f", "regexp to filter results")

	cliUpdateDomain.Key("name", config.IsStr, "n", "name of the Domain")
	cliListDomains.Key("filter", config.IsStr, "f", "regexp to filter results")

}

/*
func init() {
	rootCfg.AddConfig(cliCfg)

	cliCfg.Key("api_port", config.IsInt, "", "Address to bind (default is *:9199)")
	cliCfg.SetDefault("api_port", 9199+config.GetOffset())

	cliCfg.Key("database_mode", config.IsStr, "", "database mode (gandalf|tenant)")
	cliCfg.SetDefault("database_mode", "gandalf")
	cliCfg.SetCheck("database_mode", func(val interface{}) bool {
		strVal := strings.ToLower(strings.TrimSpace(val.(string)))
		return map[string]bool{"gandalf": true, "tenant": true}[strVal]
	})

	cliCfg.Key("tenant", config.IsStr, "", "database mode (gandalf|tenant)")
	cliCfg.SetConstraint("a tenant should be provided if database_mode == tenant",
		func() bool {
			return viper.IsSet("database_mode") && viper.GetString("database_mode") == "tenant" && viper.IsSet("tenant")
		})

	cliCfg.Key("model", config.IsStr, "", "models  gandalf(authentication|cluster|tenant|role|user) || tenant(authentication|aggregator|connector|role|user)")
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
	cliCfg.Key("command", config.IsStr, "", "command  (list|read|create|update|delete|upload)")
	cliCfg.SetCheck("command", func(val interface{}) bool {
		strVal := strings.ToLower(strings.TrimSpace(val.(string)))
		return map[string]bool{"list": true, "read": true, "create": true, "update": true, "delete": true, "upload": true}[strVal]
	})

	cliCfg.Key("token", config.IsStr, "", "")
	cliCfg.SetConstraint("a token should be provided if model != authenticaion",
		func() bool {
			return viper.IsSet("model") && viper.GetString("model") != "authentication" && viper.IsSet("token")
		})

	cliCfg.Key("id", config.IsStr, "", "id")
	cliCfg.SetConstraint("a id should be provided if command == (delete|update|read)",
		func() bool {
			return viper.IsSet("command") && (viper.GetString("command") == "delete" || viper.GetString("command") == "update" || viper.GetString("command") == "read") && viper.IsSet("id")
		})

	cliCfg.Key("value", config.IsStr, "", "json")
	cliCfg.SetConstraint("a value should be provided if command == (create|update)",
		func() bool {
			return viper.IsSet("command") && (viper.GetString("command") == "create" || viper.GetString("command") == "update") && viper.IsSet("value")
		})
}
*/
