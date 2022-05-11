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
				if databaseConnection != nil {
					var databaseClient *gorm.DB
					if componentType == "cluster" {
						databaseClient = databaseConnection.GetGandalfDatabaseClient()
					} else {
						databaseClient = databaseConnection.GetDatabaseClientByTenant(configuration.GetTenant())
					}
					if databaseClient != nil {
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
												cmdReply := cmsg.NewConfiguration("PIVOT_CONFIGURATION_REPLY", string(jsonData))
												cmdReply.Tenant = configuration.GetTenant()

												var shoset *net.ShosetConn
												connsJoin := ch.ConnsByName.Get(ch.GetLogicalName())
												if connsJoin != nil {
													shoset = connsJoin.Get(bindaddr)
												}

												fmt.Println("shoset")
												fmt.Println(shoset)
												shoset.SendMessage(cmdReply)
											}

											break
										case "aggregator":
											cmdReply := cmsg.NewConfiguration("PIVOT_CONFIGURATION_REPLY", string(jsonData))
											cmdReply.Tenant = configuration.GetTenant()

											mapshoset := ch.ConnsByName.Get(c.GetRemoteLogicalName())
											var shoset *net.ShosetConn
											if mapshoset != nil {
												shoset = mapshoset.Get(c.GetRemoteAddress())
											}

											shoset.SendMessage(cmdReply)
											break
										case "connector":
											cmdReply := cmsg.NewConfiguration("PIVOT_CONFIGURATION_REPLY", string(jsonData))
											cmdReply.Tenant = configuration.GetTenant()
											cmdReply.TargetLogicalName = configuration.GetTargetLogicalName()
											cmdReply.TargetAddress = configuration.GetTargetAddress()
											cmdReply.GetContext()["componentType"] = "connector"

											mapshoset := ch.ConnsByName.Get(c.GetRemoteLogicalName())
											var shoset *net.ShosetConn
											if mapshoset != nil {
												shoset = mapshoset.Get(c.GetRemoteAddress())
											}

											shoset.SendMessage(cmdReply)
											break
										case "admin":
											cmdReply := cmsg.NewConfiguration("PIVOT_CONFIGURATION_REPLY", string(jsonData))
											cmdReply.Tenant = configuration.GetTenant()
											cmdReply.TargetLogicalName = configuration.GetTargetLogicalName()
											cmdReply.TargetAddress = configuration.GetTargetAddress()
											cmdReply.GetContext()["componentType"] = "admin"

											mapshoset := ch.ConnsByName.Get(c.GetRemoteLogicalName())
											var shoset *net.ShosetConn
											if mapshoset != nil {
												shoset = mapshoset.Get(c.GetRemoteAddress())
											}

											shoset.SendMessage(cmdReply)
										default:
											cmdReply := cmsg.NewConfiguration("PIVOT_CONFIGURATION_REPLY", string(jsonData))
											cmdReply.Tenant = configuration.GetTenant()
											cmdReply.TargetLogicalName = configuration.GetTargetLogicalName()
											cmdReply.TargetAddress = configuration.GetTargetAddress()
											cmdReply.GetContext()["componentType"] = "worker"

											mapshoset := ch.ConnsByName.Get(c.GetRemoteLogicalName())
											var shoset *net.ShosetConn
											if mapshoset != nil {
												shoset = mapshoset.Get(c.GetRemoteAddress())
											}

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
		databaseConnection, ok := ch.Context["databaseConnection"].(*database.DatabaseConnection)
		if ok {

			if databaseConnection != nil {
				databaseClient := databaseConnection.GetDatabaseClientByTenant(configuration.GetTenant())
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
										cmdReply := cmsg.NewConfiguration("CONNECTOR_PRODUCT_CONFIGURATION_REPLY", string(jsonData))
										cmdReply.TargetLogicalName = configuration.GetTargetLogicalName()
										cmdReply.TargetAddress = configuration.GetTargetAddress()
										cmdReply.Tenant = configuration.GetTenant()

										mapshoset := ch.ConnsByName.Get(c.GetRemoteLogicalName())
										var shoset *net.ShosetConn
										if mapshoset != nil {
											shoset = mapshoset.Get(c.GetRemoteAddress())
										}

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

			if databaseConnection != nil {
				var databaseClient *gorm.DB
				databaseClient = databaseConnection.GetDatabaseClientByTenant(configuration.GetTenant())

				if databaseClient != nil {
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

				if databaseConnection != nil {
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
			}
		}
	}
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
				conf := cmsg.NewConfiguration("PIVOT_CONFIGURATION", "")
				conf.GetContext()["componentType"] = "cluster"
				conf.GetContext()["version"] = jsonVersion
				conf.GetContext()["bindAddress"] = configurationCluster.GetBindAddress()

				var shosets []*net.ShosetConn
				connsJoin := shoset.ConnsByName.Get(shoset.GetLogicalName())
				if connsJoin != nil {
					shosets = net.GetByType(connsJoin, "")

				}

				if len(shosets) != 0 {
					if conf.GetTimeout() > configurationCluster.GetMaxTimeout() {
						conf.Timeout = configurationCluster.GetMaxTimeout()
					}

					notSend := true
					for start := time.Now(); time.Since(start) < time.Duration(conf.GetTimeout())*time.Millisecond; {
						index := getConfigurationSendIndex(shosets)
						shosets[index].SendMessage(conf)
						log.Printf("%s : send command %s to %s\n", shoset.GetBindAddress(), conf.GetCommand(), shosets[index])

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
