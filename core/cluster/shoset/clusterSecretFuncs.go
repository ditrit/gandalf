//Package shoset :
package shoset

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/ditrit/gandalf/core/cluster/utils"
	cutils "github.com/ditrit/gandalf/core/cluster/utils"

	cmsg "github.com/ditrit/gandalf/core/msg"
	"github.com/ditrit/shoset/msg"

	net "github.com/ditrit/shoset"
	"github.com/jinzhu/gorm"
)

func GetSecret(c *net.ShosetConn) (msg.Message, error) {
	var conf cmsg.Secret
	err := c.ReadMessage(&conf)
	return conf, err
}

// SendConfig :
func SendSecret(c *net.Shoset, cmd msg.Message) {
	fmt.Print("Sending Secret.\n")
	c.ConnsByAddr.Iterate(
		func(key string, conn *net.ShosetConn) {
			conn.SendMessage(cmd)
		},
	)
}

// WaitConfig :
func WaitSecret(c *net.Shoset, replies *msg.Iterator, args map[string]string, timeout int) *msg.Message {
	commandName, ok := args["name"]
	if !ok {
		return nil
	}
	term := make(chan *msg.Message, 1)
	cont := true
	go func() {
		for cont {
			message := replies.Get().GetMessage()
			if message != nil {
				config := message.(cmsg.Secret)
				if config.GetCommand() == commandName {
					term <- &message
				}
			} else {
				time.Sleep(time.Duration(10) * time.Millisecond)
			}
		}
	}()
	select {
	case res := <-term:
		cont = false
		return res
	case <-time.After(time.Duration(timeout) * time.Second):
		return nil
	}
}

// HandleSecret :
func HandleSecret(c *net.ShosetConn, message msg.Message) (err error) {
	fmt.Println("SECRET")
	secret := message.(cmsg.Secret)
	ch := c.GetCh()

	err = nil

	log.Println("Handle secret")
	log.Println(secret)

	//ok := ch.Queue["secret"].Push(secret, c.ShosetType, c.GetBindAddr())

	//if ok {
	mapDatabaseClient := ch.Context["tenantDatabases"].(map[string]*gorm.DB)
	databasePath := ch.Context["databasePath"].(string)
	if mapDatabaseClient != nil {
		databaseClient := cutils.GetDatabaseClientByTenant(secret.GetTenant(), databasePath, mapDatabaseClient)
		if databaseClient != nil {
			if secret.GetCommand() == "VALIDATION" {

				var result bool
				result, err = utils.ValidateSecret(databaseClient, secret.GetContext()["componentType"].(string), secret.GetContext()["logicalName"].(string), secret.GetContext()["instanceName"].(string), secret.GetContext()["secret"].(string))

				fmt.Println("componentType")
				fmt.Println(secret.GetContext()["componentType"].(string))
				if err == nil {
					target := secret.GetTarget()
					if secret.GetContext()["componentType"] == "aggregator" {
						target = ""
					}
					if result {
						secretReply := cmsg.NewSecret(target, "VALIDATION_REPLY", "true")
						secretReply.Tenant = secret.GetTenant()
						shoset := ch.ConnsByAddr.Get(c.GetBindAddr())

						shoset.SendMessage(secretReply)
					} else {
						secretReply := cmsg.NewSecret(target, "VALIDATION_REPLY", "false")
						secretReply.Tenant = secret.GetTenant()
						shoset := ch.ConnsByAddr.Get(c.GetBindAddr())

						shoset.SendMessage(secretReply)
					}
				} else {
					log.Println("Can't validate secret")
					err = errors.New("Can't validate secret")
				}
			}
		} else {
			log.Println("Can't get database client by tenant")
			err = errors.New("Can't get database client by tenant")
		}
	} else {
		log.Println("Database client map is empty")
		err = errors.New("Database client map is empty")
	}
	/* 	} else {
		log.Println("Can't push to queue")
		err = errors.New("Can't push to queue")
	} */

	/* 	gandalfdatabaseClient := cutils.GetGandalfDatabaseClient(databasePath)
	   	if gandalfdatabaseClient != nil {

	   	} else {
	   		log.Println("Can't get gandalf database client")
	   		err = errors.New("Can't get gandalf database client")
	   	} */

	return err
}
