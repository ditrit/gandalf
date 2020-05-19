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

// HandleWorker : Cluster handle worker function.
func HandleWorker(c *net.ShosetConn, message msg.Message) (err error) {
	cmd := message.(msg.Command)
	ch := c.GetCh()
	thisOne := ch.GetBindAddr()

	err = nil

	log.Println("Handle worker")
	log.Println(cmd)

	databasePath := ch.Context["databasePath"].(string)

	databaseClient := cutils.GetGandalfDatabaseClient(databasePath)
	if databaseClient != nil {
		ok := cutils.CaptureMessage(message, "cmd", databaseClient)
		if ok {
			log.Printf("Succes capture command %s on tenant %s \n", cmd.GetCommand(), cmd.GetTenant())
		} else {
			log.Printf("Fail capture command %s on tenant %s \n", cmd.GetCommand(), cmd.GetTenant())
			err = errors.New("Fail capture command" + cmd.GetCommand() + " on tenant" + cmd.GetTenant())
		}

		conf := cutils.GetConnectorConfiguration(cmd, databaseClient)
		jsonData, err := json.Marshal(conf)

		if err == nil {
			cmdReply := msg.NewCommand(cmd.GetTarget(), "CONF_REPLY", string(jsonData))
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
