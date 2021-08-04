//Package shoset :
package shoset

import (
	"fmt"
	"log"
	"time"

	cmodels "github.com/ditrit/gandalf/core/configuration/models"

	cmsg "github.com/ditrit/gandalf/core/msg"
	net "github.com/ditrit/shoset"
	"github.com/ditrit/shoset/msg"
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

// HandleSecret :
func HandleSecret(c *net.ShosetConn, message msg.Message) (err error) {
	secret := message.(cmsg.Secret)
	ch := c.GetCh()
	dir := c.GetDir()
	err = nil
	thisOne := ch.GetBindAddress()
	log.Println("Handle secret")
	log.Println(secret)

	//if secret.GetTenant() == ch.Context["tenant"] {
	//ok := ch.Queue["secret"].Push(secret, c.GetRemoteShosetType(), c.GetBindAddress())
	//if ok {
	if dir == "in" {
		if c.GetRemoteShosetType() == "c" {
			shosets := ch.GetConnsByTypeArray("cl")
			if len(shosets) != 0 {
				secret.TargetAddress = c.GetRemoteAddress()
				secret.TargetLogicalName = c.GetRemoteLogicalName()
				configurationAggregator, ok := ch.Context["configuration"].(*cmodels.ConfigurationAggregator)
				if ok {
					secret.Tenant = configurationAggregator.GetTenant()
					index := getSecretSendIndex(shosets)
					shosets[index].SendMessage(secret)
					log.Printf("%s : send in secret %s to %s\n", thisOne, secret.GetCommand(), shosets[index])
				}
			} else {
				log.Println("Error : Can't find clusters to send")
			}
		} else {
			log.Println("Error : Wrong Shoset type")
		}
	}

	if dir == "out" {
		if c.GetRemoteShosetType() == "cl" {

			if secret.GetTargetLogicalName() == "" && secret.GetTargetAddress() == "" {
				if secret.GetCommand() == "VALIDATION_REPLY" {
					ch.Context["validation"] = secret.GetPayload()
				}
			} else {

				mapshoset := ch.ConnsByName.Get(secret.GetTargetLogicalName())
				var shoset *net.ShosetConn
				if mapshoset != nil {
					shoset = mapshoset.Get(secret.GetTargetAddress())
				}

				shoset.SendMessage(secret)
			}
		} else {
			log.Println("Error : Wrong Shoset type")
		}
	}
	/* } else {
		log.Println("can't push to queue")
		err = errors.New("can't push to queue")
	} */
	/*} else {
		log.Println("wrong tenant")
		err = errors.New("wrong tenant")
	}*/

	return err
}

//SendSecret :
func SendSecret(shoset *net.Shoset) (err error) {
	configurationAggregator, ok := shoset.Context["configuration"].(*cmodels.ConfigurationAggregator)
	if ok {
		secretMsg := cmsg.NewSecret("VALIDATION", "")

		secretMsg.Tenant = configurationAggregator.GetTenant()
		secretMsg.GetContext()["componentType"] = "aggregator"
		secretMsg.GetContext()["secret"] = configurationAggregator.GetSecret()
		secretMsg.GetContext()["bindAddress"] = configurationAggregator.GetBindAddress()
		//conf.GetContext()["product"] = shoset.Context["product"]
		fmt.Println("shoset.ConnsByName")
		fmt.Println(shoset.ConnsByName)
		shosets := shoset.GetConnsByTypeArray("cl")
		fmt.Println("shosets")
		fmt.Println(shosets)

		if len(shosets) != 0 {
			if secretMsg.GetTimeout() > configurationAggregator.GetMaxTimeout() {
				secretMsg.Timeout = configurationAggregator.GetMaxTimeout()
			}
			notSend := true
			for start := time.Now(); time.Since(start) < time.Duration(secretMsg.GetTimeout())*time.Millisecond; {
				index := getSecretSendIndex(shosets)
				shosets[index].SendMessage(secretMsg)
				log.Printf("%s : send command %s to %s\n", shoset.GetBindAddress(), secretMsg.GetCommand(), shosets[index])

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

				index := getSecretSendIndex(shosets)
				shosets[index].SendMessage(secretMsg)
				log.Printf("%s : send command %s to %s\n", shoset.GetBindAddress(), secretMsg.GetCommand(), shosets[index])

				timeoutSend := time.Duration((int(secretMsg.GetTimeout()) / len(shosets)))

				time.Sleep(timeoutSend * time.Millisecond)

				if shoset.Context["validation"] != nil {
					notSend = false
					break
				}
			}

			if notSend {
				return nil
			} */

		} else {
			log.Println("Error : Can't find cluster to send")
		}
	}

	return err
}

// getCommandSendIndex : Aggregator getSendIndex function.
func getSecretSendIndex(conns []*net.ShosetConn) int {
	if secretSendIndex >= len(conns) {
		secretSendIndex = 0
	}

	aux := secretSendIndex
	secretSendIndex++

	return aux
}
