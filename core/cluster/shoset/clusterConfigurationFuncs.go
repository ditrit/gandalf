//Package shoset :
package shoset

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	cmodels "github.com/ditrit/gandalf/core/configuration/models"
	"github.com/ditrit/gandalf/core/models"
	cmsg "github.com/ditrit/gandalf/core/msg"
	"github.com/jinzhu/gorm"

	"github.com/ditrit/gandalf/core/cluster/database"
	cutils "github.com/ditrit/gandalf/core/cluster/utils"

	net "github.com/ditrit/shoset"
	"github.com/ditrit/shoset/msg"
)

var configurationSendIndex = 0

func GetConfiguration(c *net.ShosetConn) (msg.Message, error) {
	var configuration cmsg.Configuration
	err := c.ReadMessage(&configuration)
	return configuration, err
}

// WaitConfig :
func WaitConfiguration(c *net.Shoset, replies *msg.Iterator, args map[string]string, timeout int) *msg.Message {
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
				config := message.(cmsg.Configuration)
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

// HandleConnectorConfig : Cluster handle connector config function.
func HandleConfiguration(c *net.ShosetConn, message msg.Message) (err error) {
	configuration := message.(cmsg.Configuration)
	ch := c.GetCh()

	err = nil

	log.Println("Handle configuration")
	log.Println(configuration)
	fmt.Println("HANDLE CONFIGURATION")
	fmt.Println(configuration)

	if configuration.GetCommand() == "PIVOT_CONFIGURATION" {

		componentType, ok := configuration.Context["componentType"].(string)
		if ok {
			databaseConnection, ok := ch.Context["databaseConnection"].(*database.DatabaseConnection)
			if ok {
				//mapDatabaseClient := ch.Context["tenantDatabases"].(map[string]*gorm.DB)
				//databaseBindAddr := ch.Context["databaseBindAddr"].(string)
				//configurationCluster := ch.Context["configuration"].(*cmodels.ConfigurationCluster)

				if databaseConnection != nil {
					//databaseClient := databaseConnection.GetDatabaseClientByTenant(configuration.GetTenant())
					var databaseClient *gorm.DB
					if componentType == "cluster" {
						databaseClient = databaseConnection.GetGandalfDatabaseClient()
					} else {
						databaseClient = databaseConnection.GetDatabaseClientByTenant(configuration.GetTenant())
					}
					if databaseClient != nil {
						//componentType := configuration.Context["componentType"].(string)
						var version models.Version
						jsonVersion, ok := configuration.Context["version"].([]byte)
						if ok {
							err := json.Unmarshal(jsonVersion, &version)
							if err == nil {
								pivot, err := cutils.GetPivots(databaseClient, componentType, version)
								if err == nil {
									jsonData, err := json.Marshal(pivot)

									if err == nil {
										switch componentType {
										case "cluster":
											bindaddr, ok := configuration.Context["bindAddress"].(string)
											if ok {
												cmdReply := cmsg.NewConfiguration("", "PIVOT_CONFIGURATION_REPLY", string(jsonData))
												cmdReply.Tenant = configuration.GetTenant()
												shoset := ch.ConnsJoin.Get(bindaddr)
												fmt.Println("shoset")
												fmt.Println(shoset)
												shoset.SendMessage(cmdReply)
											}

											break
										case "aggregator":
											cmdReply := cmsg.NewConfiguration("", "PIVOT_CONFIGURATION_REPLY", string(jsonData))
											cmdReply.Tenant = configuration.GetTenant()
											shoset := ch.ConnsByAddr.Get(c.GetBindAddr())

											shoset.SendMessage(cmdReply)
											break
										case "connector":
											cmdReply := cmsg.NewConfiguration(configuration.GetTarget(), "PIVOT_CONFIGURATION_REPLY", string(jsonData))
											cmdReply.Tenant = configuration.GetTenant()
											cmdReply.GetContext()["componentType"] = "connector"

											shoset := ch.ConnsByAddr.Get(c.GetBindAddr())

											shoset.SendMessage(cmdReply)
											break
										case "admin":
											cmdReply := cmsg.NewConfiguration(configuration.GetTarget(), "PIVOT_CONFIGURATION_REPLY", string(jsonData))
											cmdReply.Tenant = configuration.GetTenant()
											cmdReply.GetContext()["componentType"] = "admin"

											shoset := ch.ConnsByAddr.Get(c.GetBindAddr())

											shoset.SendMessage(cmdReply)
										default:
											cmdReply := cmsg.NewConfiguration(configuration.GetTarget(), "PIVOT_CONFIGURATION_REPLY", string(jsonData))
											cmdReply.Tenant = configuration.GetTenant()
											cmdReply.GetContext()["componentType"] = "worker"

											shoset := ch.ConnsByAddr.Get(c.GetBindAddr())

											shoset.SendMessage(cmdReply)
											break
										}

									} else {
										log.Println("Error : Can't unmarshall configuration")
									}
								} else {
									log.Println("Error : Can't find pivot")
								}
							} else {
								log.Println("Error : Can't unmarshall version")
							}
						}

					} else {
						log.Println("Error : Can't get database client by tenant")
					}
				} else {
					log.Println("Error : Can't get database clients")
				}
			}

		}

	} else if configuration.GetCommand() == "CONNECTOR_PRODUCT_CONFIGURATION" {

		/* 	componentType, ok := configuration.Context["componentType"].(string)
		if ok { */
		databaseConnection, ok := ch.Context["databaseConnection"].(*database.DatabaseConnection)
		if ok {
			//mapDatabaseClient := ch.Context["tenantDatabases"].(map[string]*gorm.DB)
			//databaseBindAddr := ch.Context["databaseBindAddr"].(string)
			//configurationCluster := ch.Context["configuration"].(*cmodels.ConfigurationCluster)

			if databaseConnection != nil {
				databaseClient := databaseConnection.GetDatabaseClientByTenant(configuration.GetTenant())
				/* 	var databaseClient *gorm.DB
				if componentType == "cluster" {
					databaseClient = databaseConnection.GetGandalfDatabaseClient()
				} else {
					databaseClient = databaseConnection.GetDatabaseClientByTenant(configuration.GetTenant())
				} */
				if databaseClient != nil {
					product, ok := configuration.Context["product"].(string)
					if ok {
						var version models.Version
						jsonVersion, ok := configuration.Context["version"].([]byte)
						if ok {
							err := json.Unmarshal(jsonVersion, &version)
							if err == nil {
								productConnectors, err := cutils.GetProductConnectors(databaseClient, product, version)
								fmt.Println("productConnectors")
								fmt.Println(productConnectors)
								if err == nil {
									jsonData, err := json.Marshal(productConnectors)

									if err == nil {
										cmdReply := cmsg.NewConfiguration(configuration.GetTarget(), "CONNECTOR_PRODUCT_CONFIGURATION_REPLY", string(jsonData))
										cmdReply.Tenant = configuration.GetTenant()
										shoset := ch.ConnsByAddr.Get(c.GetBindAddr())
										fmt.Println("shoset")
										fmt.Println(shoset)
										shoset.SendMessage(cmdReply)
									} else {
										log.Println("Error : Can't unmarshall configuration")
									}
								} else {
									log.Println("Error : Can't find product connector")
								}
							} else {
								log.Println("Error : Can't unmarshall version")
							}
						}
					}
				} else {
					log.Println("Error : Can't get database client by tenant")
				}
			} else {
				log.Println("Error : Can't get database clients")
			}
		}

		//}

	} else if configuration.GetCommand() == "PIVOT_CONFIGURATION_REPLY" {
		fmt.Println("REPLY")
		var pivot *models.Pivot
		err = json.Unmarshal([]byte(configuration.GetPayload()), &pivot)
		if err == nil {
			ch.Context["pivot"] = pivot
		}
	} else if configuration.GetCommand() == "SAVE_PIVOT_CONFIGURATION" {
		fmt.Println("SAVE")
		databaseConnection, ok := ch.Context["databaseConnection"].(*database.DatabaseConnection)
		if ok {
			//mapDatabaseClient := ch.Context["tenantDatabases"].(map[string]*gorm.DB)
			//databaseBindAddr := ch.Context["databaseBindAddr"].(string)
			//configurationCluster := ch.Context["configuration"].(*cmodels.ConfigurationCluster)

			if databaseConnection != nil {
				//databaseClient := databaseConnection.GetDatabaseClientByTenant(configuration.GetTenant())
				var databaseClient *gorm.DB
				databaseClient = databaseConnection.GetDatabaseClientByTenant(configuration.GetTenant())

				if databaseClient != nil {
					//connectorConfig := conf.GetContext()["connectorConfig"].(models.ConnectorConfig)
					var pivot models.Pivot
					err = json.Unmarshal([]byte(configuration.GetPayload()), &pivot)
					fmt.Println("PIVOT SAVE")
					fmt.Println(pivot)
					if err == nil {
						cutils.SavePivot(pivot, databaseClient)
					}
				} else {
					log.Println("Error : Can't get database client by tenant")
				}
			} else {
				log.Println("Error : Can't get database clients")
			}
		}

	} else if configuration.GetCommand() == "SAVE_PRODUCT_CONNECTOR_CONFIGURATION" {

		componentType, ok := configuration.Context["componentType"].(string)
		if ok {
			databaseConnection, ok := ch.Context["databaseConnection"].(*database.DatabaseConnection)
			if ok {
				//mapDatabaseClient := ch.Context["tenantDatabases"].(map[string]*gorm.DB)
				//databaseBindAddr := ch.Context["databaseBindAddr"].(string)
				//configurationCluster := ch.Context["configuration"].(*cmodels.ConfigurationCluster)

				if databaseConnection != nil {
					//databaseClient := databaseConnection.GetDatabaseClientByTenant(configuration.GetTenant())
					var databaseClient *gorm.DB
					if componentType == "cluster" {
						databaseClient = databaseConnection.GetGandalfDatabaseClient()
					} else {
						databaseClient = databaseConnection.GetDatabaseClientByTenant(configuration.GetTenant())
					}
					if databaseClient != nil {
						var productConnector *models.ProductConnector
						err = json.Unmarshal([]byte(configuration.GetPayload()), &productConnector)
						if err == nil {
							cutils.SaveProductConnector(productConnector, databaseClient)
						}
					} else {
						log.Println("Error : Can't get database client by tenant")
					}
				} else {
					log.Println("Error : Can't get database clients")
				}
				//connectorConfig := conf.GetContext()["connectorConfig"].(models.ConnectorConfig)
			}
		}
	}

	/* 	gandalfdatabaseClient := cutils.GetGandalfDatabaseClient(databasePath)
	   	if gandalfdatabaseClient != nil {

	   	} else {
	   		log.Println("Can't get gandalf database client")
	   		err = errors.New("Can't get gandalf database client")
	   	} */

	return err
}

//SendPivotConfiguration :
func SendClusterPivotConfiguration(shoset *net.Shoset) (err error) {
	version, ok := shoset.Context["version"].(models.Version)
	if ok {
		jsonVersion, err := json.Marshal(version)
		if err == nil {
			configurationCluster, ok := shoset.Context["configuration"].(*cmodels.ConfigurationCluster)
			if ok {
				conf := cmsg.NewConfiguration("", "PIVOT_CONFIGURATION", "")
				conf.GetContext()["componentType"] = "cluster"
				conf.GetContext()["version"] = jsonVersion
				conf.GetContext()["bindAddress"] = configurationCluster.GetBindAddress()

				//conf.GetContext()["product"] = shoset.Context["product"]

				shosets := net.GetByType(shoset.ConnsJoin, "")

				if len(shosets) != 0 {
					if conf.GetTimeout() > configurationCluster.GetMaxTimeout() {
						conf.Timeout = configurationCluster.GetMaxTimeout()
					}

					notSend := true
					for start := time.Now(); time.Since(start) < time.Duration(conf.GetTimeout())*time.Millisecond; {
						index := getConfigurationSendIndex(shosets)
						shosets[index].SendMessage(conf)
						log.Printf("%s : send command %s to %s\n", shoset.GetBindAddr(), conf.GetCommand(), shosets[index])

						timeoutSend := time.Duration((int(conf.GetTimeout()) / len(shosets)))

						time.Sleep(timeoutSend * time.Millisecond)

						if shoset.Context["pivot"] != nil {
							notSend = false
							break
						}
					}

					if notSend {
						return nil
					}
				} else {
					log.Println("Error : Can't find aggregators to send")
				}
			}
		} else {
			log.Println("Error : Can't marshall version")
		}
	}

	return err
}

// getCommandSendIndex : Aggregator getSendIndex function.
func getConfigurationSendIndex(conns []*net.ShosetConn) int {
	if configurationSendIndex >= len(conns) {
		configurationSendIndex = 0
	}

	aux := configurationSendIndex
	configurationSendIndex++

	return aux
}
