//Package shoset :
package shoset

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/ditrit/gandalf/core/cluster/database"

	"github.com/ditrit/gandalf/core/models"

	cutils "github.com/ditrit/gandalf/core/cluster/utils"

	net "github.com/ditrit/shoset"
	"github.com/ditrit/shoset/msg"
)

// HandleConnectorConfig : Cluster handle connector config function.
func HandleConnectorConfig(c *net.ShosetConn, message msg.Message) (err error) {
	conf := message.(msg.Config)
	ch := c.GetCh()

	err = nil

	log.Println("Handle connector config")
	log.Println(conf)

	databaseConnection := ch.Context["databaseConnection"].(*database.DatabaseConnection)
	//mapDatabaseClient := ch.Context["tenantDatabases"].(map[string]*gorm.DB)
	//databaseBindAddr := ch.Context["databaseBindAddr"].(string)
	//configurationCluster := ch.Context["configuration"].(*cmodels.ConfigurationCluster)

	if databaseConnection != nil {
		databaseClient := databaseConnection.GetDatabaseClientByTenant(conf.GetTenant())
		if databaseClient != nil {
			ok := cutils.CaptureMessage(message, "config", databaseClient)
			if ok {
				log.Printf("Succes capture config %s on tenant %s \n", conf.GetCommand(), conf.GetTenant())
			} else {
				log.Printf("Fail capture config %s on tenant %s \n", conf.GetCommand(), conf.GetTenant())
				err = errors.New("Fail capture command" + conf.GetCommand() + " on tenant" + conf.GetTenant())
			}
			if conf.GetCommand() == "CONFIG" {
				configurations := cutils.GetConnectorsConfiguration(databaseClient)
				jsonData, err := json.Marshal(configurations)

				if err == nil {
					cmdReply := msg.NewConfig(conf.GetTarget(), "CONFIG_REPLY", string(jsonData))
					cmdReply.Tenant = conf.GetTenant()
					shoset := ch.ConnsByAddr.Get(c.GetBindAddr())

					shoset.SendMessage(cmdReply)
				} else {
					log.Println("Can't unmarshall configuration")
					err = errors.New("Can't unmarshall configuration")
				}
			} else if conf.GetCommand() == "SAVE_CONFIG" {
				//connectorConfig := conf.GetContext()["connectorConfig"].(models.ConnectorConfig)
				var connectorConfig *models.ConnectorConfig
				err = json.Unmarshal([]byte(conf.GetPayload()), &connectorConfig)
				if err == nil {
					cutils.SaveConnectorsConfiguration(connectorConfig, databaseClient)
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
