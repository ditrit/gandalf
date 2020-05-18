//Package shoset :
package shoset

import (
	"errors"
	"fmt"
	cutils "gandalf-core/cluster/utils"
	"log"
	"shoset/msg"
	"shoset/net"
)

// HandleWorker : Cluster handle worker function.
func HandleWorker(c *net.ShosetConn, message msg.Message) (err error) {
	cmd := message.(msg.Command)
	ch := c.GetCh()
	err = nil

	log.Println("Handle worker")
	log.Println(cmd)

	databasePath := ch.Context["databasePath"].(string)

	// TODO REVOIR
	databaseClient := cutils.GetGandalfDatabaseClient(databasePath)
	if databaseClient != nil {
		ok := cutils.CaptureMessage(message, "cmd", databaseClient)
		if ok {
			log.Printf("Succes capture command %s on tenant %s \n", cmd.GetCommand(), cmd.GetTenant())
		} else {
			log.Printf("Fail capture command %s on tenant %s \n", cmd.GetCommand(), cmd.GetTenant())
			err = errors.New("Fail capture command" + cmd.GetCommand() + " on tenant" + cmd.GetTenant())
		}

		app := cutils.GetConnectorConfiguration(cmd, databaseClient)
		fmt.Println(app)
		/* 		if app != (models.Application{}) {
		   			cmd.Target = app.Connector
		   			shosets := net.GetByType(ch.ConnsByName.Get(app.Aggregator), "a")

		   			if len(shosets) != 0 {
		   				index := getSendIndex(shosets)
		   				shosets[index].SendMessage(cmd)
		   			} else {
		   				log.Println("Can't find aggregators to send")
		   				err = errors.New("Can't find aggregators to send")
		   			}
		   		} else {
		   			log.Println("Can't find application context")
		   			err = errors.New("Can't find application context")
		   		} */
	} else {
		log.Println("Can't get database client by tenant")
		err = errors.New("Can't get database client by tenant")
	}

	return err
}
