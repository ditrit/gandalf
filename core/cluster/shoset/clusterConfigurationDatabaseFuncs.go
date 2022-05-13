//Package shoset :
package shoset

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
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

	log.Info().Msg("handle configuration database")

	if configurationDb.GetCommand() == "CONFIGURATION_DATABASE" {
		log.Info().Msg("configure database")
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
						log.Error().Err(err).Str("tenant", configurationDb.Tenant).Msg("can't get tenant")
					}
				} else {
					log.Error().Err(err).Msg("can't get database client")
				}
			} else {
				log.Error().Err(err).Msg("database connection is empty")
			}
		}

	} else if configurationDb.GetCommand() == "CREATE_DATABASE" {
		log.Info().Msg("create database")

		var databaseClient *gorm.DB
		databaseConnection, ok := ch.Context["databaseConnection"].(*database.DatabaseConnection)
		if ok {
			if databaseConnection != nil {
				databaseClient = databaseConnection.GetGandalfDatabaseClient()
				if databaseClient != nil {
					tenant, err := cutils.GetTenant(configurationDb.GetPayload(), databaseClient)
					if err == nil {
						err = databaseConnection.NewDatabase(tenant.Name, tenant.Password)
						if err == nil {
							tenantDatabaseClient := databaseConnection.GetDatabaseClientByTenant(tenant.Name)

							if tenantDatabaseClient != nil {
								var login, password []string
								login, password, err = databaseConnection.InitTenantDatabase(tenantDatabaseClient)
								if err == nil {
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

										log.Info().Msg("send")
										shoset.SendMessage(creationReply)
									}
								} else {
									log.Error().Err(err).Msg("can't init database")
								}
							} else {
								log.Error().Err(err).Msg("can't get database client by tenant")
							}
						} else {
							log.Error().Err(err).Msg("can't create database")
						}
					} else {
						log.Error().Err(err).Str("payload", configurationDb.GetPayload()).Msg("can't get tenant")
					}
				} else {
					log.Error().Err(err).Msg("can't get database client")
				}
			} else {
				log.Error().Err(err).Msg("database connection is empty")
			}
		}
	}

	return err
}
