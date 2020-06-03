//Package shoset :
package shoset

import (
	"encoding/json"
	"errors"
	cutils "gandalf-core/cluster/utils"
	"log"
	"shoset/msg"
	"shoset/net"

	"github.com/jinzhu/gorm"
)

// HandleConnectorConfig : Cluster handle connector config function.
func HandleConnectorConfig(c *net.ShosetConn, message msg.Message) (err error) {
	conf := message.(msg.Config)
	ch := c.GetCh()

	err = nil

	log.Println("Handle connector config")
	log.Println(conf)

	mapDatabaseClient := ch.Context["database"].(map[string]*gorm.DB)
	databasePath := ch.Context["databasePath"].(string)
	if mapDatabaseClient != nil {
		databaseClient := cutils.GetDatabaseClientByTenant(conf.GetTenant(), databasePath, mapDatabaseClient)
		ok := cutils.CaptureMessage(message, "config", databaseClient)
		if ok {
			log.Printf("Succes capture config %s on tenant %s \n", conf.GetCommand(), conf.GetTenant())
		} else {
			log.Printf("Fail capture config %s on tenant %s \n", conf.GetCommand(), conf.GetTenant())
			err = errors.New("Fail capture command" + conf.GetCommand() + " on tenant" + conf.GetTenant())
		}

	} else {
		log.Println("Can't get database client by tenant")
		err = errors.New("Can't get database client by tenant")
	}

	gandalfdatabaseClient := cutils.GetGandalfDatabaseClient(databasePath)
	if gandalfdatabaseClient != nil {
		configuration := cutils.GetConnectorConfiguration(conf, gandalfdatabaseClient)
		jsonData, err := json.Marshal(configuration)

		if err == nil {
			cmdReply := msg.NewConfig(conf.GetTarget(), "CONF_REPLY", string(jsonData))
			shoset := ch.ConnsByAddr.Get(c.GetBindAddr())

			shoset.SendMessage(cmdReply)
		} else {
			log.Println("Can't unmarshall configuration")
			err = errors.New("Can't unmarshall configuration")
		}
	} else {
		log.Println("Can't get gandalf database client")
		err = errors.New("Can't get gandalf database client")
	}

	return err
}
