//Package shoset :
package shoset

import (
	"errors"
	"log"
	"time"

	cmsg "github.com/ditrit/gandalf/core/msg"
	net "github.com/ditrit/shoset"
	"github.com/ditrit/shoset/msg"
)

var secretSendIndex = 0

// HandleSecret :
func HandleSecret(c *net.ShosetConn, message msg.Message) (err error) {
	secret := message.(cmsg.Secret)
	ch := c.GetCh()
	dir := c.GetDir()
	err = nil
	thisOne := ch.GetBindAddr()

	log.Println("Handle secret")
	log.Println(secret)

	if secret.GetTenant() == ch.Context["tenant"] {
		ok := ch.Queue["secret"].Push(secret, c.ShosetType, c.GetBindAddr())

		if ok {
			if dir == "in" {
				if c.GetShosetType() == "c" {
					shosets := net.GetByType(ch.ConnsByAddr, "cl")
					if len(shosets) != 0 {
						index := getSecretSendIndex(shosets)
						shosets[index].SendMessage(secret)
						log.Printf("%s : send in secret %s to %s\n", thisOne, secret.GetCommand(), shosets[index])
					} else {
						log.Println("can't find clusters to send")
						err = errors.New("can't find clusters to send")
					}
				} else {
					log.Println("wrong Shoset type")
					err = errors.New("wrong Shoset type")
				}
			}

			if dir == "out" {
				if c.GetShosetType() == "cl" {
					if secret.GetTarget() == "" {
						if secret.GetCommand() == "VALIDATION_REPLY" {
							ch.Context["validation"] = secret.GetPayload()
						}
					} else {
						shosets := net.GetByType(ch.ConnsByName.Get(secret.GetTarget()), "c")
						if len(shosets) != 0 {
							index := getCommandSendIndex(shosets)
							shosets[index].SendMessage(secret)
							log.Printf("%s : send out secret %s to %s\n", thisOne, secret.GetCommand(), shosets[index])
						} else {
							log.Println("can't find connectors to send")
							err = errors.New("can't find connectors to send")
						}
					}
				} else {
					log.Println("wrong Shoset type")
					err = errors.New("wrong Shoset type")
				}
			}
		} else {
			log.Println("can't push to queue")
			err = errors.New("can't push to queue")
		}
	} else {
		log.Println("wrong tenant")
		err = errors.New("wrong tenant")
	}

	return err
}

//SendSecret :
func SendSecret(shoset *net.Shoset, timeoutMax int64, logicalName, tenant, secret string) (err error) {

	secretMsg := cmsg.NewSecret("", "VALIDATION", "")
	secretMsg.Tenant = shoset.Context["tenant"].(string)
	secretMsg.GetContext()["componentType"] = "aggregator"
	secretMsg.GetContext()["logicalName"] = logicalName
	secretMsg.GetContext()["tenant"] = tenant
	secretMsg.GetContext()["secret"] = secret
	//conf.GetContext()["product"] = shoset.Context["product"]

	shosets := net.GetByType(shoset.ConnsByAddr, "cl")

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

			if shoset.Context["validation"] != nil {
				notSend = false
				break
			}
			time.Sleep(timeoutSend)
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
