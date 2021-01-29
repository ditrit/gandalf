//Package shoset :
package shoset

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/ditrit/gandalf/core/models"

	cutils "github.com/ditrit/gandalf/core/cluster/utils"

	cmsg "github.com/ditrit/gandalf/core/msg"
	"github.com/ditrit/shoset/msg"

	net "github.com/ditrit/shoset"
	"github.com/jinzhu/gorm"
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
				config := message.(cmsg.Secret)
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
func HandleConfiguration(c *net.ShosetConn, message msg.Message) (err error) {
	configuration := message.(cmsg.Configuration)
	ch := c.GetCh()
	//dir := c.GetDir()

	err = nil

	log.Println("Handle configuration")
	log.Println(configuration)
	fmt.Println("handle configuration")
	fmt.Println(configuration)
	//ok := ch.Queue["secret"].Push(secret, c.ShosetType, c.GetBindAddr())
	//if ok {
	if configuration.GetCommand() == "CONFIGURATION" {
		fmt.Println("CONFIG")
		var databaseClient *gorm.DB
		//databasePath := ch.Context["databasePath"].(string)
		if configuration.GetContext()["componentType"].(string) == "cluster" {
			fmt.Println("CLUSTER")
			databaseClient = ch.Context["gandalfDatabase"].(*gorm.DB)
		} else {
			fmt.Println("TENANT")
			mapDatabaseClient := ch.Context["tenantDatabases"].(map[string]*gorm.DB)
			//databaseBindAddr := ch.Context["databaseBindAddr"].(string)
			configurationCluster := ch.Context["configurationCluster"].(*models.ConfigurationCluster)

			if mapDatabaseClient != nil {
				databaseClient = cutils.GetDatabaseClientByTenant(configuration.GetTenant(), configurationCluster.DatabaseBindAddress, mapDatabaseClient)
			} else {
				log.Println("Database client map is empty")
				err = errors.New("Database client map is empty")
			}
		}

		if databaseClient != nil {
			bindAddr := configuration.GetContext()["bindAddress"].(string)
			logicalName := configuration.GetContext()["logicalName"].(string)

			switch configuration.GetContext()["componentType"] {
			case "cluster":
				fmt.Println("Cluster")
				config, err := cutils.GetConfigurationCluster(logicalName, databaseClient)
				if err == nil {
					if (config == models.ConfigurationLogicalCluster{}) {
						config = configuration.GetContext()["configuration"].(models.ConfigurationLogicalCluster)
						err = cutils.SaveConfigurationCluster(config, databaseClient)
						if err != nil {
							log.Println("Can't save logical configuration Cluster")
						}
					}
					fmt.Println("Configuration")
					fmt.Println(config)
					configMarshal, err := json.Marshal(config)
					if err == nil {
						fmt.Println("MARSHALL")
						target := ""
						configurationReply := cmsg.NewConfiguration(target, "CONFIGURATION_REPLY", string(configMarshal))
						configurationReply.Tenant = configuration.GetTenant()
						shoset := ch.ConnsJoin.Get(bindAddr)
						shoset.SendMessage(configurationReply)
					}
				} else {
					log.Println("Can't get logical configuration Cluster")
				}

				break
			case "aggregator":
				fmt.Println("Aggregator")
				config, err := cutils.GetConfigurationAggregator(logicalName, databaseClient)
				if err == nil {
					if (config == models.ConfigurationLogicalAggregator{}) {
						config = configuration.GetContext()["configuration"].(models.ConfigurationLogicalAggregator)
						err = cutils.SaveConfigurationAggregator(config, databaseClient)
						if err != nil {
							log.Println("Can't save logical configuration Aggregator")
						}
					}
					fmt.Println("Configuration")
					fmt.Println(config)
					configMarshal, err := json.Marshal(config)
					if err == nil {
						fmt.Println("MARSHALL")
						target := ""
						configurationReply := cmsg.NewConfiguration(target, "CONFIGURATION_REPLY", string(configMarshal))
						configurationReply.Tenant = configuration.GetTenant()
						shoset := ch.ConnsByAddr.Get(c.GetBindAddr())
						fmt.Println("shoset")
						fmt.Println(shoset)
						shoset.SendMessage(configurationReply)
					}
				} else {
					log.Println("Can't get logical configuration Aggregator")
				}

				break
			case "connector":
				fmt.Println("Connector")
				config, err := cutils.GetConfigurationConnector(logicalName, databaseClient)
				if err == nil {
					if (config == models.ConfigurationLogicalConnector{}) {
						config = configuration.GetContext()["configuration"].(models.ConfigurationLogicalConnector)
						err = cutils.SaveConfigurationConnector(config, databaseClient)
						if err != nil {
							log.Println("Can't save logical configuration Connector")
						}
					}
					fmt.Println("Configuration")
					fmt.Println(config)
					configMarshal, err := json.Marshal(config)
					if err == nil {
						fmt.Println("MARSHALL")
						target := configuration.GetTarget()
						configurationReply := cmsg.NewConfiguration(target, "CONFIGURATION_REPLY", string(configMarshal))
						configurationReply.Tenant = configuration.GetTenant()
						shoset := ch.ConnsByAddr.Get(c.GetBindAddr())
						shoset.SendMessage(configurationReply)
					}
				} else {
					log.Println("Can't get logical configuration Connector")

				}
				break
			}

		} else {
			log.Println("Can't get database client")
			err = errors.New("Can't get database client")
		}
	} else if configuration.GetCommand() == "CONFIGURATION_REPLY" {
		var configurationCluster *models.ConfigurationLogicalCluster
		err = json.Unmarshal([]byte(configuration.GetPayload()), &configurationCluster)
		if err == nil {
			ch.Context["logicalConfiguration"] = configurationCluster
		}
	}

	/* if dir == "out" {
		if c.GetShosetType() == "cl" {
			if secret.GetCommand() == "VALIDATION_REPLY" {
				ch.Context["validation"] = secret.GetPayload()
			}
		}
	} */
	/* 	} else {
		log.Println("Can't push to queue")
		err = errors.New("Can't push to queue")
	} */

	/* 	gandalfdatabaseClient := cutils.GetGandalfDatabaseClient(databasePath)
	   	if gandalfdatabaseClient != nil {

	   	} else {
	   		log.Println("Can't get gandalf database client")
	   		err = errors.New("Can't get gandalf database client")
	   	} */

	return err
}

//SendSecret :
func SendConfiguration(shoset *net.Shoset) (err error) {
	configurationCluster := shoset.Context["configuration"].(*models.ConfigurationCluster)

	configurationLogicalCluster := models.NewConfigurationLogicalCluster(configurationCluster.LogicalName)
	configurationMsg := cmsg.NewConfiguration("", "CONFIGURATION", "")
	//secretMsg.Tenant = "cluster"
	configurationMsg.GetContext()["componentType"] = "cluster"
	configurationMsg.GetContext()["logicalName"] = configurationCluster.LogicalName
	configurationMsg.GetContext()["bindAddress"] = configurationCluster.BindAddress
	configurationMsg.GetContext()["configuration"] = configurationLogicalCluster
	//conf.GetContext()["product"] = shoset.Context["product"]

	fmt.Println("shoset.ConnsByAddr")
	fmt.Println(shoset.ConnsByAddr)

	fmt.Println("shoset.ConnsJoin")
	fmt.Println(shoset.ConnsJoin)

	shosets := net.GetByType(shoset.ConnsJoin, "")
	fmt.Println("len(shosets)")
	fmt.Println(len(shosets))
	if len(shosets) != 0 {
		if configurationMsg.GetTimeout() > configurationCluster.MaxTimeout {
			configurationMsg.Timeout = configurationCluster.MaxTimeout
		}

		notSend := true
		for notSend {

			index := getConfigurationSendIndex(shosets)
			shosets[index].SendMessage(configurationMsg)
			log.Printf("%s : send command %s to %s\n", shoset.GetBindAddr(), configurationMsg.GetCommand(), shosets[index])

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
		log.Println("can't find cluster to send")
		err = errors.New("can't find cluster to send")
	}

	return err
}

// getCommandSendIndex : Aggregator getSendIndex function.
func getConfigurationSendIndex(conns []*net.ShosetConn) int {
	aux := configurationSendIndex
	configurationSendIndex++

	if configurationSendIndex >= len(conns) {
		configurationSendIndex = 0
	}

	return aux
}
