//Package shoset :
package shoset

import (
	"errors"
	"fmt"
	"log"

	cmsg "github.com/ditrit/gandalf/core/msg"
	net "github.com/ditrit/shoset"
	"github.com/ditrit/shoset/msg"

	"time"
)

var secretSendIndex = 0

func GetSecret(c *net.ShosetConn) (msg.Message, error) {
	var conf cmsg.Secret
	err := c.ReadMessage(&conf)
	return conf, err
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
func SendSecret(shoset *net.Shoset, timeoutMax int64, logicalName, instanceName, tenant, secret string) (err error) {

	secretMsg := cmsg.NewSecret("", "VALIDATION", "")
	secretMsg.Tenant = shoset.Context["tenant"].(string)
	secretMsg.GetContext()["componentType"] = "connector"
	secretMsg.GetContext()["logicalName"] = logicalName
	secretMsg.GetContext()["instanceName"] = instanceName
	secretMsg.GetContext()["secret"] = secret
	//conf.GetContext()["product"] = shoset.Context["product"]

	shosets := net.GetByType(shoset.ConnsByAddr, "a")

	if len(shosets) != 0 {
		if secretMsg.GetTimeout() > timeoutMax {
			secretMsg.Timeout = timeoutMax
		}

		notSend := true
		for notSend {

			fmt.Println("SEND")

			index := getSecretSendIndex(shosets)
			shosets[index].SendMessage(secretMsg)
			log.Printf("%s : send command %s to %s\n", shoset.GetBindAddr(), secretMsg.GetCommand(), shosets[index])

			timeoutSend := time.Duration((int(secretMsg.GetTimeout()) / len(shosets)))
			fmt.Println("timeoutSend")
			fmt.Println(timeoutSend)

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
