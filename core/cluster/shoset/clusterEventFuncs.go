//Package shoset :
package shoset

import (
	"errors"
	"log"

	"github.com/ditrit/gandalf/core/cluster/database"

	"github.com/ditrit/gandalf/core/cluster/utils"

	net "github.com/ditrit/shoset"
	"github.com/ditrit/shoset/msg"
)

// HandleEvent : Cluster handle event function.
func HandleEvent(c *net.ShosetConn, message msg.Message) (err error) {
	evt := message.(msg.Event)
	ch := c.GetCh()
	thisOne := ch.GetBindAddress()
	err = nil

	log.Println("Handle event")
	log.Println(evt)

	//ok := ch.Queue["evt"].Push(evt, c.GetRemoteShosetType(), c.GetBindAddress())

	//if ok {
	databaseConnection, ok := ch.Context["databaseConnection"].(*database.DatabaseConnection)
	if ok {
		//mapDatabaseClient := ch.Context["tenantDatabases"].(map[string]*gorm.DB)
		//databasePath := ch.Context["databasePath"].(string)
		//configurationCluster := ch.Context["configurationCluster"].(*cmodels.ConfigurationCluster)

		if databaseConnection != nil {
			databaseClient := databaseConnection.GetDatabaseClientByTenant(evt.GetTenant())
			if databaseClient != nil {
				ok := utils.CaptureMessage(message, "evt", databaseClient)
				if ok {
					log.Printf("Succes capture event %s on tenant %s \n", evt.GetEvent(), evt.GetTenant())
				} else {
					log.Printf("Error : Fail capture event %s on tenant %s \n", evt.GetEvent(), evt.GetTenant())
					err = errors.New("Fail capture event" + evt.GetEvent() + " on tenant" + evt.GetTenant())
				}
			} else {
				log.Println("Error : Can't get database client by tenant")
			}
		} else {
			log.Println("Error : Database client map is empty")
		}

		ch.ConnsByName.IterateAll(
			func(key string, val *net.ShosetConn) {
				if key != thisOne && val.GetRemoteShosetType() == "a" && c.GetCh().Context["tenant"] == val.GetCh().Context["tenant"] {
					val.SendMessage(evt)
					log.Printf("%s : send event %s to %s\n", thisOne, evt.GetEvent(), val)
				}
			},
		)
		//}
	}

	return err
}
