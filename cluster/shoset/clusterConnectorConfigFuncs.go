//Package shoset :
package shoset

import (
	"encoding/json"
	"errors"
	cutils "gandalf-core/cluster/utils"
	"log"
	"shoset/msg"
	"shoset/net"
)

// HandleConnectorConfig : Cluster handle connector config function.
func HandleConnectorConfig(c *net.ShosetConn, message msg.Message) (err error) {
	conf := message.(msg.Config)
	ch := c.GetCh()
	thisOne := ch.GetBindAddr()

	err = nil

	log.Println("Handle connector config")
	log.Println(conf)

	databasePath := ch.Context["databasePath"].(string)

	databaseClient := cutils.GetGandalfDatabaseClient(databasePath)
	if databaseClient != nil {
		ok := cutils.CaptureMessage(message, "cmd", databaseClient)
		if ok {
			log.Printf("Succes capture command %s on tenant %s \n", conf.GetCommand(), conf.GetTenant())
		} else {
			log.Printf("Fail capture command %s on tenant %s \n", conf.GetCommand(), conf.GetTenant())
			err = errors.New("Fail capture command" + conf.GetCommand() + " on tenant" + conf.GetTenant())
		}

		configuration := cutils.GetConnectorConfiguration(conf, databaseClient)
		jsonData, err := json.Marshal(configuration)

		if err == nil {
			cmdReply := msg.NewConfig(conf.GetTarget(), "CONF_REPLY", string(jsonData))
			shoset := ch.ConnsByAddr.Get(thisOne)
			shoset.SendMessage(cmdReply)
		} else {
			log.Println("Can't unmarshall configuration")
			err = errors.New("Can't unmarshall configuration")
		}

	} else {
		log.Println("Can't get database client by tenant")
		err = errors.New("Can't get database client by tenant")
	}

	return err
}
