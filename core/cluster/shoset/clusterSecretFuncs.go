//Package shoset :
package shoset

import (
	"errors"
	"log"

	"github.com/ditrit/gandalf/core/cluster/utils"

	"github.com/ditrit/gandalf/core/msg"

	net "github.com/ditrit/shoset"
	"github.com/jinzhu/gorm"
)

// HandleSecret :
func HandleSecret(c *net.ShosetConn, message msg.Message) (err error) {
	secret := message.(msg.Secret)
	ch := c.GetCh()

	err = nil

	log.Println("Handle secret")
	log.Println(secret)

	gandalfDatabaseClient := ch.Context["gandalfDatabase"].(*gorm.DB)
	if gandalfDatabaseClient != nil {
		if secret.GetCommand() == "VALIDATION" {

			var result bool
			result, err = utils.ValidateSecret(gandalfDatabaseClient, secret.GetContext()["componentType"], secret.GetContext()["logicalName"], secret.GetContext()["tenant"], secret.GetContext()["secret"])
			if err == nil {
				if secret.GetContext()["componentType"] == "aggregator" {
					secret.GetTarget() = ""
				}
				if result {
					cmdReply := msg.Secret(secret.GetTarget(), "VALIDATION_REPLY", "true")
					shoset := ch.ConnsByAddr.Get(c.GetBindAddr())

					shoset.SendMessage(cmdReply)
				} else {
					cmdReply := msg.Secret(secret.GetTarget(), "VALIDATION_REPLY", "false")
					shoset := ch.ConnsByAddr.Get(c.GetBindAddr())

					shoset.SendMessage(cmdReply)
				}
			} else {
				log.Println("Can't validate secret")
				err = errors.New("Can't validate secret")
			}
		}
	} else {
		log.Println("Can't get gandalf database")
		err = errors.New("Can't get gandalf database")
	}

	/* 	gandalfdatabaseClient := cutils.GetGandalfDatabaseClient(databasePath)
	   	if gandalfdatabaseClient != nil {

	   	} else {
	   		log.Println("Can't get gandalf database client")
	   		err = errors.New("Can't get gandalf database client")
	   	} */

	return err
}
