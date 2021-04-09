//Package shoset :
package shoset

import (
	"errors"
	"log"

	cmodels "github.com/ditrit/gandalf/core/configuration/models"

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
		//ch.Context["tenant"] = secret.GetTenant()
		configurationConnector := ch.Context["configuration"].(*cmodels.ConfigurationConnector)
		configurationConnector.SetTenant(secret.GetTenant())
		//ch.Context["configurationConnector"] = configurationConnector
		ch.Context["validation"] = secret.GetPayload()
	}

	return err
}

//SendSecret :
func SendSecret(shoset *net.Shoset) (err error) {
	configurationConnector := shoset.Context["configuration"].(*cmodels.ConfigurationConnector)

	secretMsg := cmsg.NewSecret("", "VALIDATION", "")
	//secretMsg.Tenant = shoset.Context["tenant"].(string)
	secretMsg.GetContext()["componentType"] = "connector"
	secretMsg.GetContext()["secret"] = configurationConnector.GetSecret()
	secretMsg.GetContext()["bindAddress"] = configurationConnector.GetBindAddress()
	//conf.GetContext()["product"] = shoset.Context["product"]

	shosets := net.GetByType(shoset.ConnsByAddr, "a")

	if len(shosets) != 0 {
		if secretMsg.GetTimeout() > configurationConnector.GetMaxTimeout() {
			secretMsg.Timeout = configurationConnector.GetMaxTimeout()
		}

		notSend := true
		for start := time.Now(); time.Since(start) < time.Duration(secretMsg.GetTimeout())*time.Millisecond; {
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
		/* notSend := true
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
		*/
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
