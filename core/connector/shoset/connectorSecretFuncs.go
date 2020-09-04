//Package shoset :
package shoset

import (
	"errors"
	"log"

	cmsg "github.com/ditrit/gandalf/core/msg"
	net "github.com/ditrit/shoset"
	"github.com/ditrit/shoset/msg"

	"time"
)

var secretSendIndex = 0

// HandleConnectorConfig : Connector handle connector config.
func HandleSecret(c *net.ShosetConn, message msg.Message) (err error) {
	secret := message.(cmsg.Secret)
	ch := c.GetCh()
	err = nil

	log.Println("Handle secret")
	log.Println(secret)

	if secret.GetCommand() == "VALIDATION_REPLY" {
		ch.Context["validation"] = secret.GetPayload()
	}

	return err
}

//SendSecret :
func SendSecret(shoset *net.Shoset, timeoutMax int64, logicalName, tenant, secret string) (err error) {

	secretMsg := cmsg.NewSecret("", "VALIDATION", "")
	secretMsg.Tenant = shoset.Context["tenant"].(string)
	secretMsg.GetContext()["componentType"] = "connector"
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
		log.Println("can't find aggregator to send")
		err = errors.New("can't find aggregator to send")
	}

	return err
}

// getSendIndex : Cluster getSendIndex function.
func getSecretSendIndex(conns []*net.ShosetConn) int {
	aux := secretSendIndex
	secretSendIndex++

	if secretSendIndex >= len(conns) {
		secretSendIndex = 0
	}

	return aux
}
