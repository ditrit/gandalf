//Package shoset :
package shoset

import (
	"errors"
	"log"

	"github.com/ditrit/gandalf/core/cluster/database"

	"github.com/ditrit/gandalf/core/models"

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
	//ok := ch.Queue["cmd"].Push(cmd, c.ShosetType, c.GetBindAddr())

	//if ok {
	//mapDatabaseClient := ch.Context["tenantDatabases"].(map[string]*gorm.DB)
	databaseConnection := ch.Context["databaseConnection"].(*database.DatabaseConnection)
	//databasePath := ch.Context["databasePath"].(string)
	//configurationCluster := ch.Context["configuration"].(*cmodels.ConfigurationCluster)
	if databaseConnection != nil {
		databaseClient := databaseConnection.GetDatabaseClientByTenant(cmd.GetTenant())
		if databaseClient != nil {
			ok := cutils.CaptureMessage(message, "cmd", databaseClient)
			if ok {
				log.Printf("Succes capture command %s on tenant %s \n", cmd.GetCommand(), cmd.GetTenant())
			} else {
				log.Printf("Fail capture command %s on tenant %s \n", cmd.GetCommand(), cmd.GetTenant())
				err = errors.New("Fail capture command" + cmd.GetCommand() + " on tenant" + cmd.GetTenant())
			}

			app := cutils.GetApplicationContext(cmd, databaseClient)

			if app != (models.Application{}) {
				mapConn := ch.ConnsByName.Get(app.Aggregator.LogicalName)
				if mapConn != nil {
					cmd.Target = app.Connector.LogicalName
					shosets := net.GetByType(ch.ConnsByName.Get(app.Aggregator.LogicalName), "a")

					if len(shosets) != 0 {
						index := getSendIndex(shosets)
						shosets[index].SendMessage(cmd)
					} else {
						log.Println("Can't find aggregators to send")
						err = errors.New("Can't find aggregators to send")
					}
				} else {
					log.Printf("Can't find connection with name %s \n", app.Aggregator.LogicalName)
					err = errors.New("Can't find connection with name " + app.Aggregator.LogicalName)
				}
			} else {
				log.Println("Can't find application context")
				err = errors.New("Can't find application context")
			}
		} else {
			log.Println("Can't get database client by tenant")
			err = errors.New("Can't get database client by tenant")
		}
	} else {
		log.Println("Database connection is empty")
		err = errors.New("Database connection is empty")
	}
	/* 	} else {
		log.Println("Can't push to queue")
		err = errors.New("Can't push to queue")
	} */

	return err
}

// getSendIndex : Cluster getSendIndex function.
func getSendIndex(conns []*net.ShosetConn) int {
	aux := sendIndex
	sendIndex++

	if sendIndex >= len(conns) {
		sendIndex = 0
	}

	return aux
}
