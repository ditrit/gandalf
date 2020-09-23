package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"strconv"

	"gandalf/cli/client"

	cmodels "github.com/ditrit/gandalf/core/models"
)

func main() {

	var agent string
	flag.StringVar(&agent, "agent", "gandalf", "a string var")

	var typeDB string
	flag.StringVar(&typeDB, "typeDB", "", "a string var")

	var tenant string
	flag.StringVar(&tenant, "tenant", "", "a string var")

	var models string
	flag.StringVar(&models, "models", "", "a string var")

	var command string
	flag.StringVar(&command, "command", "", "a string var")

	var token string
	flag.StringVar(&token, "token", "", "a string var")

	var id string
	flag.StringVar(&id, "id", "", "a string var")

	var value string
	flag.StringVar(&value, "value", "", "a string var")

	flag.Parse()

	fmt.Println("agent:", agent)
	fmt.Println("typeDB:", typeDB)
	fmt.Println("tenant:", tenant)
	fmt.Println("models:", models)
	fmt.Println("command:", command)
	fmt.Println("token:", token)
	fmt.Println("value:", value)

	cliClient := client.NewClient(agent)

	if typeDB == "gandalf" {
		switch models {
		case "authentication":
			var user cmodels.User
			err := json.Unmarshal([]byte(value), &user)
			if err == nil {
				result, _ := cliClient.GandalfAuthenticationService.Login(user)
				fmt.Println(result)
			} else {
				fmt.Println(err)
			}
		case "cluster":
			switch command {
			case "list":
				clusters, err := cliClient.GandalfClusterService.List(token)
				if err == nil {
					fmt.Println(clusters)
				} else {
					fmt.Println(err)
				}
			case "read":
				var cluster *cmodels.Cluster
				intId, err := strconv.Atoi(id)
				if err == nil {
					cluster, err = cliClient.GandalfClusterService.Read(token, intId)
					if err == nil {
						fmt.Println(cluster)
					} else {
						fmt.Println(err)
					}
				} else {
					fmt.Println(err)
				}
			case "create":
				var cluster cmodels.Cluster
				err := json.Unmarshal([]byte(value), &cluster)
				if err == nil {
					err = cliClient.GandalfClusterService.Create(token, cluster)
					if err != nil {
						fmt.Println(err)
					}
				} else {
					fmt.Println(err)
				}
			case "update":
				var cluster cmodels.Cluster
				err := json.Unmarshal([]byte(value), &cluster)
				if err == nil {
					intId, err := strconv.Atoi(id)
					if err == nil {
						err = cliClient.GandalfClusterService.Update(token, intId, cluster)
						if err != nil {
							fmt.Println(err)
						}
					} else {
						fmt.Println(err)
					}
				} else {
					fmt.Println(err)
				}
			case "delete":
				intId, err := strconv.Atoi(id)
				if err == nil {
					err := cliClient.GandalfClusterService.Delete(token, intId)
					if err != nil {
						fmt.Println(err)
					}
				} else {
					fmt.Println(err)
				}
			}
		case "tenant":
			switch command {
			case "list":
				tenants, err := cliClient.GandalfTenantService.List(token)
				if err == nil {
					fmt.Println(tenants)
				} else {
					fmt.Println(err)
				}
			case "read":
				var tenant *cmodels.Tenant
				intId, err := strconv.Atoi(id)
				if err == nil {
					tenant, err = cliClient.GandalfTenantService.Read(token, intId)
					if err == nil {
						fmt.Println(tenant)
					} else {
						fmt.Println(err)
					}
				} else {
					fmt.Println(err)
				}
			case "create":
				var tenant cmodels.Tenant
				err := json.Unmarshal([]byte(value), &tenant)
				if err == nil {
					var login, password string
					login, password, err = cliClient.GandalfTenantService.Create(token, tenant)
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
				var tenant cmodels.Tenant
				err := json.Unmarshal([]byte(value), &tenant)
				if err == nil {
					intId, err := strconv.Atoi(id)
					if err == nil {
						err = cliClient.GandalfTenantService.Update(token, intId, tenant)
						if err != nil {
							fmt.Println(err)
						}
					} else {
						fmt.Println(err)
					}
				} else {
					fmt.Println(err)
				}
			case "delete":
				intId, err := strconv.Atoi(id)
				if err == nil {
					err := cliClient.GandalfTenantService.Delete(token, intId)
					if err != nil {
						fmt.Println(err)
					}
				} else {
					fmt.Println(err)
				}
			}
		case "role":
			switch command {
			case "list":
				roles, err := cliClient.GandalfRoleService.List(token)
				if err == nil {
					fmt.Println(roles)
				} else {
					fmt.Println(err)
				}
			case "read":
				var role *cmodels.Role
				intId, err := strconv.Atoi(id)
				if err == nil {
					role, err = cliClient.GandalfRoleService.Read(token, intId)
					if err == nil {
						fmt.Println(role)
					} else {
						fmt.Println(err)
					}
				} else {
					fmt.Println(err)
				}
			case "create":
				var role cmodels.Role
				err := json.Unmarshal([]byte(value), &role)
				if err == nil {
					err = cliClient.GandalfRoleService.Create(token, role)
					if err != nil {
						fmt.Println(err)
					}
				} else {
					fmt.Println(err)
				}
			case "update":
				var role cmodels.Role
				err := json.Unmarshal([]byte(value), &role)
				if err == nil {
					intId, err := strconv.Atoi(id)
					if err == nil {
						err = cliClient.GandalfRoleService.Update(token, intId, role)
						if err != nil {
							fmt.Println(err)
						}
					} else {
						fmt.Println(err)
					}
				} else {
					fmt.Println(err)
				}
			case "delete":
				intId, err := strconv.Atoi(id)
				if err == nil {
					err := cliClient.GandalfRoleService.Delete(token, intId)
					if err != nil {
						fmt.Println(err)
					}
				} else {
					fmt.Println(err)
				}
			}
		case "user":
			switch command {
			case "list":
				users, err := cliClient.GandalfUserService.List(token)
				if err == nil {
					fmt.Println(users)
				} else {
					fmt.Println(err)
				}
			case "read":
				var user *cmodels.User
				intId, err := strconv.Atoi(id)
				if err == nil {
					user, err = cliClient.GandalfUserService.Read(token, intId)
					if err == nil {
						fmt.Println(user)
					} else {
						fmt.Println(err)
					}
				} else {
					fmt.Println(err)
				}
			case "create":
				var user cmodels.User
				err := json.Unmarshal([]byte(value), &user)
				if err == nil {
					err = cliClient.GandalfUserService.Create(token, user)
					if err != nil {
						fmt.Println(err)
					}
				} else {
					fmt.Println(err)
				}
			case "update":
				var user cmodels.User
				err := json.Unmarshal([]byte(value), &user)
				if err == nil {
					intId, err := strconv.Atoi(id)
					if err == nil {
						err = cliClient.GandalfUserService.Update(token, intId, user)
						if err != nil {
							fmt.Println(err)
						}
					} else {
						fmt.Println(err)
					}
				} else {
					fmt.Println(err)
				}
			case "delete":
				intId, err := strconv.Atoi(id)
				if err == nil {
					err := cliClient.GandalfUserService.Delete(token, intId)
					if err != nil {
						fmt.Println(err)
					}
				} else {
					fmt.Println(err)
				}
			}
		}
	} else if typeDB == "tenants" {
		switch models {
		case "authentication":
			var user cmodels.User
			err := json.Unmarshal([]byte(value), &user)
			if err == nil {
				result, _ := cliClient.TenantsAuthenticationService.Login(tenant, user)
				fmt.Println(result)
			} else {
				fmt.Println(err)
			}
		case "aggregator":
			switch command {
			case "list":
				aggregators, err := cliClient.TenantsAggregatorService.List(token, tenant)
				if err == nil {
					fmt.Println(aggregators)
				} else {
					fmt.Println(err)
				}
			case "read":
				var aggregator *cmodels.Aggregator
				intId, err := strconv.Atoi(id)
				if err == nil {
					aggregator, err = cliClient.TenantsAggregatorService.Read(token, tenant, intId)
					if err == nil {
						fmt.Println(aggregator)
					} else {
						fmt.Println(err)
					}
				} else {
					fmt.Println(err)
				}
			case "create":
				var aggregator cmodels.Aggregator
				err := json.Unmarshal([]byte(value), &aggregator)
				if err == nil {
					err = cliClient.TenantsAggregatorService.Create(token, tenant, aggregator)
					if err != nil {
						fmt.Println(err)
					}
				} else {
					fmt.Println(err)
				}
			case "update":
				var aggregator cmodels.Aggregator
				err := json.Unmarshal([]byte(value), &aggregator)
				if err == nil {
					intId, err := strconv.Atoi(id)
					if err == nil {
						err = cliClient.TenantsAggregatorService.Update(token, tenant, intId, aggregator)
						if err != nil {
							fmt.Println(err)
						}
					} else {
						fmt.Println(err)
					}
				} else {
					fmt.Println(err)
				}
			case "delete":
				intId, err := strconv.Atoi(id)
				if err == nil {
					err := cliClient.TenantsAggregatorService.Delete(token, tenant, intId)
					if err != nil {
						fmt.Println(err)
					}
				} else {
					fmt.Println(err)
				}
			}
		case "connector":
			switch command {
			case "list":
				connectors, err := cliClient.TenantsConnectorService.List(token, tenant)
				if err == nil {
					fmt.Println(connectors)
				} else {
					fmt.Println(err)
				}
			case "read":
				var connector *cmodels.Connector
				intId, err := strconv.Atoi(id)
				if err == nil {
					connector, err = cliClient.TenantsConnectorService.Read(token, tenant, intId)
					if err == nil {
						fmt.Println(connector)
					} else {
						fmt.Println(err)
					}
				} else {
					fmt.Println(err)
				}
			case "create":
				var connector cmodels.Connector
				err := json.Unmarshal([]byte(value), &connector)
				if err == nil {
					err = cliClient.TenantsConnectorService.Create(token, tenant, connector)
					if err != nil {
						fmt.Println(err)
					}
				} else {
					fmt.Println(err)
				}
			case "update":
				var connector cmodels.Connector
				err := json.Unmarshal([]byte(value), &connector)
				if err == nil {
					intId, err := strconv.Atoi(id)
					if err == nil {
						err = cliClient.TenantsConnectorService.Update(token, tenant, intId, connector)
						if err != nil {
							fmt.Println(err)
						}
					} else {
						fmt.Println(err)
					}
				} else {
					fmt.Println(err)
				}
			case "delete":
				intId, err := strconv.Atoi(id)
				if err == nil {
					err := cliClient.TenantsConnectorService.Delete(token, tenant, intId)
					if err != nil {
						fmt.Println(err)
					}
				} else {
					fmt.Println(err)
				}
			}
		case "role":
			switch command {
			case "list":
				roles, err := cliClient.TenantsRoleService.List(token, tenant)
				if err == nil {
					fmt.Println(roles)
				} else {
					fmt.Println(err)
				}
			case "read":
				var role *cmodels.Role
				intId, err := strconv.Atoi(id)
				if err == nil {
					role, err = cliClient.TenantsRoleService.Read(token, tenant, intId)
					if err == nil {
						fmt.Println(role)
					} else {
						fmt.Println(err)
					}
				} else {
					fmt.Println(err)
				}
			case "create":
				var role cmodels.Role
				err := json.Unmarshal([]byte(value), &role)
				if err == nil {
					err = cliClient.TenantsRoleService.Create(token, tenant, role)
					if err != nil {
						fmt.Println(err)
					}
				} else {
					fmt.Println(err)
				}
			case "update":
				var role cmodels.Role
				err := json.Unmarshal([]byte(value), &role)
				if err == nil {
					intId, err := strconv.Atoi(id)
					if err == nil {
						err = cliClient.TenantsRoleService.Update(token, tenant, intId, role)
						if err != nil {
							fmt.Println(err)
						}
					} else {
						fmt.Println(err)
					}
				} else {
					fmt.Println(err)
				}
			case "delete":
				intId, err := strconv.Atoi(id)
				if err == nil {
					err := cliClient.TenantsRoleService.Delete(token, tenant, intId)
					if err != nil {
						fmt.Println(err)
					}
				} else {
					fmt.Println(err)
				}
			}
		case "user":
			switch command {
			case "list":
				users, err := cliClient.TenantsUserService.List(token, tenant)
				if err == nil {
					fmt.Println(users)
				} else {
					fmt.Println(err)
				}
			case "read":
				var user *cmodels.User
				intId, err := strconv.Atoi(id)
				if err == nil {
					user, err = cliClient.TenantsUserService.Read(token, tenant, intId)
					if err == nil {
						fmt.Println(user)
					} else {
						fmt.Println(err)
					}
				} else {
					fmt.Println(err)
				}
			case "create":
				var user cmodels.User
				err := json.Unmarshal([]byte(value), &user)
				if err == nil {
					err = cliClient.TenantsUserService.Create(token, tenant, user)
					if err != nil {
						fmt.Println(err)
					}
				} else {
					fmt.Println(err)
				}
			case "update":
				var user cmodels.User
				err := json.Unmarshal([]byte(value), &user)
				if err == nil {
					intId, err := strconv.Atoi(id)
					if err == nil {
						err = cliClient.TenantsUserService.Update(token, tenant, intId, user)
						if err != nil {
							fmt.Println(err)
						}
					} else {
						fmt.Println(err)
					}
				} else {
					fmt.Println(err)
				}
			case "delete":
				intId, err := strconv.Atoi(id)
				if err == nil {
					err := cliClient.TenantsUserService.Delete(token, tenant, intId)
					if err != nil {
						fmt.Println(err)
					}
				} else {
					fmt.Println(err)
				}
			}
		}
	}
}
