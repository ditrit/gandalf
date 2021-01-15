//Package shoset :
package shoset

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/ditrit/gandalf/core/cluster/utils"
	cutils "github.com/ditrit/gandalf/core/cluster/utils"

	cmsg "github.com/ditrit/gandalf/core/msg"
	"github.com/ditrit/shoset/msg"

	net "github.com/ditrit/shoset"
	"github.com/jinzhu/gorm"
)

var secretSendIndex = 0

func GetSecret(c *net.ShosetConn) (msg.Message, error) {
	var secret cmsg.Secret
	err := c.ReadMessage(&secret)
	return secret, err
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
	secret := message.(cmsg.Secret)
	ch := c.GetCh()
	dir := c.GetDir()

	err = nil

	log.Println("Handle secret")
	log.Println(secret)
	fmt.Println("dir")
	fmt.Println(dir)
	//ok := ch.Queue["secret"].Push(secret, c.ShosetType, c.GetBindAddr())
	//if ok {
	if secret.GetCommand() == "VALIDATION" {

		var databaseClient *gorm.DB
		//databasePath := ch.Context["databasePath"].(string)
		if secret.GetContext()["componentType"].(string) == "cluster" {
			fmt.Println("CLUSTER")
			databaseClient = ch.Context["gandalfDatabase"].(*gorm.DB)
		} else {
			mapDatabaseClient := ch.Context["tenantDatabases"].(map[string]*gorm.DB)
			databaseBindAddr := ch.Context["databaseBindAddr"].(string)
			if mapDatabaseClient != nil {
				databaseClient = cutils.GetDatabaseClientByTenant(secret.GetTenant(), databaseBindAddr, mapDatabaseClient)
			} else {
				log.Println("Database client map is empty")
				err = errors.New("Database client map is empty")
			}
		}

		if databaseClient != nil {
			bindAddr := secret.GetContext()["bindAddress"].(string)

			var result bool
			result, err = utils.ValidateSecret(databaseClient, secret.GetContext()["componentType"].(string), secret.GetContext()["logicalName"].(string), secret.GetContext()["secret"].(string), secret.GetContext()["bindAddress"].(string))
			fmt.Println("RESULT")
			fmt.Println(result)
			if err == nil {
				target := secret.GetTarget()
				if secret.GetContext()["componentType"] == "aggregator" || secret.GetContext()["componentType"] == "cluster" {
					target = ""
				}
				secretReply := cmsg.NewSecret(target, "VALIDATION_REPLY", strconv.FormatBool(result))
				secretReply.Tenant = secret.GetTenant()
				fmt.Println("strconv.FormatBool(result)")
				fmt.Println(strconv.FormatBool(result))
				fmt.Println("secretReply")
				fmt.Println(secretReply)

				var shoset *net.ShosetConn
				if secret.GetContext()["componentType"].(string) == "cluster" {
					fmt.Println("ch.ConnsJoin")
					fmt.Println(ch.ConnsJoin)
					fmt.Println("c.GetBindAddr()")
					fmt.Println(c.GetBindAddr())
					shoset = ch.ConnsJoin.Get(bindAddr)
				} else {
					shoset = ch.ConnsByAddr.Get(c.GetBindAddr())

				}
				fmt.Println("shoset")
				fmt.Println(shoset)
				shoset.SendMessage(secretReply)

			} else {
				log.Println("Can't validate secret")
				err = errors.New("Can't validate secret")
			}

		} else {
			log.Println("Can't get database client")
			err = errors.New("Can't get database client")
		}
	} else if secret.GetCommand() == "VALIDATION_REPLY" {
		ch.Context["validation"] = secret.GetPayload()
	}

	/* if dir == "out" {
		if c.GetShosetType() == "cl" {
			if secret.GetCommand() == "VALIDATION_REPLY" {
				ch.Context["validation"] = secret.GetPayload()
			}
		}
	} */
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

//SendSecret :
func SendSecret(shoset *net.Shoset, timeoutMax int64, logicalName, secret, bindAddress string) (err error) {
	secretMsg := cmsg.NewSecret("", "VALIDATION", "")
	//secretMsg.Tenant = "cluster"
	secretMsg.GetContext()["componentType"] = "cluster"
	secretMsg.GetContext()["logicalName"] = logicalName
	secretMsg.GetContext()["secret"] = secret
	secretMsg.GetContext()["bindAddress"] = bindAddress
	//conf.GetContext()["product"] = shoset.Context["product"]

	fmt.Println("shoset.ConnsByAddr")
	fmt.Println(shoset.ConnsByAddr)

	fmt.Println("shoset.ConnsJoin")
	fmt.Println(shoset.ConnsJoin)

	shosets := net.GetByType(shoset.ConnsJoin, "")
	fmt.Println("len(shosets)")
	fmt.Println(len(shosets))
	if len(shosets) != 0 {
		if secretMsg.GetTimeout() > timeoutMax {
			secretMsg.Timeout = timeoutMax
		}

		notSend := true
		for notSend {

			index := getSecretSendIndex(shosets)
			shosets[index].SendMessage(secretMsg)
			log.Printf("%s : send command %s to %s\n", shoset.GetBindAddr(), secretMsg.GetCommand(), shosets[index])

			timeoutSend := time.Duration((int(secretMsg.GetTimeout()) / len(shosets)))

			time.Sleep(timeoutSend * time.Millisecond)

			if shoset.Context["validation"] != nil {
				notSend = false
				break
			}
		}

		if notSend {
			return nil
		}

	} else {
		log.Println("can't find cluster to send")
		err = errors.New("can't find cluster to send")
	}

	return err
}

// getCommandSendIndex : Aggregator getSendIndex function.
func getSecretSendIndex(conns []*net.ShosetConn) int {
	aux := secretSendIndex
	secretSendIndex++

	if secretSendIndex >= len(conns) {
		secretSendIndex = 0
	}

	return aux
}
