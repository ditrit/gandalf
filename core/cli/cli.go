package cli

import (
	"encoding/json"
	"fmt"

	"github.com/ditrit/gandalf/core/cli/client"
	cmodels "github.com/ditrit/gandalf/core/configuration/models"

	models "github.com/ditrit/gandalf/core/models"
)

func Cli(configurationCli *cmodels.ConfigurationCli) {

	fmt.Println("APIBindAddress:", configurationCli.GetAPIBindAddress())
	fmt.Println("Database Mode:", configurationCli.GetDatabaseMode())
	fmt.Println("Tenant:", configurationCli.GetTenant())
	fmt.Println("Model:", configurationCli.GetModel())
	fmt.Println("Command:", configurationCli.GetCommand())
	fmt.Println("Token:", configurationCli.GetToken())
	fmt.Println("ID:", configurationCli.GetID())
	fmt.Println("Value:", configurationCli.GetValue())

	cliClient := client.NewClient(configurationCli.GetAPIBindAddress())

	if configurationCli.GetDatabaseMode() == "gandalf" {
		switch configurationCli.GetModel() {
		case "authentication":
			var user models.User
			err := json.Unmarshal([]byte(configurationCli.GetValue()), &user)
			if err == nil {
				result, _ := cliClient.GandalfAuthenticationService.Login(user)
				fmt.Println(result)
			} else {
				fmt.Println(err)
			}
		case "cluster":
			switch configurationCli.GetCommand() {
			case "list":
				clusters, err := cliClient.GandalfClusterService.List(configurationCli.GetToken())
				if err == nil {
					fmt.Println(clusters)
				} else {
					fmt.Println(err)
				}
			case "read":
				var cluster *models.Cluster
				cluster, err := cliClient.GandalfClusterService.Read(configurationCli.GetToken(), configurationCli.GetID())
				if err == nil {
					fmt.Println(cluster)
				} else {
					fmt.Println(err)
				}
			case "create":
				var cluster models.Cluster
				err := json.Unmarshal([]byte(configurationCli.GetValue()), &cluster)
				if err == nil {
					err = cliClient.GandalfClusterService.Create(configurationCli.GetToken(), cluster)
					if err != nil {
						fmt.Println(err)
					}
				} else {
					fmt.Println(err)
				}
			case "update":
				var cluster models.Cluster
				err := json.Unmarshal([]byte(configurationCli.GetValue()), &cluster)
				if err == nil {
					err = cliClient.GandalfClusterService.Update(configurationCli.GetToken(), configurationCli.GetID(), cluster)
					if err != nil {
						fmt.Println(err)
					}
				} else {
					fmt.Println(err)
				}
			case "delete":
				err := cliClient.GandalfClusterService.Delete(configurationCli.GetToken(), configurationCli.GetID())
				if err != nil {
					fmt.Println(err)
				}
			}
		case "tenant":
			switch configurationCli.GetCommand() {
			case "list":
				tenants, err := cliClient.GandalfTenantService.List(configurationCli.GetToken())
				if err == nil {
					fmt.Println(tenants)
				} else {
					fmt.Println(err)
				}
			case "read":
				var tenant *models.Tenant
				tenant, err := cliClient.GandalfTenantService.Read(configurationCli.GetToken(), configurationCli.GetID())
				if err == nil {
					fmt.Println(tenant)
				} else {
					fmt.Println(err)
				}
			case "create":
				var tenant models.Tenant
				err := json.Unmarshal([]byte(configurationCli.GetValue()), &tenant)
				fmt.Println(err)
				if err == nil {
					var login, password string
					login, password, err = cliClient.GandalfTenantService.Create(configurationCli.GetToken(), tenant)
					if err == nil {
						fmt.Println(login)
						fmt.Println(password)
					} else {
						fmt.Println(err)
					}
				} else {
					fmt.Println(err)
				}
			case "update":
				var tenant models.Tenant
				err := json.Unmarshal([]byte(configurationCli.GetValue()), &tenant)
				if err == nil {
					err = cliClient.GandalfTenantService.Update(configurationCli.GetToken(), configurationCli.GetID(), tenant)
					if err != nil {
						fmt.Println(err)
					}
				} else {
					fmt.Println(err)
				}
			case "delete":
				err := cliClient.GandalfTenantService.Delete(configurationCli.GetToken(), configurationCli.GetID())
				if err != nil {
					fmt.Println(err)
				}
			}
		case "role":
			switch configurationCli.GetCommand() {
			case "list":
				roles, err := cliClient.GandalfRoleService.List(configurationCli.GetToken())
				if err == nil {
					fmt.Println(roles)
				} else {
					fmt.Println(err)
				}
			case "read":
				var role *models.Role
				role, err := cliClient.GandalfRoleService.Read(configurationCli.GetToken(), configurationCli.GetID())
				if err == nil {
					fmt.Println(role)
				} else {
					fmt.Println(err)
				}
			case "create":
				var role models.Role
				err := json.Unmarshal([]byte(configurationCli.GetValue()), &role)
				fmt.Println("role")
				fmt.Println(role)
				if err == nil {
					err = cliClient.GandalfRoleService.Create(configurationCli.GetToken(), role)
					fmt.Println("err")
					fmt.Println(err)
					if err != nil {
						fmt.Println(err)
					}
				} else {
					fmt.Println(err)
				}
			case "update":
				var role models.Role
				err := json.Unmarshal([]byte(configurationCli.GetValue()), &role)
				if err == nil {
					err = cliClient.GandalfRoleService.Update(configurationCli.GetToken(), configurationCli.GetID(), role)
					if err != nil {
						fmt.Println(err)
					}
				} else {
					fmt.Println(err)
				}
			case "delete":
				err := cliClient.GandalfRoleService.Delete(configurationCli.GetToken(), configurationCli.GetID())
				if err != nil {
					fmt.Println(err)
				}
			}
		case "user":
			switch configurationCli.GetCommand() {
			case "list":
				users, err := cliClient.GandalfUserService.List(configurationCli.GetToken())
				if err == nil {
					fmt.Println(users)
				} else {
					fmt.Println(err)
				}
			case "read":
				var user *models.User
				user, err := cliClient.GandalfUserService.Read(configurationCli.GetToken(), configurationCli.GetID())
				if err == nil {
					fmt.Println(user)
				} else {
					fmt.Println(err)
				}
			case "create":
				var user models.User
				err := json.Unmarshal([]byte(configurationCli.GetValue()), &user)
				if err == nil {
					err = cliClient.GandalfUserService.Create(configurationCli.GetToken(), user)
					if err != nil {
						fmt.Println(err)
					}
				} else {
					fmt.Println(err)
				}
			case "update":
				var user models.User
				err := json.Unmarshal([]byte(configurationCli.GetValue()), &user)
				if err == nil {
					err = cliClient.GandalfUserService.Update(configurationCli.GetToken(), configurationCli.GetID(), user)
					if err != nil {
						fmt.Println(err)
					}
				} else {
					fmt.Println(err)
				}
			case "delete":
				err := cliClient.GandalfUserService.Delete(configurationCli.GetToken(), configurationCli.GetID())
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	} else if configurationCli.GetDatabaseMode() == "tenants" {
		switch configurationCli.GetModel() {
		case "authentication":
			var user models.User
			err := json.Unmarshal([]byte(configurationCli.GetValue()), &user)
			if err == nil {
				result, _ := cliClient.TenantsAuthenticationService.Login(configurationCli.GetTenant(), user)
				fmt.Println(result)
			} else {
				fmt.Println(err)
			}
		case "aggregator":
			switch configurationCli.GetCommand() {
			case "list":
				aggregators, err := cliClient.TenantsAggregatorService.List(configurationCli.GetToken(), configurationCli.GetTenant())
				if err == nil {
					fmt.Println(aggregators)
				} else {
					fmt.Println(err)
				}
			case "read":
				var aggregator *models.Aggregator
				aggregator, err := cliClient.TenantsAggregatorService.Read(configurationCli.GetToken(), configurationCli.GetTenant(), configurationCli.GetID())
				if err == nil {
					fmt.Println(aggregator)
				} else {
					fmt.Println(err)
				}
			case "create":
				var aggregator models.Aggregator
				err := json.Unmarshal([]byte(configurationCli.GetValue()), &aggregator)
				if err == nil {
					err = cliClient.TenantsAggregatorService.Create(configurationCli.GetToken(), configurationCli.GetTenant(), aggregator)
					if err != nil {
						fmt.Println(err)
					}
				} else {
					fmt.Println(err)
				}
			case "update":
				var aggregator models.Aggregator
				err := json.Unmarshal([]byte(configurationCli.GetValue()), &aggregator)
				if err == nil {
					err = cliClient.TenantsAggregatorService.Update(configurationCli.GetToken(), configurationCli.GetTenant(), configurationCli.GetID(), aggregator)
					if err != nil {
						fmt.Println(err)
					}
				} else {
					fmt.Println(err)
				}
			case "delete":
				err := cliClient.TenantsAggregatorService.Delete(configurationCli.GetToken(), configurationCli.GetTenant(), configurationCli.GetID())
				if err != nil {
					fmt.Println(err)
				}
			}
		case "connector":
			switch configurationCli.GetCommand() {
			case "list":
				connectors, err := cliClient.TenantsConnectorService.List(configurationCli.GetToken(), configurationCli.GetTenant())
				if err == nil {
					fmt.Println(connectors)
				} else {
					fmt.Println(err)
				}
			case "read":
				var connector *models.Connector
				connector, err := cliClient.TenantsConnectorService.Read(configurationCli.GetToken(), configurationCli.GetTenant(), configurationCli.GetID())
				if err == nil {
					fmt.Println(connector)
				} else {
					fmt.Println(err)
				}
			case "create":
				var connector models.Connector
				err := json.Unmarshal([]byte(configurationCli.GetValue()), &connector)
				if err == nil {
					err = cliClient.TenantsConnectorService.Create(configurationCli.GetToken(), configurationCli.GetTenant(), connector)
					if err != nil {
						fmt.Println(err)
					}
				} else {
					fmt.Println(err)
				}
			case "update":
				var connector models.Connector
				err := json.Unmarshal([]byte(configurationCli.GetValue()), &connector)
				if err == nil {
					err = cliClient.TenantsConnectorService.Update(configurationCli.GetToken(), configurationCli.GetTenant(), configurationCli.GetID(), connector)
					if err != nil {
						fmt.Println(err)
					}
				} else {
					fmt.Println(err)
				}
			case "delete":
				err := cliClient.TenantsConnectorService.Delete(configurationCli.GetToken(), configurationCli.GetTenant(), configurationCli.GetID())
				if err != nil {
					fmt.Println(err)
				}
			}
		case "role":
			switch configurationCli.GetCommand() {
			case "list":
				roles, err := cliClient.TenantsRoleService.List(configurationCli.GetToken(), configurationCli.GetTenant())
				if err == nil {
					fmt.Println(roles)
				} else {
					fmt.Println(err)
				}
			case "read":
				var role *models.Role
				role, err := cliClient.TenantsRoleService.Read(configurationCli.GetToken(), configurationCli.GetTenant(), configurationCli.GetID())
				if err == nil {
					fmt.Println(role)
				} else {
					fmt.Println(err)
				}
			case "create":
				var role models.Role
				err := json.Unmarshal([]byte(configurationCli.GetValue()), &role)
				if err == nil {
					err = cliClient.TenantsRoleService.Create(configurationCli.GetToken(), configurationCli.GetTenant(), role)
					if err != nil {
						fmt.Println(err)
					}
				} else {
					fmt.Println(err)
				}
			case "update":
				var role models.Role
				err := json.Unmarshal([]byte(configurationCli.GetValue()), &role)
				if err == nil {
					err = cliClient.TenantsRoleService.Update(configurationCli.GetToken(), configurationCli.GetTenant(), configurationCli.GetID(), role)
					if err != nil {
						fmt.Println(err)
					}
				} else {
					fmt.Println(err)
				}
			case "delete":
				err := cliClient.TenantsRoleService.Delete(configurationCli.GetToken(), configurationCli.GetTenant(), configurationCli.GetID())
				if err != nil {
					fmt.Println(err)
				}
			}
		case "user":
			switch configurationCli.GetCommand() {
			case "list":
				users, err := cliClient.TenantsUserService.List(configurationCli.GetToken(), configurationCli.GetTenant())
				if err == nil {
					fmt.Println(users)
				} else {
					fmt.Println(err)
				}
			case "read":
				var user *models.User
				user, err := cliClient.TenantsUserService.Read(configurationCli.GetToken(), configurationCli.GetTenant(), configurationCli.GetID())
				if err == nil {
					fmt.Println(user)
				} else {
					fmt.Println(err)
				}
			case "create":
				var user models.User
				err := json.Unmarshal([]byte(configurationCli.GetValue()), &user)
				if err == nil {
					err = cliClient.TenantsUserService.Create(configurationCli.GetToken(), configurationCli.GetTenant(), user)
					if err != nil {
						fmt.Println(err)
					}
				} else {
					fmt.Println(err)
				}
			case "update":
				var user models.User
				err := json.Unmarshal([]byte(configurationCli.GetValue()), &user)
				if err == nil {
					err = cliClient.TenantsUserService.Update(configurationCli.GetToken(), configurationCli.GetTenant(), configurationCli.GetID(), user)
					if err != nil {
						fmt.Println(err)
					}
				} else {
					fmt.Println(err)
				}
			case "delete":
				err := cliClient.TenantsUserService.Delete(configurationCli.GetToken(), configurationCli.GetTenant(), configurationCli.GetID())
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}
