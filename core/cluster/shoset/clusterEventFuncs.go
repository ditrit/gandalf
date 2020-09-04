//Package shoset :
package shoset

import (
	"errors"
	"log"

	"github.com/ditrit/gandalf/core/cluster/utils"

	net "github.com/ditrit/shoset"
	"github.com/ditrit/shoset/msg"

	"github.com/jinzhu/gorm"
)

// HandleEvent : Cluster handle event function.
func HandleEvent(c *net.ShosetConn, message msg.Message) (err error) {
	evt := message.(msg.Event)
	ch := c.GetCh()
	thisOne := ch.GetBindAddr()
	err = nil

	log.Println("Handle event")
	log.Println(evt)

	ok := ch.Queue["evt"].Push(evt, c.ShosetType, c.GetBindAddr())

	if ok {
		mapDatabaseClient := ch.Context["tenantDatabases"].(map[string]*gorm.DB)
		databasePath := ch.Context["databasePath"].(string)
		if mapDatabaseClient != nil {
			databaseClient := utils.GetDatabaseClientByTenant(evt.GetTenant(), databasePath, mapDatabaseClient)
			if databaseClient != nil {
				ok := utils.CaptureMessage(message, "evt", databaseClient)
				if ok {
					log.Printf("Succes capture event %s on tenant %s \n", evt.GetEvent(), evt.GetTenant())
				} else {
					log.Printf("Fail capture event %s on tenant %s \n", evt.GetEvent(), evt.GetTenant())
					err = errors.New("Fail capture event" + evt.GetEvent() + " on tenant" + evt.GetTenant())
				}
			} else {
				log.Println("Can't get database client by tenant")
				err = errors.New("Can't get database client by tenant")
			}
		} else {
			log.Println("Database client map is empty")
			err = errors.New("Database client map is empty")
		}

		ch.ConnsByAddr.Iterate(
			func(key string, val *net.ShosetConn) {
				if key != c.GetBindAddr() && key != thisOne && val.ShosetType == "a" && c.GetCh().Context["tenant"] == val.GetCh().Context["tenant"] {
					val.SendMessage(evt)
					log.Printf("%s : send event %s to %s\n", thisOne, evt.GetEvent(), val)
				}
			},
		)
	}

	return err
}
