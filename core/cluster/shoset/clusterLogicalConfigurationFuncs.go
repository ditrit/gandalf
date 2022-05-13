//Package shoset :
package shoset

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"time"

	"github.com/ditrit/gandalf/core/cluster/database"

	cmodels "github.com/ditrit/gandalf/core/configuration/models"
	"github.com/ditrit/gandalf/core/models"

	cutils "github.com/ditrit/gandalf/core/cluster/utils"

	cmsg "github.com/ditrit/gandalf/core/msg"
	"github.com/ditrit/shoset/msg"

	net "github.com/ditrit/shoset"
	"github.com/jinzhu/gorm"
)

var logicalConfigurationSendIndex = 0

func GetLogicalConfiguration(c *net.ShosetConn) (msg.Message, error) {
	var logicalConfiguration cmsg.LogicalConfiguration
	err := c.ReadMessage(&logicalConfiguration)
	return logicalConfiguration, err
}

// WaitConfig :
func WaitLogicalConfiguration(c *net.Shoset, replies *msg.Iterator, args map[string]string, timeout int) *msg.Message {
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
				config := message.(cmsg.LogicalConfiguration)
				if config.GetCommand() == commandName {
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
func HandleLogicalConfiguration(c *net.ShosetConn, message msg.Message) (err error) {
	logicalConfiguration := message.(cmsg.LogicalConfiguration)
	ch := c.GetCh()

	err = nil

	log.Info().Msg("Handle logical configuration")

	if logicalConfiguration.GetCommand() == "LOGICAL_CONFIGURATION" {
		var databaseClient *gorm.DB
		componentType, ok := logicalConfiguration.GetContext()["componentType"].(string)
		if ok {
			databaseConnection, ok := ch.Context["databaseConnection"].(*database.DatabaseConnection)
			if ok {
				if databaseConnection != nil {
					if componentType == "cluster" {
						databaseClient = databaseConnection.GetGandalfDatabaseClient()
					} else {
						databaseClient = databaseConnection.GetDatabaseClientByTenant(logicalConfiguration.GetTenant())
					}

					if databaseClient != nil {
						logicalName, ok := logicalConfiguration.GetContext()["logicalName"].(string)
						if ok {
							logicalComponent, err := cutils.GetLogicalComponents(databaseClient, logicalName)
							if err == nil {
								jsonData, err := json.Marshal(logicalComponent)
								if err == nil {
									switch componentType {
									case "cluster":
										bindAddr, ok := logicalConfiguration.GetContext()["bindAddress"].(string)
										if ok {
											configurationReply := cmsg.NewLogicalConfiguration("LOGICAL_CONFIGURATION_REPLY", string(jsonData))
											configurationReply.Tenant = logicalConfiguration.GetTenant()

											var shoset *net.ShosetConn
											connsJoin := ch.ConnsByName.Get(ch.GetLogicalName())
											if connsJoin != nil {
												shoset = connsJoin.Get(bindAddr)
											}

											shoset.SendMessage(configurationReply)
										}
										break
									case "aggregator":
										configurationReply := cmsg.NewLogicalConfiguration("LOGICAL_CONFIGURATION_REPLY", string(jsonData))
										configurationReply.Tenant = logicalConfiguration.GetTenant()

										mapshoset := ch.ConnsByName.Get(c.GetRemoteLogicalName())
										var shoset *net.ShosetConn
										if mapshoset != nil {
											shoset = mapshoset.Get(c.GetRemoteAddress())
										}

										shoset.SendMessage(configurationReply)
										break
									case "connector":
										configurationReply := cmsg.NewLogicalConfiguration("LOGICAL_CONFIGURATION_REPLY", string(jsonData))
										configurationReply.TargetLogicalName = logicalConfiguration.GetTargetLogicalName()
										configurationReply.TargetAddress = logicalConfiguration.GetTargetAddress()
										configurationReply.Tenant = logicalConfiguration.GetTenant()

										mapshoset := ch.ConnsByName.Get(c.GetRemoteLogicalName())
										var shoset *net.ShosetConn
										if mapshoset != nil {
											shoset = mapshoset.Get(c.GetRemoteAddress())
										}
										shoset.SendMessage(configurationReply)
										break
									}
								} else {
									log.Error().Err(err).Msg("can't unmarshall configuration")
								}
							} else {
								log.Error().Err(err).Msg("can't find logical component")
							}
						}
					} else {
						log.Error().Err(err).Msg("can't get database client")
					}
				} else {
					log.Error().Err(err).Msg("Error : Database connection is empty")
				}
			}
		}

	} else if logicalConfiguration.GetCommand() == "LOGICAL_CONFIGURATION_REPLY" {
		var logicalComponents *models.LogicalComponent
		err = json.Unmarshal([]byte(logicalConfiguration.GetPayload()), &logicalComponents)
		if err == nil {
			ch.Context["logicalConfiguration"] = logicalComponents
		}
	}

	return err
}

//SendSecret :
func SendLogicalConfiguration(shoset *net.Shoset) (err error) {
	configurationCluster := shoset.Context["configuration"].(*cmodels.ConfigurationCluster)

	if err == nil {
		configurationMsg := cmsg.NewLogicalConfiguration("LOGICAL_CONFIGURATION", "")
		configurationMsg.GetContext()["componentType"] = "cluster"
		configurationMsg.GetContext()["logicalName"] = configurationCluster.GetLogicalName()
		configurationMsg.GetContext()["bindAddress"] = configurationCluster.GetBindAddress()

		var shosets []*net.ShosetConn
		connsJoin := shoset.ConnsByName.Get(shoset.GetLogicalName())
		if connsJoin != nil {
			shosets = net.GetByType(connsJoin, "")

		}

		if len(shosets) != 0 {
			if configurationMsg.GetTimeout() > configurationCluster.GetMaxTimeout() {
				configurationMsg.Timeout = configurationCluster.GetMaxTimeout()
			}

			notSend := true
			for start := time.Now(); time.Since(start) < time.Duration(configurationMsg.GetTimeout())*time.Millisecond; {
				index := getLogicalConfigurationSendIndex(shosets)
				shosets[index].SendMessage(configurationMsg)
				log.Printf("%s : send command %s to %s\n", shoset.GetBindAddress(), configurationMsg.GetCommand(), shosets[index])

				timeoutSend := time.Duration((int(configurationMsg.GetTimeout()) / len(shosets)))

				time.Sleep(timeoutSend * time.Millisecond)

				if shoset.Context["logicalConfiguration"] != nil {
					notSend = false
					break
				}
			}

			if notSend {
				return nil
			}

		} else {
			log.Error().Err(err).Msg("can't find cluster to send")
		}
	}

	return err
}

// getCommandSendIndex : Aggregator getSendIndex function.
func getLogicalConfigurationSendIndex(conns []*net.ShosetConn) int {
	if logicalConfigurationSendIndex >= len(conns) {
		logicalConfigurationSendIndex = 0
	}

	aux := logicalConfigurationSendIndex
	logicalConfigurationSendIndex++

	return aux
}
