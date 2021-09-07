//Package shoset :
package shoset

import (
	"errors"
	"fmt"
	"log"

	"github.com/ditrit/gandalf/core/cluster/database"

	cutils "github.com/ditrit/gandalf/core/cluster/utils"

	net "github.com/ditrit/shoset"
	"github.com/ditrit/shoset/msg"
)

var sendIndex = 0

// HandleCommand : Cluster handle command function.
func HandleCommand(c *net.ShosetConn, message msg.Message) (err error) {
	cmd := message.(msg.Command)
	ch := c.GetCh()
	err = nil

	log.Println("Handle command")
	log.Println(cmd)

	fmt.Println("Handle command")
	fmt.Println(cmd)
	//ok := ch.Queue["cmd"].Push(cmd, c.GetRemoteShosetType(), c.GetBindAddress())

	//if ok {
	//mapDatabaseClient := ch.Context["tenantDatabases"].(map[string]*gorm.DB)
	databaseConnection, ok := ch.Context["databaseConnection"].(*database.DatabaseConnection)
	if ok {
		//databasePath := ch.Context["databasePath"].(string)
		//configurationCluster := ch.Context["configuration"].(*cmodels.ConfigurationCluster)
		if databaseConnection != nil {
			databaseClient := databaseConnection.GetDatabaseClientByTenant(cmd.GetTenant())
			if databaseClient != nil {
				ok := cutils.CaptureMessage(message, "cmd", databaseClient)
				if ok {
					log.Printf("Succes capture command %s on tenant %s \n", cmd.GetCommand(), cmd.GetTenant())
				} else {
					log.Printf("Error : Fail capture command %s on tenant %s \n", cmd.GetCommand(), cmd.GetTenant())
					err = errors.New("Fail capture command" + cmd.GetCommand() + " on tenant" + cmd.GetTenant())
				}

				app := cutils.GetApplicationContext(cmd, databaseClient)

				if app.LogicalName != "" {
					mapConn := ch.ConnsByName.Get(app.Aggregator)
					if mapConn != nil {
						cmd.Target = app.LogicalName
						shosets := net.GetByType(ch.ConnsByName.Get(app.Aggregator), "a")

						if len(shosets) != 0 {
							index := getSendIndex(shosets)
							shosets[index].SendMessage(cmd)
						} else {
							log.Println("Error : Can't find aggregators to send")
						}
					} else {
						log.Printf("Error : Can't find connection with name %s \n", app.Aggregator)
					}
				} else {
					log.Println("Error : Can't find application context")
				}
			} else {
				log.Println("Error : Can't get database client by tenant")
			}
		} else {
			log.Println("Error : Database connection is empty")
		}
		/* 	} else {
			log.Println("Can't push to queue")
			err = errors.New("Can't push to queue")
		} */
	}

	return err
}

// getSendIndex : Cluster getSendIndex function.
func getSendIndex(conns []*net.ShosetConn) int {
	if sendIndex >= len(conns) {
		sendIndex = 0
	}

	aux := sendIndex
	sendIndex++
	return aux
}
