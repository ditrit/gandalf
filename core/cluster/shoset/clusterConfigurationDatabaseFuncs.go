//Package shoset :
package shoset

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/ditrit/gandalf/core/cluster/database"

	"github.com/ditrit/gandalf/core/models"

	cutils "github.com/ditrit/gandalf/core/cluster/utils"

	cmsg "github.com/ditrit/gandalf/core/msg"
	"github.com/ditrit/shoset/msg"

	net "github.com/ditrit/shoset"
	"github.com/jinzhu/gorm"
)

func GetConfigurationDatabase(c *net.ShosetConn) (msg.Message, error) {
	var configurationDatabase cmsg.ConfigurationDatabase
	err := c.ReadMessage(&configurationDatabase)
	return configurationDatabase, err
}

// WaitConfig :
func WaitConfigurationDatabase(c *net.Shoset, replies *msg.Iterator, args map[string]string, timeout int) *msg.Message {
	commandName, ok := args["name"]
	if !ok {
		return nil
	}
	term := make(chan *msg.Message, 1)
	cont := true
	go func() {
		for cont {
			message := replies.Get().GetMessage()
			if message != nil {
				configurationDatabase := message.(cmsg.ConfigurationDatabase)
				if configurationDatabase.GetCommand() == commandName {
					term <- &message
				}
			} else {
				time.Sleep(time.Duration(10) * time.Millisecond)
			}
		}
	}()
	select {
	case res := <-term:
		cont = false
		return res
	case <-time.After(time.Duration(timeout) * time.Second):
		return nil
	}
}

// HandleSecret :
func HandleConfigurationDatabase(c *net.ShosetConn, message msg.Message) (err error) {
	configurationDb := message.(cmsg.ConfigurationDatabase)
	ch := c.GetCh()

	err = nil

	log.Println("Handle configuration database")
	log.Println(configurationDb)

	fmt.Println("CONFIGURATION_DATABASE")
	if configurationDb.GetCommand() == "CONFIGURATION_DATABASE" {
		var databaseClient *gorm.DB
		databaseConnection, ok := ch.Context["databaseConnection"].(*database.DatabaseConnection)
		if ok {
			if databaseConnection != nil {
				databaseClient = databaseConnection.GetGandalfDatabaseClient()
				if databaseClient != nil {
					tenant, err := cutils.GetTenant(configurationDb.GetTenant(), databaseClient)

					if err == nil {
						configurationDatabaseAggregator := models.NewConfigurationDatabaseAggregator(tenant.Name, tenant.Password, databaseConnection.GetConfigurationCluster().GetDatabaseBindAddress())
						configMarshal, err := json.Marshal(configurationDatabaseAggregator)
						if err == nil {
							target := ""
							configurationReply := cmsg.NewConfigurationDatabase(target, "CONFIGURATION_DATABASE_REPLY", string(configMarshal))
							configurationReply.Tenant = configurationDb.GetTenant()

							mapshoset := ch.ConnsByName.Get(c.GetRemoteLogicalName())
							var shoset *net.ShosetConn
							if mapshoset != nil {
								shoset = mapshoset.Get(c.GetRemoteAddress())
							}
							shoset.SendMessage(configurationReply)
						}
					} else {
						log.Println("Error : Can't get tenant " + configurationDb.Tenant)
					}
				} else {
					log.Println("Error : Can't get database client")
				}
			} else {
				log.Println("Error : Database connection is empty")
			}
		}

	} else if configurationDb.GetCommand() == "CREATE_DATABASE" {
		fmt.Println("CREATE")

		var databaseClient *gorm.DB
		databaseConnection, ok := ch.Context["databaseConnection"].(*database.DatabaseConnection)
		if ok {
			if databaseConnection != nil {
				fmt.Println("CREATE1")
				databaseClient = databaseConnection.GetGandalfDatabaseClient()
				if databaseClient != nil {
					fmt.Println("CREATE2")
					tenant, err := cutils.GetTenant(configurationDb.GetPayload(), databaseClient)
					fmt.Println(err)
					if err == nil {
						fmt.Println("CREATE3")
						err = databaseConnection.NewDatabase(tenant.Name, tenant.Password)
						fmt.Println(err)
						if err == nil {
							fmt.Println("CREATE4")
							tenantDatabaseClient := databaseConnection.GetDatabaseClientByTenant(tenant.Name)

							if tenantDatabaseClient != nil {
								fmt.Println("CREATE5")
								var login, password []string
								login, password, err = databaseConnection.InitTenantDatabase(tenantDatabaseClient)
								if err == nil {
									fmt.Println("CREATE6")
									createDatabase := models.NewCreateDatabase(login, password)
									configMarshal, err := json.Marshal(createDatabase)
									if err == nil {
										target := ""
										creationReply := cmsg.NewConfigurationDatabase(target, "CREATE_DATABASE_REPLY", string(configMarshal))
										creationReply.Tenant = configurationDb.GetTenant()

										mapshoset := ch.ConnsByName.Get(c.GetRemoteLogicalName())
										var shoset *net.ShosetConn
										if mapshoset != nil {
											shoset = mapshoset.Get(c.GetRemoteAddress())
										}

										fmt.Println("SEND")
										shoset.SendMessage(creationReply)
									}
								} else {
									log.Println("Error : Can't init database")
								}
							} else {
								log.Println("Error : Can't get database client by tenant")
							}
						} else {
							log.Println("Error : Can't create database")
						}
					} else {
						log.Println("Error : Can't get tenant " + configurationDb.GetPayload())
					}
				} else {
					log.Println("Error : Can't get database client")
				}
			} else {
				log.Println("Error : Database connection is empty")
			}
		}
	}

	return err
}
