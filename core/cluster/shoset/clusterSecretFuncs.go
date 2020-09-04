//Package shoset :
package shoset

import (
	"errors"
	"log"

	"github.com/ditrit/gandalf/core/cluster/utils"

	cmsg "github.com/ditrit/gandalf/core/msg"
	"github.com/ditrit/shoset/msg"

	net "github.com/ditrit/shoset"
	"github.com/jinzhu/gorm"
)

// HandleSecret :
func HandleSecret(c *net.ShosetConn, message msg.Message) (err error) {
	secret := message.(cmsg.Secret)
	ch := c.GetCh()

	err = nil

	log.Println("Handle secret")
	log.Println(secret)

	gandalfDatabaseClient := ch.Context["gandalfDatabase"].(*gorm.DB)
	if gandalfDatabaseClient != nil {
		if secret.GetCommand() == "VALIDATION" {

			var result bool
			result, err = utils.ValidateSecret(gandalfDatabaseClient, secret.GetContext()["componentType"].(string), secret.GetContext()["logicalName"].(string), secret.GetContext()["tenant"].(string), secret.GetContext()["secret"].(string))
			if err == nil {
				target := secret.GetTarget()
				if secret.GetContext()["componentType"] == "aggregator" {
					target = ""
				}
				if result {
					secretReply := cmsg.NewSecret(target, "VALIDATION_REPLY", "true")
					shoset := ch.ConnsByAddr.Get(c.GetBindAddr())

					shoset.SendMessage(secretReply)
				} else {
					secretReply := cmsg.NewSecret(target, "VALIDATION_REPLY", "false")
					shoset := ch.ConnsByAddr.Get(c.GetBindAddr())

					shoset.SendMessage(secretReply)
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
