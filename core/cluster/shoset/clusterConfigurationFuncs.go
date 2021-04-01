//Package shoset :
package shoset

import (
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/ditrit/gandalf/core/cluster/database"
	"github.com/ditrit/gandalf/core/models"
	cmsg "github.com/ditrit/gandalf/core/msg"

	cutils "github.com/ditrit/gandalf/core/cluster/utils"

	net "github.com/ditrit/shoset"
	"github.com/ditrit/shoset/msg"
)

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

	databaseConnection := ch.Context["databaseConnection"].(*database.DatabaseConnection)
	//mapDatabaseClient := ch.Context["tenantDatabases"].(map[string]*gorm.DB)
	//databaseBindAddr := ch.Context["databaseBindAddr"].(string)
	//configurationCluster := ch.Context["configuration"].(*cmodels.ConfigurationCluster)

	if databaseConnection != nil {
		databaseClient := databaseConnection.GetDatabaseClientByTenant(configuration.GetTenant())
		if databaseClient != nil {
			ok := cutils.CaptureMessage(message, "config", databaseClient)
			if ok {
				log.Printf("Succes capture config %s on tenant %s \n", configuration.GetCommand(), configuration.GetTenant())
			} else {
				log.Printf("Fail capture config %s on tenant %s \n", configuration.GetCommand(), configuration.GetTenant())
				err = errors.New("Fail capture command" + configuration.GetCommand() + " on tenant" + configuration.GetTenant())
			}
			if configuration.GetCommand() == "PIVOT_CONFIGURATION" {
				connectorType := configuration.Context["connectorType"].(string)
				version := configuration.Context["version"].(models.Version)
				pivots := cutils.GetPivots(databaseClient, connectorType, version)
				jsonData, err := json.Marshal(pivots)

				if err == nil {
					cmdReply := msg.NewConfig(configuration.GetTarget(), "PIVOT_CONFIGURATION_REPLY", string(jsonData))
					cmdReply.Tenant = configuration.GetTenant()
					cmdReply.Context["connectorType"] = connectorType

					shoset := ch.ConnsByAddr.Get(c.GetBindAddr())

					shoset.SendMessage(cmdReply)
				} else {
					log.Println("Can't unmarshall configuration")
					err = errors.New("Can't unmarshall configuration")
				}
			} else if configuration.GetCommand() == "CONNECTOR_PRODUCT_CONFIGURATION" {
				product := configuration.Context["product"].(string)
				version := configuration.Context["version"].(models.Version)
				productConnectors := cutils.GetProductConnectors(databaseClient, product, version)
				jsonData, err := json.Marshal(productConnectors)

				if err == nil {
					cmdReply := msg.NewConfig(configuration.GetTarget(), "CONNECTOR_PRODUCT_CONFIGURATION_REPLY", string(jsonData))
					cmdReply.Tenant = configuration.GetTenant()
					shoset := ch.ConnsByAddr.Get(c.GetBindAddr())

					shoset.SendMessage(cmdReply)
				} else {
					log.Println("Can't unmarshall configuration")
					err = errors.New("Can't unmarshall configuration")
				}
			} else if configuration.GetCommand() == "SAVE_PIVOT_CONFIGURATION" {
				//connectorConfig := conf.GetContext()["connectorConfig"].(models.ConnectorConfig)
				var pivot *models.Pivot
				err = json.Unmarshal([]byte(configuration.GetPayload()), &pivot)
				if err == nil {
					cutils.SavePivot(pivot, databaseClient)
				}
			} else if configuration.GetCommand() == "SAVE_PRODUCT_CONNECTOR_CONFIGURATION" {
				//connectorConfig := conf.GetContext()["connectorConfig"].(models.ConnectorConfig)
				var productConnector *models.ProductConnector
				err = json.Unmarshal([]byte(configuration.GetPayload()), &productConnector)
				if err == nil {
					cutils.SaveProductConnector(productConnector, databaseClient)
				}
			}
		} else {
			log.Println("Can't get database client by tenant")
			err = errors.New("Can't get database client by tenant")
		}
	} else {
		log.Println("Can't get database clients")
		err = errors.New("Can't get database clients")
	}

	/* 	gandalfdatabaseClient := cutils.GetGandalfDatabaseClient(databasePath)
	   	if gandalfdatabaseClient != nil {

	   	} else {
	   		log.Println("Can't get gandalf database client")
	   		err = errors.New("Can't get gandalf database client")
	   	} */

	return err
}
